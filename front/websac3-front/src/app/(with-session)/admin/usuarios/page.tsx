"use client";

import { useState } from "react";
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
  Mail,
  User,
  CheckCircle,
  XCircle,
  MoreHorizontal
} from "lucide-react";
import { 
  useListUsersQuery,
  useListRolesQuery,
  useActivateUserMutation,
  useDeactivateUserMutation
} from "@/services/api";
import { ProgramsPagination } from "@/components/websac3/program/ProgramsPagination";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { useDispatch } from "react-redux";
import { showError, showSuccess } from "@/store/errorSlice";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";

export default function GestionarUsuariosPage() {
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
    const email = searchParams.get('email') || '';
    const name = searchParams.get('name') || '';
    const lastname = searchParams.get('lastname') || '';
    const role = searchParams.get('role') || '';
    const isActive = searchParams.get('is_active') || 'all';
    return { email, name, lastname, role, isActive };
  });
  
  const [showFilters, setShowFilters] = useState(false);
  const [selectedUserId, setSelectedUserId] = useState<number | null>(null);
  const [selectedUserName, setSelectedUserName] = useState<string>('');
  const [isActionModalOpen, setIsActionModalOpen] = useState(false);
  const [actionType, setActionType] = useState<'activate' | 'deactivate' | null>(null);

  // Build query parameters with proper filter format
  const buildFilters = () => {
    const filterParams: Record<string, string> = {};
    
    if (filters.email.trim()) {
      filterParams['email[cont]'] = filters.email.trim();
    }
    
    if (filters.name.trim()) {
      filterParams['Person.name[cont]'] = filters.name.trim();
    }
    
    if (filters.lastname.trim()) {
      filterParams['Person.lastname[cont]'] = filters.lastname.trim();
    }
    
    if (filters.role.trim()) {
      filterParams['Role.name[cont]'] = filters.role.trim();
    }
    
    if (filters.isActive !== '' && filters.isActive !== 'all') {
      filterParams['is_active[eq]'] = filters.isActive;
    }
    
    return filterParams;
  };

  const queryParams = {
    current_page: currentPage,
    items_per_page: itemsPerPage,
    filters: buildFilters()
  };

  const { data, isLoading, error, refetch } = useListUsersQuery(queryParams);
  const [activateUser, { isLoading: isActivating }] = useActivateUserMutation();
  const [deactivateUser, { isLoading: isDeactivating }] = useDeactivateUserMutation();

  // Function to update URL with current parameters
  const updateURL = (newParams: {
    page?: number;
    per_page?: number;
    email?: string;
    name?: string;
    lastname?: string;
    role?: string;
    is_active?: string;
  }) => {
    const params = new URLSearchParams();
    const page = newParams.page ?? currentPage;
    const perPage = newParams.per_page ?? itemsPerPage;
    const email = newParams.email ?? filters.email;
    const name = newParams.name ?? filters.name;
    const lastname = newParams.lastname ?? filters.lastname;
    const role = newParams.role ?? filters.role;
    const isActive = newParams.is_active ?? filters.isActive;
    
    if (page !== 1) params.set('page', page.toString());
    if (perPage !== 10) params.set('per_page', perPage.toString());
    if (email.trim()) params.set('email', email.trim());
    if (name.trim()) params.set('name', name.trim());
    if (lastname.trim()) params.set('lastname', lastname.trim());
    if (role.trim()) params.set('role', role.trim());
    if (isActive !== '') params.set('is_active', isActive);
    
    const newURL = params.toString() 
      ? `/admin/usuarios?${params.toString()}`
      : '/admin/usuarios';
    
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
    setFilters({ email: '', name: '', lastname: '', role: '', isActive: 'all' });
    setCurrentPage(1);
    updateURL({ email: '', name: '', lastname: '', role: '', isActive: 'all', page: 1 });
  };

  const handleActionClick = (userId: number, userName: string, action: 'activate' | 'deactivate') => {
    setSelectedUserId(userId);
    setSelectedUserName(userName);
    setActionType(action);
    setIsActionModalOpen(true);
  };

  const handleConfirmAction = async () => {
    if (!selectedUserId || !actionType) return;

    try {
      if (actionType === 'activate') {
        await activateUser({ id: selectedUserId }).unwrap();
        dispatch(showSuccess({ 
          message: 'Usuario activado exitosamente.',
          title: 'Usuario Activado'
        }));
      } else {
        await deactivateUser({ id: selectedUserId }).unwrap();
        dispatch(showSuccess({ 
          message: 'Usuario desactivado exitosamente.',
          title: 'Usuario Desactivado'
        }));
      }
      
      setIsActionModalOpen(false);
      setSelectedUserId(null);
      setSelectedUserName('');
      setActionType(null);
      refetch(); // Refresh the list
    } catch (error: any) {
      console.log('Error details:', error);
      console.log('Error data:', error?.data);
      console.log('Error status:', error?.status);
      
      // Handle different error structures
      let errorMessage = `Error al ${actionType === 'activate' ? 'activar' : 'desactivar'} el usuario`;
      
      if (error?.data?.Errors && Array.isArray(error.data.Errors) && error.data.Errors.length > 0) {
        // Backend error structure: { Errors: ["message"] }
        errorMessage = error.data.Errors[0];
      } else if (error?.data?.errors && Array.isArray(error.data.errors) && error.data.errors.length > 0) {
        // Alternative error structure: { errors: ["message"] }
        errorMessage = error.data.errors[0];
      } else if (error?.data?.message) {
        // Single message structure: { message: "text" }
        errorMessage = error.data.message;
      } else if (error?.message) {
        // Direct error message
        errorMessage = error.message;
      }
      
      dispatch(showError(errorMessage));
    }
  };

  const getStatusBadge = (isActive: boolean) => {
    return isActive ? (
      <Badge variant="outline" className="bg-green-50 text-green-700 border-green-200">
        <CheckCircle className="h-3 w-3 mr-1" />
        Activo
      </Badge>
    ) : (
      <Badge variant="outline" className="bg-red-50 text-red-700 border-red-200">
        <XCircle className="h-3 w-3 mr-1" />
        Inactivo
      </Badge>
    );
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
            <h1 className="text-3xl font-bold text-gray-900">Gestionar Usuarios</h1>
            <p className="text-gray-600 mt-2">Administra los usuarios registrados en el sistema</p>
          </div>
        </div>
        
        <Card>
          <CardContent className="p-6">
            <div className="text-center text-gray-500">
              <Users className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <p>No se pudieron cargar los usuarios</p>
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
          <h1 className="text-3xl font-bold text-gray-900">Gestionar Usuarios</h1>
          <p className="text-gray-600 mt-2">Administra los usuarios registrados en el sistema</p>
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
                  {data.total_count} usuario{data.total_count !== 1 ? 's' : ''} registrado{data.total_count !== 1 ? 's' : ''}
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
                    {data.data.filter(user => user.is_active).length} activos
                  </span>
                </div>
                <div className="flex items-center gap-2">
                  <XCircle className="h-4 w-4 text-red-600" />
                  <span className="text-sm text-gray-600">
                    {data.data.filter(user => !user.is_active).length} inactivos
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
              Personaliza tu búsqueda y visualización de usuarios
            </CardDescription>
          </CardHeader>
          <CardContent className="p-6">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {/* Email Filter */}
              <div className="space-y-2">
                <Label htmlFor="email" className="text-sm font-medium text-gray-700">
                  Email
                </Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="Buscar por email..."
                  value={filters.email}
                  onChange={(e) => handleFilterChange('email', e.target.value)}
                  className="w-full"
                />
              </div>

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

              {/* Role Filter */}
              <div className="space-y-2">
                <Label htmlFor="role" className="text-sm font-medium text-gray-700">
                  Rol
                </Label>
                <Input
                  id="role"
                  type="text"
                  placeholder="Buscar por rol..."
                  value={filters.role}
                  onChange={(e) => handleFilterChange('role', e.target.value)}
                  className="w-full"
                />
              </div>

              {/* Status Filter */}
              <div className="space-y-2">
                <Label htmlFor="isActive" className="text-sm font-medium text-gray-700">
                  Estado
                </Label>
                <Select value={filters.isActive} onValueChange={(value) => handleFilterChange('isActive', value)}>
                  <SelectTrigger>
                    <SelectValue placeholder="Todos los estados" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="all">Todos los estados</SelectItem>
                    <SelectItem value="true">Activos</SelectItem>
                    <SelectItem value="false">Inactivos</SelectItem>
                  </SelectContent>
                </Select>
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

      {/* Users List */}
      <div className="space-y-4">
        {data && data.data.length > 0 && !(error && error.status === 404) ? (
          <div className="grid grid-cols-1 gap-6">
            {data.data.map((user) => (
              <Card key={user.id} className="bg-white border border-gray-200 shadow-lg hover:shadow-xl transition-all duration-300">
                <CardContent className="p-6">
                  <div className="flex items-start justify-between">
                    <div className="flex-1">
                      <div className="flex items-center gap-3 mb-3">
                        <h3 className="text-lg font-semibold text-gray-900">
                          {user.full_name}
                        </h3>
                        {getStatusBadge(user.is_active)}
                      </div>
                      
                      <div className="space-y-2">
                        <div className="flex items-center gap-2">
                          <Mail className="h-4 w-4 text-gray-500" />
                          <span className="text-sm text-gray-600">{user.email}</span>
                        </div>
                        <div className="flex items-center gap-2">
                          <User className="h-4 w-4 text-gray-500" />
                          <span className="text-sm text-gray-600">ID: {user.id}</span>
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
                          {user.is_active ? (
                            <DropdownMenuItem
                              onClick={() => handleActionClick(user.id, user.full_name, 'deactivate')}
                              className="text-red-600"
                            >
                              <XCircle className="h-4 w-4 mr-2" />
                              Desactivar
                            </DropdownMenuItem>
                          ) : (
                            <DropdownMenuItem
                              onClick={() => handleActionClick(user.id, user.full_name, 'activate')}
                              className="text-green-600"
                            >
                              <CheckCircle className="h-4 w-4 mr-2" />
                              Activar
                            </DropdownMenuItem>
                          )}
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
                  No hay usuarios disponibles
                </h3>
                <p className="text-gray-600 mb-6">
                  No se encontraron usuarios registrados
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
              {actionType === 'activate' ? (
                <CheckCircle className="h-5 w-5 text-green-600" />
              ) : (
                <XCircle className="h-5 w-5 text-red-600" />
              )}
              {actionType === 'activate' ? 'Activar Usuario' : 'Desactivar Usuario'}
            </DialogTitle>
            <DialogDescription>
              {actionType === 'activate' 
                ? `¿Estás seguro de que deseas activar el usuario ${selectedUserName}?`
                : `¿Estás seguro de que deseas desactivar el usuario ${selectedUserName}?`
              }
            </DialogDescription>
          </DialogHeader>
          
          <div className="space-y-4">
            <div className={`border rounded-lg p-3 ${
              actionType === 'activate' 
                ? 'bg-green-50 border-green-200' 
                : 'bg-red-50 border-red-200'
            }`}>
              <div className="flex items-start gap-2">
                {actionType === 'activate' ? (
                  <CheckCircle className="h-4 w-4 text-green-600 mt-0.5" />
                ) : (
                  <XCircle className="h-4 w-4 text-red-600 mt-0.5" />
                )}
                <div className={`text-sm ${
                  actionType === 'activate' ? 'text-green-800' : 'text-red-800'
                }`}>
                  <p className="font-medium">
                    {actionType === 'activate' ? 'Activación' : 'Desactivación'}
                  </p>
                  <p className="text-xs mt-1">
                    {actionType === 'activate' 
                      ? 'El usuario podrá acceder nuevamente al sistema y recibirá una notificación por correo.'
                      : 'El usuario no podrá acceder al sistema hasta que sea reactivado.'
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
                setSelectedUserId(null);
                setSelectedUserName('');
                setActionType(null);
              }}
            >
              Cancelar
            </Button>
            <Button
              onClick={handleConfirmAction}
              disabled={isActivating || isDeactivating}
              className={actionType === 'activate' 
                ? 'bg-green-600 hover:bg-green-700' 
                : 'bg-red-600 hover:bg-red-700'
              }
            >
              {(isActivating || isDeactivating) ? (
                <>
                  <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                  {actionType === 'activate' ? 'Activando...' : 'Desactivando...'}
                </>
              ) : (
                <>
                  {actionType === 'activate' ? (
                    <CheckCircle className="h-4 w-4 mr-2" />
                  ) : (
                    <XCircle className="h-4 w-4 mr-2" />
                  )}
                  {actionType === 'activate' ? 'Activar' : 'Desactivar'}
                </>
              )}
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}
