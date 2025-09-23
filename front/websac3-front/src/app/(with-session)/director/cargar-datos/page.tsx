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
import { Upload, ArrowLeft, CheckCircle, BookOpen, Clock, GraduationCap, Target, Search, Settings, Calendar } from "lucide-react";
import { useListDegreeProgramsQuery } from "@/services/api";
import { useRouter, useSearchParams } from "next/navigation";
import { ProgramsPagination } from "@/components/websac3/program/ProgramsPagination";

export default function CargarDatosPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [selectedProgramId, setSelectedProgramId] = useState<number | null>(null);
  
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
      ? `/director/cargar-datos?${params.toString()}`
      : '/director/cargar-datos';
    
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

  const handleLoadData = (programId: number) => {
    router.push(`/director/cargar-datos/formulario?program_id=${programId}`);
  };

  if (isLoading) {
    return <LoadingSkeleton />;
  }

  if (error) {
  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
          <div className="text-center text-gray-500">
            <BookOpen className="h-12 w-12 text-gray-400 mx-auto mb-4" />
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
          <h1 className="text-3xl font-bold text-gray-900">Cargar Datos</h1>
          <p className="text-gray-600 mt-2">Selecciona un programa de grado para cargar cursos</p>
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
                  {data.total_count} programa{data.total_count !== 1 ? 's' : ''} encontrado{data.total_count !== 1 ? 's' : ''}
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

              {/* Load Data Button */}
              <div className="pt-4 border-t border-gray-100">
                <Button
                  onClick={() => handleLoadData(program.id)}
                  className="w-full bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 shadow-lg hover:shadow-xl transition-all duration-200"
                >
                  <Upload className="h-4 w-4 mr-2" />
                  Cargar Datos
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
                <BookOpen className="h-12 w-12 text-gray-400 mx-auto mb-4" />
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
    </div>
  );
}

// Loading skeleton component
function LoadingSkeleton() {
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <Skeleton className="h-8 w-48 mb-2" />
          <Skeleton className="h-4 w-64" />
        </div>
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
