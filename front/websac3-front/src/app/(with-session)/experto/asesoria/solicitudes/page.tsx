"use client";

import React, { useState, useCallback, useMemo, memo } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { 
  MessageSquare, 
  BookOpen, 
  Clock, 
  GraduationCap, 
  Search, 
  Settings, 
  Calendar,
  User,
  FileText,
  CheckCircle,
  AlertTriangle,
  Eye,
  Filter,
  Mail
} from "lucide-react";
import { useListPendingExpertConsultationsQuery } from "@/services/api";

interface PendingConsultation {
  id: number;
  requester_id: number;
  requester_name: string;
  requester_email: string;
  requester_institution_snies: number;
  requester_institution_name: string;
  requester_institution_ownership: string;
  requester_job_position: string;
  degree_program_id: number;
  degree_program_name: string;
  degree_program_snies: number;
  report_id: number;
  report_score: number;
  request_message: string;
  status_id: number;
  status_name: string;
  created_at: string;
  updated_at: string;
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

export default function PendingConsultationsPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(10);
  const [filters, setFilters] = useState({
    programName: "",
    requesterName: "",
    requesterEmail: "",
    institutionName: ""
  });
  
  const [showFilters, setShowFilters] = useState(false);

  // Initialize filters from URL on component mount
  React.useEffect(() => {
    const urlFilters = {
      programName: searchParams.get('programName') || "",
      requesterName: searchParams.get('requesterName') || "",
      requesterEmail: searchParams.get('requesterEmail') || "",
      institutionName: searchParams.get('institutionName') || ""
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
    if (newFilters.programName.trim()) {
      params.set('programName', newFilters.programName.trim());
    }
    if (newFilters.requesterName.trim()) {
      params.set('requesterName', newFilters.requesterName.trim());
    }
    if (newFilters.requesterEmail.trim()) {
      params.set('requesterEmail', newFilters.requesterEmail.trim());
    }
    if (newFilters.institutionName.trim()) {
      params.set('institutionName', newFilters.institutionName.trim());
    }
    
    // Update URL without causing a page reload
    const newURL = `${window.location.pathname}?${params.toString()}`;
    router.replace(newURL, { scroll: false });
  }, [router]);

  const queryParams = useMemo(() => ({
    current_page: currentPage,
    items_per_page: itemsPerPage,
    lang: 'es'
  }), [currentPage, itemsPerPage]);

  const { data: consultationsData, isLoading: loading, error: queryError, refetch } = useListPendingExpertConsultationsQuery(queryParams);

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
  const handleProgramNameChange = useCallback((value: string) => {
    handleFilterChange('programName', value);
  }, [handleFilterChange]);

  const handleRequesterNameChange = useCallback((value: string) => {
    handleFilterChange('requesterName', value);
  }, [handleFilterChange]);

  const handleRequesterEmailChange = useCallback((value: string) => {
    handleFilterChange('requesterEmail', value);
  }, [handleFilterChange]);

  const handleInstitutionNameChange = useCallback((value: string) => {
    handleFilterChange('institutionName', value);
  }, [handleFilterChange]);

  // Clear all filters - memoized to prevent unnecessary re-renders
  const clearFilters = useCallback(() => {
    const clearedFilters = {
      programName: "",
      requesterName: "",
      requesterEmail: "",
      institutionName: ""
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
  const consultations = consultationsData?.data || [];
  const totalCount = consultationsData?.total_count || 0;
  const currentPageFromAPI = consultationsData?.current_page || 1;
  const itemsPerPageFromAPI = consultationsData?.items_per_page || 10;
  
  // Calculate total pages based on total count and items per page
  const totalPages = Math.ceil(totalCount / itemsPerPageFromAPI);
  
  // Check if it's a 404 error (no consultations found) vs a real error
  const is404Error = queryError && 'status' in queryError && queryError.status === 404;
  const error = queryError && !is404Error ? "Error al cargar las solicitudes" : null;

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
          Página {currentPage} de {totalPages} ({totalCount} solicitudes)
        </div>
      </div>
    );
  };

  const ConsultationCard = ({ consultation }: { consultation: PendingConsultation }) => (
    <Card className="hover:shadow-md transition-shadow">
      <CardHeader className="pb-3">
        <div className="flex items-start justify-between">
          <div className="flex items-center space-x-2">
            <AlertTriangle className="h-5 w-5 text-orange-500" />
            <CardTitle className="text-lg">Solicitud #{consultation.id}</CardTitle>
          </div>
          <Badge className="bg-orange-100 text-orange-800">
            {consultation.status_name}
          </Badge>
        </div>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="flex items-center space-x-2">
            <GraduationCap className="h-4 w-4 text-gray-500" />
            <div>
              <p className="text-sm font-medium">{consultation.degree_program_name}</p>
              <p className="text-xs text-gray-500">SNIES: {consultation.degree_program_snies}</p>
            </div>
          </div>
          
          <div className="flex items-center space-x-2">
            <User className="h-4 w-4 text-gray-500" />
            <div>
              <p className="text-sm font-medium">{consultation.requester_name}</p>
              <p className="text-xs text-gray-500">{consultation.requester_email}</p>
            </div>
          </div>
          
          <div className="flex items-center space-x-2">
            <FileText className="h-4 w-4 text-gray-500" />
            <div>
              <p className="text-sm">Reporte #{consultation.report_id}</p>
              <p className="text-xs text-gray-500">Puntaje: {formatScore(consultation.report_score)}%</p>
            </div>
          </div>
          
          <div className="flex items-center space-x-2">
            <Calendar className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{formatDate(consultation.created_at)}</p>
          </div>
        </div>

        {/* Institution Information */}
        <div className="bg-blue-50 rounded-lg p-3">
          <h4 className="text-xs font-medium text-gray-700 mb-2 flex items-center gap-1">
            <GraduationCap className="h-3 w-3" />
            Información Institucional
          </h4>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-2 text-xs">
            <div>
              <p className="font-medium text-gray-800">{consultation.requester_institution_name}</p>
              <p className="text-gray-600">SNIES: {consultation.requester_institution_snies}</p>
            </div>
            <div>
              <p className="text-gray-600">Tipo: {consultation.requester_institution_ownership}</p>
              <p className="text-gray-600">Cargo: {consultation.requester_job_position}</p>
            </div>
          </div>
        </div>

        {/* Request Message Preview */}
        <div className="bg-gray-50 rounded-lg p-3">
          <h4 className="text-xs font-medium text-gray-700 mb-1 flex items-center gap-1">
            <MessageSquare className="h-3 w-3" />
            Solicitud
          </h4>
          <p className="text-xs text-gray-600 line-clamp-3">
            {consultation.request_message}
          </p>
        </div>
        
        <div className="pt-3 border-t">
          <Button 
            className="w-full" 
            variant="outline"
            onClick={() => router.push(`/experto/asesoria/solicitudes/${consultation.id}`)}
          >
            <Eye className="h-4 w-4 mr-2" />
            Responder Solicitud
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
          <h1 className="text-3xl font-bold text-gray-900">Solicitudes de Asesoría</h1>
          <p className="text-gray-600 mt-2">Revisa las solicitudes de asesoría pendientes de respuesta</p>
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
      {consultationsData && !(queryError && 'status' in queryError && queryError.status === 404) && (
        <div className="bg-white border border-gray-200 rounded-lg shadow-sm p-4">
          <div className="flex flex-wrap items-center justify-between gap-4">
            {/* Left side - Main info */}
            <div className="flex items-center gap-6">
              <div className="flex items-center gap-2">
                <MessageSquare className="h-4 w-4 text-orange-600" />
                <span className="text-sm font-medium text-gray-700">
                  {consultationsData.total_count} solicitud{consultationsData.total_count !== 1 ? 'es' : ''} pendiente{consultationsData.total_count !== 1 ? 's' : ''}
                </span>
              </div>
              
              <div className="flex items-center gap-2">
                <Calendar className="h-4 w-4 text-gray-500" />
                <span className="text-sm text-gray-600">
                  Página {consultationsData.current_page} de {Math.ceil(consultationsData.total_count / itemsPerPage)}
                </span>
              </div>
            </div>

            {/* Right side - Active filters info */}
            {(filters.programName || filters.requesterName || filters.requesterEmail || filters.institutionName) && (
              <div className="flex items-center gap-3">
                <div className="flex items-center gap-2">
                  <Filter className="h-4 w-4 text-indigo-600" />
                  <span className="text-sm font-medium text-gray-700">
                    Filtros activos
                  </span>
                </div>
                <div className="flex gap-2 flex-wrap">
                  {filters.programName && (
                    <Badge className="bg-blue-100 text-blue-800 border-blue-200 text-xs">
                      Programa: {filters.programName}
                    </Badge>
                  )}
                  {filters.requesterName && (
                    <Badge className="bg-green-100 text-green-800 border-green-200 text-xs">
                      Solicitante: {filters.requesterName}
                    </Badge>
                  )}
                  {filters.requesterEmail && (
                    <Badge className="bg-purple-100 text-purple-800 border-purple-200 text-xs">
                      Email: {filters.requesterEmail}
                    </Badge>
                  )}
                  {filters.institutionName && (
                    <Badge className="bg-orange-100 text-orange-800 border-orange-200 text-xs">
                      Institución: {filters.institutionName}
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
              <div className="p-2 bg-orange-100 rounded-lg">
                <Search className="h-5 w-5 text-orange-600" />
              </div>
              Filtros y Configuración
            </CardTitle>
            <CardDescription className="text-gray-600 mt-1">
              Personaliza tu búsqueda y visualización de solicitudes
            </CardDescription>
          </CardHeader>
          <CardContent className="p-6">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {/* Program Name Filter */}
              <FilterInput
                id="program-filter"
                label="Programa Académico"
                placeholder="Buscar por programa..."
                value={filters.programName}
                onChange={handleProgramNameChange}
                icon={GraduationCap}
                className="focus:border-blue-500 focus:ring-blue-500"
              />

              {/* Requester Name Filter */}
              <FilterInput
                id="requester-filter"
                label="Solicitante"
                placeholder="Buscar por nombre..."
                value={filters.requesterName}
                onChange={handleRequesterNameChange}
                icon={User}
                className="focus:border-green-500 focus:ring-green-500"
              />

              {/* Requester Email Filter */}
              <FilterInput
                id="email-filter"
                label="Email del Solicitante"
                placeholder="Buscar por email..."
                value={filters.requesterEmail}
                onChange={handleRequesterEmailChange}
                icon={Mail}
                className="focus:border-purple-500 focus:ring-purple-500"
              />

              {/* Institution Name Filter */}
              <FilterInput
                id="institution-filter"
                label="Institución Educativa"
                placeholder="Buscar por institución..."
                value={filters.institutionName}
                onChange={handleInstitutionNameChange}
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
  ), [showFilters, filters, consultationsData, queryError, itemsPerPage, toggleFilters, clearFilters, handleItemsPerPageChange, handleProgramNameChange, handleRequesterNameChange, handleRequesterEmailChange, handleInstitutionNameChange]);

  if (error) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6">
          <div className="text-center">
            <AlertTriangle className="h-12 w-12 text-red-500 mx-auto mb-4" />
            <h2 className="text-xl font-semibold text-red-600 mb-2">Error al cargar solicitudes</h2>
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
        ) : (consultations.length === 0 || is404Error) ? (
          <div className="text-center py-12">
            <div className="bg-green-50 rounded-full w-20 h-20 flex items-center justify-center mx-auto mb-6">
              <CheckCircle className="h-10 w-10 text-green-500" />
            </div>
            <h3 className="text-xl font-semibold text-gray-800 mb-3">
              ¡Excelente trabajo!
            </h3>
            <p className="text-gray-600 mb-2">
              No hay solicitudes de asesoría pendientes en este momento.
            </p>
            <p className="text-sm text-gray-500">
              Todas las solicitudes han sido respondidas o no hay nuevas solicitudes disponibles.
            </p>
          </div>
        ) : (
          <>
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              {consultations.map((consultation) => (
                <ConsultationCard key={consultation.id} consultation={consultation} />
              ))}
            </div>
            
            {totalCount > 0 && <Pagination />}
          </>
        )}
      </div>
    </div>
  );
}
