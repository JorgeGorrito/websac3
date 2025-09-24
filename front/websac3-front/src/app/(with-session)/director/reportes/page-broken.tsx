"use client";

import React, { useState } from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { 
  FileText, 
  BookOpen, 
  Clock, 
  GraduationCap, 
  Target, 
  Search, 
  Settings, 
  Calendar,
  Eye,
  Download,
  Star,
  Calendar as CalendarIcon
} from "lucide-react";
import jsPDF from 'jspdf';
import html2canvas from 'html2canvas';
import { useListDegreeProgramsQuery, useListDegreeProgramReportsQuery, useGetReportDetailQuery } from "@/services/api";
import { useRouter, useSearchParams } from "next/navigation";
import { ProgramsPagination } from "@/components/websac3/program/ProgramsPagination";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { useDispatch } from "react-redux";
import { showError } from "@/store/errorSlice";

// Component for handling images with fallbacks
const ReportImage = ({ 
  src, 
  alt, 
  className, 
  fallbackText 
}: { 
  src: string; 
  alt: string; 
  className: string; 
  fallbackText?: string; 
}) => {
  const [imageError, setImageError] = useState(false);
  
  if (imageError) {
    return (
      <div className={`flex items-center justify-center ${className}`}>
        <span className="text-sm font-medium text-gray-600">
          {fallbackText || alt}
        </span>
      </div>
    );
  }
  
  return (
    <img 
      src={src} 
      alt={alt} 
      className={className}
      onError={() => setImageError(true)}
    />
  );
};

export default function ReportesPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const dispatch = useDispatch();
  const [selectedProgramId, setSelectedProgramId] = useState<number | null>(null);
  const [selectedProgramName, setSelectedProgramName] = useState<string>("");
  const [isReportsModalOpen, setIsReportsModalOpen] = useState(false);
  const [reportsCurrentPage, setReportsCurrentPage] = useState(1);
  const [reportsItemsPerPage, setReportsItemsPerPage] = useState(10);
  const [selectedReportId, setSelectedReportId] = useState<number | null>(null);
  const [isReportDetailModalOpen, setIsReportDetailModalOpen] = useState(false);
  
  // Get parameters from URL or use defaults
  const [currentPage, setCurrentPage] = useState(() => {
    const page = searchParams.get('page');
    return page ? parseInt(page, 10) : 1;
  });
  
  const [itemsPerPage, setItemsPerPage] = useState(() => {
    const perPage = searchParams.get('per_page');
    return perPage ? parseInt(perPage, 10) : 6;
  });
  
  const [filters, setFilters] = useState(() => {
    const name = searchParams.get('name') || '';
    const snies = searchParams.get('snies') || '';
    return { name, snies };
  });
  
  const [showFilters, setShowFilters] = useState(false);

  // Build query parameters with proper filter format
  const buildFilters = () => {
    const filterParams: Record<string, string> = {};
    
    if (filters.name.trim()) {
      filterParams['name[cont]'] = filters.name.trim();
    }
    
    if (filters.snies.trim()) {
      filterParams['snies[cont]'] = filters.snies.trim();
    }
    
    return filterParams;
  };

  const queryParams = {
    current_page: currentPage,
    items_per_page: itemsPerPage,
    filters: buildFilters()
  };

  // Fetch degree programs
  const { data, isLoading, error } = useListDegreeProgramsQuery(queryParams);
  
  // Fetch reports for selected program
  const { data: reportsData, isLoading: reportsLoading, error: reportsError } = useListDegreeProgramReportsQuery(
    { 
      degree_program_id: selectedProgramId!, 
      current_page: reportsCurrentPage, 
      items_per_page: reportsItemsPerPage 
    },
    { skip: !selectedProgramId }
  );

  // Fetch report detail
  const { data: reportDetail, isLoading: reportDetailLoading, error: reportDetailError } = useGetReportDetailQuery(
    selectedReportId!,
    { skip: !selectedReportId }
  );

  // Function to update URL with current parameters
  const updateURL = (newParams: {
    page?: number;
    per_page?: number;
    name?: string;
    snies?: string;
  }) => {
    const params = new URLSearchParams();
    
    const page = newParams.page ?? currentPage;
    const perPage = newParams.per_page ?? itemsPerPage;
    const name = newParams.name ?? filters.name;
    const snies = newParams.snies ?? filters.snies;
    
    if (page > 1) params.set('page', page.toString());
    if (perPage !== 6) params.set('per_page', perPage.toString());
    if (name.trim()) params.set('name', name.trim());
    if (snies.trim()) params.set('snies', snies.trim());
    
    const newURL = params.toString() 
      ? `/director/reportes?${params.toString()}`
      : '/director/reportes';
    
    router.push(newURL, { scroll: false });
  };

  // Handle page change
  const handlePageChange = (page: number) => {
    setCurrentPage(page);
    updateURL({ page });
  };

  // Handle items per page change
  const handleItemsPerPageChange = (perPage: number) => {
    setItemsPerPage(perPage);
    setCurrentPage(1); // Reset to first page
    updateURL({ per_page: perPage, page: 1 });
  };

  // Handle filter changes
  const handleFilterChange = (key: string, value: string) => {
    const newFilters = { ...filters, [key]: value };
    setFilters(newFilters);
    setCurrentPage(1); // Reset to first page
    updateURL({ 
      name: newFilters.name, 
      snies: newFilters.snies, 
      page: 1 
    });
  };

  // Clear all filters
  const clearFilters = () => {
    setFilters({ name: '', snies: '' });
    setCurrentPage(1);
    updateURL({ name: '', snies: '', page: 1 });
  };

  const handleViewReports = (programId: number, programName: string) => {
    setSelectedProgramId(programId);
    setSelectedProgramName(programName);
    setReportsCurrentPage(1);
    setIsReportsModalOpen(true);
  };

  const handleViewReport = (reportId: number) => {
    setSelectedReportId(reportId);
    setIsReportDetailModalOpen(true);
  };

  const handleReportsPageChange = (page: number) => {
    setReportsCurrentPage(page);
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('es-ES', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  const getScoreColor = (score: number) => {
    // Convert decimal score to percentage for color logic
    const percentage = score * 100;
    if (percentage >= 80) return 'text-green-600 bg-green-100 border-green-200';
    if (percentage >= 60) return 'text-yellow-600 bg-yellow-100 border-yellow-200';
    return 'text-red-600 bg-red-100 border-red-200';
  };

  const getScoreLabel = (score: number) => {
    // Convert decimal score to percentage for label logic
    const percentage = score * 100;
    if (percentage >= 80) return 'Excelente';
    if (percentage >= 60) return 'Bueno';
    return 'Necesita Mejora';
  };

  const formatScore = (score: number) => {
    // Show score as decimal with 2 decimal places (e.g., 0.09)
    return score.toFixed(2);
  };

  const downloadReportAsPDF = async () => {
    if (!reportDetail) return;

    try {
      // Create a simple text-based PDF
      const pdf = new jsPDF('p', 'mm', 'a4');
      
      // Add title
      pdf.setFontSize(16);
      pdf.text('Reporte de Evaluación - WebSAC3', 20, 20);
      
      // Add program info
      pdf.setFontSize(12);
      let yPosition = 40;
      
      if (reportDetail.degree_program?.name) {
        pdf.text(`Programa: ${reportDetail.degree_program.name}`, 20, yPosition);
        yPosition += 10;
      }
      
      if (reportDetail.degree_program?.snies) {
        pdf.text(`SNIES: ${reportDetail.degree_program.snies}`, 20, yPosition);
        yPosition += 10;
      }
      
      if (reportDetail.professional_role?.name) {
        pdf.text(`Rol Profesional: ${reportDetail.professional_role.name}`, 20, yPosition);
        yPosition += 10;
      }
      
      pdf.text(`Puntaje Total: ${formatScore(reportDetail.score || 0)}`, 20, yPosition);
      yPosition += 20;
      
      // Add knowledge areas
      pdf.text('Áreas de Conocimiento:', 20, yPosition);
      yPosition += 10;
      
      reportDetail.knowledge_area_reports?.forEach((area, index) => {
        if (yPosition > 270) {
          pdf.addPage();
          yPosition = 20;
        }
        
        pdf.setFontSize(10);
        pdf.text(`${index + 1}. ${area.name}`, 25, yPosition);
        yPosition += 8;
        pdf.text(`   Horas esperadas: ${area.total_learn_hours_expected}`, 25, yPosition);
        yPosition += 6;
        pdf.text(`   Horas alcanzadas: ${area.total_learn_hours_actual}`, 25, yPosition);
        yPosition += 6;
        pdf.text(`   Puntaje: ${formatScore(area.score_got)}`, 25, yPosition);
        yPosition += 10;
      });
      
      // Generate filename
      const programName = reportDetail.degree_program?.name || 'Programa';
      const date = new Date(reportDetail.created_at).toLocaleDateString('es-ES');
      const filename = `Reporte_${programName.replace(/[^a-zA-Z0-9]/g, '_')}_${date.replace(/\//g, '-')}.pdf`;
      
      // Download the PDF
      pdf.save(filename);
    } catch (error) {
      console.error('Error generating PDF:', error);
      dispatch(showError('Error al generar el PDF. Inténtalo de nuevo.'));
    }
  };

  if (isLoading) {
    return <LoadingSkeleton />;
  }

  if (error) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
          <div className="text-center text-gray-500">
            <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
            <p>No se pudieron cargar los programas de grado</p>
            <p className="text-sm mt-2">Los errores se mostrarán en un modal automáticamente</p>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Consultar Reportes</h1>
          <p className="text-gray-600 mt-2">Consulta los reportes de evaluación de ciberseguridad de los programas de grado</p>
        </div>
        <Button 
          onClick={() => setShowFilters(!showFilters)}
          variant="outline"
          className={`flex items-center border-2 transition-all duration-200 ${
            showFilters 
              ? 'border-slate-500 bg-slate-100 text-slate-700' 
              : 'border-slate-300 hover:border-slate-400 hover:bg-slate-50'
          }`}
        >
          <Settings className="h-4 w-4 mr-2" />
          {showFilters ? 'Ocultar Filtros' : 'Mostrar Filtros'}
        </Button>
      </div>

      {/* Compact Stats Bar */}
      {data && (
        <div className="bg-white border border-gray-200 rounded-lg shadow-sm p-4">
          <div className="flex flex-wrap items-center justify-between gap-4">
            {/* Left side - Main info */}
            <div className="flex items-center gap-6">
              <div className="flex items-center gap-2">
                <BookOpen className="h-4 w-4 text-blue-600" />
                <span className="text-sm font-medium text-gray-700">
                  {data.total_count} programa{data.total_count !== 1 ? 's' : ''} disponible{data.total_count !== 1 ? 's' : ''}
                </span>
              </div>
              
              <div className="flex items-center gap-2">
                <Calendar className="h-4 w-4 text-gray-500" />
                <span className="text-sm text-gray-600">
                  Página {data.current_page} de {Math.ceil(data.total_count / itemsPerPage)}
                </span>
              </div>
            </div>

            {/* Right side - Institution info */}
            {data.data.length > 0 && (
              <div className="flex items-center gap-3">
                <div className="flex items-center gap-2">
                  <GraduationCap className="h-4 w-4 text-indigo-600" />
                  <span className="text-sm font-medium text-gray-700">
                    {data.data[0].higher_education_institution.name}
                  </span>
                </div>
                <Badge className="bg-indigo-100 text-indigo-800 border-indigo-200 text-xs">
                  SNIES: {data.data[0].higher_education_institution.snies}
                </Badge>
              </div>
            )}
          </div>
        </div>
      )}

      {/* Active Filters Info */}
      {(filters.name || filters.snies) && (
        <div className="bg-amber-50 border border-amber-200 rounded-lg p-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3">
              <div className="p-2 bg-amber-500 rounded-lg">
                <Search className="h-4 w-4 text-white" />
              </div>
              <div>
                <span className="text-sm font-semibold text-amber-800">Filtros activos:</span>
                <div className="flex items-center gap-2 mt-1">
                  {filters.name && (
                    <Badge className="bg-blue-100 text-blue-800 border-blue-200 font-medium text-xs">
                      Nombre: "{filters.name}"
                    </Badge>
                  )}
                  {filters.snies && (
                    <Badge className="bg-green-100 text-green-800 border-green-200 font-medium text-xs">
                      SNIES: "{filters.snies}"
                    </Badge>
                  )}
                </div>
              </div>
            </div>
            <Button 
              onClick={clearFilters} 
              variant="outline" 
              size="sm"
              className="border-amber-300 text-amber-700 hover:bg-amber-100 text-xs"
            >
              Limpiar
            </Button>
          </div>
        </div>
      )}

      {/* Filters Panel */}
      {showFilters && (
        <Card className="bg-white border border-gray-200 shadow-lg py-0">
          <CardHeader className="bg-gray-50 border-b border-gray-200 pb-4 pt-4">
            <CardTitle className="flex items-center gap-3 text-lg font-semibold text-gray-800">
              <div className="p-2 bg-blue-100 rounded-lg">
                <Search className="h-5 w-5 text-blue-600" />
              </div>
              Filtros y Configuración
            </CardTitle>
            <CardDescription className="text-gray-600 mt-1">
              Personaliza tu búsqueda y visualización de programas
            </CardDescription>
          </CardHeader>
          <CardContent className="p-6">
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              {/* Name Filter */}
              <div className="space-y-2">
                <Label htmlFor="name-filter" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                  <BookOpen className="h-4 w-4 text-blue-500" />
                  Nombre del Programa
                </Label>
                <Input
                  id="name-filter"
                  placeholder="Buscar por nombre..."
                  value={filters.name}
                  onChange={(e) => handleFilterChange('name', e.target.value)}
                  className="h-10 border-gray-300 focus:border-blue-500 focus:ring-1 focus:ring-blue-500"
                />
              </div>
              
              {/* SNIES Filter */}
              <div className="space-y-2">
                <Label htmlFor="snies-filter" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                  <Target className="h-4 w-4 text-green-500" />
                  Código SNIES
                </Label>
                <Input
                  id="snies-filter"
                  placeholder="Buscar por SNIES..."
                  value={filters.snies}
                  onChange={(e) => handleFilterChange('snies', e.target.value)}
                  className="h-10 border-gray-300 focus:border-green-500 focus:ring-1 focus:ring-green-500"
                />
              </div>
              
              {/* Items Per Page */}
              <div className="space-y-2">
                <Label htmlFor="per-page" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                  <Clock className="h-4 w-4 text-purple-500" />
                  Elementos por página
                </Label>
                <select
                  id="per-page"
                  value={itemsPerPage}
                  onChange={(e) => handleItemsPerPageChange(parseInt(e.target.value))}
                  className="w-full h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-purple-500 focus:border-purple-500 bg-white"
                >
                  <option value={3}>3 por página</option>
                  <option value={6}>6 por página</option>
                  <option value={9}>9 por página</option>
                  <option value={12}>12 por página</option>
                  <option value={15}>15 por página</option>
                </select>
              </div>
            </div>
            
            {/* Clear Filters */}
            <div className="flex justify-end mt-6 pt-4 border-t border-gray-100">
              <Button 
                onClick={clearFilters} 
                variant="outline" 
                size="sm"
                className="text-gray-600 border-gray-300 hover:bg-gray-50"
              >
                Limpiar Filtros
              </Button>
            </div>
          </CardContent>
        </Card>
      )}

      {/* Programs List */}
      <div>
        {isLoading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {Array.from({ length: 6 }).map((_, index) => (
              <Card key={index} className="py-0">
                <CardHeader>
                  <Skeleton className="h-6 w-3/4" />
                  <Skeleton className="h-4 w-1/2" />
                </CardHeader>
                <CardContent>
                  <div className="space-y-2">
                    <Skeleton className="h-4 w-full" />
                    <Skeleton className="h-4 w-2/3" />
                    <Skeleton className="h-4 w-1/2" />
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        ) : data && data.data.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {data.data.map((program) => (
              <Card key={program.id} className="bg-white border border-gray-200 shadow-lg hover:shadow-xl transition-all duration-300 hover:-translate-y-1 py-0">
                <CardHeader className="pb-3 pt-4 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-gray-100">
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <CardTitle className="text-xl font-bold text-gray-900 line-clamp-1 mb-2">
                        {program.name}
                      </CardTitle>
                      <div className="flex items-center gap-2">
                        <Badge className="bg-blue-100 text-blue-800 border-blue-200 font-semibold">
                          SNIES: {program.snies}
                        </Badge>
                      </div>
                    </div>
                    <Badge className="bg-gradient-to-r from-green-500 to-emerald-500 text-white font-semibold px-3 py-1">
                      {program.duration_unit.name}
                    </Badge>
                  </div>
                </CardHeader>
                
                <CardContent className="p-5 space-y-4">
                  {/* Program Details */}
                  <div className="grid grid-cols-2 gap-4">
                    <div className="flex items-center p-3 bg-blue-50 rounded-lg">
                      <div className="p-2 bg-blue-500 rounded-lg mr-3">
                        <GraduationCap className="h-4 w-4 text-white" />
                      </div>
                      <div>
                        <p className="text-xs font-medium text-blue-600 uppercase tracking-wide">Créditos</p>
                        <p className="text-lg font-bold text-blue-900">{program.total_credits}</p>
                      </div>
                    </div>
                    <div className="flex items-center p-3 bg-green-50 rounded-lg">
                      <div className="p-2 bg-green-500 rounded-lg mr-3">
                        <Clock className="h-4 w-4 text-white" />
                      </div>
                      <div>
                        <p className="text-xs font-medium text-green-600 uppercase tracking-wide">Duración</p>
                        <p className="text-lg font-bold text-green-900">{program.duration_value} {program.duration_unit.name.toLowerCase()}</p>
                      </div>
                    </div>
                  </div>

                  {/* Program Focus */}
                  <div className="p-4 bg-purple-50 rounded-lg border border-purple-100">
                    <div className="flex items-center mb-2">
                      <div className="p-2 bg-purple-500 rounded-lg mr-3">
                        <Target className="h-4 w-4 text-white" />
                      </div>
                      <span className="text-sm font-semibold text-purple-700 uppercase tracking-wide">Enfoque</span>
                    </div>
                    <p className="text-sm text-gray-700 line-clamp-2 pl-11">
                      {program.program_focus}
                    </p>
                  </div>

                  {/* Institution */}
                  <div className="p-4 bg-indigo-50 rounded-lg border border-indigo-100">
                    <div className="flex items-center mb-2">
                      <div className="p-2 bg-indigo-500 rounded-lg mr-3">
                        <GraduationCap className="h-4 w-4 text-white" />
                      </div>
                      <span className="text-sm font-semibold text-indigo-700 uppercase tracking-wide">Institución</span>
                    </div>
                    <p className="text-sm font-medium text-gray-800 pl-11 mb-1">
                      {program.higher_education_institution.name}
                    </p>
                    <Badge className="ml-11 bg-indigo-100 text-indigo-800 border-indigo-200 font-medium">
                      SNIES: {program.higher_education_institution.snies}
                    </Badge>
                  </div>

                  {/* View Reports Button */}
                  <div className="pt-4 border-t border-gray-100">
                    <Button
                      onClick={() => handleViewReports(program.id, program.name)}
                      className="w-full bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 shadow-lg hover:shadow-xl transition-all duration-200"
                    >
                      <FileText className="h-4 w-4 mr-2" />
                      Ver Reportes
                    </Button>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        ) : (
          <Card>
            <CardContent className="p-12">
              <div className="text-center">
                <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                <h3 className="text-lg font-medium text-gray-900 mb-2">
                  No hay programas registrados
                </h3>
                <p className="text-gray-600 mb-6">
                  No se encontraron programas que coincidan con los filtros aplicados
                </p>
                <Button 
                  onClick={clearFilters} 
                  variant="outline"
                  className="text-gray-600 border-gray-300 hover:bg-gray-50"
                >
                  Limpiar Filtros
                </Button>
              </div>
            </CardContent>
          </Card>
        )}
        
        {/* Pagination - Always visible when not loading */}
        {!isLoading && (
          <ProgramsPagination
            currentPage={currentPage}
            totalPages={data ? Math.ceil(data.total_count / itemsPerPage) : 1}
            onPageChange={handlePageChange}
          />
        )}
      </div>

      {/* Reports Modal */}
      <Dialog open={isReportsModalOpen} onOpenChange={setIsReportsModalOpen}>
        <DialogContent className="sm:max-w-4xl max-h-[80vh] overflow-hidden flex flex-col">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <FileText className="h-5 w-5 text-blue-600" />
              Reportes de Evaluación - {selectedProgramName}
            </DialogTitle>
            <DialogDescription>
              Lista de reportes de evaluación de ciberseguridad para este programa de grado.
            </DialogDescription>
          </DialogHeader>
          
          <div className="flex-1 overflow-hidden flex flex-col">
            {reportsLoading ? (
              <div className="space-y-3 flex-1 overflow-y-auto">
                {Array.from({ length: 5 }).map((_, index) => (
                  <div key={index} className="flex items-center space-x-4 p-4 border rounded-lg">
                    <Skeleton className="h-12 w-12 rounded-lg" />
                    <div className="flex-1 space-y-2">
                      <Skeleton className="h-4 w-1/4" />
                      <Skeleton className="h-3 w-1/2" />
                    </div>
                    <Skeleton className="h-6 w-16" />
                  </div>
                ))}
              </div>
            ) : reportsError ? (
              <div className="flex-1 flex items-center justify-center">
                <div className="text-center text-gray-500">
                  <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <p>Error al cargar los reportes</p>
                  <p className="text-sm mt-2">Inténtalo de nuevo más tarde</p>
                </div>
              </div>
            ) : reportsData && reportsData.data && reportsData.data.length > 0 ? (
              <>
                <div className="flex-1 overflow-y-auto space-y-3">
                  {reportsData.data.map((report) => (
                    <Card key={report.id} className="hover:shadow-md transition-shadow">
                      <CardContent className="p-4">
                        <div className="flex items-center justify-between">
                          <div className="flex items-center space-x-4">
                            <div className="p-3 bg-blue-100 rounded-lg">
                              <FileText className="h-6 w-6 text-blue-600" />
                            </div>
                            <div>
                              <div className="flex items-center gap-2 mb-1">
                                <h4 className="font-semibold text-gray-900">Reporte #{report.id}</h4>
                                <Badge className={`text-xs font-medium ${getScoreColor(report.score)}`}>
                                  {getScoreLabel(report.score)}
                                </Badge>
                              </div>
                              <div className="flex items-center gap-4 text-sm text-gray-600">
                                <div className="flex items-center gap-1">
                                  <CalendarIcon className="h-4 w-4" />
                                  {formatDate(report.created_at)}
                                </div>
                                <div className="flex items-center gap-1">
                                  <Star className="h-4 w-4" />
                                  Puntuación: {formatScore(report.score)}
                                </div>
                                <div className="flex items-center gap-1">
                                  <span className="text-xs bg-gray-100 px-2 py-1 rounded">
                                    {report.lang.toUpperCase()}
                                  </span>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div className="flex items-center gap-2">
                            <Button 
                              variant="outline" 
                              size="sm"
                              onClick={() => handleViewReport(report.id)}
                            >
                              <Eye className="h-4 w-4 mr-2" />
                              Ver
                            </Button>
                          </div>
                        </div>
                      </CardContent>
                    </Card>
                  ))}
                </div>
                
                {/* Reports Pagination */}
                {reportsData.total_count > reportsItemsPerPage && (
                  <div className="pt-4 border-t">
                    <ProgramsPagination
                      currentPage={reportsCurrentPage}
                      totalPages={Math.ceil(reportsData.total_count / reportsItemsPerPage)}
                      onPageChange={handleReportsPageChange}
                    />
                  </div>
                )}
              </>
            ) : (
              <div className="flex-1 flex items-center justify-center">
                <div className="text-center text-gray-500">
                  <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <p>No hay reportes disponibles</p>
                  <p className="text-sm mt-2">Este programa aún no tiene reportes de evaluación</p>
                </div>
              </div>
            )}
          </div>
          
          <div className="flex justify-end pt-4 border-t">
            <Button
              variant="outline"
              onClick={() => setIsReportsModalOpen(false)}
            >
              Cerrar
            </Button>
          </div>
        </DialogContent>
      </Dialog>

      {/* Report Detail Modal */}
      <Dialog open={isReportDetailModalOpen} onOpenChange={setIsReportDetailModalOpen}>
        <DialogContent className="sm:max-w-6xl max-h-[90vh] overflow-hidden flex flex-col">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <FileText className="h-5 w-5 text-blue-600" />
              Reporte de Evaluación - WebSAC3
            </DialogTitle>
            <DialogDescription>
              Detalle completo del reporte de evaluación de ciberseguridad
            </DialogDescription>
          </DialogHeader>
          
          <div className="flex-1 overflow-y-auto">
            {reportDetailLoading ? (
              <div className="space-y-4">
                <Skeleton className="h-8 w-full" />
                <Skeleton className="h-32 w-full" />
                <Skeleton className="h-64 w-full" />
              </div>
            ) : reportDetailError ? (
              <div className="flex items-center justify-center py-12">
                <div className="text-center text-gray-500">
                  <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <p>Error al cargar el reporte</p>
                  <p className="text-sm mt-2">Inténtalo de nuevo más tarde</p>
                </div>
              </div>
            ) : reportDetail ? (
              <div id="report-content" className="bg-white rounded-lg shadow-sm border p-8 max-w-4xl mx-auto">
                {/* Header with Logo */}
                <div className="text-center mb-8">
                  <ReportImage 
                    src="/websac3/logo-websac3.png" 
                    alt="WebSAC3" 
                    className="h-16 mx-auto mb-6"
                    fallbackText="WebSAC3"
                  />
                </div>

                {/* Title */}
                <div className="text-center mb-8">
                  <h2 className="text-3xl font-bold text-gray-800 mb-3">Reporte de Evaluación</h2>
                  <p className="text-lg text-gray-600">Resultado de la evaluación del componente de ciberseguridad</p>
                </div>

                {/* Main Content with Two Columns */}
                <div className="flex flex-col lg:flex-row gap-8 mb-8">
                  {/* Left Column - Information */}
                  <div className="flex-1">
                    {/* Institution Info */}
                    <div className="bg-gray-50 border border-gray-200 rounded-lg p-6 mb-6">
                      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                        <div className="text-center">
                          <div className="text-xs text-gray-500 font-medium mb-2">SNIES Institución</div>
                          <div className="text-base font-bold text-gray-800">
                            {reportDetail.degree_program?.higher_education_institution?.snies || 'N/A'}
                          </div>
                        </div>
                        <div className="text-center">
                          <div className="text-xs text-gray-500 font-medium mb-2">Institución</div>
                          <div className="text-base font-bold text-gray-800">
                            {reportDetail.degree_program?.higher_education_institution?.name || 'No disponible'}
                          </div>
                        </div>
                        <div className="text-center">
                          <div className="text-xs text-gray-500 font-medium mb-2">Fecha del reporte</div>
                          <div className="text-base font-bold text-gray-800">
                            {new Date(reportDetail.created_at).toLocaleDateString('es-ES')}
                          </div>
                        </div>
                      </div>
                    </div>

                    {/* Program Info */}
                    <div className="bg-gray-50 border border-gray-200 rounded-lg p-6">
                      <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
                        <div className="text-center">
                          <div className="text-xs text-gray-500 font-medium mb-2">SNIES Programa</div>
                          <div className="text-base font-bold text-gray-800">{reportDetail.degree_program?.snies || 'N/A'}</div>
                        </div>
                        <div className="text-center">
                          <div className="text-xs text-gray-500 font-medium mb-2">Programa</div>
                          <div className="text-base font-bold text-gray-800">{reportDetail.degree_program?.name || 'No disponible'}</div>
                        </div>
                        <div className="text-center">
                          <div className="text-xs text-gray-500 font-medium mb-2">Rol profesional</div>
                          <div className="text-base font-bold text-gray-800">{reportDetail.professional_role?.name || 'No disponible'}</div>
                        </div>
                        <div className="text-center">
                          <div className="text-xs text-gray-500 font-medium mb-2">Puntaje total</div>
                          <div className="text-base font-bold text-gray-800">{formatScore(reportDetail.score || 0)}</div>
                        </div>
                      </div>
                    </div>
                  </div>

                  {/* Right Column - Score Chart */}
                  <div className="flex items-center justify-center lg:w-80">
                    <div className="flex flex-col items-center">
                      <div className="w-32 h-32 relative mb-3">
                        <svg className="w-32 h-32 transform -rotate-90" viewBox="0 0 120 120">
                          <circle
                            cx="60"
                            cy="60"
                            r="50"
                            fill="none"
                            stroke="#e5e7eb"
                            strokeWidth="8"
                          />
                          <circle
                            cx="60"
                            cy="60"
                            r="50"
                            fill="none"
                            stroke={getScoreColor(reportDetail.score || 0).includes('green') ? '#10b981' : 
                                   getScoreColor(reportDetail.score || 0).includes('yellow') ? '#f59e0b' : '#ef4444'}
                            strokeWidth="8"
                            strokeLinecap="round"
                            strokeDasharray={`${((reportDetail.score || 0) * 314.16)} 314.16`}
                          />
                          <text
                            x="60"
                            y="65"
                            textAnchor="middle"
                            className="text-2xl font-bold fill-gray-800"
                            style={{ transform: 'rotate(90deg) translate(0px, -120px)' }}
                          >
                            {Math.round((reportDetail.score || 0) * 100)}%
                          </text>
                        </svg>
                      </div>
                      <div className={`px-3 py-1 rounded-full text-sm font-semibold border ${getScoreColor(reportDetail.score || 0)}`}>
                        {getScoreLabel(reportDetail.score || 0)}
                      </div>
                    </div>
                  </div>
                </div>

                <hr className="border-gray-200 my-8" />

                {/* Knowledge Areas */}
                <div className="space-y-8">
                  {reportDetail.knowledge_area_reports?.map((area, index) => (
                    <div key={index} className="bg-white border border-gray-200 rounded-lg p-6">
                      <div className="flex items-start justify-between mb-4">
                        <div className="flex-1">
                          <h3 className="text-xl font-semibold text-purple-700 mb-3 text-center">
                            Área de conocimiento: {area.name}
                          </h3>
                          <div className="text-center mb-4">
                            <p className="text-sm text-gray-600">
                              <strong>Horas esperadas:</strong> {area.total_learn_hours_expected.toFixed(2)} · 
                              <strong>Horas alcanzadas:</strong> {area.total_learn_hours_actual.toFixed(2)} · 
                              <strong>Peso (0–1):</strong> {area.score_expected.toFixed(2)} · 
                              <strong>Ponderado (obtenido):</strong> {area.score_got.toFixed(2)}
                            </p>
                          </div>
                        </div>
                      <div className="ml-4">
                        <div className="w-16 h-16 relative">
                          <svg className="w-16 h-16 transform -rotate-90" viewBox="0 0 60 60">
                            <circle
                              cx="30"
                              cy="30"
                              r="22"
                              fill="none"
                              stroke="#e0e0e0"
                              strokeWidth="4"
                            />
                            <circle
                              cx="30"
                              cy="30"
                              r="22"
                              fill="none"
                              stroke={getScoreColor(area.score_got / area.score_expected).includes('green') ? '#28a745' : 
                                     getScoreColor(area.score_got / area.score_expected).includes('yellow') ? '#ffc107' : '#dc3545'}
                              strokeWidth="4"
                              strokeDasharray={`${(area.score_got / area.score_expected) * 138.23} 138.23`}
                              strokeLinecap="round"
                            />
                            <text
                              x="30"
                              y="33"
                              textAnchor="middle"
                              className="text-xs font-bold fill-gray-800"
                              style={{ transform: 'rotate(90deg) translate(0px, -60px)' }}
                            >
                              {Math.round((area.score_got / area.score_expected) * 100)}%
                            </text>
                          </svg>
                        </div>
                        <div className={`text-center text-xs font-semibold mt-1 ${
                          getScoreColor(area.score_got / area.score_expected).includes('green') ? 'text-green-600' : 
                          getScoreColor(area.score_got / area.score_expected).includes('yellow') ? 'text-yellow-600' : 'text-red-600'
                        }`}>
                          {getScoreLabel(area.score_got / area.score_expected)}
                        </div>
                      </div>
                    </div>

                    {/* Topics Table */}
                      <div className="overflow-x-auto">
                        <table className="w-full border-collapse border border-gray-300 mx-auto">
                          <thead>
                            <tr className="bg-gray-100">
                              <th className="border border-gray-300 px-4 py-3 text-center text-sm font-medium text-gray-700">
                                Temática
                              </th>
                              <th className="border border-gray-300 px-4 py-3 text-center text-sm font-medium text-gray-700">
                                Horas esperadas
                              </th>
                              <th className="border border-gray-300 px-4 py-3 text-center text-sm font-medium text-gray-700">
                                Horas alcanzadas
                              </th>
                            </tr>
                          </thead>
                          <tbody>
                            {area.topic_reports?.map((topic, topicIndex) => (
                              <tr key={topicIndex}>
                                <td className="border border-gray-300 px-4 py-3 text-sm text-gray-800 text-center">
                                  {topic.name}
                                </td>
                                <td className="border border-gray-300 px-4 py-3 text-sm text-gray-800 text-center">
                                  {topic.learn_hours_expected.toFixed(2)}
                                </td>
                                <td className="border border-gray-300 px-4 py-3 text-sm text-gray-800 text-center">
                                  {topic.learn_hours_actual.toFixed(2)}
                                </td>
                              </tr>
                            ))}
                          </tbody>
                        </table>
                      </div>
                  </div>
                ))}

                {/* Unexpected Knowledge Areas */}
                {reportDetail.unexpected_knowledge_area_reports && reportDetail.unexpected_knowledge_area_reports.length > 0 && (
                  <>
                    <hr className="border-green-500 my-8" />
                    <h2 className="text-2xl font-bold text-green-600 text-center mb-3">
                      📚 Áreas de Conocimiento Adicionales
                    </h2>
                    <p className="text-lg text-gray-600 text-center mb-8">
                      Temáticas cubiertas en el programa que no se esperaban en el rol profesional
                    </p>

                    <div className="space-y-6">
                      {reportDetail.unexpected_knowledge_area_reports.map((area, index) => (
                        <div key={index} className="border border-green-200 rounded-lg bg-green-50 p-6">
                          <h3 className="text-xl font-semibold text-green-700 mb-3 text-center">
                            🌱 {area.name}
                          </h3>
                          <p className="text-sm text-gray-600 mb-4 text-center">
                            <strong>Total de horas cubiertas:</strong> {area.total_learn_hours.toFixed(2)}
                          </p>

                          <div className="overflow-x-auto">
                            <table className="w-full border-collapse border border-green-300 mx-auto">
                              <thead>
                                <tr className="bg-green-100">
                                  <th className="border border-green-300 px-4 py-3 text-center text-sm font-medium text-green-800">
                                    Temática
                                  </th>
                                  <th className="border border-green-300 px-4 py-3 text-center text-sm font-medium text-green-800">
                                    Horas cubiertas
                                  </th>
                                </tr>
                              </thead>
                              <tbody>
                                {area.topic_reports?.map((topic, topicIndex) => (
                                  <tr key={topicIndex}>
                                    <td className="border border-green-300 px-4 py-3 text-sm text-gray-800 text-center">
                                      {topic.name}
                                    </td>
                                    <td className="border border-green-300 px-4 py-3 text-sm text-gray-800 text-center">
                                      {topic.learn_hours_actual.toFixed(2)}
                                    </td>
                                  </tr>
                                ))}
                              </tbody>
                            </table>
                          </div>
                        </div>
                      ))}
                    </div>
                  </>
                )}

                {/* Footer */}
                <hr className="border-gray-200 my-8" />
                <div className="text-center">
                  <div className="flex justify-center items-center gap-6 mb-4">
                    <ReportImage 
                      src="/unillanos/logo-unillanos.png" 
                      alt="Universidad de los Llanos" 
                      className="h-10"
                      fallbackText="Unillanos"
                    />
                    <ReportImage 
                      src="/unillanos/logo-fcbi.png" 
                      alt="Facultad de Ciencias Básicas e Ingeniería" 
                      className="h-10"
                      fallbackText="FCBI"
                    />
                  </div>
                  <p className="text-sm text-gray-600 font-medium">
                    Universidad de los Llanos • Facultad de Ciencias Básicas e Ingeniería
                  </p>
                  <p className="text-xs text-gray-500 mt-2">
                    Sistema WebSAC3 – Web System for Analysis of Curricula Cybersecurity Component
                  </p>
                </div>
              </div>
            ) : null}
          </div>
          
          <div className="flex justify-between pt-4 border-t">
            <Button
              variant="outline"
              onClick={downloadReportAsPDF}
              disabled={!reportDetail}
              className="flex items-center gap-2"
            >
              <Download className="h-4 w-4" />
              Descargar PDF
            </Button>
            <Button
              variant="outline"
              onClick={() => setIsReportDetailModalOpen(false)}
            >
              Cerrar
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}

// Loading skeleton component
function LoadingSkeleton() {
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <Skeleton className="h-8 w-64 mb-2" />
          <Skeleton className="h-4 w-80" />
        </div>
        <Skeleton className="h-10 w-32" />
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {Array.from({ length: 6 }).map((_, i) => (
          <Card key={i} className="bg-white border border-gray-200 shadow-lg py-0">
            <CardHeader className="pb-3 pt-4 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-gray-100">
              <Skeleton className="h-6 w-3/4 mb-2" />
              <Skeleton className="h-4 w-1/2" />
            </CardHeader>
            <CardContent className="p-5 space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <Skeleton className="h-16 w-full" />
                <Skeleton className="h-16 w-full" />
              </div>
              <Skeleton className="h-20 w-full" />
              <Skeleton className="h-20 w-full" />
              <Skeleton className="h-10 w-full" />
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
