"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { useGetStatisticsQuery } from "@/services/api";
import { 
  MessageSquare, 
  FileText, 
  GraduationCap,
  Activity,
  BookOpen,
  PlusSquare,
  Upload,
  Headphones,
  PieChart,
  Cog,
  UserCog
} from "lucide-react";

interface DirectorStatsProps {
  className?: string;
}

export function DirectorStats({ className = "" }: DirectorStatsProps) {
  const router = useRouter();
  const { data: statistics, isLoading, error } = useGetStatisticsQuery({}, {
    skip: false, // Allow the query to run
    refetchOnMountOrArgChange: false, // Don't refetch on mount
  });

  if (error) {
    console.error("Error loading director statistics:", error);
  }

  // Provide default stats when API is not available or returns error
  const stats = statistics?.program_lead_stats || {
    active_consultations: 0,
    generated_reports: 0,
    registered_programs: 0
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
          <div className="p-2 bg-green-100 rounded-lg">
            <Activity className="h-6 w-6 text-green-600" />
          </div>
          <div>
            <h2 className="text-xl font-semibold text-gray-900">Estadísticas del Programa</h2>
            <p className="text-gray-600 text-sm">Métricas de gestión académica</p>
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
  //           <h2 className="text-xl font-semibold text-gray-900">Estadísticas del Programa</h2>
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
        <div className="p-2 bg-green-100 rounded-lg">
          <Activity className="h-6 w-6 text-green-600" />
        </div>
        <div>
          <h2 className="text-xl font-semibold text-gray-900">Estadísticas del Programa</h2>
          <p className="text-gray-600 text-sm">Métricas de gestión académica</p>
        </div>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <StatCard
          title="Consultas Activas"
          value={stats.active_consultations}
          icon={MessageSquare}
          color="blue"
        />
        <StatCard
          title="Reportes Generados"
          value={stats.generated_reports}
          icon={FileText}
          color="green"
        />
        <StatCard
          title="Programas Registrados"
          value={stats.registered_programs}
          icon={GraduationCap}
          color="purple"
        />
      </div>

      {/* Summary */}
      <div className="bg-gradient-to-r from-green-50 to-blue-50 border border-green-200 rounded-lg p-4">
        <div className="flex items-center gap-2 mb-2">
          <BookOpen className="h-4 w-4 text-green-600" />
          <span className="text-sm font-semibold text-green-900">Resumen Académico</span>
        </div>
        <p className="text-sm text-green-800">
          Actualmente hay <strong>{stats.active_consultations}</strong> consultas activas en curso, 
          se han generado <strong>{stats.generated_reports}</strong> reportes académicos y 
          están registrados <strong>{stats.registered_programs}</strong> programas de grado.
        </p>
      </div>

      {/* Quick Actions */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-blue-500" onClick={() => router.push('/director/programas')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <PlusSquare className="h-8 w-8 text-blue-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Programas de Grado</h3>
                <p className="text-sm text-gray-600">Administrar programas académicos</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-green-500" onClick={() => router.push('/director/cargar-datos')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <Upload className="h-8 w-8 text-green-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Cargar Datos</h3>
                <p className="text-sm text-gray-600">Subir información del programa</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-purple-500" onClick={() => router.push('/director/asesoria-experto')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <Headphones className="h-8 w-8 text-purple-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Asesoría con Experto</h3>
                <p className="text-sm text-gray-600">Solicitar consultoría especializada</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-orange-500" onClick={() => router.push('/director/reportes')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <PieChart className="h-8 w-8 text-orange-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Consultar Reportes</h3>
                <p className="text-sm text-gray-600">Revisar reportes generados</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-red-500" onClick={() => router.push('/director/operar-modelo')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <Cog className="h-8 w-8 text-red-600" />
              <div>
                <h3 className="font-semibold text-gray-900">Evaluar Ciberseguridad</h3>
                <p className="text-sm text-gray-600">Operar modelo de evaluación</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card className="hover:shadow-md transition-shadow duration-200 cursor-pointer border-l-4 border-l-gray-500" onClick={() => router.push('/director/perfil')}>
          <CardContent className="p-4">
            <div className="flex items-center gap-3">
              <UserCog className="h-8 w-8 text-gray-600" />
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
