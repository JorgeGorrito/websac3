"use client";

import { useState, useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useListDegreeProgramsQuery } from "@/services/api";
import { Plus, BookOpen, Calendar, Clock, GraduationCap, Target, Search, Settings } from "lucide-react";
import { DegreeProgramCard } from "@/components/websac3/program/DegreeProgramCard";
import { ProgramsPagination } from "@/components/websac3/program/ProgramsPagination";
import { useDispatch } from "react-redux";
import { showError } from "@/store/errorSlice";

export default function ProgramasPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const dispatch = useDispatch();
  
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

  const { data, isLoading, error, refetch } = useListDegreeProgramsQuery(queryParams);

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
      ? `/director/programas?${params.toString()}`
      : '/director/programas';
    
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

  const handleCreateProgram = () => {
    router.push("/director/registrar-programa");
  };

  // Note: Error handling is now done globally via ErrorHandler component
  // This error state is only for UI fallback
  if (error) {
    return (
      <div className="space-y-6">
        <div className="flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Programas de Grado</h1>
            <p className="text-gray-600 mt-2">Gestiona los programas de grado registrados</p>
          </div>
          <Button onClick={handleCreateProgram} className="bg-blue-600 hover:bg-blue-700">
            <Plus className="h-4 w-4 mr-2" />
            Registrar Programa de Grado
          </Button>
        </div>
        
        <Card>
          <CardContent className="p-6">
            <div className="text-center text-gray-500">
              <BookOpen className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <p>No se pudieron cargar los programas de grado</p>
              <p className="text-sm mt-2">Los errores se mostrarán en un modal automáticamente</p>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Programas de Grado</h1>
          <p className="text-gray-600 mt-2">Gestiona los programas de grado registrados</p>
        </div>
        <div className="flex gap-3">
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
          <Button 
            onClick={handleCreateProgram} 
            className="bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 shadow-lg hover:shadow-xl transition-all duration-200"
          >
            <Plus className="h-4 w-4 mr-2" />
            Registrar Programa de Grado
          </Button>
        </div>
      </div>

      {/* Stats */}
      {data && (
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
          <Card className="bg-gradient-to-br from-blue-50 to-blue-100 border-blue-200 shadow-lg hover:shadow-xl transition-shadow duration-300">
            <CardContent className="p-6">
              <div className="flex items-center">
                <div className="p-3 bg-blue-500 rounded-xl shadow-md">
                  <BookOpen className="h-8 w-8 text-white" />
                </div>
                <div className="ml-4">
                  <p className="text-sm font-semibold text-blue-700 uppercase tracking-wide">Total Programas</p>
                  <p className="text-3xl font-bold text-blue-900 mt-1">{data.total_count}</p>
                </div>
              </div>
            </CardContent>
          </Card>
          
          <Card className="bg-gradient-to-br from-green-50 to-green-100 border-green-200 shadow-lg hover:shadow-xl transition-shadow duration-300">
            <CardContent className="p-6">
              <div className="flex items-center">
                <div className="p-3 bg-green-500 rounded-xl shadow-md">
                  <Calendar className="h-8 w-8 text-white" />
                </div>
                <div className="ml-4">
                  <p className="text-sm font-semibold text-green-700 uppercase tracking-wide">Página Actual</p>
                  <p className="text-3xl font-bold text-green-900 mt-1">{data.current_page}</p>
                </div>
              </div>
            </CardContent>
          </Card>
          
          <Card className="bg-gradient-to-br from-orange-50 to-orange-100 border-orange-200 shadow-lg hover:shadow-xl transition-shadow duration-300">
            <CardContent className="p-6">
              <div className="flex items-center">
                <div className="p-3 bg-orange-500 rounded-xl shadow-md">
                  <Clock className="h-8 w-8 text-white" />
                </div>
                <div className="ml-4">
                  <p className="text-sm font-semibold text-orange-700 uppercase tracking-wide">Por Página</p>
                  <p className="text-3xl font-bold text-orange-900 mt-1">{itemsPerPage}</p>
                </div>
              </div>
            </CardContent>
          </Card>
          
          <Card className="bg-gradient-to-br from-purple-50 to-purple-100 border-purple-200 shadow-lg hover:shadow-xl transition-shadow duration-300">
            <CardContent className="p-6">
              <div className="flex items-center">
                <div className="p-3 bg-purple-500 rounded-xl shadow-md">
                  <Target className="h-8 w-8 text-white" />
                </div>
                <div className="ml-4">
                  <p className="text-sm font-semibold text-purple-700 uppercase tracking-wide">Total Páginas</p>
                  <p className="text-3xl font-bold text-purple-900 mt-1">
                    {Math.ceil(data.total_count / itemsPerPage)}
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>
          
          {/* Active Filters Info */}
          {(filters.name || filters.snies) && (
            <Card className="md:col-span-4 bg-gradient-to-r from-amber-50 to-yellow-50 border-amber-200 shadow-lg">
              <CardContent className="p-5">
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-3">
                    <div className="p-2 bg-amber-500 rounded-lg">
                      <Search className="h-5 w-5 text-white" />
                    </div>
                    <div>
                      <span className="text-sm font-semibold text-amber-800">Filtros activos:</span>
                      <div className="flex items-center gap-2 mt-1">
                        {filters.name && (
                          <Badge className="bg-blue-100 text-blue-800 border-blue-200 font-medium">
                            Nombre: "{filters.name}"
                          </Badge>
                        )}
                        {filters.snies && (
                          <Badge className="bg-green-100 text-green-800 border-green-200 font-medium">
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
                    className="border-amber-300 text-amber-700 hover:bg-amber-100"
                  >
                    Limpiar
                  </Button>
                </div>
              </CardContent>
            </Card>
          )}
          
          {data.data.length > 0 && (
            <Card className="md:col-span-4 bg-gradient-to-r from-indigo-50 to-blue-50 border-indigo-200 shadow-lg py-0">
              <CardContent className="p-4">
                <div className="flex items-center justify-between">
                  <div className="flex items-center">
                    <div className="p-3 bg-indigo-500 rounded-xl shadow-md">
                      <GraduationCap className="h-8 w-8 text-white" />
                    </div>
                    <div className="ml-4">
                      <p className="text-sm font-semibold text-indigo-700 uppercase tracking-wide">Institución</p>
                      <p className="text-xl font-bold text-gray-900 mt-1">
                        {data.data[0].higher_education_institution.name}
                      </p>
                    </div>
                  </div>
                  <div className="text-right">
                    <p className="text-sm font-semibold text-indigo-700 uppercase tracking-wide">Código SNIES</p>
                    <div className="mt-2">
                      <Badge className="bg-indigo-100 text-indigo-800 border-indigo-200 font-bold text-lg px-4 py-2">
                        {data.data[0].higher_education_institution.snies}
                      </Badge>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          )}
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
      <div className="space-y-4">
        {isLoading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {Array.from({ length: 6 }).map((_, index) => (
              <Card key={index}>
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
          <>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {data.data.map((program) => (
                <DegreeProgramCard key={program.id} program={program} />
              ))}
            </div>
          </>
        ) : (
          <Card>
            <CardContent className="p-12">
              <div className="text-center">
                <BookOpen className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                <h3 className="text-lg font-medium text-gray-900 mb-2">
                  No hay programas registrados
                </h3>
                <p className="text-gray-600 mb-6">
                  Comienza registrando tu primer programa de grado
                </p>
                <Button onClick={handleCreateProgram} className="bg-blue-600 hover:bg-blue-700">
                  <Plus className="h-4 w-4 mr-2" />
                  Registrar Programa de Grado
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
