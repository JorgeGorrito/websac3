"use client";

import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { BookOpen, Shield, ShieldCheck, FileText, GraduationCap } from "lucide-react";

export interface Course {
  id: number;
  code: string;
  name: string;
  credits: number;
  period_number: number;
  type_id: number;
  nature_id: number;
  is_cybersecurity: boolean;
  contains_cybersecurity_topics: boolean;
  nature_name: string;
  type_name: string;
  degree_program_name: string;
  created_by: number;
  creator_name: string;
}

interface AcademicSemaphoreProps {
  courses: Course[];
  programName: string;
  totalCredits: number;
}

type CourseType = 'cybersecurity' | 'with_topics' | 'regular';

const courseTypeConfig = {
  cybersecurity: {
    label: 'Ciberseguridad',
    color: 'bg-gradient-to-br from-red-500 to-red-600',
    textColor: 'text-white',
    icon: Shield,
    description: 'Curso especializado en ciberseguridad'
  },
  with_topics: {
    label: 'Con Tópicos CS',
    color: 'bg-gradient-to-br from-purple-500 to-purple-600',
    textColor: 'text-white',
    icon: ShieldCheck,
    description: 'Contiene tópicos de ciberseguridad'
  },
  regular: {
    label: 'Regular',
    color: 'bg-gradient-to-br from-blue-500 to-blue-600',
    textColor: 'text-white',
    icon: BookOpen,
    description: 'Curso regular del programa'
  }
};

// Ya no necesitamos estos mapeos porque ahora recibimos los nombres directamente del API

export function AcademicSemaphore({ courses, programName, totalCredits }: AcademicSemaphoreProps) {

  // Agrupar cursos por período
  const coursesByPeriod = courses.reduce((acc, course) => {
    const period = course.period_number;
    if (!acc[period]) {
      acc[period] = [];
    }
    acc[period].push(course);
    return acc;
  }, {} as Record<number, Course[]>);

  // Obtener todos los períodos y ordenarlos
  const periods = Object.keys(coursesByPeriod)
    .map(Number)
    .sort((a, b) => a - b);

  // Calcular estadísticas de ciberseguridad
  const totalCourses = courses.length;
  const cybersecurityCourses = courses.filter(c => c.is_cybersecurity).length;
  const coursesWithTopics = courses.filter(c => c.contains_cybersecurity_topics && !c.is_cybersecurity).length;
  const regularCourses = courses.filter(c => !c.is_cybersecurity && !c.contains_cybersecurity_topics).length;

  // Calcular créditos registrados
  const registeredCredits = courses.reduce((sum, c) => sum + c.credits, 0);
  const progressPercentage = totalCredits > 0 ? (registeredCredits / totalCredits) * 100 : 0;

  // Función para determinar el tipo de curso
  const getCourseType = (course: Course): CourseType => {
    if (course.is_cybersecurity) return 'cybersecurity';
    if (course.contains_cybersecurity_topics) return 'with_topics';
    return 'regular';
  };

  return (
    <div className="space-y-6">
      {/* Header con estadísticas */}
      <Card className="bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200">
        <CardHeader className="pb-4">
          <CardTitle className="text-2xl font-bold text-gray-900 flex items-center gap-3">
            <GraduationCap className="h-8 w-8 text-blue-600" />
            Cursos del Programa
          </CardTitle>
          <div className="flex flex-wrap items-center gap-4 text-sm text-gray-600">
            <span className="font-semibold">{programName}</span>
            <span>•</span>
            <span>{totalCredits} créditos totales</span>
            <span>•</span>
            <span>{registeredCredits} créditos registrados</span>
            <span>•</span>
            <span>{Math.round(progressPercentage)}% registrado</span>
          </div>
        </CardHeader>
        <CardContent className="pt-0">
          {/* Barra de progreso de créditos */}
          <div className="w-full bg-gray-200 rounded-full h-3 mb-4">
            <div 
              className="bg-gradient-to-r from-blue-500 to-indigo-500 h-3 rounded-full transition-all duration-500"
              style={{ width: `${progressPercentage}%` }}
            />
          </div>
          
          {/* Estadísticas por tipo de curso */}
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
            <div className="text-center">
              <div className="text-2xl font-bold text-blue-600">{totalCourses}</div>
              <div className="text-sm text-gray-600">Total Cursos</div>
            </div>
            <div className="text-center">
              <div className="text-2xl font-bold text-red-600">{cybersecurityCourses}</div>
              <div className="text-sm text-gray-600">Ciberseguridad</div>
            </div>
            <div className="text-center">
              <div className="text-2xl font-bold text-purple-600">{coursesWithTopics}</div>
              <div className="text-sm text-gray-600">Con Tópicos CS</div>
            </div>
            <div className="text-center">
              <div className="text-2xl font-bold text-green-600">{regularCourses}</div>
              <div className="text-sm text-gray-600">Regulares</div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Leyenda de tipos de cursos */}
      <Card>
        <CardHeader className="pb-3">
          <CardTitle className="text-lg">Tipos de Cursos</CardTitle>
        </CardHeader>
        <CardContent className="pt-0">
          <div className="flex flex-wrap gap-4">
            {Object.entries(courseTypeConfig).map(([type, config]) => {
              const Icon = config.icon;
              return (
                <div key={type} className="flex items-center gap-2">
                  <div className={`w-4 h-4 rounded ${config.color}`} />
                  <Icon className="h-4 w-4 text-gray-600" />
                  <span className="text-sm font-medium">{config.label}</span>
                  <span className="text-xs text-gray-500">- {config.description}</span>
                </div>
              );
            })}
          </div>
        </CardContent>
      </Card>

      {/* Grid de cursos por período */}
      <Card>
        <CardHeader className="pb-3">
          <CardTitle className="text-lg">Cursos por Período</CardTitle>
        </CardHeader>
        <CardContent className="pt-0">
          {periods.length === 0 ? (
            <div className="text-center py-8 text-gray-500">
              <BookOpen className="h-12 w-12 mx-auto mb-4 text-gray-400" />
              <p className="text-lg font-medium">No hay cursos registrados</p>
              <p className="text-sm">Los cursos aparecerán aquí conforme se vayan registrando</p>
            </div>
          ) : (
            <div className="overflow-x-auto">
              <div className="flex gap-4 min-w-max">
                {periods.map((period) => (
                  <div key={period} className="flex-shrink-0 w-80">
                    <div className="bg-gray-50 rounded-lg p-4 border border-gray-200">
                      <h3 className="text-lg font-semibold text-gray-800 mb-4 text-center">
                        Período {period}
                      </h3>
                      <div className="space-y-3">
                        {coursesByPeriod[period].map((course) => {
                          const courseType = getCourseType(course);
                          const config = courseTypeConfig[courseType];
                          const Icon = config.icon;
                          
                          return (
                            <div
                              key={course.id}
                              className={`${config.color} ${config.textColor} rounded-lg p-3 shadow-sm hover:shadow-md transition-all duration-200 transform hover:scale-105`}
                            >
                              <div className="flex items-start justify-between mb-2">
                                <div className="flex-1">
                                  <div className="flex items-center gap-2 mb-1">
                                    <Icon className="h-4 w-4" />
                                    <span className="text-xs font-medium">{course.code}</span>
                                  </div>
                                  <h4 className="text-sm font-semibold leading-tight mb-1">
                                    {course.name}
                                  </h4>
                                  <div className="flex items-center gap-2 text-xs opacity-90">
                                    <span>{course.credits} créditos</span>
                                    <span>•</span>
                                    <span>{course.type_name}</span>
                                    <span>•</span>
                                    <span>{course.nature_name}</span>
                                  </div>
                                </div>
                                <div className="text-right">
                                  <div className="text-xs opacity-90">
                                    {config.label}
                                  </div>
                                </div>
                              </div>
                              
                              {/* Información adicional */}
                              <div className="flex items-center justify-between mt-2 pt-2 border-t border-white/20">
                                <div className="text-xs opacity-75">
                                  Creado por: {course.creator_name}
                                </div>
                                {course.is_cybersecurity && (
                                  <Badge variant="secondary" className="text-xs bg-white/20 text-white border-white/30">
                                    <Shield className="h-3 w-3 mr-1" />
                                    Especializado
                                  </Badge>
                                )}
                                {course.contains_cybersecurity_topics && !course.is_cybersecurity && (
                                  <Badge variant="secondary" className="text-xs bg-white/20 text-white border-white/30">
                                    <ShieldCheck className="h-3 w-3 mr-1" />
                                    Con Tópicos
                                  </Badge>
                                )}
                              </div>
                            </div>
                          );
                        })}
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}

