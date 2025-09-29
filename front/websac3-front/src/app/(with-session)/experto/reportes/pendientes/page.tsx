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
  AlertTriangle, 
  Calendar, 
  GraduationCap, 
  Building, 
  User, 
  Star,
  Search,
  Filter,
  Settings,
  BookOpen,
  Clock,
  ArrowUpDown
} from "lucide-react";
import { useListPendingReportsQuery } from "@/services/api";
import { useRouter, useSearchParams } from "next/navigation";

interface PendingReport {
  id: number;
  created_at: string;
  score: number;
  has_feedback: boolean;
  degree_program: {
    id: number;
    name: string;
    snies: number;
  };
  higher_education_institution: {
    id: number;
    name: string;
  };
  professional_role: {
    id: number;
    name: string;
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

export default function PendingReportsPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(10);
  const [sortBy, setSortBy] = useState('created_at');
  const [sortOrder, setSortOrder] = useState('desc');
  const [filters, setFilters] = useState({
    institutionName: "",
    institutionSnies: ""
  });
  
  const [showFilters, setShowFilters] = useState(false);

  // Initialize filters from URL on component mount
  React.useEffect(() => {
    const urlFilters = {
      institutionName: searchParams.get('institutionName') || "",
      institutionSnies: searchParams.get('institutionSnies') || ""
    };
    
    const urlPage = parseInt(searchParams.get('page') || '1');
    const urlItemsPerPage = parseInt(searchParams.get('itemsPerPage') || '10');
    const urlSortBy = searchParams.get('sortBy') || 'created_at';
    const urlSortOrder = searchParams.get('sortOrder') || 'desc';
    
    setFilters(urlFilters);
    setCurrentPage(urlPage);
    setItemsPerPage(urlItemsPerPage);
    setSortBy(urlSortBy);
    setSortOrder(urlSortOrder);
  }, [searchParams]);

  // Function to update URL with current filters and pagination
  const updateURL = useCallback((newFilters: typeof filters, newPage: number, newItemsPerPage: number, newSortBy: string, newSortOrder: string) => {
    const params = new URLSearchParams();
    
    // Add pagination params
    params.set('page', newPage.toString());
    params.set('itemsPerPage', newItemsPerPage.toString());
    params.set('sortBy', newSortBy);
    params.set('sortOrder', newSortOrder);
    
    // Add filter params only if they have values
    if (newFilters.institutionName.trim()) {
      params.set('institutionName', newFilters.institutionName.trim());
    }
    if (newFilters.institutionSnies.trim()) {
      params.set('institutionSnies', newFilters.institutionSnies.trim());
    }
    
    // Update URL without causing a page reload
    const newURL = `${window.location.pathname}?${params.toString()}`;
    router.replace(newURL, { scroll: false });
  }, [router]);

  // Build query parameters with proper filter format - memoized to prevent unnecessary re-renders
  const buildFilters = useMemo(() => {
    const filterParams: Record<string, string> = {};
    
    if (filters.institutionName.trim()) {
      filterParams['DegreeProgram.UserCreator.Person.HigherEducationInstitution.name[cont]'] = filters.institutionName.trim();
    }
    
    if (filters.institutionSnies.trim()) {
      filterParams['DegreeProgram.UserCreator.Person.HigherEducationInstitution.snies[cont]'] = filters.institutionSnies.trim();
    }
    
    return filterParams;
  }, [filters]);

  const queryParams = useMemo(() => ({
    current_page: currentPage,
    items_per_page: itemsPerPage,
    sort_by: sortBy,
    sort_order: sortOrder,
    filters: buildFilters,
    lang: 'es'
  }), [currentPage, itemsPerPage, sortBy, sortOrder, buildFilters]);

  const {
    data: reportsData,
    isLoading: loading,
    error: queryError,
    refetch
  } = useListPendingReportsQuery(queryParams);

  const handlePageChange = useCallback((page: number) => {
    setCurrentPage(page);
    updateURL(filters, page, itemsPerPage, sortBy, sortOrder);
  }, [filters, itemsPerPage, sortBy, sortOrder, updateURL]);

  const handleItemsPerPageChange = useCallback((perPage: number) => {
    setItemsPerPage(perPage);
    setCurrentPage(1); // Reset to first page
    updateURL(filters, 1, perPage, sortBy, sortOrder);
  }, [filters, sortBy, sortOrder, updateURL]);

  const handleSortByChange = useCallback((newSortBy: string) => {
    setSortBy(newSortBy);
    setCurrentPage(1); // Reset to first page
    updateURL(filters, 1, itemsPerPage, newSortBy, sortOrder);
  }, [filters, itemsPerPage, sortOrder, updateURL]);

  const handleSortOrderChange = useCallback((newSortOrder: string) => {
    setSortOrder(newSortOrder);
    setCurrentPage(1); // Reset to first page
    updateURL(filters, 1, itemsPerPage, sortBy, newSortOrder);
  }, [filters, itemsPerPage, sortBy, updateURL]);

  // Handle filter changes - memoized to prevent unnecessary re-renders
  const handleFilterChange = useCallback((key: string, value: string) => {
    const newFilters = { ...filters, [key]: value };
    setFilters(newFilters);
    setCurrentPage(1); // Reset to first page
    updateURL(newFilters, 1, itemsPerPage, sortBy, sortOrder);
  }, [filters, itemsPerPage, sortBy, sortOrder, updateURL]);

  // Specific handlers for each filter to prevent unnecessary re-renders
  const handleInstitutionNameChange = useCallback((value: string) => {
    handleFilterChange('institutionName', value);
  }, [handleFilterChange]);

  const handleInstitutionSniesChange = useCallback((value: string) => {
    handleFilterChange('institutionSnies', value);
  }, [handleFilterChange]);

  // Clear all filters - memoized to prevent unnecessary re-renders
  const clearFilters = useCallback(() => {
    const clearedFilters = {
      institutionName: "",
      institutionSnies: ""
    };
    setFilters(clearedFilters);
    setCurrentPage(1);
    updateURL(clearedFilters, 1, itemsPerPage, sortBy, sortOrder);
  }, [itemsPerPage, sortBy, sortOrder, updateURL]);

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

  // Extract data from the query result
  const reports = reportsData?.data || [];
  const totalCount = reportsData?.total_count || 0;
  const currentPageFromAPI = reportsData?.current_page || 1;
  const itemsPerPageFromAPI = reportsData?.items_per_page || 10;
  
  // Calculate total pages based on total count and items per page
  const totalPages = Math.ceil(totalCount / itemsPerPageFromAPI);
  
  // Check if it's a 404 error (no reports found) vs a real error
  const is404Error = queryError && 'status' in queryError && queryError.status === 404;
  const error = queryError && !is404Error ? "Error al cargar los reportes pendientes" : null;

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
          Página {currentPage} de {totalPages} ({totalCount} reportes)
        </div>
      </div>
    );
  };

  const ReportCard = ({ report }: { report: PendingReport }) => (
    <Card className="hover:shadow-md transition-shadow">
      <CardHeader className="pb-3">
        <div className="flex items-start justify-between">
          <div className="flex items-center space-x-2">
            <AlertTriangle className="h-5 w-5 text-orange-500" />
            <CardTitle className="text-lg">Reporte #{report.id}</CardTitle>
          </div>
          <Badge className={getScoreColor(report.score * 100)}>
            <Star className="h-3 w-3 mr-1" />
            {(() => {
              const percentage = Math.floor(report.score * 10000) / 100;
              return percentage % 1 === 0 ? percentage.toFixed(0) : percentage.toFixed(2);
            })()}%
          </Badge>
        </div>
      </CardHeader>
      <CardContent className="space-y-3">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="flex items-center space-x-2">
            <GraduationCap className="h-4 w-4 text-gray-500" />
            <div>
              <p className="text-sm font-medium">{report.degree_program.name}</p>
              <p className="text-xs text-gray-500">SNIES: {report.degree_program.snies}</p>
            </div>
          </div>
          
          <div className="flex items-center space-x-2">
            <Building className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{report.higher_education_institution.name}</p>
          </div>
          
          <div className="flex items-center space-x-2">
            <User className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{report.professional_role.name}</p>
          </div>
          
          <div className="flex items-center space-x-2">
            <Calendar className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{formatDate(report.created_at)}</p>
          </div>
        </div>
        
        <div className="pt-3 border-t">
          <Button 
            className="w-full" 
            variant="outline"
            onClick={() => router.push(`/experto/reportes/${report.id}`)}
          >
            Revisar Reporte
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
          <h1 className="text-3xl font-bold text-gray-900">Reportes Pendientes</h1>
          <p className="text-gray-600 mt-2">Revisa los reportes que requieren retroalimentación</p>
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
      {reportsData && !(queryError && queryError.status === 404) && (
        <div className="bg-white border border-gray-200 rounded-lg shadow-sm p-4">
          <div className="flex flex-wrap items-center justify-between gap-4">
            {/* Left side - Main info */}
            <div className="flex items-center gap-6">
              <div className="flex items-center gap-2">
                <AlertTriangle className="h-4 w-4 text-orange-600" />
                <span className="text-sm font-medium text-gray-700">
                  {reportsData.total_count} reporte{reportsData.total_count !== 1 ? 's' : ''} pendiente{reportsData.total_count !== 1 ? 's' : ''}
                </span>
              </div>
              
              <div className="flex items-center gap-2">
                <Calendar className="h-4 w-4 text-gray-500" />
                <span className="text-sm text-gray-600">
                  Página {reportsData.current_page} de {Math.ceil(reportsData.total_count / itemsPerPage)}
                </span>
              </div>
            </div>

            {/* Right side - Active filters info */}
            {(filters.institutionName || filters.institutionSnies) && (
              <div className="flex items-center gap-3">
                <div className="flex items-center gap-2">
                  <Filter className="h-4 w-4 text-indigo-600" />
                  <span className="text-sm font-medium text-gray-700">
                    Filtros activos
                  </span>
                </div>
                <div className="flex gap-2 flex-wrap">
                  {filters.institutionName && (
                    <Badge className="bg-blue-100 text-blue-800 border-blue-200 text-xs">
                      Institución: {filters.institutionName}
                    </Badge>
                  )}
                  {filters.institutionSnies && (
                    <Badge className="bg-green-100 text-green-800 border-green-200 text-xs">
                      SNIES Inst.: {filters.institutionSnies}
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
              Personaliza tu búsqueda y visualización de reportes pendientes
            </CardDescription>
          </CardHeader>
          <CardContent className="p-6">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {/* Institution Name Filter */}
              <FilterInput
                id="institution-filter"
                label="Institución Educativa"
                placeholder="Buscar por institución..."
                value={filters.institutionName}
                onChange={handleInstitutionNameChange}
                icon={Building}
                className="focus:border-blue-500 focus:ring-blue-500"
              />

              {/* Institution SNIES Filter */}
              <FilterInput
                id="institution-snies-filter"
                label="SNIES Institución Educativa"
                placeholder="Buscar por SNIES de institución..."
                value={filters.institutionSnies}
                onChange={handleInstitutionSniesChange}
                icon={Building}
                className="focus:border-green-500 focus:ring-green-500"
              />
              
              {/* Sort By */}
              <div className="space-y-2">
                <Label htmlFor="sort-by" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                  <ArrowUpDown className="h-4 w-4 text-indigo-500" />
                  Ordenar por
                </Label>
                <Select value={sortBy} onValueChange={handleSortByChange}>
                  <SelectTrigger className="h-10">
                    <SelectValue placeholder="Seleccionar campo" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="created_at">Fecha de creación</SelectItem>
                  </SelectContent>
                </Select>
              </div>

              {/* Sort Order */}
              <div className="space-y-2">
                <Label htmlFor="sort-order" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                  <ArrowUpDown className="h-4 w-4 text-indigo-500" />
                  Orden
                </Label>
                <Select value={sortOrder} onValueChange={handleSortOrderChange}>
                  <SelectTrigger className="h-10">
                    <SelectValue placeholder="Seleccionar orden" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="desc">Descendente (Más recientes primero)</SelectItem>
                    <SelectItem value="asc">Ascendente (Más antiguos primero)</SelectItem>
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
  ), [showFilters, filters, reportsData, queryError, itemsPerPage, sortBy, sortOrder, toggleFilters, clearFilters, handleItemsPerPageChange, handleSortByChange, handleSortOrderChange, handleInstitutionNameChange, handleInstitutionSniesChange]);

  if (error) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6">
          <div className="text-center">
            <AlertTriangle className="h-12 w-12 text-red-500 mx-auto mb-4" />
            <h2 className="text-xl font-semibold text-red-600 mb-2">Error al cargar reportes</h2>
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
        ) : (reports.length === 0 || is404Error) ? (
          <div className="text-center py-12">
            <div className="bg-green-50 rounded-full w-20 h-20 flex items-center justify-center mx-auto mb-6">
              <AlertTriangle className="h-10 w-10 text-green-500" />
            </div>
            <h3 className="text-xl font-semibold text-gray-800 mb-3">
              ¡Excelente trabajo!
            </h3>
            <p className="text-gray-600 mb-2">
              No hay reportes pendientes de revisión en este momento.
            </p>
            <p className="text-sm text-gray-500">
              Todos los reportes han sido procesados o no hay nuevos reportes disponibles.
            </p>
          </div>
        ) : (
          <>
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              {reports.map((report) => (
                <ReportCard key={report.id} report={report} />
              ))}
            </div>
            
            
            {totalCount > 0 && <Pagination />}
          </>
        )}
      </div>
    </div>
  );
}
