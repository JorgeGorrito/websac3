export default function AdminDashboardPage() {
  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <h2 className="text-2xl font-semibold text-center">
            Dashboard Administrador
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div className="bg-gradient-to-br from-blue-50 to-blue-100 p-6 rounded-xl border border-blue-200">
            <h3 className="text-lg font-semibold text-blue-800 mb-2">Usuarios Activos</h3>
            <p className="text-3xl font-bold text-blue-600">156</p>
            <p className="text-blue-600 text-sm">Total de usuarios</p>
          </div>
          
          <div className="bg-gradient-to-br from-green-50 to-green-100 p-6 rounded-xl border border-green-200">
            <h3 className="text-lg font-semibold text-green-800 mb-2">Reportes</h3>
            <p className="text-3xl font-bold text-green-600">89</p>
            <p className="text-green-600 text-sm">Este mes</p>
          </div>
          
          <div className="bg-gradient-to-br from-orange-50 to-orange-100 p-6 rounded-xl border border-orange-200">
            <h3 className="text-lg font-semibold text-orange-800 mb-2">Solicitudes</h3>
            <p className="text-3xl font-bold text-orange-600">23</p>
            <p className="text-orange-600 text-sm">Pendientes</p>
          </div>
          
          <div className="bg-gradient-to-br from-purple-50 to-purple-100 p-6 rounded-xl border border-purple-200">
            <h3 className="text-lg font-semibold text-purple-800 mb-2">Accesos</h3>
            <p className="text-3xl font-bold text-purple-600">12</p>
            <p className="text-purple-600 text-sm">Por revisar</p>
          </div>
        </div>
        
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Gestión del Sistema</h3>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Gestionar Usuarios</h4>
              <p className="text-sm text-gray-600 mt-1">Administrar usuarios del sistema</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Gestionar Reportes</h4>
              <p className="text-sm text-gray-600 mt-1">Revisar y aprobar reportes</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Solicitudes de Asesoría</h4>
              <p className="text-sm text-gray-600 mt-1">Gestionar solicitudes pendientes</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Solicitudes de Acceso</h4>
              <p className="text-sm text-gray-600 mt-1">Revisar solicitudes de acceso</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
