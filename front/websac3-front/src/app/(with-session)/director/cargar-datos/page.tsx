"use client";

import React, { useState } from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Upload, ArrowLeft, CheckCircle, BookOpen, Clock, GraduationCap, Target } from "lucide-react";
import { useListDegreeProgramsQuery } from "@/services/api";
import { useRouter } from "next/navigation";

export default function CargarDatosPage() {
  const router = useRouter();
  const [selectedProgramId, setSelectedProgramId] = useState<number | null>(null);

  // Fetch degree programs
  const { data, isLoading, error } = useListDegreeProgramsQuery({
    current_page: 1,
    items_per_page: 50, // Get more programs for selection
  });

  const handleLoadData = (programId: number) => {
    router.push(`/director/cargar-datos/formulario?program_id=${programId}`);
  };

  if (isLoading) {
    return <LoadingSkeleton />;
  }

  if (error) {
  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
          <div className="text-center text-gray-500">
            <BookOpen className="h-12 w-12 text-gray-400 mx-auto mb-4" />
            <p>No se pudieron cargar los programas de grado</p>
            <p className="text-sm mt-2">Los errores se mostrarán en un modal automáticamente</p>
        </div>
      </div>
    </div>
  );
}

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Cargar Datos</h1>
          <p className="text-gray-600 mt-2">Selecciona un programa de grado para cargar cursos</p>
        </div>
                  </div>

      {/* Programs List */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {data?.data.map((program) => (
          <Card key={program.id} className="bg-white border border-gray-200 shadow-lg hover:shadow-xl transition-all duration-300 hover:-translate-y-1">
            <CardHeader className="pb-3 pt-4 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-gray-100">
              <div className="flex items-start justify-between">
                <div className="flex-1">
                  <CardTitle className="text-xl font-bold text-gray-900 line-clamp-1 mb-2">
                    {program.name}
                  </CardTitle>
                  <div className="flex items-center gap-2">
                    <Badge className="bg-blue-100 text-blue-800 border-blue-200 font-semibold">
                      SNIES: {program.snies}
                    </Badge>
                  </div>
                </div>
                <Badge className="bg-gradient-to-r from-green-500 to-emerald-500 text-white font-semibold px-3 py-1">
                  {program.duration_unit.name}
                </Badge>
              </div>
            </CardHeader>
            
            <CardContent className="p-5 space-y-4">
              {/* Program Details */}
              <div className="grid grid-cols-2 gap-4">
                <div className="flex items-center p-3 bg-blue-50 rounded-lg">
                  <div className="p-2 bg-blue-500 rounded-lg mr-3">
                    <GraduationCap className="h-4 w-4 text-white" />
                  </div>
                  <div>
                    <p className="text-xs font-medium text-blue-600 uppercase tracking-wide">Créditos</p>
                    <p className="text-lg font-bold text-blue-900">{program.total_credits}</p>
                  </div>
                </div>
                <div className="flex items-center p-3 bg-green-50 rounded-lg">
                  <div className="p-2 bg-green-500 rounded-lg mr-3">
                    <Clock className="h-4 w-4 text-white" />
                  </div>
                  <div>
                    <p className="text-xs font-medium text-green-600 uppercase tracking-wide">Duración</p>
                    <p className="text-lg font-bold text-green-900">{program.duration_value} {program.duration_unit.name.toLowerCase()}</p>
                  </div>
                </div>
              </div>

              {/* Program Focus */}
              <div className="p-4 bg-purple-50 rounded-lg border border-purple-100">
                <div className="flex items-center mb-2">
                  <div className="p-2 bg-purple-500 rounded-lg mr-3">
                    <Target className="h-4 w-4 text-white" />
                  </div>
                  <span className="text-sm font-semibold text-purple-700 uppercase tracking-wide">Enfoque</span>
                </div>
                <p className="text-sm text-gray-700 line-clamp-2 pl-11">
                  {program.program_focus}
                </p>
              </div>

              {/* Institution */}
              <div className="p-4 bg-indigo-50 rounded-lg border border-indigo-100">
                <div className="flex items-center mb-2">
                  <div className="p-2 bg-indigo-500 rounded-lg mr-3">
                    <GraduationCap className="h-4 w-4 text-white" />
                  </div>
                  <span className="text-sm font-semibold text-indigo-700 uppercase tracking-wide">Institución</span>
                </div>
                <p className="text-sm font-medium text-gray-800 pl-11 mb-1">
                  {program.higher_education_institution.name}
                </p>
                <Badge className="ml-11 bg-indigo-100 text-indigo-800 border-indigo-200 font-medium">
                  SNIES: {program.higher_education_institution.snies}
                </Badge>
              </div>

              {/* Load Data Button */}
              <div className="pt-4 border-t border-gray-100">
                <Button
                  onClick={() => handleLoadData(program.id)}
                  className="w-full bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 shadow-lg hover:shadow-xl transition-all duration-200"
                >
                  <Upload className="h-4 w-4 mr-2" />
                  Cargar Datos
                </Button>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      {data?.data.length === 0 && (
        <Card className="bg-white border border-gray-200 shadow-lg">
          <CardContent className="p-8">
            <div className="text-center text-gray-500">
              <BookOpen className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <p className="text-lg font-medium">No hay programas de grado registrados</p>
              <p className="text-sm mt-2">Registra un programa de grado primero para poder cargar datos</p>
            </div>
          </CardContent>
        </Card>
      )}
    </div>
  );
}

// Loading skeleton component
function LoadingSkeleton() {
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <Skeleton className="h-8 w-48 mb-2" />
          <Skeleton className="h-4 w-64" />
        </div>
                </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {Array.from({ length: 6 }).map((_, i) => (
          <Card key={i} className="bg-white border border-gray-200 shadow-lg">
            <CardHeader className="pb-3 pt-4 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-gray-100">
              <Skeleton className="h-6 w-3/4 mb-2" />
              <Skeleton className="h-4 w-1/2" />
            </CardHeader>
            <CardContent className="p-5 space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <Skeleton className="h-16 w-full" />
                <Skeleton className="h-16 w-full" />
              </div>
              <Skeleton className="h-20 w-full" />
              <Skeleton className="h-20 w-full" />
              <Skeleton className="h-10 w-full" />
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
