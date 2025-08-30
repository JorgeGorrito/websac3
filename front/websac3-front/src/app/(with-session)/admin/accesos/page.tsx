"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  useApproveAccessRequestMutation,
  useListAccessRequestsQuery,
  useRejectAccessRequestMutation,
} from "@/services/api";
import { useMemo, useState } from "react";

export default function AdminAccessRequestsPage() {
  const [search, setSearch] = useState("");
  const { data, isLoading, isError, refetch, isFetching } = useListAccessRequestsQuery({
    current_page: 1,
    items_per_page: 20,
    filters: search
      ? { "Applicant.name[cont]": search, "Applicant.lastname[cont]": search }
      : undefined,
  });
  const [approve, { isLoading: approving }] = useApproveAccessRequestMutation();
  const [reject, { isLoading: rejecting }] = useRejectAccessRequestMutation();

  const rows = data?.data ?? [];
  const loading = isLoading || isFetching;

  const totals = useMemo(() => {
    const total = data?.total_count ?? 0;
    const pending = rows.filter((r) => r.status_name?.toLowerCase().includes("pend")).length;
    const approved = rows.filter((r) => r.status_name?.toLowerCase().includes("aprob")).length;
    const rejected = rows.filter((r) => r.status_name?.toLowerCase().includes("rech")).length;
    return { total, pending, approved, rejected };
  }, [data, rows]);

  const onApprove = async (id: number) => {
    await approve({ id }).unwrap();
    refetch();
  };

  const onReject = async (id: number) => {
    await reject({ id }).unwrap();
    refetch();
  };

  return (
    <div className="space-y-6">
      <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
        <h2 className="text-2xl font-semibold">Solicitudes de Acceso</h2>
        <div className="flex items-center gap-2">
          <Input
            placeholder="Buscar por nombre o apellido"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-64"
          />
          <Button variant="outline" onClick={() => refetch()} disabled={loading}>
            Buscar
          </Button>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <Card>
          <CardHeader>
            <CardTitle>Total</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-bold">{totals.total}</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle>Pendientes</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-bold">{totals.pending}</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle>Aprobadas</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-bold">{totals.approved}</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle>Rechazadas</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-bold">{totals.rejected}</p>
          </CardContent>
        </Card>
      </div>

      <div className="bg-white rounded-xl shadow-sm border">
        <div className="overflow-x-auto">
          <table className="min-w-full text-left text-sm">
            <thead className="bg-gray-50 text-gray-700">
              <tr>
                <th className="px-4 py-3">Nombre</th>
                <th className="px-4 py-3">Correo</th>
                <th className="px-4 py-3">Documento</th>
                <th className="px-4 py-3">Institución</th>
                <th className="px-4 py-3">Estado</th>
                <th className="px-4 py-3 text-right">Acciones</th>
              </tr>
            </thead>
            <tbody>
              {loading && (
                <tr>
                  <td className="px-4 py-6" colSpan={6}>Cargando...</td>
                </tr>
              )}
              {isError && !loading && (
                <tr>
                  <td className="px-4 py-6 text-red-600" colSpan={6}>
                    Error al cargar solicitudes. Intenta nuevamente.
                  </td>
                </tr>
              )}
              {!loading && !isError && rows.length === 0 && (
                <tr>
                  <td className="px-4 py-6" colSpan={6}>No hay solicitudes</td>
                </tr>
              )}
              {rows.map((r) => (
                <tr key={r.id} className="border-t">
                  <td className="px-4 py-3">
                    {r.name} {r.lastname}
                    <div className="text-xs text-gray-500">{r.job_position}</div>
                  </td>
                  <td className="px-4 py-3">{r.email}</td>
                  <td className="px-4 py-3">
                    {r.identification_type} {r.identification_number}
                  </td>
                  <td className="px-4 py-3">
                    <div>{r.higher_education_institution_name}</div>
                    <div className="text-xs text-gray-500">
                      {r.municipality_name} - {r.department_name}
                    </div>
                  </td>
                  <td className="px-4 py-3">{r.status_name}</td>
                  <td className="px-4 py-3 text-right">
                    <div className="inline-flex gap-2">
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => onApprove(r.id)}
                        disabled={approving || rejecting}
                      >
                        Aprobar
                      </Button>
                      <Button
                        variant="destructive"
                        size="sm"
                        onClick={() => onReject(r.id)}
                        disabled={approving || rejecting}
                      >
                        Rechazar
                      </Button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}


