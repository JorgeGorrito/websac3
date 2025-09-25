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
  Shield, 
  BookOpen, 
  Clock, 
  GraduationCap, 
  Target, 
  Search, 
  Settings, 
  Calendar,
  CheckCircle,
  AlertCircle,
  Mail
} from "lucide-react";
import { useListDegreeProgramsQuery, useListProfessionalRolesQuery, useEvaluateDegreeProgramMutation } from "@/services/api";
import { useRouter, useSearchParams } from "next/navigation";
import { ProgramsPagination } from "@/components/websac3/program/ProgramsPagination";
import { DegreeProgramCard, ActionButton } from "@/components/websac3/program/DegreeProgramCard";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { useDispatch } from "react-redux";
import { showError } from "@/store/errorSlice";

export default function OperarModeloPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const dispatch = useDispatch();
  const [selectedProgramId, setSelectedProgramId] = useState<number | null>(null);
  const [selectedRoleId, setSelectedRoleId] = useState<number | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isResultModalOpen, setIsResultModalOpen] = useState(false);
  const [evaluationResult, setEvaluationResult] = useState<{ success: boolean; message: string } | null>(null);
  
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
  
  // Fetch professional roles
  const { data: rolesData, isLoading: rolesLoading } = useListProfessionalRolesQuery();
  
  // Evaluation mutation
  const [evaluateProgram, { isLoading: isEvaluating }] = useEvaluateDegreeProgramMutation();

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
      ? `/director/operar-modelo?${params.toString()}`
      : '/director/operar-modelo';
    
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

  const handleEvaluateProgram = (programId: number) => {
    setSelectedProgramId(programId);
    setIsModalOpen(true);
  };

  const handleRoleSelection = async (roleId: number) => {
    if (!selectedProgramId) return;
    
    setSelectedRoleId(roleId);
    
    try {
      const result = await evaluateProgram({
        degree_program_id: selectedProgramId,
        professional_role_id: roleId
      }).unwrap();
      
      // Success - close modal and show success message
      setIsModalOpen(false);
      setSelectedProgramId(null);
      setSelectedRoleId(null);
      
      // Show success result modal
      setEvaluationResult({
        success: true,
        message: result.result || 'El reporte de evaluación ha sido enviado a su correo exitosamente.'
      });
      setIsResultModalOpen(true);
      
    } catch (error: any) {
      console.error('Evaluation failed:', error);
      
      // Close role selection modal
      setIsModalOpen(false);
      setSelectedProgramId(null);
      setSelectedRoleId(null);
      
      // Show error result modal
      const errorMessage = error?.data?.errors?.[0] || 
                          error?.data?.result || 
                          error?.message || 
                          'Error al evaluar el programa de grado. Por favor, inténtalo de nuevo.';
      
      setEvaluationResult({
        success: false,
        message: errorMessage
      });
      setIsResultModalOpen(true);
    }
  };

  if (isLoading) {
    return <LoadingSkeleton />;
  }

  if (error && error.status !== 404) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
          <div className="text-center text-gray-500">
            <Shield className="h-12 w-12 text-gray-400 mx-auto mb-4" />
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
          <h1 className="text-3xl font-bold text-gray-900">Evaluar Componente de Ciberseguridad</h1>
          <p className="text-gray-600 mt-2">Selecciona un programa de grado para evaluar su componente de ciberseguridad</p>
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
      {data && !(error && error.status === 404) && (
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
            {data.data && data.data.length > 0 && (
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
        ) : data && data.data.length > 0 && !(error && error.status === 404) ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {data.data.map((program) => {
              const actions: ActionButton[] = [
                {
                  label: "Evaluar Componente de Ciberseguridad",
                  icon: Shield,
                  variant: "default",
                  className: "w-full bg-gradient-to-r from-purple-600 to-purple-700 hover:from-purple-700 hover:to-purple-800 shadow-lg hover:shadow-xl transition-all duration-200",
                  onClick: () => handleEvaluateProgram(program.id)
                }
              ];

              return (
                <DegreeProgramCard
                  key={program.id}
                  program={program}
                  actions={actions}
                />
              );
            })}
          </div>
        ) : (
          <Card>
            <CardContent className="p-12">
              <div className="text-center">
                <Shield className="h-12 w-12 text-gray-400 mx-auto mb-4" />
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
        {!isLoading && !(error && error.status === 404) && (
          <ProgramsPagination
            currentPage={currentPage}
            totalPages={data ? Math.ceil(data.total_count / itemsPerPage) : 1}
            onPageChange={handlePageChange}
          />
        )}
      </div>

      {/* Professional Role Selection Modal */}
      <Dialog open={isModalOpen} onOpenChange={setIsModalOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <Shield className="h-5 w-5 text-purple-600" />
              Seleccionar Rol Profesional
            </DialogTitle>
            <DialogDescription>
              Selecciona el rol profesional con el cual quieres evaluar el componente de ciberseguridad del programa de grado.
            </DialogDescription>
          </DialogHeader>
          
          <div className="space-y-4">
            {rolesLoading ? (
              <div className="space-y-3">
                {Array.from({ length: 3 }).map((_, index) => (
                  <div key={index} className="flex items-center space-x-3 p-3 border rounded-lg">
                    <Skeleton className="h-4 w-4 rounded" />
                    <Skeleton className="h-4 flex-1" />
                  </div>
                ))}
              </div>
            ) : rolesData && rolesData.data.length > 0 ? (
              <div className="space-y-2 max-h-60 overflow-y-auto">
                {rolesData.data.map((role) => (
                  <button
                    key={role.id}
                    onClick={() => handleRoleSelection(role.id)}
                    disabled={isEvaluating}
                    className={`w-full flex items-center space-x-3 p-3 border rounded-lg text-left transition-all duration-200 ${
                      selectedRoleId === role.id
                        ? 'border-purple-500 bg-purple-50 text-purple-700'
                        : 'border-gray-200 hover:border-purple-300 hover:bg-purple-50'
                    } ${isEvaluating ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer'}`}
                  >
                    <div className={`w-4 h-4 rounded-full border-2 flex items-center justify-center ${
                      selectedRoleId === role.id
                        ? 'border-purple-500 bg-purple-500'
                        : 'border-gray-300'
                    }`}>
                      {selectedRoleId === role.id && (
                        <CheckCircle className="h-3 w-3 text-white" />
                      )}
                    </div>
                    <span className="font-medium">{role.name}</span>
                  </button>
                ))}
              </div>
            ) : (
              <div className="text-center py-8">
                <AlertCircle className="h-8 w-8 text-gray-400 mx-auto mb-2" />
                <p className="text-gray-500">No hay roles profesionales disponibles</p>
              </div>
            )}
          </div>
          
          <div className="flex justify-end space-x-2 pt-4 border-t">
            <Button
              variant="outline"
              onClick={() => setIsModalOpen(false)}
              disabled={isEvaluating}
            >
              Cancelar
            </Button>
            {selectedRoleId && (
              <Button
                onClick={() => handleRoleSelection(selectedRoleId)}
                disabled={isEvaluating}
                className="bg-purple-600 hover:bg-purple-700"
              >
                {isEvaluating ? (
                  <>
                    <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                    Evaluando...
                  </>
                ) : (
                  <>
                    <Shield className="h-4 w-4 mr-2" />
                    Evaluar
                  </>
                )}
              </Button>
            )}
          </div>
        </DialogContent>
      </Dialog>

      {/* Evaluation Result Modal */}
      <Dialog open={isResultModalOpen} onOpenChange={setIsResultModalOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              {evaluationResult?.success ? (
                <>
                  <CheckCircle className="h-5 w-5 text-green-600" />
                  Evaluación Completada
                </>
              ) : (
                <>
                  <AlertCircle className="h-5 w-5 text-red-600" />
                  Error en la Evaluación
                </>
              )}
            </DialogTitle>
            <DialogDescription>
              {evaluationResult?.success 
                ? "La evaluación del componente de ciberseguridad se ha completado exitosamente."
                : "Ha ocurrido un error durante la evaluación del programa de grado."
              }
            </DialogDescription>
          </DialogHeader>
          
          <div className="space-y-4">
            <div className={`p-4 rounded-lg border ${
              evaluationResult?.success 
                ? 'bg-green-50 border-green-200' 
                : 'bg-red-50 border-red-200'
            }`}>
              <div className="flex items-start gap-3">
                <div className={`p-2 rounded-lg ${
                  evaluationResult?.success 
                    ? 'bg-green-100' 
                    : 'bg-red-100'
                }`}>
                  {evaluationResult?.success ? (
                    <Mail className="h-5 w-5 text-green-600" />
                  ) : (
                    <AlertCircle className="h-5 w-5 text-red-600" />
                  )}
                </div>
                <div className="flex-1">
                  <p className={`text-sm mt-1 ${
                    evaluationResult?.success 
                      ? 'text-green-700' 
                      : 'text-red-700'
                  }`}>
                    {evaluationResult?.message}
                  </p>
                </div>
              </div>
            </div>
          </div>
          
          <div className="flex justify-end pt-4 border-t">
            <Button
              onClick={() => setIsResultModalOpen(false)}
              className={`${
                evaluationResult?.success 
                  ? 'bg-green-600 hover:bg-green-700' 
                  : 'bg-red-600 hover:bg-red-700'
              }`}
            >
              {evaluationResult?.success ? 'Entendido' : 'Cerrar'}
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
