"use client";

import Link from "next/link";
import { useListAccessRequestsQuery } from "@/services/api";

export default function AdminDashboardPage() {
  const { data, isLoading, isError, error } = useListAccessRequestsQuery({ current_page: 1, items_per_page: 10 });
  
  // Handle 404 as a normal case (no access requests), not an error
  const hasAccessRequests = data && data.data && data.data.length > 0;
  const accessRequests = data?.data ?? [];
  const pendingAccessCount = accessRequests.filter((r) => r.status_name?.toLowerCase().includes("pend")).length;
  
  // Only consider it an error if it's not a 404
  const isActualError = isError && error && (error as any)?.status !== 404;
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
            <p className="text-3xl font-bold text-blue-600">—</p>
            <p className="text-blue-600 text-sm">Total de usuarios</p>
          </div>
          
          <div className="bg-gradient-to-br from-green-50 to-green-100 p-6 rounded-xl border border-green-200">
            <h3 className="text-lg font-semibold text-green-800 mb-2">Reportes</h3>
            <p className="text-3xl font-bold text-green-600">—</p>
            <p className="text-green-600 text-sm">Este mes</p>
          </div>
          
          <div className="bg-gradient-to-br from-orange-50 to-orange-100 p-6 rounded-xl border border-orange-200">
            <h3 className="text-lg font-semibold text-orange-800 mb-2">Solicitudes</h3>
            <p className="text-3xl font-bold text-orange-600">—</p>
            <p className="text-orange-600 text-sm">Pendientes</p>
          </div>
          
          <div className="bg-gradient-to-br from-purple-50 to-purple-100 p-6 rounded-xl border border-purple-200">
            <h3 className="text-lg font-semibold text-purple-800 mb-2">Accesos</h3>
            <p className="text-3xl font-bold text-purple-600">
              {isLoading ? "…" : isActualError ? "!" : pendingAccessCount}
            </p>
            <p className="text-purple-600 text-sm">Por revisar (últimos 10)</p>
          </div>
        </div>
        
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Gestión del Sistema</h3>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <Link href="/admin/usuarios" className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors">
              <h4 className="font-medium text-gray-900">Gestionar Usuarios</h4>
              <p className="text-sm text-gray-600 mt-1">Administrar usuarios del sistema</p>
            </Link>
            
            <Link href="/admin/reportes" className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors">
              <h4 className="font-medium text-gray-900">Gestionar Reportes</h4>
              <p className="text-sm text-gray-600 mt-1">Revisar y aprobar reportes</p>
            </Link>
            
            <Link href="/admin/asesoria" className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors">
              <h4 className="font-medium text-gray-900">Solicitudes de Asesoría</h4>
              <p className="text-sm text-gray-600 mt-1">Gestionar solicitudes pendientes</p>
            </Link>
            
            <Link href="/admin/accesos" className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors">
              <h4 className="font-medium text-gray-900">Solicitudes de Acceso</h4>
              <p className="text-sm text-gray-600 mt-1">Revisar solicitudes de acceso</p>
            </Link>
          </div>
          <div className="mt-6">
            <h4 className="font-medium text-gray-900 mb-2">Últimas solicitudes de acceso</h4>
            {isLoading && <p className="text-sm text-gray-500">Cargando…</p>}
            {isActualError && <p className="text-sm text-red-600">Error al cargar</p>}
            {!isLoading && !isActualError && (
              <ul className="list-disc pl-5 space-y-1 max-h-48 overflow-auto">
                {accessRequests.slice(0, 5).map((r) => (
                  <li key={r.id} className="text-sm text-gray-700">
                    {r.name} {r.lastname} • {r.higher_education_institution_name} • {r.status_name}
                  </li>
                ))}
                {accessRequests.length === 0 && <li className="text-sm text-gray-500">No hay solicitudes de acceso registradas</li>}
              </ul>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
