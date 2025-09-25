"use client";

import { useState, useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { 
  Users, 
  Search, 
  Settings, 
  Calendar,
  Building2,
  MapPin,
  User,
  Mail,
  IdCard,
  Briefcase,
  CheckCircle,
  XCircle,
  MoreHorizontal
} from "lucide-react";
import { 
  useListAccessRequestsQuery, 
  useListRolesQuery,
  useApproveAccessRequestMutation,
  useRejectAccessRequestMutation
} from "@/services/api";
import { ProgramsPagination } from "@/components/websac3/program/ProgramsPagination";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { useDispatch } from "react-redux";
import { showError, showSuccess } from "@/store/errorSlice";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";

export default function GestionarAccesosPage() {
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
    return perPage ? parseInt(perPage, 10) : 10;
  });
  
  const [filters, setFilters] = useState(() => {
    const name = searchParams.get('name') || '';
    const lastname = searchParams.get('lastname') || '';
    const institution = searchParams.get('institution') || '';
    const municipality = searchParams.get('municipality') || '';
    const department = searchParams.get('department') || '';
    return { name, lastname, institution, municipality, department };
  });
  
  const [showFilters, setShowFilters] = useState(false);
  const [selectedRequestId, setSelectedRequestId] = useState<number | null>(null);
  const [selectedRequestName, setSelectedRequestName] = useState<string>('');
  const [isActionModalOpen, setIsActionModalOpen] = useState(false);
  const [actionType, setActionType] = useState<'approve' | 'reject' | null>(null);
  const [selectedRoleId, setSelectedRoleId] = useState<string>('');

  // Build query parameters with proper filter format
  const buildFilters = () => {
    const filterParams: Record<string, string> = {};
    
    if (filters.name.trim()) {
      filterParams['Applicant.name[cont]'] = filters.name.trim();
    }
    
    if (filters.lastname.trim()) {
      filterParams['Applicant.lastname[cont]'] = filters.lastname.trim();
    }
    
    if (filters.institution.trim()) {
      filterParams['Applicant.HigherEducationInstitution.name[cont]'] = filters.institution.trim();
    }
    
    if (filters.municipality.trim()) {
      filterParams['Applicant.HigherEducationInstitution.Municipality.name[cont]'] = filters.municipality.trim();
    }
    
    if (filters.department.trim()) {
      filterParams['Applicant.HigherEducationInstitution.Department.name[cont]'] = filters.department.trim();
    }
    
    return filterParams;
  };

  const queryParams = {
    current_page: currentPage,
    items_per_page: itemsPerPage,
    filters: buildFilters()
  };

  const { data, isLoading, error, refetch } = useListAccessRequestsQuery(queryParams);
  const { data: rolesData, isLoading: rolesLoading } = useListRolesQuery({});
  const [approveRequest, { isLoading: isApproving }] = useApproveAccessRequestMutation();
  const [rejectRequest, { isLoading: isRejecting }] = useRejectAccessRequestMutation();

  // Function to update URL with current parameters
  const updateURL = (newParams: {
    page?: number;
    per_page?: number;
    name?: string;
    lastname?: string;
    institution?: string;
    municipality?: string;
    department?: string;
  }) => {
    const params = new URLSearchParams();
    const page = newParams.page ?? currentPage;
    const perPage = newParams.per_page ?? itemsPerPage;
    const name = newParams.name ?? filters.name;
    const lastname = newParams.lastname ?? filters.lastname;
    const institution = newParams.institution ?? filters.institution;
    const municipality = newParams.municipality ?? filters.municipality;
    const department = newParams.department ?? filters.department;
    
    if (page !== 1) params.set('page', page.toString());
    if (perPage !== 10) params.set('per_page', perPage.toString());
    if (name.trim()) params.set('name', name.trim());
    if (lastname.trim()) params.set('lastname', lastname.trim());
    if (institution.trim()) params.set('institution', institution.trim());
    if (municipality.trim()) params.set('municipality', municipality.trim());
    if (department.trim()) params.set('department', department.trim());
    
    const newURL = params.toString() 
      ? `/admin/accesos?${params.toString()}`
      : '/admin/accesos';
    
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
    setFilters({ name: '', lastname: '', institution: '', municipality: '', department: '' });
    setCurrentPage(1);
    updateURL({ name: '', lastname: '', institution: '', municipality: '', department: '', page: 1 });
  };

  const handleActionClick = (requestId: number, requestName: string, action: 'approve' | 'reject') => {
    setSelectedRequestId(requestId);
    setSelectedRequestName(requestName);
    setActionType(action);
    setSelectedRoleId('');
    setIsActionModalOpen(true);
  };

  const handleConfirmAction = async () => {
    if (!selectedRequestId || !actionType) return;

    if (actionType === 'approve' && !selectedRoleId) {
      dispatch(showError('Por favor selecciona un rol para el usuario.'));
      return;
    }

    try {
      if (actionType === 'approve') {
        await approveRequest({
          id: selectedRequestId,
          role_id: parseInt(selectedRoleId)
        }).unwrap();
        dispatch(showSuccess({ 
          message: 'Solicitud aprobada exitosamente.',
          title: 'Solicitud Aprobada'
        }));
      } else {
        await rejectRequest({ id: selectedRequestId }).unwrap();
        dispatch(showSuccess({ 
          message: 'Solicitud rechazada exitosamente.',
          title: 'Solicitud Rechazada'
        }));
      }
      
      setIsActionModalOpen(false);
      setSelectedRequestId(null);
      setSelectedRequestName('');
      setActionType(null);
      setSelectedRoleId('');
      refetch(); // Refresh the list
    } catch (error: any) {
      const errorMessage = error?.data?.errors?.[0] || `Error al ${actionType === 'approve' ? 'aprobar' : 'rechazar'} la solicitud`;
      dispatch(showError(errorMessage));
    }
  };

  const getStatusBadge = (statusName: string) => {
    switch (statusName.toLowerCase()) {
      case 'pending':
      case 'pendiente':
        return <Badge variant="outline" className="bg-yellow-50 text-yellow-700 border-yellow-200">Pendiente</Badge>;
      case 'approved':
      case 'aprobado':
        return <Badge variant="outline" className="bg-green-50 text-green-700 border-green-200">Aprobado</Badge>;
      case 'rejected':
      case 'rechazado':
        return <Badge variant="outline" className="bg-red-50 text-red-700 border-red-200">Rechazado</Badge>;
      default:
        return <Badge variant="outline">{statusName}</Badge>;
    }
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
      
      <div className="grid grid-cols-1 gap-6">
        {Array.from({ length: 5 }).map((_, index) => (
          <Card key={index}>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div className="flex-1">
                  <Skeleton className="h-6 w-48 mb-2" />
                  <Skeleton className="h-4 w-64" />
                </div>
                <Skeleton className="h-8 w-24" />
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
            <h1 className="text-3xl font-bold text-gray-900">Gestionar Solicitudes de Acceso</h1>
            <p className="text-gray-600 mt-2">Administra las solicitudes de acceso al sistema</p>
          </div>
        </div>
        
        <Card>
          <CardContent className="p-6">
            <div className="text-center text-gray-500">
              <Users className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <p>No se pudieron cargar las solicitudes de acceso</p>
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
          <h1 className="text-3xl font-bold text-gray-900">Gestionar Solicitudes de Acceso</h1>
          <p className="text-gray-600 mt-2">Administra las solicitudes de acceso al sistema</p>
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
                <Users className="h-4 w-4 text-blue-600" />
                <span className="text-sm font-medium text-gray-700">
                  {data.total_count} solicitud{data.total_count !== 1 ? 'es' : ''} disponible{data.total_count !== 1 ? 's' : ''}
                </span>
              </div>
              
              <div className="flex items-center gap-2">
                <Calendar className="h-4 w-4 text-gray-500" />
                <span className="text-sm text-gray-600">
                  Página {data.current_page} de {Math.ceil(data.total_count / itemsPerPage)}
                </span>
              </div>
            </div>

            {/* Right side - Status summary */}
            {data.data && data.data.length > 0 && (
              <div className="flex items-center gap-3">
                <div className="flex items-center gap-2">
                  <CheckCircle className="h-4 w-4 text-green-600" />
                  <span className="text-sm text-gray-600">
                    {data.data.filter(item => item.status_name.toLowerCase().includes('aprobado') || item.status_name.toLowerCase().includes('approved')).length} aprobadas
                  </span>
                </div>
                <div className="flex items-center gap-2">
                  <XCircle className="h-4 w-4 text-red-600" />
                  <span className="text-sm text-gray-600">
                    {data.data.filter(item => item.status_name.toLowerCase().includes('rechazado') || item.status_name.toLowerCase().includes('rejected')).length} rechazadas
                  </span>
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
              Personaliza tu búsqueda y visualización de solicitudes
            </CardDescription>
          </CardHeader>
          <CardContent className="p-6">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {/* Name Filter */}
              <div className="space-y-2">
                <Label htmlFor="name" className="text-sm font-medium text-gray-700">
                  Nombre
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

              {/* Lastname Filter */}
              <div className="space-y-2">
                <Label htmlFor="lastname" className="text-sm font-medium text-gray-700">
                  Apellido
                </Label>
                <Input
                  id="lastname"
                  type="text"
                  placeholder="Buscar por apellido..."
                  value={filters.lastname}
                  onChange={(e) => handleFilterChange('lastname', e.target.value)}
                  className="w-full"
                />
              </div>

              {/* Institution Filter */}
              <div className="space-y-2">
                <Label htmlFor="institution" className="text-sm font-medium text-gray-700">
                  Institución
                </Label>
                <Input
                  id="institution"
                  type="text"
                  placeholder="Buscar por institución..."
                  value={filters.institution}
                  onChange={(e) => handleFilterChange('institution', e.target.value)}
                  className="w-full"
                />
              </div>

              {/* Municipality Filter */}
              <div className="space-y-2">
                <Label htmlFor="municipality" className="text-sm font-medium text-gray-700">
                  Municipio
                </Label>
                <Input
                  id="municipality"
                  type="text"
                  placeholder="Buscar por municipio..."
                  value={filters.municipality}
                  onChange={(e) => handleFilterChange('municipality', e.target.value)}
                  className="w-full"
                />
              </div>

              {/* Department Filter */}
              <div className="space-y-2">
                <Label htmlFor="department" className="text-sm font-medium text-gray-700">
                  Departamento
                </Label>
                <Input
                  id="department"
                  type="text"
                  placeholder="Buscar por departamento..."
                  value={filters.department}
                  onChange={(e) => handleFilterChange('department', e.target.value)}
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
                  <option value={10}>10 por página</option>
                  <option value={20}>20 por página</option>
                  <option value={50}>50 por página</option>
                  <option value={100}>100 por página</option>
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

      {/* Requests List */}
      <div className="space-y-4">
        {isLoading ? (
          <div className="grid grid-cols-1 gap-6">
            {Array.from({ length: 5 }).map((_, index) => (
              <Card key={index}>
                <CardContent className="p-6">
                  <div className="flex items-center justify-between">
                    <div className="flex-1">
                      <Skeleton className="h-6 w-48 mb-2" />
                      <Skeleton className="h-4 w-64" />
                    </div>
                    <Skeleton className="h-8 w-24" />
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        ) : data && data.data.length > 0 && !(error && error.status === 404) ? (
          <div className="grid grid-cols-1 gap-6">
            {data.data.map((request) => (
              <Card key={request.id} className="bg-white border border-gray-200 shadow-lg hover:shadow-xl transition-all duration-300">
                <CardContent className="p-6">
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <div className="flex items-center gap-3 mb-3">
                        <h3 className="text-lg font-semibold text-gray-900">
                          {request.name} {request.lastname}
                        </h3>
                        {getStatusBadge(request.status_name)}
                      </div>
                      
                      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                        <div className="space-y-2">
                          <div className="flex items-center gap-2">
                            <Mail className="h-4 w-4 text-gray-500" />
                            <span className="text-sm text-gray-600">{request.email}</span>
                          </div>
                          <div className="flex items-center gap-2">
                            <IdCard className="h-4 w-4 text-gray-500" />
                            <span className="text-sm text-gray-600">
                              {request.identification_type}: {request.identification_number}
                            </span>
                          </div>
                          <div className="flex items-center gap-2">
                            <Briefcase className="h-4 w-4 text-gray-500" />
                            <span className="text-sm text-gray-600">{request.job_position}</span>
                          </div>
                        </div>
                        
                        <div className="space-y-2">
                          <div className="flex items-center gap-2">
                            <Building2 className="h-4 w-4 text-gray-500" />
                            <span className="text-sm text-gray-600">{request.higher_education_institution_name}</span>
                          </div>
                          <div className="flex items-center gap-2">
                            <MapPin className="h-4 w-4 text-gray-500" />
                            <span className="text-sm text-gray-600">
                              {request.municipality_name}, {request.department_name}
                            </span>
                          </div>
                          <div className="flex items-center gap-2">
                            <Building2 className="h-4 w-4 text-gray-500" />
                            <span className="text-sm text-gray-600">
                              SNIES: {request.higher_education_institution_snies} • {request.higher_education_institution_ownership}
                            </span>
                          </div>
                        </div>
                      </div>
                    </div>
                    
                    <div className="flex items-center gap-2">
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button variant="outline" size="sm">
                            <MoreHorizontal className="h-4 w-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end">
                          <DropdownMenuItem
                            onClick={() => handleActionClick(request.id, `${request.name} ${request.lastname}`, 'approve')}
                            className="text-green-600"
                          >
                            <CheckCircle className="h-4 w-4 mr-2" />
                            Aprobar
                          </DropdownMenuItem>
                          <DropdownMenuItem
                            onClick={() => handleActionClick(request.id, `${request.name} ${request.lastname}`, 'reject')}
                            className="text-red-600"
                          >
                            <XCircle className="h-4 w-4 mr-2" />
                            Rechazar
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </div>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        ) : (
          <Card>
            <CardContent className="p-12">
              <div className="text-center">
                <Users className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                <h3 className="text-lg font-medium text-gray-900 mb-2">
                  No hay solicitudes disponibles
                </h3>
                <p className="text-gray-600 mb-6">
                  No se encontraron solicitudes de acceso
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

      {/* Action Modal */}
      <Dialog open={isActionModalOpen} onOpenChange={setIsActionModalOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              {actionType === 'approve' ? (
                <CheckCircle className="h-5 w-5 text-green-600" />
              ) : (
                <XCircle className="h-5 w-5 text-red-600" />
              )}
              {actionType === 'approve' ? 'Aprobar Solicitud' : 'Rechazar Solicitud'}
            </DialogTitle>
            <DialogDescription>
              {actionType === 'approve' 
                ? `¿Estás seguro de que deseas aprobar la solicitud de ${selectedRequestName}?`
                : `¿Estás seguro de que deseas rechazar la solicitud de ${selectedRequestName}?`
              }
            </DialogDescription>
          </DialogHeader>
          
          <div className="space-y-4">
            {actionType === 'approve' && (
              <div className="space-y-2">
                <Label htmlFor="role-select" className="text-sm font-medium text-gray-700">
                  Rol del usuario *
                </Label>
                <Select value={selectedRoleId} onValueChange={setSelectedRoleId}>
                  <SelectTrigger>
                    <SelectValue placeholder="Selecciona un rol" />
                  </SelectTrigger>
                  <SelectContent>
                    {rolesLoading ? (
                      <SelectItem value="loading" disabled>
                        Cargando roles...
                      </SelectItem>
                    ) : rolesData?.data ? (
                      rolesData.data.map((role) => (
                        <SelectItem key={role.id} value={role.id.toString()}>
                          {role.name}
                        </SelectItem>
                      ))
                    ) : (
                      <SelectItem value="no-roles" disabled>
                        No hay roles disponibles
                      </SelectItem>
                    )}
                  </SelectContent>
                </Select>
              </div>
            )}
            
            <div className={`border rounded-lg p-3 ${
              actionType === 'approve' 
                ? 'bg-green-50 border-green-200' 
                : 'bg-red-50 border-red-200'
            }`}>
              <div className="flex items-start gap-2">
                {actionType === 'approve' ? (
                  <CheckCircle className="h-4 w-4 text-green-600 mt-0.5" />
                ) : (
                  <XCircle className="h-4 w-4 text-red-600 mt-0.5" />
                )}
                <div className={`text-sm ${
                  actionType === 'approve' ? 'text-green-800' : 'text-red-800'
                }`}>
                  <p className="font-medium">
                    {actionType === 'approve' ? 'Aprobación' : 'Rechazo'}
                  </p>
                  <p className="text-xs mt-1">
                    {actionType === 'approve' 
                      ? 'El usuario recibirá acceso al sistema con el rol seleccionado y se le enviará un correo de confirmación.'
                      : 'La solicitud será marcada como rechazada y se le notificará al usuario por correo.'
                    }
                  </p>
                </div>
              </div>
            </div>
          </div>

          <div className="flex gap-3 justify-end">
            <Button
              variant="outline"
              onClick={() => {
                setIsActionModalOpen(false);
                setSelectedRequestId(null);
                setSelectedRequestName('');
                setActionType(null);
                setSelectedRoleId('');
              }}
            >
              Cancelar
            </Button>
            <Button
              onClick={handleConfirmAction}
              disabled={(actionType === 'approve' && !selectedRoleId) || isApproving || isRejecting}
              className={actionType === 'approve' 
                ? 'bg-green-600 hover:bg-green-700' 
                : 'bg-red-600 hover:bg-red-700'
              }
            >
              {(isApproving || isRejecting) ? (
                <>
                  <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                  {actionType === 'approve' ? 'Aprobando...' : 'Rechazando...'}
                </>
              ) : (
                <>
                  {actionType === 'approve' ? (
                    <CheckCircle className="h-4 w-4 mr-2" />
                  ) : (
                    <XCircle className="h-4 w-4 mr-2" />
                  )}
                  {actionType === 'approve' ? 'Aprobar' : 'Rechazar'}
                </>
              )}
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}