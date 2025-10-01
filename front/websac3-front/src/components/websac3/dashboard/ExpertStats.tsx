"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { useGetStatisticsQuery } from "@/services/api";
import { 
  MessageSquare, 
  FileCheck, 
  AlertTriangle,
  Activity,
  Shield,
  FolderCheck,
  Headphones,
  UserCog
} from "lucide-react";

interface ExpertStatsProps {
  className?: string;
}

export function ExpertStats({ className = "" }: ExpertStatsProps) {
  const router = useRouter();
  const { data: statistics, isLoading, error } = useGetStatisticsQuery({}, {
    skip: false,
    refetchOnMountOrArgChange: false,
  });

  if (error) {
    console.error("Error loading expert statistics:", error);
  }

  // Provide default stats when API is not available or returns error
  const stats = statistics?.cybersecurity_auditor_stats || {
    consultation_requests: 0,
    feedback_reports: 0,
    pending_reports: 0
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
    color?: "blue" | "green" | "orange" | "red";
  }) => {
    const colorClasses = {
      blue: "bg-blue-500 text-white",
      green: "bg-green-500 text-white", 
      orange: "bg-orange-500 text-white",
      red: "bg-red-500 text-white"
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
          <div className="p-2 bg-purple-100 rounded-lg">
            <Activity className="h-6 w-6 text-purple-600" />
          </div>
          <div>
            <h2 className="text-xl font-semibold text-gray-900">Estadísticas de Auditoría</h2>
            <p className="text-gray-600 text-sm">Métricas de ciberseguridad</p>
          </div>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
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
  //           <h2 className="text-xl font-semibold text-gray-900">Estadísticas de Auditoría</h2>
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
        <div className="p-2 bg-purple-100 rounded-lg">
          <Activity className="h-6 w-6 text-purple-600" />
        </div>
        <div>
          <h2 className="text-xl font-semibold text-gray-900">Estadísticas de Auditoría</h2>
          <p className="text-gray-600 text-sm">Métricas de ciberseguridad</p>
        </div>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <StatCard
          title="Solicitudes de Consulta"
          value={stats.consultation_requests}
          icon={MessageSquare}
          color="blue"
        />
        <StatCard
          title="Reportes con Retroalimentación"
          value={stats.feedback_reports}
          icon={FileCheck}
          color="green"
        />
        <StatCard
          title="Reportes Pendientes"
          value={stats.pending_reports}
          icon={AlertTriangle}
          color={stats.pending_reports > 5 ? "red" : "orange"}
        />
      </div>

      {/* Summary */}
      <div className="bg-gradient-to-r from-purple-50 to-blue-50 border border-purple-200 rounded-lg p-4">
        <div className="flex items-center gap-2 mb-2">
          <Shield className="h-4 w-4 text-purple-600" />
          <span className="text-sm font-semibold text-purple-900">Resumen de Auditoría</span>
        </div>
        <p className="text-sm text-purple-800">
          Se han recibido <strong>{stats.consultation_requests}</strong> solicitudes de consulta, 
          se han completado <strong>{stats.feedback_reports}</strong> reportes con retroalimentación y 
          quedan <strong>{stats.pending_reports}</strong> reportes pendientes de revisión.
        </p>
      </div>

      {/* Priority Actions */}
      {stats.pending_reports > 0 && (
        <div className="bg-orange-50 border border-orange-200 rounded-lg p-4">
          <div className="flex items-center gap-2 mb-2">
            <AlertTriangle className="h-4 w-4 text-orange-600" />
            <span className="text-sm font-semibold text-orange-900">Acción Requerida</span>
          </div>
          <p className="text-sm text-orange-800">
            Tienes <strong>{stats.pending_reports}</strong> reportes pendientes que requieren tu atención. 
            Revisa la sección de reportes pendientes para continuar con el proceso de auditoría.
          </p>
        </div>
      )}

      {/* Quick Actions */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-blue-500" onClick={() => router.push('/experto/asesoria/solicitudes')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <Headphones className="h-8 w-8 text-blue-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Solicitudes de Asesoría</h3>
                <p className="text-sm text-gray-600">Revisar solicitudes de consultoría</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-green-500" onClick={() => router.push('/experto/reportes/retroalimentados')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <FolderCheck className="h-8 w-8 text-green-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Reportes Retroalimentados</h3>
                <p className="text-sm text-gray-600">Ver reportes completados</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-orange-500" onClick={() => router.push('/experto/reportes/pendientes')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <AlertTriangle className="h-8 w-8 text-orange-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Reportes Pendientes</h3>
                <p className="text-sm text-gray-600">Revisar reportes en espera</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-purple-500" onClick={() => router.push('/experto/perfil')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <UserCog className="h-8 w-8 text-purple-600" />
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
