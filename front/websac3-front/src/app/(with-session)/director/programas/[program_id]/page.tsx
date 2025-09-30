"use client";

import React from "react";
import { useParams, useRouter } from "next/navigation";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { 
  ArrowLeft, 
  BookOpen, 
  Calendar, 
  Clock, 
  GraduationCap, 
  MapPin, 
  Building, 
  Target, 
  Users, 
  Award,
  FileText,
  Edit
} from "lucide-react";
import { useGetDegreeProgramQuery, useGetDegreeProgramCoursesQuery } from "@/services/api";
import { AcademicSemaphore } from "@/components/websac3/academic/AcademicSemaphore";

export default function ProgramDetailPage() {
  const params = useParams();
  const router = useRouter();
  const programId = parseInt(params.program_id as string, 10);

  // Fetch program details
  const { 
    data: programData, 
    isLoading: programLoading, 
    error: programError 
  } = useGetDegreeProgramQuery({ degree_program_id: programId });

  // Fetch program courses
  const { 
    data: coursesData, 
    isLoading: coursesLoading, 
    error: coursesError 
  } = useGetDegreeProgramCoursesQuery({ degree_program_id: programId });

  const handleBack = () => {
    router.push("/director/programas");
  };

  const handleEditProgram = () => {
    router.push(`/director/editar-programa/${programId}`);
  };

  if (programLoading) {
    return <LoadingSkeleton />;
  }

  if (programError || !programData) {
    return (
      <div className="space-y-6">
        <div className="flex items-center gap-4">
          <Button 
            onClick={handleBack}
            variant="outline"
            className="flex items-center gap-2"
          >
            <ArrowLeft className="h-4 w-4" />
            Volver a Programas
          </Button>
        </div>
        
        <Card>
          <CardContent className="p-12">
            <div className="text-center text-gray-500">
              <BookOpen className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <p className="text-lg font-medium">No se pudo cargar el programa</p>
              <p className="text-sm mt-2">Los errores se mostrarán en un modal automáticamente</p>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  const program = programData;
  const courses = coursesData?.data || [];

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          <Button 
            onClick={handleBack}
            variant="outline"
            className="flex items-center gap-2"
          >
            <ArrowLeft className="h-4 w-4" />
            Volver a Programas
          </Button>
          <div>
            <h1 className="text-3xl font-bold text-gray-900">{program.name}</h1>
            <p className="text-gray-600 mt-1">Detalle del programa de grado</p>
          </div>
        </div>
        <Button 
          onClick={handleEditProgram}
          className="bg-blue-600 hover:bg-blue-700"
        >
          <Edit className="h-4 w-4 mr-2" />
          Editar Programa
        </Button>
      </div>

      {/* Program Overview */}
      <Card className="bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200">
        <CardHeader>
          <CardTitle className="text-2xl font-bold text-gray-900 flex items-center gap-3">
            <GraduationCap className="h-8 w-8 text-blue-600" />
            Información General
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-6">
          {/* Basic Info Grid */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div className="flex items-center gap-3">
              <div className="p-2 bg-blue-100 rounded-lg">
                <Target className="h-5 w-5 text-blue-600" />
              </div>
              <div>
                <p className="text-sm text-gray-600">Código SNIES</p>
                <p className="font-semibold text-gray-900">{program.snies}</p>
              </div>
            </div>
            
            <div className="flex items-center gap-3">
              <div className="p-2 bg-green-100 rounded-lg">
                <Award className="h-5 w-5 text-green-600" />
              </div>
              <div>
                <p className="text-sm text-gray-600">Créditos Totales</p>
                <p className="font-semibold text-gray-900">{program.total_credits}</p>
              </div>
            </div>
            
            <div className="flex items-center gap-3">
              <div className="p-2 bg-purple-100 rounded-lg">
                <Clock className="h-5 w-5 text-purple-600" />
              </div>
              <div>
                <p className="text-sm text-gray-600">Duración</p>
                <p className="font-semibold text-gray-900">
                  {program.duration_value} {program.duration_unit.name}
                </p>
              </div>
            </div>
            
            <div className="flex items-center gap-3">
              <div className="p-2 bg-orange-100 rounded-lg">
                <BookOpen className="h-5 w-5 text-orange-600" />
              </div>
              <div>
                <p className="text-sm text-gray-600">Nivel de Formación</p>
                <p className="font-semibold text-gray-900">{program.formation_level.name}</p>
              </div>
            </div>
          </div>

          {/* Institution Info */}
          <div className="bg-white rounded-lg p-4 border border-gray-200">
            <div className="flex items-center gap-3 mb-3">
              <Building className="h-5 w-5 text-indigo-600" />
              <h3 className="text-lg font-semibold text-gray-900">Institución Educativa</h3>
            </div>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <p className="text-sm text-gray-600">Nombre</p>
                <p className="font-medium text-gray-900">{program.higher_education_institution.name}</p>
              </div>
              <div>
                <p className="text-sm text-gray-600">SNIES</p>
                <p className="font-medium text-gray-900">{program.higher_education_institution.snies}</p>
              </div>
              <div>
                <p className="text-sm text-gray-600">Ubicación</p>
                <p className="font-medium text-gray-900">
                  {program.higher_education_institution.municipality}, {program.higher_education_institution.department}
                </p>
              </div>
              <div>
                <p className="text-sm text-gray-600">Características</p>
                <div className="flex gap-2 mt-1">
                  <Badge className="bg-indigo-100 text-indigo-800 border-indigo-200 text-xs">
                    {program.higher_education_institution.ownership}
                  </Badge>
                  <Badge className="bg-blue-100 text-blue-800 border-blue-200 text-xs">
                    {program.higher_education_institution.institutional_category}
                  </Badge>
                </div>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Program Details */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Entry Profile */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Users className="h-5 w-5 text-green-600" />
              Perfil de Ingreso
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-gray-700 leading-relaxed">{program.entry_profile}</p>
          </CardContent>
        </Card>

        {/* Graduate Profile */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <GraduationCap className="h-5 w-5 text-blue-600" />
              Perfil de Egreso
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-gray-700 leading-relaxed">{program.graduate_profile}</p>
          </CardContent>
        </Card>

        {/* Professional Profile */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Target className="h-5 w-5 text-purple-600" />
              Perfil Profesional
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-gray-700 leading-relaxed">{program.professional_profile}</p>
          </CardContent>
        </Card>

        {/* Program Focus */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <FileText className="h-5 w-5 text-orange-600" />
              Enfoque del Programa
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-gray-700 leading-relaxed">{program.program_focus}</p>
          </CardContent>
        </Card>
      </div>

      {/* Professional Roles */}
      {program.professional_roles && program.professional_roles.length > 0 && (
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Award className="h-5 w-5 text-indigo-600" />
              Roles Profesionales
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-wrap gap-2">
              {program.professional_roles.map((role) => (
                <Badge 
                  key={role.id} 
                  className="bg-indigo-100 text-indigo-800 border-indigo-200"
                >
                  {role.name}
                </Badge>
              ))}
            </div>
          </CardContent>
        </Card>
      )}

      {/* Academic Semaphore */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <BookOpen className="h-5 w-5 text-blue-600" />
            Cursos del Programa
          </CardTitle>
          <CardDescription>
            Visualización de cursos organizados por período con indicadores de ciberseguridad
          </CardDescription>
        </CardHeader>
        <CardContent>
          {coursesLoading ? (
            <div className="space-y-4">
              <Skeleton className="h-8 w-1/3" />
              <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
                {Array.from({ length: 4 }).map((_, i) => (
                  <Skeleton key={i} className="h-20" />
                ))}
              </div>
              <Skeleton className="h-64 w-full" />
            </div>
          ) : coursesError ? (
            <div className="text-center py-8 text-gray-500">
              <BookOpen className="h-12 w-12 mx-auto mb-4 text-gray-400" />
              <p className="text-lg font-medium">No se pudieron cargar los cursos</p>
              <p className="text-sm">Los errores se mostrarán en un modal automáticamente</p>
            </div>
          ) : (
            <AcademicSemaphore
              courses={courses}
              programName={program.name}
              totalCredits={program.total_credits}
            />
          )}
        </CardContent>
      </Card>

      {/* Program Metadata */}
      <Card className="bg-gray-50 border-gray-200">
        <CardContent className="p-4">
          <div className="flex flex-wrap items-center justify-between gap-4 text-sm text-gray-600">
            <div className="flex items-center gap-2">
              <Calendar className="h-4 w-4" />
              <span>Creado: {new Date(program.created_at).toLocaleDateString('es-ES')}</span>
            </div>
            <div className="flex items-center gap-2">
              <Clock className="h-4 w-4" />
              <span>Actualizado: {new Date(program.updated_at).toLocaleDateString('es-ES')}</span>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

// Loading skeleton component
function LoadingSkeleton() {
  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Skeleton className="h-10 w-32" />
        <div>
          <Skeleton className="h-8 w-64 mb-2" />
          <Skeleton className="h-4 w-48" />
        </div>
      </div>
      
      <Card>
        <CardHeader>
          <Skeleton className="h-8 w-48" />
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            {Array.from({ length: 4 }).map((_, i) => (
              <div key={i} className="flex items-center gap-3">
                <Skeleton className="h-9 w-9 rounded-lg" />
                <div>
                  <Skeleton className="h-4 w-20 mb-1" />
                  <Skeleton className="h-5 w-16" />
                </div>
              </div>
            ))}
          </div>
          <Skeleton className="h-32 w-full" />
        </CardContent>
      </Card>
      
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {Array.from({ length: 4 }).map((_, i) => (
          <Card key={i}>
            <CardHeader>
              <Skeleton className="h-6 w-32" />
            </CardHeader>
            <CardContent>
              <Skeleton className="h-20 w-full" />
            </CardContent>
          </Card>
        ))}
      </div>
      
      <Skeleton className="h-64 w-full" />
    </div>
  );
}
