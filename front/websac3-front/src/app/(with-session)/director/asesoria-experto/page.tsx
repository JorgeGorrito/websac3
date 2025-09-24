"use client";

import { useState, useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { 
  MessageSquare, 
  BookOpen, 
  Clock, 
  GraduationCap, 
  Target, 
  Search, 
  Settings, 
  Calendar,
  Plus,
  User,
  FileText
} from "lucide-react";
import { useListDegreeProgramsQuery, useListDegreeProgramReportsQuery, useCreateExpertConsultationMutation } from "@/services/api";
import { ProgramsPagination } from "@/components/websac3/program/ProgramsPagination";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { useDispatch } from "react-redux";
import { showError } from "@/store/errorSlice";

export default function AsesoriaExpertoPage() {
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
  const [selectedProgramId, setSelectedProgramId] = useState<number | null>(null);
  const [selectedProgramName, setSelectedProgramName] = useState<string>('');
  const [isReportsModalOpen, setIsReportsModalOpen] = useState(false);
  const [isConsultationModalOpen, setIsConsultationModalOpen] = useState(false);
  const [consultationMessage, setConsultationMessage] = useState('');
  const [selectedReportId, setSelectedReportId] = useState<number | null>(null);
  const [reportsCurrentPage, setReportsCurrentPage] = useState(1);
  const [reportsItemsPerPage, setReportsItemsPerPage] = useState(10);

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
  const [createConsultation, { isLoading: isCreatingConsultation }] = useCreateExpertConsultationMutation();

  // Fetch reports for selected program
  const { data: reportsData, isLoading: reportsLoading, error: reportsError } = useListDegreeProgramReportsQuery(
    { 
      degree_program_id: selectedProgramId!,
      current_page: reportsCurrentPage,
      items_per_page: reportsItemsPerPage
    },
    { skip: !selectedProgramId }
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
    
    if (page !== 1) params.set('page', page.toString());
    if (perPage !== 6) params.set('per_page', perPage.toString());
    if (name.trim()) params.set('name', name.trim());
    if (snies.trim()) params.set('snies', snies.trim());
    
    const newURL = params.toString() 
      ? `/director/asesoria-experto?${params.toString()}`
      : '/director/asesoria-experto';
    
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
    setCurrentPage(1);
    updateURL({ per_page: perPage, page: 1 });
  };

  // Handle filter changes
  const handleFilterChange = (key: keyof typeof filters, value: string) => {
    setFilters(prev => ({ ...prev, [key]: value }));
    setCurrentPage(1);
    updateURL({ [key]: value, page: 1 });
  };

  // Clear all filters
  const clearFilters = () => {
    setFilters({ name: '', snies: '' });
    setCurrentPage(1);
    updateURL({ name: '', snies: '', page: 1 });
  };

  const handleProgramClick = (programId: number, programName: string) => {
    setSelectedProgramId(programId);
    setSelectedProgramName(programName);
    setReportsCurrentPage(1); // Reset reports pagination
    setIsReportsModalOpen(true);
  };

  const handleReportClick = (reportId: number) => {
    setSelectedReportId(reportId);
    setIsReportsModalOpen(false);
    setIsConsultationModalOpen(true);
  };

  // Handle reports pagination
  const handleReportsPageChange = (page: number) => {
    setReportsCurrentPage(page);
  };

  const handleReportsItemsPerPageChange = (itemsPerPage: number) => {
    setReportsItemsPerPage(itemsPerPage);
    setReportsCurrentPage(1); // Reset to first page when changing items per page
  };

  const handleCreateConsultation = async () => {
    if (!selectedProgramId || !selectedReportId || !consultationMessage.trim()) {
      dispatch(showError('Por favor completa todos los campos requeridos.'));
      return;
    }

    try {
      await createConsultation({
        degree_program_id: selectedProgramId,
        report_id: selectedReportId,
        request_message: consultationMessage.trim(),
        requester_id: 1 // TODO: Get from auth context
      }).unwrap();

      dispatch(showError('Solicitud de asesoría enviada exitosamente.'));
      setIsConsultationModalOpen(false);
      setConsultationMessage('');
      setSelectedReportId(null);
    } catch (error: any) {
      const errorMessage = error?.data?.errors?.[0] || 'Error al enviar la solicitud de asesoría';
      dispatch(showError(errorMessage));
    }
  };

  const getScoreColor = (score: number) => {
    const percentage = score * 100;
    if (percentage >= 80) return 'text-green-600';
    if (percentage >= 60) return 'text-yellow-600';
    return 'text-red-600';
  };

  const getScoreLabel = (score: number) => {
    const percentage = score * 100;
    if (percentage >= 80) return 'Excelente';
    if (percentage >= 60) return 'Aceptable';
    return 'Por mejorar';
  };

  const formatScore = (score: number) => {
    return `${(score * 100).toFixed(1)}%`;
  };

  // Loading skeleton component
  const LoadingSkeleton = () => (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <Skeleton className="h-8 w-64" />
          <Skeleton className="h-4 w-96 mt-2" />
        </div>
        <Skeleton className="h-10 w-32" />
      </div>
      
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
    </div>
  );

  if (isLoading) {
    return <LoadingSkeleton />;
  }

  if (error && error.status !== 404) {
    return (
      <div className="space-y-6">
        <div className="flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Asesoría con Experto</h1>
            <p className="text-gray-600 mt-2">Solicita asesoría especializada en ciberseguridad</p>
          </div>
        </div>
        
        <Card>
          <CardContent className="p-6">
            <div className="text-center text-gray-500">
              <MessageSquare className="h-12 w-12 text-gray-400 mx-auto mb-4" />
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
          <h1 className="text-3xl font-bold text-gray-900">Asesoría con Experto</h1>
          <p className="text-gray-600 mt-2">Solicita asesoría especializada en ciberseguridad para tu programa de grado</p>
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
              {/* Program Name Filter */}
              <div className="space-y-2">
                <Label htmlFor="name" className="text-sm font-medium text-gray-700">
                  Nombre del Programa
                </Label>
                <Input
                  id="name"
                  type="text"
                  placeholder="Buscar por nombre..."
                  value={filters.name}
                  onChange={(e) => handleFilterChange('name', e.target.value)}
                  className="w-full"
                />
              </div>

              {/* SNIES Filter */}
              <div className="space-y-2">
                <Label htmlFor="snies" className="text-sm font-medium text-gray-700">
                  Código SNIES
                </Label>
                <Input
                  id="snies"
                  type="text"
                  placeholder="Buscar por SNIES..."
                  value={filters.snies}
                  onChange={(e) => handleFilterChange('snies', e.target.value)}
                  className="w-full"
                />
              </div>

              {/* Items per page */}
              <div className="space-y-2">
                <Label htmlFor="itemsPerPage" className="text-sm font-medium text-gray-700">
                  Elementos por página
                </Label>
                <select
                  id="itemsPerPage"
                  value={itemsPerPage}
                  onChange={(e) => handleItemsPerPageChange(parseInt(e.target.value))}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                >
                  <option value={6}>6 por página</option>
                  <option value={12}>12 por página</option>
                  <option value={24}>24 por página</option>
                  <option value={48}>48 por página</option>
                </select>
              </div>
            </div>
            
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
        ) : data && data.data.length > 0 && !(error && error.status === 404) ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {data.data.map((program) => (
              <Card key={program.id} className="bg-white border border-gray-200 shadow-lg hover:shadow-xl transition-all duration-300 hover:-translate-y-1 py-0">
                <CardHeader className="pb-3 pt-4 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-gray-100">
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <CardTitle className="text-lg font-bold text-gray-900 mb-2 line-clamp-2">
                        {program.name}
                      </CardTitle>
                      <div className="flex items-center gap-2 mb-2">
                        <Badge className="bg-blue-100 text-blue-800 border-blue-200 text-xs">
                          SNIES: {program.snies}
                        </Badge>
                        <Badge className="bg-green-100 text-green-800 border-green-200 text-xs">
                          {program.duration_unit.name}
                        </Badge>
                      </div>
                    </div>
                  </div>
                </CardHeader>
                <CardContent className="p-4">
                  <div className="grid grid-cols-2 gap-4 mb-4">
                    <div className="flex items-center gap-2">
                      <GraduationCap className="h-4 w-4 text-blue-600" />
                      <div>
                        <div className="text-xs text-gray-500">CRÉDITOS</div>
                        <div className="text-sm font-semibold text-gray-900">{program.total_credits}</div>
                      </div>
                    </div>
                    <div className="flex items-center gap-2">
                      <Clock className="h-4 w-4 text-green-600" />
                      <div>
                        <div className="text-xs text-gray-500">DURACIÓN</div>
                        <div className="text-sm font-semibold text-gray-900">
                          {program.duration_value} {program.duration_unit.name.toLowerCase()}
                        </div>
                      </div>
                    </div>
                  </div>
                  
                  <div className="space-y-2 mb-4">
                    <div className="flex items-start gap-2">
                      <Target className="h-4 w-4 text-purple-600 mt-0.5" />
                      <div>
                        <div className="text-xs text-gray-500">ENFOQUE</div>
                        <div className="text-sm text-gray-700 line-clamp-2">{program.program_focus}</div>
                      </div>
                    </div>
                    <div className="flex items-start gap-2">
                      <BookOpen className="h-4 w-4 text-orange-600 mt-0.5" />
                      <div>
                        <div className="text-xs text-gray-500">PERFIL DE INGRESO</div>
                        <div className="text-sm text-gray-700 line-clamp-2">{program.entry_profile}</div>
                      </div>
                    </div>
                  </div>

                  <div className="flex items-center gap-2 mb-4">
                    <GraduationCap className="h-4 w-4 text-indigo-600" />
                    <div>
                      <div className="text-xs text-gray-500">INSTITUCIÓN</div>
                      <div className="text-sm font-medium text-gray-900">{program.higher_education_institution.name}</div>
                    </div>
                  </div>

                  <Button
                    onClick={() => handleProgramClick(program.id, program.name)}
                    className="w-full bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 text-white font-medium py-2 px-4 rounded-lg transition-all duration-200 flex items-center justify-center gap-2"
                  >
                    <FileText className="h-4 w-4" />
                    Ver Reportes
                  </Button>
                </CardContent>
              </Card>
            ))}
          </div>
        ) : (
          <Card>
            <CardContent className="p-12">
              <div className="text-center">
                <MessageSquare className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                <h3 className="text-lg font-medium text-gray-900 mb-2">
                  No hay programas disponibles
                </h3>
                <p className="text-gray-600 mb-6">
                  No se encontraron programas de grado para solicitar asesoría
                </p>
                <Button onClick={clearFilters} className="bg-blue-600 hover:bg-blue-700">
                  <Search className="h-4 w-4 mr-2" />
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

      {/* Reports Modal */}
      <Dialog open={isReportsModalOpen} onOpenChange={setIsReportsModalOpen}>
        <DialogContent className="sm:max-w-4xl max-h-[80vh] overflow-hidden flex flex-col">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <FileText className="h-5 w-5 text-blue-600" />
              Reportes de {selectedProgramName}
            </DialogTitle>
            <DialogDescription>
              Selecciona un reporte para solicitar asesoría especializada
            </DialogDescription>
          </DialogHeader>
          
          <div className="flex-1 overflow-y-auto">
            {reportsLoading ? (
              <div className="space-y-4">
                {Array.from({ length: 3 }).map((_, index) => (
                  <Card key={index}>
                    <CardContent className="p-6">
                      <div className="flex items-center justify-between">
                        <div className="flex-1">
                          <Skeleton className="h-6 w-32 mb-2" />
                          <Skeleton className="h-4 w-48" />
                        </div>
                        <Skeleton className="h-8 w-24" />
                      </div>
                    </CardContent>
                  </Card>
                ))}
              </div>
            ) : reportsError ? (
              <div className="flex items-center justify-center py-12">
                <div className="text-center text-gray-500">
                  <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <p>Error al cargar los reportes</p>
                  <p className="text-sm mt-2">Inténtalo de nuevo más tarde</p>
                </div>
              </div>
            ) : reportsData?.data && reportsData.data.length > 0 ? (
              <div className="space-y-4">
                {reportsData.data.map((report) => (
                  <Card key={report.id} className="hover:shadow-md transition-shadow">
                    <CardContent className="p-6">
                      <div className="flex items-center justify-between">
                        <div className="flex-1">
                          <div className="flex items-center gap-3 mb-2">
                            <h3 className="text-lg font-semibold text-gray-900">
                              Reporte #{report.id}
                            </h3>
                            <Badge 
                              variant="outline" 
                              className={getScoreColor(report.score)}
                            >
                              {getScoreLabel(report.score)}
                            </Badge>
                          </div>
                          <div className="space-y-1 text-sm text-gray-600">
                            <p><strong>Rol profesional:</strong> {report.professional_role?.name || 'N/A'}</p>
                            <p><strong>Puntaje:</strong> {formatScore(report.score)}</p>
                            <p><strong>Fecha:</strong> {new Date(report.created_at).toLocaleDateString('es-ES')}</p>
                          </div>
                        </div>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => handleReportClick(report.id)}
                          className="flex items-center gap-2"
                        >
                          <MessageSquare className="h-4 w-4" />
                          Solicitar Asesoría
                        </Button>
                      </div>
                    </CardContent>
                  </Card>
                ))}
              </div>
            ) : (
              <div className="flex items-center justify-center py-12">
                <div className="text-center text-gray-500">
                  <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <p>No hay reportes disponibles</p>
                  <p className="text-sm mt-2">Este programa no tiene reportes de evaluación</p>
                </div>
              </div>
            )}

            {/* Reports Pagination */}
            {reportsData && (
              <div className="mt-6">
                <ProgramsPagination
                  currentPage={reportsCurrentPage}
                  totalCount={reportsData.total_count || 0}
                  itemsPerPage={reportsItemsPerPage}
                  onPageChange={handleReportsPageChange}
                  onItemsPerPageChange={handleReportsItemsPerPageChange}
                />
              </div>
            )}
          </div>
        </DialogContent>
      </Dialog>

      {/* Consultation Request Modal */}
      <Dialog open={isConsultationModalOpen} onOpenChange={setIsConsultationModalOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <MessageSquare className="h-5 w-5 text-blue-600" />
              Solicitar Asesoría
            </DialogTitle>
            <DialogDescription>
              Envía una solicitud de asesoría especializada para: <strong>{selectedProgramName}</strong>
              {selectedReportId && (
                <span className="block mt-1">
                  Reporte seleccionado: <strong>#{selectedReportId}</strong>
                </span>
              )}
            </DialogDescription>
          </DialogHeader>
          
          <div className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="consultation-message" className="text-sm font-medium text-gray-700">
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
            
            <div className="bg-blue-50 border border-blue-200 rounded-lg p-3">
              <div className="flex items-start gap-2">
                <User className="h-4 w-4 text-blue-600 mt-0.5" />
                <div className="text-sm text-blue-800">
                  <p className="font-medium">Información adicional:</p>
                  <p className="text-xs mt-1">
                    • Tu solicitud será revisada por un experto en ciberseguridad<br/>
                    • Se te notificará al correo la respuesta del experto a tu solicitud
                  </p>
                </div>
              </div>
            </div>
          </div>

          <div className="flex gap-3 justify-end">
            <Button
              variant="outline"
              onClick={() => {
                setIsConsultationModalOpen(false);
                setConsultationMessage('');
              }}
            >
              Cancelar
            </Button>
            <Button
              onClick={handleCreateConsultation}
              disabled={isCreatingConsultation || !consultationMessage.trim()}
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
