export default function DirectorDashboardPage() {
  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <h2 className="text-2xl font-semibold text-center">
            Dashboard Director
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="bg-gradient-to-br from-blue-50 to-blue-100 p-6 rounded-xl border border-blue-200">
            <h3 className="text-lg font-semibold text-blue-800 mb-2">Programas Registrados</h3>
            <p className="text-3xl font-bold text-blue-600">12</p>
            <p className="text-blue-600 text-sm">Total de programas</p>
          </div>
          
          <div className="bg-gradient-to-br from-green-50 to-green-100 p-6 rounded-xl border border-green-200">
            <h3 className="text-lg font-semibold text-green-800 mb-2">Reportes Generados</h3>
            <p className="text-3xl font-bold text-green-600">45</p>
            <p className="text-green-600 text-sm">Este mes</p>
          </div>
          
          <div className="bg-gradient-to-br from-purple-50 to-purple-100 p-6 rounded-xl border border-purple-200">
            <h3 className="text-lg font-semibold text-purple-800 mb-2">Asesorías Activas</h3>
            <p className="text-3xl font-bold text-purple-600">8</p>
            <p className="text-purple-600 text-sm">En curso</p>
          </div>
        </div>
        
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Acciones Rápidas</h3>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Registrar Nuevo Programa</h4>
              <p className="text-sm text-gray-600 mt-1">Crear un nuevo programa académico</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Cargar Datos</h4>
              <p className="text-sm text-gray-600 mt-1">Subir información del programa</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Generar Reporte</h4>
              <p className="text-sm text-gray-600 mt-1">Crear reporte de análisis</p>
            </div>
            
            <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
              <h4 className="font-medium text-gray-900">Solicitar Asesoría</h4>
              <p className="text-sm text-gray-600 mt-1">Contactar con experto</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
