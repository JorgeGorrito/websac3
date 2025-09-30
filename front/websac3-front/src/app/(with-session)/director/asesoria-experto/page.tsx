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
  Plus
} from "lucide-react";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Textarea } from "@/components/ui/textarea";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { useListDegreeProgramsQuery, useListDegreeProgramReportsQuery } from "@/services/api";
import { useDispatch } from "react-redux";
import { showError } from "@/store/errorSlice";
import { useListUserExpertConsultationsQuery, useCreateExpertConsultationMutation, useListExpertConsultationStatusesQuery } from "@/services/api";

interface ExpertConsultation {
  id: number;
  degree_program_id: number;
  degree_program_name: string;
  degree_program_snies: number;
  report_id: number;
  report_score: number;
  requester_id: number;
  requester_name: string;
  requester_email: string;
  expert_id: number;
  expert_name: string;
  expert_email: string;
  request_message: string;
  expert_response: string;
  status_id: number;
  status_name: string;
  created_at: string;
  updated_at: string;
  answered_at: string;
  closed_at: string;
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

export default function AsesoriaExpertoPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const dispatch = useDispatch();
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(10);
  const [filters, setFilters] = useState({
    programName: "",
    status: ""
  });
  
  const [showFilters, setShowFilters] = useState(false);
  
  // Estados para el modal de creación
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
  const [selectedProgramId, setSelectedProgramId] = useState<number | null>(null);
  const [selectedReportId, setSelectedReportId] = useState<number | null>(null);
  const [consultationMessage, setConsultationMessage] = useState('');
  const [reportsCurrentPage, setReportsCurrentPage] = useState(1);
  const [reportsItemsPerPage, setReportsItemsPerPage] = useState(10);

  // Initialize filters from URL on component mount
  React.useEffect(() => {
    const urlFilters = {
      programName: searchParams.get('programName') || "",
      status: searchParams.get('status') || ""
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
    if (newFilters.status.trim()) {
      params.set('status', newFilters.status.trim());
    }
    
    // Update URL without causing a page reload
    const newURL = `${window.location.pathname}?${params.toString()}`;
    router.replace(newURL, { scroll: false });
  }, [router]);

  // Build query parameters with proper filter format - memoized to prevent unnecessary re-renders
  const buildFilters = useMemo(() => {
    const filterParams: Record<string, string> = {};
    
    if (filters.programName.trim()) {
      filterParams['DegreeProgram.name[cont]'] = filters.programName.trim();
    }
    
    if (filters.status.trim()) {
      filterParams['status_id[eq]'] = filters.status.trim();
    }
    
    return filterParams;
  }, [filters]);

  const queryParams = useMemo(() => ({
    current_page: currentPage,
    items_per_page: itemsPerPage,
    filters: buildFilters,
    lang: 'es'
  }), [currentPage, itemsPerPage, buildFilters]);

  const { data: consultationsData, isLoading: loading, error: queryError, refetch } = useListUserExpertConsultationsQuery(queryParams);
  const [createConsultation, { isLoading: isCreatingConsultation }] = useCreateExpertConsultationMutation();
  
  // Get consultation statuses
  const { data: statusesData, isLoading: statusesLoading } = useListExpertConsultationStatusesQuery({ lang: 'es' });

  // Query para obtener programas de grado
  const { data: programsData, isLoading: programsLoading } = useListDegreeProgramsQuery({
    current_page: 1,
    items_per_page: 100, // Obtener todos los programas
    filters: {}
  });

  // Query para obtener reportes del programa seleccionado
  const { data: reportsData, isLoading: reportsLoading, error: reportsError } = useListDegreeProgramReportsQuery(
    { 
      degree_program_id: selectedProgramId!,
      current_page: reportsCurrentPage,
      items_per_page: reportsItemsPerPage
    },
    { skip: !selectedProgramId }
  );

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


  const handleStatusChange = useCallback((value: string) => {
    // Convert "all" to empty string for filtering logic
    const filterValue = value === "all" ? "" : value;
    handleFilterChange('status', filterValue);
  }, [handleFilterChange]);

  // Clear all filters - memoized to prevent unnecessary re-renders
  const clearFilters = useCallback(() => {
    const clearedFilters = {
      programName: "",
      status: ""
    };
    setFilters(clearedFilters);
    setCurrentPage(1);
    updateURL(clearedFilters, 1, itemsPerPage);
  }, [itemsPerPage, updateURL]);

  const toggleFilters = useCallback(() => {
    setShowFilters(prev => !prev);
  }, []);

  // Funciones para el modal de creación
  const handleCreateConsultation = async () => {
    if (!selectedProgramId || !selectedReportId || !consultationMessage.trim()) {
      dispatch(showError({ type: 'error', message: 'Por favor completa todos los campos requeridos.' }));
      return;
    }

    try {
      await createConsultation({
        degree_program_id: selectedProgramId,
        report_id: selectedReportId,
        request_message: consultationMessage.trim(),
        requester_id: 1 // TODO: Get from auth context
      }).unwrap();

      dispatch(showError({ type: 'success', message: 'Solicitud de asesoría enviada exitosamente.' }));
      setIsCreateModalOpen(false);
      setConsultationMessage('');
      setSelectedProgramId(null);
      setSelectedReportId(null);
      refetch(); // Refrescar la lista de asesorías
    } catch (error: any) {
      const errorMessage = error?.data?.errors?.[0] || 'Error al enviar la solicitud de asesoría';
      dispatch(showError({ type: 'error', message: errorMessage }));
    }
  };

  const handleProgramChange = (programId: string) => {
    setSelectedProgramId(parseInt(programId));
    setSelectedReportId(null); // Reset report selection
    setReportsCurrentPage(1); // Reset reports pagination
  };

  const handleReportChange = (reportId: string) => {
    setSelectedReportId(parseInt(reportId));
  };

  const handleReportsPageChange = (page: number) => {
    setReportsCurrentPage(page);
  };

  const handleReportsItemsPerPageChange = (itemsPerPage: number) => {
    setReportsItemsPerPage(itemsPerPage);
    setReportsCurrentPage(1);
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("es-ES", {
      year: "numeric",
      month: "long",
      day: "numeric",
    });
  };

  const getStatusColor = (statusName: string) => {
    switch (statusName.toLowerCase()) {
      case 'pendiente':
        return "bg-yellow-100 text-yellow-800";
      case 'respondida':
        return "bg-green-100 text-green-800";
      case 'cerrada':
        return "bg-gray-100 text-gray-800";
      default:
        return "bg-blue-100 text-blue-800";
    }
  };

  const getStatusIcon = (statusName: string) => {
    switch (statusName.toLowerCase()) {
      case 'pendiente':
        return AlertTriangle;
      case 'respondida':
        return CheckCircle;
      case 'cerrada':
        return Clock;
      default:
        return MessageSquare;
    }
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
  const error = queryError && !is404Error ? "Error al cargar las asesorías" : null;

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
          Página {currentPage} de {totalPages} ({totalCount} asesorías)
        </div>
      </div>
    );
  };

  const ConsultationCard = ({ consultation }: { consultation: ExpertConsultation }) => {
    const StatusIcon = getStatusIcon(consultation.status_name);
    
    return (
      <Card className="hover:shadow-md transition-shadow">
        <CardHeader className="pb-3">
          <div className="flex items-start justify-between">
            <div className="flex items-center space-x-2">
              <StatusIcon className="h-5 w-5 text-blue-500" />
              <CardTitle className="text-lg">Asesoría #{consultation.id}</CardTitle>
            </div>
            <Badge className={getStatusColor(consultation.status_name)}>
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
              <p className="text-sm">{consultation.expert_name}</p>
            </div>
            
            <div className="flex items-center space-x-2">
              <FileText className="h-4 w-4 text-gray-500" />
              <p className="text-sm">Reporte #{consultation.report_id}</p>
            </div>
            
            <div className="flex items-center space-x-2">
              <Calendar className="h-4 w-4 text-gray-500" />
              <p className="text-sm">{formatDate(consultation.created_at)}</p>
            </div>
          </div>

          {/* Request Message Preview */}
          <div className="bg-gray-50 rounded-lg p-3">
            <h4 className="text-xs font-medium text-gray-700 mb-1 flex items-center gap-1">
              <MessageSquare className="h-3 w-3" />
              Solicitud
            </h4>
            <p className="text-xs text-gray-600 line-clamp-2">
              {consultation.request_message}
            </p>
          </div>

          {/* Expert Response Preview (if exists) */}
          {consultation.expert_response && (
            <div className="bg-green-50 rounded-lg p-3">
              <h4 className="text-xs font-medium text-gray-700 mb-1 flex items-center gap-1">
                <CheckCircle className="h-3 w-3" />
                Respuesta del Experto
              </h4>
              <p className="text-xs text-gray-600 line-clamp-2">
                {consultation.expert_response}
              </p>
            </div>
          )}
          
          <div className="pt-3 border-t">
            <Button 
              className="w-full" 
              variant="outline"
              onClick={() => router.push(`/director/asesoria-experto/${consultation.id}`)}
            >
              <Eye className="h-4 w-4 mr-2" />
              Ver Detalles
            </Button>
          </div>
        </CardContent>
      </Card>
    );
  };

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
          <h1 className="text-3xl font-bold text-gray-900">Asesorías con Experto</h1>
          <p className="text-gray-600 mt-2">Revisa las asesorías solicitadas a expertos en ciberseguridad</p>
        </div>
        <div className="flex gap-3">
          <Button 
            onClick={() => setIsCreateModalOpen(true)}
            className="bg-blue-600 hover:bg-blue-700 text-white"
          >
            <Plus className="h-4 w-4 mr-2" />
            Nueva Asesoría
          </Button>
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
                <MessageSquare className="h-4 w-4 text-blue-600" />
                <span className="text-sm font-medium text-gray-700">
                  {consultationsData.total_count} asesoría{consultationsData.total_count !== 1 ? 's' : ''}
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
            {(filters.programName || filters.status) && (
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
                  {filters.status && (
                    <Badge className="bg-purple-100 text-purple-800 border-purple-200 text-xs">
                      Estado: {statusesData?.data?.find(s => s.id.toString() === filters.status)?.name || filters.status}
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
              Personaliza tu búsqueda y visualización de asesorías
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


              {/* Status Filter */}
              <div className="space-y-2">
                <Label htmlFor="status-filter" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                  <Clock className="h-4 w-4 text-purple-500" />
                  Estado
                </Label>
                <Select value={filters.status || "all"} onValueChange={handleStatusChange}>
                  <SelectTrigger className="h-10">
                    <SelectValue placeholder="Seleccionar estado" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="all">Todos los estados</SelectItem>
                    {statusesLoading ? (
                      <SelectItem value="loading" disabled>Cargando estados...</SelectItem>
                    ) : statusesData?.data ? (
                      statusesData.data.map((status) => (
                        <SelectItem key={status.id} value={status.id.toString()}>
                          {status.name}
                        </SelectItem>
                      ))
                    ) : (
                      <SelectItem value="no-data" disabled>No hay estados disponibles</SelectItem>
                    )}
                  </SelectContent>
                </Select>
              </div>
              
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
  ), [showFilters, filters, consultationsData, queryError, itemsPerPage, toggleFilters, clearFilters, handleItemsPerPageChange, handleProgramNameChange, handleStatusChange]);

  if (error) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6">
          <div className="text-center">
            <AlertTriangle className="h-12 w-12 text-red-500 mx-auto mb-4" />
            <h2 className="text-xl font-semibold text-red-600 mb-2">Error al cargar asesorías</h2>
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
            <div className="bg-blue-50 rounded-full w-20 h-20 flex items-center justify-center mx-auto mb-6">
              <MessageSquare className="h-10 w-10 text-blue-500" />
            </div>
            <h3 className="text-xl font-semibold text-gray-800 mb-3">
              No hay asesorías
            </h3>
            <p className="text-gray-600 mb-2">
              No se encontraron asesorías con los filtros aplicados.
            </p>
            <p className="text-sm text-gray-500">
              Intenta ajustar los filtros o verifica si hay asesorías disponibles.
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

      {/* Modal para crear nueva asesoría */}
      <Dialog open={isCreateModalOpen} onOpenChange={setIsCreateModalOpen}>
        <DialogContent className="sm:max-w-4xl max-h-[80vh] overflow-hidden flex flex-col">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <Plus className="h-5 w-5 text-blue-600" />
              Nueva Solicitud de Asesoría
            </DialogTitle>
            <DialogDescription>
              Solicita asesoría especializada en ciberseguridad para tu programa de grado
            </DialogDescription>
          </DialogHeader>
          
          <div className="flex-1 overflow-y-auto space-y-6">
            {/* Selección de Programa */}
            <div className="space-y-2">
              <Label htmlFor="program-select" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                <GraduationCap className="h-4 w-4 text-blue-500" />
                Programa Académico *
              </Label>
              <Select value={selectedProgramId?.toString() || ""} onValueChange={handleProgramChange}>
                <SelectTrigger className="h-10">
                  <SelectValue placeholder="Selecciona un programa académico" />
                </SelectTrigger>
                <SelectContent>
                  {programsLoading ? (
                    <SelectItem value="" disabled>Cargando programas...</SelectItem>
                  ) : programsData?.data ? (
                    programsData.data.map((program) => (
                      <SelectItem key={program.id} value={program.id.toString()}>
                        {program.name} (SNIES: {program.snies})
                      </SelectItem>
                    ))
                  ) : (
                    <SelectItem value="" disabled>No hay programas disponibles</SelectItem>
                  )}
                </SelectContent>
              </Select>
            </div>

            {/* Selección de Reporte */}
            {selectedProgramId && (
              <div className="space-y-2">
                <Label htmlFor="report-select" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                  <FileText className="h-4 w-4 text-green-500" />
                  Reporte de Evaluación *
                </Label>
                {reportsLoading ? (
                  <div className="flex items-center justify-center py-8">
                    <div className="text-center text-gray-500">
                      <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto mb-2"></div>
                      <p>Cargando reportes...</p>
                    </div>
                  </div>
                ) : reportsError ? (
                  <div className="flex items-center justify-center py-8">
                    <div className="text-center text-gray-500">
                      <AlertTriangle className="h-8 w-8 text-red-500 mx-auto mb-2" />
                      <p>Error al cargar los reportes</p>
                    </div>
                  </div>
                ) : reportsData?.data && reportsData.data.length > 0 ? (
                  <div className="space-y-4">
                    <Select value={selectedReportId?.toString() || ""} onValueChange={handleReportChange}>
                      <SelectTrigger className="h-10">
                        <SelectValue placeholder="Selecciona un reporte" />
                      </SelectTrigger>
                      <SelectContent>
                        {reportsData.data.map((report) => (
                          <SelectItem key={report.id} value={report.id.toString()}>
                            Reporte #{report.id} - {(report as any).professional_role?.name || 'N/A'} 
                            ({(report.score * 100).toFixed(1)}%)
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                    
                    {/* Paginación de reportes */}
                    {reportsData.total_count > reportsItemsPerPage && (
                      <div className="flex items-center justify-between">
                        <div className="text-sm text-gray-600">
                          Mostrando {((reportsCurrentPage - 1) * reportsItemsPerPage) + 1} - {Math.min(reportsCurrentPage * reportsItemsPerPage, reportsData.total_count)} de {reportsData.total_count} reportes
                        </div>
                        <div className="flex gap-2">
                          <Button
                            variant="outline"
                            size="sm"
                            onClick={() => handleReportsPageChange(reportsCurrentPage - 1)}
                            disabled={reportsCurrentPage === 1}
                          >
                            Anterior
                          </Button>
                          <Button
                            variant="outline"
                            size="sm"
                            onClick={() => handleReportsPageChange(reportsCurrentPage + 1)}
                            disabled={reportsCurrentPage >= Math.ceil(reportsData.total_count / reportsItemsPerPage)}
                          >
                            Siguiente
                          </Button>
                        </div>
                      </div>
                    )}
                  </div>
                ) : (
                  <div className="flex items-center justify-center py-8">
                    <div className="text-center text-gray-500">
                      <FileText className="h-8 w-8 text-gray-400 mx-auto mb-2" />
                      <p>No hay reportes disponibles para este programa</p>
                    </div>
                  </div>
                )}
              </div>
            )}

            {/* Mensaje de solicitud */}
            <div className="space-y-2">
              <Label htmlFor="consultation-message" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                <MessageSquare className="h-4 w-4 text-purple-500" />
                Mensaje de solicitud *
              </Label>
              <Textarea
                id="consultation-message"
                placeholder="Describe tu consulta o solicitud de asesoría..."
                value={consultationMessage}
                onChange={(e) => setConsultationMessage(e.target.value)}
                rows={4}
                className="w-full"
              />
            </div>
            
            {/* Información adicional */}
            <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
              <div className="flex items-start gap-3">
                <User className="h-5 w-5 text-blue-600 mt-0.5" />
                <div className="text-sm text-blue-800">
                  <p className="font-medium mb-2">Información importante:</p>
                  <ul className="space-y-1 text-xs">
                    <li>• Tu solicitud será revisada por un experto en ciberseguridad</li>
                    <li>• Se te notificará al correo la respuesta del experto</li>
                    <li>• El tiempo de respuesta puede variar según la complejidad de la consulta</li>
                    <li>• Asegúrate de proporcionar información clara y específica</li>
                  </ul>
                </div>
              </div>
            </div>
          </div>

          {/* Botones de acción */}
          <div className="flex gap-3 justify-end pt-4 border-t">
            <Button
              variant="outline"
              onClick={() => {
                setIsCreateModalOpen(false);
                setConsultationMessage('');
                setSelectedProgramId(null);
                setSelectedReportId(null);
              }}
            >
              Cancelar
            </Button>
            <Button
              onClick={handleCreateConsultation}
              disabled={isCreatingConsultation || !selectedProgramId || !selectedReportId || !consultationMessage.trim()}
              className="bg-blue-600 hover:bg-blue-700"
            >
              {isCreatingConsultation ? (
                <>
                  <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                  Enviando...
                </>
              ) : (
                <>
                  <MessageSquare className="h-4 w-4 mr-2" />
                  Enviar Solicitud
                </>
              )}
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}
