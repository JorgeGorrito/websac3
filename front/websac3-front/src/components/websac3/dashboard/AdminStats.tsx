"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { useGetStatisticsQuery } from "@/services/api";
import { 
  Users, 
  Inbox, 
  MessageSquare, 
  FileText,
  Activity,
  UserCog
} from "lucide-react";

interface AdminStatsProps {
  className?: string;
}

export function AdminStats({ className = "" }: AdminStatsProps) {
  const router = useRouter();
  const { data: statistics, isLoading, error } = useGetStatisticsQuery({}, {
    skip: false,
    refetchOnMountOrArgChange: false,
  });

  if (error) {
    console.error("Error loading admin statistics:", error);
  }

  // Provide default stats when API is not available or returns error
  const stats = statistics?.admin_stats || {
    active_users: 0,
    total_access_requests: 0,
    total_consultations: 0,
    total_reports: 0
  };

  const StatCard = ({ 
    title, 
    value, 
    icon: Icon, 
    color = "blue"
  }: {
    title: string;
    value: number;
    icon: React.ElementType;
    color?: "blue" | "green" | "orange" | "purple";
  }) => {
    const colorClasses = {
      blue: "bg-blue-500 text-white",
      green: "bg-green-500 text-white", 
      orange: "bg-orange-500 text-white",
      purple: "bg-purple-500 text-white"
    };

    return (
      <Card className="hover:shadow-md transition-shadow duration-200">
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-sm font-medium text-gray-600">
            {title}
          </CardTitle>
          <div className={`p-2 rounded-lg ${colorClasses[color]}`}>
            <Icon className="h-4 w-4" />
          </div>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold text-gray-900">{value}</div>
        </CardContent>
      </Card>
    );
  };

  const StatSkeleton = () => (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <Skeleton className="h-4 w-24" />
        <Skeleton className="h-8 w-8 rounded-lg" />
      </CardHeader>
      <CardContent>
        <Skeleton className="h-8 w-16 mb-2" />
      </CardContent>
    </Card>
  );

  if (isLoading) {
    return (
      <div className={`space-y-6 ${className}`}>
        <div className="flex items-center gap-3">
          <div className="p-2 bg-blue-100 rounded-lg">
            <Activity className="h-6 w-6 text-blue-600" />
          </div>
          <div>
            <h2 className="text-xl font-semibold text-gray-900">Estadísticas del Sistema</h2>
            <p className="text-gray-600 text-sm">Métricas generales de la plataforma</p>
          </div>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          <StatSkeleton />
          <StatSkeleton />
          <StatSkeleton />
          <StatSkeleton />
        </div>
      </div>
    );
  }

  // Don't show error state, just use default values
  // if (error || !stats) {
  //   return (
  //     <div className={`space-y-6 ${className}`}>
  //       <div className="flex items-center gap-3">
  //         <div className="p-2 bg-red-100 rounded-lg">
  //           <Activity className="h-6 w-6 text-red-600" />
  //         </div>
  //         <div>
  //           <h2 className="text-xl font-semibold text-gray-900">Estadísticas del Sistema</h2>
  //           <p className="text-gray-600 text-sm">Error al cargar las métricas</p>
  //         </div>
  //       </div>
  //       
  //       <div className="bg-red-50 border border-red-200 rounded-lg p-4">
  //         <p className="text-red-800 text-sm">
  //           No se pudieron cargar las estadísticas. Por favor, intente nuevamente.
  //         </p>
  //       </div>
  //     </div>
  //   );
  // }

  return (
    <div className={`space-y-6 ${className}`}>
      {/* Header */}
      <div className="flex items-center gap-3">
        <div className="p-2 bg-blue-100 rounded-lg">
          <Activity className="h-6 w-6 text-blue-600" />
        </div>
        <div>
          <h2 className="text-xl font-semibold text-gray-900">Estadísticas del Sistema</h2>
          <p className="text-gray-600 text-sm">Métricas generales de la plataforma</p>
        </div>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <StatCard
          title="Usuarios Activos"
          value={stats.active_users}
          icon={Users}
          color="blue"
        />
        <StatCard
          title="Solicitudes de Acceso"
          value={stats.total_access_requests}
          icon={Inbox}
          color="orange"
        />
        <StatCard
          title="Consultas Totales"
          value={stats.total_consultations}
          icon={MessageSquare}
          color="green"
        />
        <StatCard
          title="Reportes Generados"
          value={stats.total_reports}
          icon={FileText}
          color="purple"
        />
      </div>

      {/* Summary */}
      <div className="bg-gradient-to-r from-blue-50 to-purple-50 border border-blue-200 rounded-lg p-4">
        <div className="flex items-center gap-2 mb-2">
          <Activity className="h-4 w-4 text-blue-600" />
          <span className="text-sm font-semibold text-blue-900">Resumen General</span>
        </div>
        <p className="text-sm text-blue-800">
          La plataforma cuenta con <strong>{stats.active_users}</strong> usuarios activos, 
          ha procesado <strong>{stats.total_access_requests}</strong> solicitudes de acceso, 
          realizado <strong>{stats.total_consultations}</strong> consultas y generado 
          <strong> {stats.total_reports}</strong> reportes en total.
        </p>
      </div>

      {/* Quick Actions */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-blue-500" onClick={() => router.push('/admin/usuarios')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <Users className="h-8 w-8 text-blue-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Gestionar Usuarios</h3>
                <p className="text-sm text-gray-600">Administrar usuarios del sistema</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-orange-500" onClick={() => router.push('/admin/accesos')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <Inbox className="h-8 w-8 text-orange-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Solicitudes de Acceso</h3>
                <p className="text-sm text-gray-600">Revisar solicitudes de acceso</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-green-500" onClick={() => router.push('/admin/perfil')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <UserCog className="h-8 w-8 text-green-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Mi Perfil</h3>
                <p className="text-sm text-gray-600">Gestionar información personal</p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
