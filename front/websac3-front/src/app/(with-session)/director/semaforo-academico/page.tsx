"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { ArrowLeft, BookOpen, GraduationCap, TrendingUp } from "lucide-react";
import { AcademicSemaphore, Course } from "@/components/websac3/academic/AcademicSemaphore";
import { useListDegreeProgramsQuery, useListCoursesQuery, DegreeProgramItem } from "@/services/api";

export default function SemáforoAcadémicoPage() {
  const router = useRouter();
  const [selectedProgramId, setSelectedProgramId] = useState<number | null>(null);

  // Obtener lista de programas
  const { data: programsData, isLoading: programsLoading } = useListDegreeProgramsQuery({
    current_page: 1,
    items_per_page: 50,
  });

  // Obtener cursos del programa seleccionado
  const { data: coursesData, isLoading: coursesLoading } = useListCoursesQuery(
    {
      current_page: 1,
      items_per_page: 100,
      degree_program_id: selectedProgramId || undefined,
    },
    {
      skip: !selectedProgramId,
    }
  );

  const selectedProgram = programsData?.data.find(p => p.id === selectedProgramId);

  // Simular estados de cursos (en un sistema real, esto vendría del backend)
  const coursesWithStatus: Course[] = (coursesData?.data || []).map(course => ({
    ...course,
    status: getRandomStatus(),
    grade: getRandomGrade(),
  }));

  function getRandomStatus(): 'completed' | 'current' | 'pending' | 'failed' {
    const statuses = ['completed', 'current', 'pending', 'failed'] as const;
    const weights = [0.6, 0.1, 0.2, 0.1]; // 60% completados, 10% actuales, 20% pendientes, 10% reprobados
    const random = Math.random();
    let cumulative = 0;
    
    for (let i = 0; i < statuses.length; i++) {
      cumulative += weights[i];
      if (random <= cumulative) {
        return statuses[i];
      }
    }
    return 'pending';
  }

  function getRandomGrade(): number {
    return Math.round((Math.random() * 2 + 3) * 10) / 10; // Entre 3.0 y 5.0
  }

  const handleProgramSelect = (programId: number) => {
    setSelectedProgramId(programId);
  };

  if (programsLoading) {
    return <LoadingSkeleton />;
  }

  if (!selectedProgramId) {
    return (
      <div className="space-y-6">
        <div className="flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Semáforo Académico</h1>
            <p className="text-gray-600 mt-2">Selecciona un programa para visualizar el progreso académico</p>
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {programsData?.data.map((program) => (
            <Card 
              key={program.id} 
              className="bg-white border border-gray-200 shadow-lg hover:shadow-xl transition-all duration-300 hover:-translate-y-1 cursor-pointer"
              onClick={() => handleProgramSelect(program.id)}
            >
              <CardHeader className="pb-3">
                <div className="flex items-center gap-3">
                  <div className="p-3 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-lg">
                    <GraduationCap className="h-6 w-6 text-white" />
                  </div>
                  <div className="flex-1">
                    <CardTitle className="text-lg font-semibold text-gray-900">
                      {program.name}
                    </CardTitle>
                    <p className="text-sm text-gray-600 mt-1">
                      {program.higher_education_institution.name}
                    </p>
                  </div>
                </div>
              </CardHeader>
              <CardContent className="pt-0">
                <div className="space-y-3">
                  <div className="flex items-center justify-between">
                    <span className="text-sm text-gray-600">SNIES:</span>
                    <Badge className="bg-blue-100 text-blue-800 border-blue-200">
                      {program.snies}
                    </Badge>
                  </div>
                  <div className="flex items-center justify-between">
                    <span className="text-sm text-gray-600">Créditos:</span>
                    <Badge className="bg-indigo-100 text-indigo-800 border-indigo-200">
                      {program.total_credits}
                    </Badge>
                  </div>
                  <div className="flex items-center justify-between">
                    <span className="text-sm text-gray-600">Duración:</span>
                    <span className="text-sm font-medium text-gray-800">
                      {program.duration_value} {program.duration_unit.name}
                    </span>
                  </div>
                </div>
                <div className="mt-4 pt-4 border-t border-gray-100">
                  <Button 
                    className="w-full bg-gradient-to-r from-green-600 to-green-700 hover:from-green-700 hover:to-green-800"
                    onClick={(e) => {
                      e.stopPropagation();
                      handleProgramSelect(program.id);
                    }}
                  >
                    <TrendingUp className="h-4 w-4 mr-2" />
                    Ver Semáforo
                  </Button>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>

        {programsData?.data.length === 0 && (
          <Card className="text-center py-12">
            <CardContent>
              <BookOpen className="h-12 w-12 mx-auto mb-4 text-gray-400" />
              <h3 className="text-lg font-medium text-gray-900 mb-2">
                No hay programas registrados
              </h3>
              <p className="text-gray-600 mb-4">
                Los programas aparecerán aquí conforme se vayan registrando
              </p>
              <Button 
                onClick={() => router.push("/director/registrar-programa")}
                className="bg-blue-600 hover:bg-blue-700"
              >
                Registrar Programa
              </Button>
            </CardContent>
          </Card>
        )}
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Button 
          variant="ghost" 
          onClick={() => setSelectedProgramId(null)}
          className="text-gray-600 hover:text-gray-800"
        >
          <ArrowLeft className="w-4 h-4 mr-2" />
          Volver a Programas
        </Button>
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Semáforo Académico</h1>
          <p className="text-gray-600 mt-1">
            {selectedProgram?.name} - {selectedProgram?.higher_education_institution.name}
          </p>
        </div>
      </div>

      {coursesLoading ? (
        <div className="space-y-6">
          <Skeleton className="h-32 w-full" />
          <Skeleton className="h-24 w-full" />
          <Skeleton className="h-96 w-full" />
        </div>
      ) : (
        <AcademicSemaphore
          courses={coursesWithStatus}
          programName={selectedProgram?.name || ""}
          totalCredits={selectedProgram?.total_credits || 0}
        />
      )}
    </div>
  );
}

function LoadingSkeleton() {
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <Skeleton className="h-8 w-64 mb-2" />
          <Skeleton className="h-4 w-96" />
        </div>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {Array.from({ length: 6 }).map((_, i) => (
          <Card key={i} className="p-6">
            <div className="flex items-center gap-3 mb-4">
              <Skeleton className="h-12 w-12 rounded-lg" />
              <div className="flex-1">
                <Skeleton className="h-5 w-3/4 mb-2" />
                <Skeleton className="h-4 w-1/2" />
              </div>
            </div>
            <div className="space-y-3">
              <div className="flex justify-between">
                <Skeleton className="h-4 w-16" />
                <Skeleton className="h-6 w-20" />
              </div>
              <div className="flex justify-between">
                <Skeleton className="h-4 w-20" />
                <Skeleton className="h-6 w-16" />
              </div>
              <div className="flex justify-between">
                <Skeleton className="h-4 w-18" />
                <Skeleton className="h-4 w-24" />
              </div>
            </div>
            <Skeleton className="h-10 w-full mt-4" />
          </Card>
        ))}
      </div>
    </div>
  );
}

