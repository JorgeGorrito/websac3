export default function ExpertoDashboardPage() {
  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <h2 className="text-2xl font-semibold text-center">
            Dashboard Experto en Ciberseguridad
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="bg-gradient-to-br from-green-50 to-green-100 p-6 rounded-xl border border-green-200">
            <h3 className="text-lg font-semibold text-green-800 mb-2">Reportes Retroalimentados</h3>
            <p className="text-3xl font-bold text-green-600">28</p>
            <p className="text-green-600 text-sm">Este mes</p>
          </div>
          
          <div className="bg-gradient-to-br from-orange-50 to-orange-100 p-6 rounded-xl border border-orange-200">
            <h3 className="text-lg font-semibold text-orange-800 mb-2">Reportes Pendientes</h3>
            <p className="text-3xl font-bold text-orange-600">15</p>
            <p className="text-orange-600 text-sm">Por revisar</p>
          </div>
          
          <div className="bg-gradient-to-br from-blue-50 to-blue-100 p-6 rounded-xl border border-blue-200">
            <h3 className="text-lg font-semibold text-blue-800 mb-2">Solicitudes de Asesoría</h3>
            <p className="text-3xl font-bold text-blue-600">9</p>
            <p className="text-blue-600 text-sm">Activas</p>
          </div>
        </div>
        
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Actividades del Experto</h3>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Revisar Reportes</h4>
              <p className="text-sm text-gray-600 mt-1">Analizar y retroalimentar reportes</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Asesorías</h4>
              <p className="text-sm text-gray-600 mt-1">Gestionar solicitudes de asesoría</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Gestionar Modelo</h4>
              <p className="text-sm text-gray-600 mt-1">Configurar y optimizar el modelo</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Perfil</h4>
              <p className="text-sm text-gray-600 mt-1">Gestionar información personal</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
