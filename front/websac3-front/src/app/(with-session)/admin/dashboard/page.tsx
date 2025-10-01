"use client";

import Link from "next/link";
import { useListAccessRequestsQuery } from "@/services/api";
import { AdminStats } from "@/components/websac3/dashboard/AdminStats";

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
      {/* Statistics Section */}
      <AdminStats />

      {/* Recent Access Requests */}
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <h2 className="text-2xl font-semibold text-center">
            Dashboard Administrador
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>
        
        <div className="mt-8">
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
  );
}
