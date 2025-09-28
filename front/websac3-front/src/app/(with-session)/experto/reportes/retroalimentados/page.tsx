"use client";

import React, { useState, useCallback, useMemo, memo } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { Input } from "@/components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import { CardDescription } from "@/components/ui/card";
import { 
  MessageSquare, 
  User, 
  Calendar, 
  Target, 
  Lightbulb, 
  AlertTriangle,
  Star,
  CheckCircle,
  FileText,
  GraduationCap,
  Building,
  Search,
  Filter,
  Eye,
  ChevronDown,
  ChevronUp,
  Settings,
  BookOpen,
  Clock
} from "lucide-react";
import { useListReportFeedbacksQuery } from "@/services/api";
import { useRouter, useSearchParams } from "next/navigation";

interface ReportFeedback {
  id: number;
  report_id: number;
  general_comments: string;
  recommendations: string;
  created_at: string;
  updated_at: string;
  auditor: {
    id: number;
    email: string;
    name: string;
  };
  knowledge_area_feedbacks: Array<{
    id: number;
    knowledge_area_name: string;
    knowledge_area_report_id: number;
    comments: string;
  }>;
  report: {
    id: number;
    created_at: string;
    score: number;
    degree_program: {
      id: number;
      name: string;
      snies: number;
    };
    higher_education_institution: {
      id: number;
      name: string;
    };
  };
}

// Componente memoizado para los inputs de filtro
const FilterInput = memo(({ 
  id, 
  label, 
  placeholder, 
  value, 
  onChange, 
  icon: Icon, 
  className 
}: {
  id: string;
  label: string;
  placeholder: string;
  value: string;
  onChange: (value: string) => void;
  icon: React.ComponentType<{ className?: string }>;
  className?: string;
}) => (
  <div className="space-y-2">
    <Label htmlFor={id} className="text-sm font-medium text-gray-700 flex items-center gap-2">
      <Icon className="h-4 w-4 text-blue-500" />
      {label}
    </Label>
    <Input
      id={id}
      placeholder={placeholder}
      value={value}
      onChange={(e) => onChange(e.target.value)}
      className={`h-10 border-gray-300 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 ${className || ''}`}
    />
  </div>
));

export default function ReportFeedbacksPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(10);
  const [filters, setFilters] = useState({
    auditorName: "",
    auditorLastName: "",
    programName: "",
    creatorInstitutionName: ""
  });
  
  const [showFilters, setShowFilters] = useState(false);

  // Initialize filters from URL on component mount
  React.useEffect(() => {
    const urlFilters = {
      auditorName: searchParams.get('auditorName') || "",
      auditorLastName: searchParams.get('auditorLastName') || "",
      programName: searchParams.get('programName') || "",
      creatorInstitutionName: searchParams.get('creatorInstitutionName') || ""
    };
    
    const urlPage = parseInt(searchParams.get('page') || '1');
    const urlItemsPerPage = parseInt(searchParams.get('itemsPerPage') || '10');
    
    setFilters(urlFilters);
    setCurrentPage(urlPage);
    setItemsPerPage(urlItemsPerPage);
  }, [searchParams]);

  // Function to update URL with current filters and pagination
  const updateURL = useCallback((newFilters: typeof filters, newPage: number, newItemsPerPage: number) => {
    const params = new URLSearchParams();
    
    // Add pagination params
    params.set('page', newPage.toString());
    params.set('itemsPerPage', newItemsPerPage.toString());
    
    // Add filter params only if they have values
    if (newFilters.auditorName.trim()) {
      params.set('auditorName', newFilters.auditorName.trim());
    }
    if (newFilters.auditorLastName.trim()) {
      params.set('auditorLastName', newFilters.auditorLastName.trim());
    }
    if (newFilters.programName.trim()) {
      params.set('programName', newFilters.programName.trim());
    }
    if (newFilters.creatorInstitutionName.trim()) {
      params.set('creatorInstitutionName', newFilters.creatorInstitutionName.trim());
    }
    
    // Update URL without causing a page reload
    const newURL = `${window.location.pathname}?${params.toString()}`;
    router.replace(newURL, { scroll: false });
  }, [router]);

  // Build query parameters with proper filter format - memoized to prevent unnecessary re-renders
  const buildFilters = useMemo(() => {
    const filterParams: Record<string, string> = {};
    
    if (filters.auditorName.trim()) {
      filterParams['Auditor.Person.name[cont]'] = filters.auditorName.trim();
    }
    
    if (filters.auditorLastName.trim()) {
      filterParams['Auditor.Person.lastname[cont]'] = filters.auditorLastName.trim();
    }
    
    if (filters.programName.trim()) {
      filterParams['Report.DegreeProgram.name[cont]'] = filters.programName.trim();
    }
    
    if (filters.creatorInstitutionName.trim()) {
      filterParams['Report.DegreeProgram.UserCreator.Person.HigherEducationInstitution.name[cont]'] = filters.creatorInstitutionName.trim();
    }
    
    return filterParams;
  }, [filters]);

  const queryParams = useMemo(() => ({
    current_page: currentPage,
    items_per_page: itemsPerPage,
    filters: buildFilters
  }), [currentPage, itemsPerPage, buildFilters]);

  const {
    data: feedbacksData,
    isLoading: loading,
    error: queryError,
    refetch
  } = useListReportFeedbacksQuery(queryParams);

  const handlePageChange = useCallback((page: number) => {
    setCurrentPage(page);
    updateURL(filters, page, itemsPerPage);
  }, [filters, itemsPerPage, updateURL]);

  const handleItemsPerPageChange = useCallback((perPage: number) => {
    setItemsPerPage(perPage);
    setCurrentPage(1); // Reset to first page
    updateURL(filters, 1, perPage);
  }, [filters, updateURL]);

  // Handle filter changes - memoized to prevent unnecessary re-renders
  const handleFilterChange = useCallback((key: string, value: string) => {
    const newFilters = { ...filters, [key]: value };
    setFilters(newFilters);
    setCurrentPage(1); // Reset to first page
    updateURL(newFilters, 1, itemsPerPage);
  }, [filters, itemsPerPage, updateURL]);

  // Specific handlers for each filter to prevent unnecessary re-renders
  const handleAuditorNameChange = useCallback((value: string) => {
    handleFilterChange('auditorName', value);
  }, [handleFilterChange]);

  const handleAuditorLastNameChange = useCallback((value: string) => {
    handleFilterChange('auditorLastName', value);
  }, [handleFilterChange]);

  const handleProgramNameChange = useCallback((value: string) => {
    handleFilterChange('programName', value);
  }, [handleFilterChange]);

  const handleCreatorInstitutionNameChange = useCallback((value: string) => {
    handleFilterChange('creatorInstitutionName', value);
  }, [handleFilterChange]);

  // Clear all filters - memoized to prevent unnecessary re-renders
  const clearFilters = useCallback(() => {
    const clearedFilters = {
      auditorName: "",
      auditorLastName: "",
      programName: "",
      creatorInstitutionName: ""
    };
    setFilters(clearedFilters);
    setCurrentPage(1);
    updateURL(clearedFilters, 1, itemsPerPage);
  }, [itemsPerPage, updateURL]);

  const toggleFilters = useCallback(() => {
    setShowFilters(prev => !prev);
  }, []);


  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("es-ES", {
      year: "numeric",
      month: "long",
      day: "numeric",
    });
  };

  const getScoreColor = (score: number) => {
    if (score >= 80) return "bg-green-100 text-green-800";
    if (score >= 60) return "bg-yellow-100 text-yellow-800";
    return "bg-red-100 text-red-800";
  };

  const formatScore = (score: number) => {
    const percentage = Math.floor(score * 10000) / 100;
    return percentage % 1 === 0 ? percentage.toFixed(0) : percentage.toFixed(2);
  };

  // Extract data from the query result
  const feedbacks = feedbacksData?.data || [];
  const totalCount = feedbacksData?.total_count || 0;
  const currentPageFromAPI = feedbacksData?.current_page || 1;
  const itemsPerPageFromAPI = feedbacksData?.items_per_page || 10;
  
  // Calculate total pages based on total count and items per page
  const totalPages = Math.ceil(totalCount / itemsPerPageFromAPI);
  
  // Check if it's a 404 error (no feedbacks found) vs a real error
  const is404Error = queryError && 'status' in queryError && queryError.status === 404;
  const error = queryError && !is404Error ? "Error al cargar las retroalimentaciones" : null;

  const Pagination = () => {
    const pages = [];
    const maxVisiblePages = 5;
    let startPage = Math.max(1, currentPage - Math.floor(maxVisiblePages / 2));
    let endPage = Math.min(totalPages, startPage + maxVisiblePages - 1);

    if (endPage - startPage + 1 < maxVisiblePages) {
      startPage = Math.max(1, endPage - maxVisiblePages + 1);
    }

    for (let i = startPage; i <= endPage; i++) {
      pages.push(
        <Button
          key={i}
          variant={i === currentPage ? "default" : "outline"}
          size="sm"
          onClick={() => handlePageChange(i)}
          className="mx-1"
        >
          {i}
        </Button>
      );
    }

    return (
      <div className="flex items-center justify-center space-x-2 mt-6">
        <Button
          variant="outline"
          size="sm"
          onClick={() => handlePageChange(currentPage - 1)}
          disabled={currentPage === 1}
        >
          Anterior
        </Button>
        {pages}
        <Button
          variant="outline"
          size="sm"
          onClick={() => handlePageChange(currentPage + 1)}
          disabled={currentPage === totalPages}
        >
          Siguiente
        </Button>
        <div className="ml-4 text-sm text-gray-600">
          Página {currentPage} de {totalPages} ({totalCount} retroalimentaciones)
        </div>
      </div>
    );
  };

  const FeedbackCard = ({ feedback }: { feedback: ReportFeedback }) => (
    <Card className="hover:shadow-md transition-shadow">
      <CardHeader className="pb-3">
        <div className="flex items-start justify-between">
          <div className="flex items-center space-x-2">
            <CheckCircle className="h-5 w-5 text-green-500" />
            <CardTitle className="text-lg">Retroalimentación #{feedback.id}</CardTitle>
          </div>
          <Badge className={getScoreColor(feedback.report.score * 100)}>
            <Star className="h-3 w-3 mr-1" />
            {formatScore(feedback.report.score)}%
          </Badge>
        </div>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="flex items-center space-x-2">
            <GraduationCap className="h-4 w-4 text-gray-500" />
            <div>
              <p className="text-sm font-medium">{feedback.report.degree_program.name}</p>
              <p className="text-xs text-gray-500">SNIES: {feedback.report.degree_program.snies}</p>
            </div>
          </div>
          
          <div className="flex items-center space-x-2">
            <Building className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{feedback.report.higher_education_institution.name}</p>
          </div>
          
          <div className="flex items-center space-x-2">
            <User className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{feedback.auditor.name}</p>
          </div>
          
          <div className="flex items-center space-x-2">
            <Calendar className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{formatDate(feedback.created_at)}</p>
          </div>
        </div>

        <div className="flex items-center space-x-2">
          <FileText className="h-4 w-4 text-gray-500" />
          <p className="text-sm">Reporte #{feedback.report.id}</p>
        </div>

        {/* Feedback Preview */}
        <div className="space-y-2">
          {feedback.general_comments && (
            <div className="bg-gray-50 rounded-lg p-3">
              <h4 className="text-xs font-medium text-gray-700 mb-1 flex items-center gap-1">
                <MessageSquare className="h-3 w-3" />
                Comentarios Generales
              </h4>
              <p className="text-xs text-gray-600 line-clamp-2">
                {feedback.general_comments}
              </p>
            </div>
          )}
          
          {feedback.recommendations && (
            <div className="bg-gray-50 rounded-lg p-3">
              <h4 className="text-xs font-medium text-gray-700 mb-1 flex items-center gap-1">
                <Lightbulb className="h-3 w-3" />
                Recomendaciones
              </h4>
              <p className="text-xs text-gray-600 line-clamp-2">
                {feedback.recommendations}
              </p>
            </div>
          )}
        </div>
        
        <div className="pt-3 border-t">
          <Button 
            className="w-full" 
            variant="outline"
            onClick={() => router.push(`/experto/reportes/${feedback.report_id}`)}
          >
            <Eye className="h-4 w-4 mr-2" />
            Ver Detalles
          </Button>
        </div>
      </CardContent>
    </Card>
  );

  const LoadingSkeleton = () => (
    <div className="space-y-4">
      {Array.from({ length: 3 }).map((_, index) => (
        <Card key={index}>
          <CardHeader>
            <div className="flex items-center justify-between">
              <Skeleton className="h-6 w-32" />
              <Skeleton className="h-6 w-16" />
            </div>
          </CardHeader>
          <CardContent className="space-y-3">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <Skeleton className="h-4 w-full" />
              <Skeleton className="h-4 w-full" />
              <Skeleton className="h-4 w-full" />
              <Skeleton className="h-4 w-full" />
            </div>
            <Skeleton className="h-20 w-full" />
            <Skeleton className="h-10 w-full" />
          </CardContent>
        </Card>
      ))}
    </div>
  );

  const FiltersSection = useMemo(() => (
    <div className="space-y-6">
      {/* Header with Toggle Button */}
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Reportes Retroalimentados</h1>
          <p className="text-gray-600 mt-2">Revisa las retroalimentaciones de reportes de evaluación</p>
        </div>
        <div className="flex gap-3">
          <Button 
            onClick={toggleFilters}
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
      </div>

      {/* Compact Stats Bar */}
      {feedbacksData && !(queryError && queryError.status === 404) && (
        <div className="bg-white border border-gray-200 rounded-lg shadow-sm p-4">
          <div className="flex flex-wrap items-center justify-between gap-4">
            {/* Left side - Main info */}
            <div className="flex items-center gap-6">
              <div className="flex items-center gap-2">
                <MessageSquare className="h-4 w-4 text-blue-600" />
                <span className="text-sm font-medium text-gray-700">
                  {feedbacksData.total_count} retroalimentacion{feedbacksData.total_count !== 1 ? 'es' : ''}
                </span>
              </div>
              
              <div className="flex items-center gap-2">
                <Calendar className="h-4 w-4 text-gray-500" />
                <span className="text-sm text-gray-600">
                  Página {feedbacksData.current_page} de {Math.ceil(feedbacksData.total_count / itemsPerPage)}
                </span>
              </div>
            </div>

            {/* Right side - Active filters info */}
            {(filters.auditorName || filters.auditorLastName || filters.programName || filters.creatorInstitutionName) && (
              <div className="flex items-center gap-3">
                <div className="flex items-center gap-2">
                  <Filter className="h-4 w-4 text-indigo-600" />
                  <span className="text-sm font-medium text-gray-700">
                    Filtros activos
                  </span>
                </div>
                <div className="flex gap-2 flex-wrap">
                  {filters.auditorName && (
                    <Badge className="bg-blue-100 text-blue-800 border-blue-200 text-xs">
                      Auditor: {filters.auditorName}
                    </Badge>
                  )}
                  {filters.auditorLastName && (
                    <Badge className="bg-blue-100 text-blue-800 border-blue-200 text-xs">
                      Apellido: {filters.auditorLastName}
                    </Badge>
                  )}
                  {filters.programName && (
                    <Badge className="bg-green-100 text-green-800 border-green-200 text-xs">
                      Programa: {filters.programName}
                    </Badge>
                  )}
                  {filters.creatorInstitutionName && (
                    <Badge className="bg-orange-100 text-orange-800 border-orange-200 text-xs">
                      Inst. Creador: {filters.creatorInstitutionName}
                    </Badge>
                  )}
                </div>
              </div>
            )}
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
              Personaliza tu búsqueda y visualización de retroalimentaciones
            </CardDescription>
          </CardHeader>
          <CardContent className="p-6">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {/* Auditor Name Filter */}
              <FilterInput
                id="auditor-filter"
                label="Nombre del Auditor"
                placeholder="Buscar por nombre..."
                value={filters.auditorName}
                onChange={handleAuditorNameChange}
                icon={User}
              />

              {/* Auditor Last Name Filter */}
              <FilterInput
                id="auditor-lastname-filter"
                label="Apellido del Auditor"
                placeholder="Buscar por apellido..."
                value={filters.auditorLastName}
                onChange={handleAuditorLastNameChange}
                icon={User}
              />
              
              {/* Program Name Filter */}
              <FilterInput
                id="program-filter"
                label="Programa Académico"
                placeholder="Buscar por programa..."
                value={filters.programName}
                onChange={handleProgramNameChange}
                icon={BookOpen}
                className="focus:border-green-500 focus:ring-green-500"
              />

              {/* Creator Institution Filter */}
              <FilterInput
                id="creator-institution-filter"
                label="Institución del Creador"
                placeholder="Buscar por institución del creador..."
                value={filters.creatorInstitutionName}
                onChange={handleCreatorInstitutionNameChange}
                icon={GraduationCap}
                className="focus:border-orange-500 focus:ring-orange-500"
              />
              
              {/* Items Per Page */}
              <div className="space-y-2">
                <Label htmlFor="per-page" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                  <Clock className="h-4 w-4 text-indigo-500" />
                  Elementos por página
                </Label>
                <select
                  id="per-page"
                  value={itemsPerPage}
                  onChange={(e) => handleItemsPerPageChange(parseInt(e.target.value))}
                  className="w-full h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-indigo-500 focus:border-indigo-500 bg-white"
                >
                  <option value={5}>5 por página</option>
                  <option value={10}>10 por página</option>
                  <option value={15}>15 por página</option>
                  <option value={20}>20 por página</option>
                  <option value={25}>25 por página</option>
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
    </div>
  ), [showFilters, filters, feedbacksData, queryError, itemsPerPage, toggleFilters, clearFilters, handleItemsPerPageChange, handleAuditorNameChange, handleAuditorLastNameChange, handleProgramNameChange, handleCreatorInstitutionNameChange]);

  if (error) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6">
          <div className="text-center">
            <AlertTriangle className="h-12 w-12 text-red-500 mx-auto mb-4" />
            <h2 className="text-xl font-semibold text-red-600 mb-2">Error al cargar retroalimentaciones</h2>
            <p className="text-gray-600 mb-4">{error}</p>
            <Button onClick={() => refetch()}>
              Reintentar
            </Button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {FiltersSection}

      {/* Content */}
      <div className="bg-white rounded-xl shadow-sm p-6">
        {loading ? (
          <LoadingSkeleton />
        ) : (feedbacks.length === 0 || is404Error) ? (
          <div className="text-center py-12">
            <div className="bg-blue-50 rounded-full w-20 h-20 flex items-center justify-center mx-auto mb-6">
              <MessageSquare className="h-10 w-10 text-blue-500" />
            </div>
            <h3 className="text-xl font-semibold text-gray-800 mb-3">
              No hay retroalimentaciones
            </h3>
            <p className="text-gray-600 mb-2">
              No se encontraron retroalimentaciones con los filtros aplicados.
            </p>
            <p className="text-sm text-gray-500">
              Intenta ajustar los filtros o verifica si hay retroalimentaciones disponibles.
            </p>
          </div>
        ) : (
          <>
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              {feedbacks.map((feedback) => (
                <FeedbackCard key={feedback.id} feedback={feedback} />
              ))}
            </div>
            
            {totalCount > 0 && <Pagination />}
          </>
        )}
      </div>
    </div>
  );
}

