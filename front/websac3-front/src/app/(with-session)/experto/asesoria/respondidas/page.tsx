"use client";

import React, { useState, useCallback, useMemo, memo } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Label } from "@/components/ui/label";
import { 
  MessageSquare, 
  GraduationCap, 
  Calendar,
  FileText,
  CheckCircle,
  Eye,
  ChevronLeft,
  ChevronRight,
  Search,
  Filter,
  User,
  Mail,
  BookOpen
} from "lucide-react";
import { useListAnsweredExpertConsultationsQuery } from "@/services/api";
import { Input } from "@/components/ui/input";

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

FilterInput.displayName = 'FilterInput';

export default function AnsweredConsultationsPage() {
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

  const { data: consultationsData, isLoading, error } = useListAnsweredExpertConsultationsQuery(queryParams);

  const handlePageChange = useCallback((page: number) => {
    setCurrentPage(page);
    updateURL(filters, page, itemsPerPage);
  }, [filters, itemsPerPage, updateURL]);

  const handleFilterChange = useCallback((field: keyof typeof filters, value: string) => {
    const newFilters = { ...filters, [field]: value };
    setFilters(newFilters);
    setCurrentPage(1); // Reset to first page when filters change
    updateURL(newFilters, 1, itemsPerPage);
  }, [filters, itemsPerPage, updateURL]);

  const clearFilters = useCallback(() => {
    const emptyFilters = {
      programName: "",
      requesterName: "",
      requesterEmail: "",
      institutionName: ""
    };
    setFilters(emptyFilters);
    setCurrentPage(1);
    updateURL(emptyFilters, 1, itemsPerPage);
  }, [itemsPerPage, updateURL]);

  const hasActiveFilters = useMemo(() => 
    Object.values(filters).some(value => value.trim() !== ""),
    [filters]
  );

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("es-ES", {
      year: "numeric",
      month: "short",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit"
    });
  };

  const formatScore = (score: number) => {
    const percentage = Math.floor(score * 10000) / 100;
    return percentage % 1 === 0 ? percentage.toFixed(0) : percentage.toFixed(2);
  };

  const getStatusBadgeColor = (statusName: string) => {
    const status = statusName.toLowerCase();
    if (status.includes('aceptada') || status.includes('accepted')) {
      return 'bg-green-100 text-green-800';
    } else if (status.includes('rechazada') || status.includes('rejected')) {
      return 'bg-red-100 text-red-800';
    } else if (status.includes('cerrada') || status.includes('closed')) {
      return 'bg-gray-100 text-gray-800';
    }
    return 'bg-blue-100 text-blue-800';
  };

  // Filter consultations based on filters
  const filteredConsultations = useMemo(() => {
    if (!consultationsData?.data) return [];
    
    return consultationsData.data.filter(consultation => {
      const matchesProgram = !filters.programName.trim() || 
        consultation.degree_program_name.toLowerCase().includes(filters.programName.toLowerCase());
      
      const matchesRequester = !filters.requesterName.trim() || 
        consultation.requester_name.toLowerCase().includes(filters.requesterName.toLowerCase());
      
      const matchesEmail = !filters.requesterEmail.trim() || 
        consultation.requester_email.toLowerCase().includes(filters.requesterEmail.toLowerCase());
      
      const matchesInstitution = !filters.institutionName.trim() || 
        consultation.requester_institution_name.toLowerCase().includes(filters.institutionName.toLowerCase());
      
      return matchesProgram && matchesRequester && matchesEmail && matchesInstitution;
    });
  }, [consultationsData?.data, filters]);

  const totalPages = consultationsData ? Math.ceil(consultationsData.total_count / itemsPerPage) : 0;

  if (isLoading) {
    return (
      <div className="space-y-6">
        <div className="flex items-center justify-between">
          <Skeleton className="h-10 w-64" />
          <Skeleton className="h-10 w-32" />
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[1, 2, 3, 4, 5, 6].map((i) => (
            <Card key={i}>
              <CardHeader>
                <Skeleton className="h-6 w-full" />
              </CardHeader>
              <CardContent className="space-y-3">
                <Skeleton className="h-4 w-3/4" />
                <Skeleton className="h-4 w-1/2" />
                <Skeleton className="h-20 w-full" />
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="space-y-6">
        <h1 className="text-3xl font-bold text-gray-900">Solicitudes Respondidas</h1>
        <Card>
          <CardContent className="p-6">
            <div className="text-center text-red-600">
              <p className="text-lg font-semibold mb-2">Error al cargar las solicitudes</p>
              <p className="text-sm">Por favor, intenta nuevamente más tarde.</p>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Solicitudes Respondidas</h1>
          <p className="text-gray-600 mt-1">Revisa las asesorías que ya has respondido</p>
        </div>
        <div className="flex items-center gap-2">
          <Badge variant="outline" className="text-lg px-4 py-2">
            Total: {consultationsData?.total_count || 0}
          </Badge>
        </div>
      </div>

      {/* Filters Section */}
      <Card>
        <CardHeader>
          <div className="flex items-center justify-between">
            <CardTitle className="flex items-center gap-2">
              <Filter className="h-5 w-5 text-blue-600" />
              Filtros de Búsqueda
            </CardTitle>
            <div className="flex items-center gap-2">
              {hasActiveFilters && (
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={clearFilters}
                  className="text-red-600 hover:text-red-700 hover:bg-red-50"
                >
                  Limpiar Filtros
                </Button>
              )}
              <Button
                variant="outline"
                size="sm"
                onClick={() => setShowFilters(!showFilters)}
              >
                {showFilters ? 'Ocultar' : 'Mostrar'} Filtros
              </Button>
            </div>
          </div>
        </CardHeader>
        {showFilters && (
          <CardContent>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
              <FilterInput
                id="program-filter"
                label="Programa"
                placeholder="Buscar por programa..."
                value={filters.programName}
                onChange={(value) => handleFilterChange('programName', value)}
                icon={BookOpen}
              />
              
              <FilterInput
                id="requester-filter"
                label="Solicitante"
                placeholder="Buscar por nombre..."
                value={filters.requesterName}
                onChange={(value) => handleFilterChange('requesterName', value)}
                icon={User}
              />
              
              <FilterInput
                id="email-filter"
                label="Email"
                placeholder="Buscar por email..."
                value={filters.requesterEmail}
                onChange={(value) => handleFilterChange('requesterEmail', value)}
                icon={Mail}
              />
              
              <FilterInput
                id="institution-filter"
                label="Institución"
                placeholder="Buscar por institución..."
                value={filters.institutionName}
                onChange={(value) => handleFilterChange('institutionName', value)}
                icon={GraduationCap}
              />
            </div>
          </CardContent>
        )}
      </Card>

      {/* Consultations Grid */}
      {filteredConsultations.length === 0 ? (
        <Card>
          <CardContent className="p-12">
            <div className="text-center text-gray-500">
              <CheckCircle className="h-12 w-12 mx-auto mb-4 text-gray-400" />
              <h2 className="text-xl font-semibold mb-2">
                {hasActiveFilters ? 'No se encontraron resultados' : 'No hay solicitudes respondidas'}
              </h2>
              <p className="text-sm mb-4">
                {hasActiveFilters 
                  ? 'Intenta ajustar los filtros de búsqueda' 
                  : 'Las solicitudes que respondas aparecerán aquí'}
              </p>
              {hasActiveFilters && (
                <Button onClick={clearFilters} variant="outline">
                  Limpiar Filtros
                </Button>
              )}
            </div>
          </CardContent>
        </Card>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {filteredConsultations.map((consultation) => (
              <Card key={consultation.id} className="hover:shadow-lg transition-shadow">
                <CardHeader>
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <CardTitle className="text-lg line-clamp-2">
                        {consultation.degree_program_name}
                      </CardTitle>
                      <p className="text-xs text-gray-500 mt-1">
                        SNIES: {consultation.degree_program_snies}
                      </p>
                    </div>
                    <Badge className={getStatusBadgeColor(consultation.status_name)}>
                      {consultation.status_name}
                    </Badge>
                  </div>
                </CardHeader>
                <CardContent className="space-y-3">
                  {/* Requester Info */}
                  <div className="flex items-center space-x-2 text-sm">
                    <GraduationCap className="h-4 w-4 text-gray-500 flex-shrink-0" />
                    <div className="min-w-0 flex-1">
                      <p className="font-medium truncate">{consultation.requester_name}</p>
                      <p className="text-xs text-gray-500 truncate">
                        {consultation.requester_institution_name}
                      </p>
                    </div>
                  </div>
                  
                  {/* Report Info */}
                  <div className="flex items-center space-x-2 text-sm">
                    <FileText className="h-4 w-4 text-gray-500" />
                    <div>
                      <span className="text-gray-600">Reporte #{consultation.report_id}</span>
                      <span className="ml-2 text-xs text-gray-500">
                        Score: {formatScore(consultation.report_score)}%
                      </span>
                    </div>
                  </div>
                  
                  {/* Dates */}
                  <div className="flex items-center space-x-2 text-xs text-gray-500">
                    <Calendar className="h-3 w-3" />
                    <div>
                      <p>Respondida: {formatDate(consultation.answered_at)}</p>
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

                  {/* Expert Response Preview */}
                  <div className="bg-green-50 rounded-lg p-3">
                    <h4 className="text-xs font-medium text-green-800 mb-1 flex items-center gap-1">
                      <CheckCircle className="h-3 w-3" />
                      Tu Respuesta
                    </h4>
                    <p className="text-xs text-green-900 line-clamp-2">
                      {consultation.expert_response}
                    </p>
                  </div>
                  
                  {/* View Details Button */}
                  <div className="pt-2">
                    <Button 
                      className="w-full" 
                      variant="outline"
                      onClick={() => router.push(`/experto/asesoria/respondidas/${consultation.id}`)}
                    >
                      <Eye className="h-4 w-4 mr-2" />
                      Ver Detalles
                    </Button>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>

          {/* Pagination */}
          {totalPages > 1 && (
            <Card>
              <CardContent className="p-4">
                <div className="flex items-center justify-between">
                  <Button
                    variant="outline"
                    onClick={() => handlePageChange(Math.max(1, currentPage - 1))}
                    disabled={currentPage === 1}
                  >
                    <ChevronLeft className="h-4 w-4 mr-2" />
                    Anterior
                  </Button>
                  
                  <div className="flex items-center gap-2">
                    <span className="text-sm text-gray-600">
                      Página {currentPage} de {totalPages}
                    </span>
                    <span className="text-xs text-gray-500">
                      ({filteredConsultations.length} de {consultationsData?.total_count || 0} resultados)
                    </span>
                  </div>
                  
                  <Button
                    variant="outline"
                    onClick={() => handlePageChange(Math.min(totalPages, currentPage + 1))}
                    disabled={currentPage === totalPages}
                  >
                    Siguiente
                    <ChevronRight className="h-4 w-4 ml-2" />
                  </Button>
                </div>
              </CardContent>
            </Card>
          )}
        </>
      )}
    </div>
  );
}
