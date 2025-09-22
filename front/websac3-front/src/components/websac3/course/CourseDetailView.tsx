"use client";

import React, { useState, useCallback } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { 
  BookOpen, 
  Clock, 
  GraduationCap, 
  Shield, 
  ShieldCheck, 
  User, 
  Calendar,
  Edit,
  ArrowLeft,
  ChevronLeft,
  ChevronRight
} from "lucide-react";
import { CourseItem, useGetCourseTopicsQuery } from "@/services/api";

interface CourseDetailViewProps {
  course: CourseItem | undefined;
  programName: string;
  onEdit?: () => void;
  onBack?: () => void;
  isLoading?: boolean;
}

export function CourseDetailView({ 
  course, 
  programName, 
  onEdit, 
  onBack, 
  isLoading = false 
}: CourseDetailViewProps) {
  // Pagination state
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(5);

  // Fetch course topics
  const { data: courseTopics, isLoading: topicsLoading, error: topicsError } = useGetCourseTopicsQuery(
    { course_id: course?.id || 0 },
    { skip: !course?.id }
  );

  // Pagination logic
  const totalTopics = courseTopics?.length || 0;
  const totalPages = Math.ceil(totalTopics / itemsPerPage);
  const startIndex = (currentPage - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;
  const currentTopics = courseTopics?.slice(startIndex, endIndex) || [];

  const handlePageChange = useCallback((page: number) => {
    setCurrentPage(page);
  }, []);

  const handleItemsPerPageChange = useCallback((e: React.ChangeEvent<HTMLSelectElement>) => {
    const items = parseInt(e.target.value);
    setItemsPerPage(items);
    setCurrentPage(1); // Reset to first page immediately
  }, []);


  if (isLoading || !course) {
    return <CourseDetailSkeleton />;
  }

  const getCourseTypeInfo = () => {
    if (course.is_cybersecurity) {
      return {
        label: 'Ciberseguridad',
        color: 'bg-gradient-to-br from-red-500 to-red-600',
        textColor: 'text-white',
        icon: Shield,
        description: 'Curso especializado en ciberseguridad'
      };
    }
    if (course.contains_cybersecurity_topics) {
      return {
        label: 'Con Tópicos CS',
        color: 'bg-gradient-to-br from-purple-500 to-purple-600',
        textColor: 'text-white',
        icon: ShieldCheck,
        description: 'Contiene tópicos de ciberseguridad'
      };
    }
    return {
      label: 'Regular',
      color: 'bg-gradient-to-br from-blue-500 to-blue-600',
      textColor: 'text-white',
      icon: BookOpen,
      description: 'Curso regular del programa'
    };
  };

  const courseTypeInfo = getCourseTypeInfo();
  const Icon = courseTypeInfo.icon;

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          {onBack && (
            <Button 
              variant="ghost" 
              onClick={onBack}
              className="text-gray-600 hover:text-gray-800"
            >
              <ArrowLeft className="w-4 h-4 mr-2" />
              Volver
            </Button>
          )}
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Detalle del Curso</h1>
            <p className="text-gray-600 mt-1">{programName}</p>
          </div>
        </div>
        {onEdit && (
          <Button 
            onClick={onEdit}
            className="bg-blue-600 hover:bg-blue-700"
          >
            <Edit className="w-4 h-4 mr-2" />
            Editar Curso
          </Button>
        )}
      </div>

      {/* Course Type Badge */}
      <Card className="bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200">
        <CardContent className="p-6">
          <div className="flex items-center gap-4">
            <div className={`p-4 rounded-xl ${courseTypeInfo.color}`}>
              <Icon className="h-8 w-8 text-white" />
            </div>
            <div>
              <h2 className="text-2xl font-bold text-gray-900">{course.name}</h2>
              <p className="text-lg text-gray-600">{course.code}</p>
              <Badge className={`mt-2 ${courseTypeInfo.color} ${courseTypeInfo.textColor} border-0`}>
                {courseTypeInfo.label}
              </Badge>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Course Information Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {/* Basic Information */}
        <Card>
          <CardHeader className="pb-3">
            <CardTitle className="text-lg flex items-center gap-2">
              <BookOpen className="h-5 w-5 text-blue-600" />
              Información Básica
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">Código</p>
              <p className="text-lg font-semibold text-gray-800">{course.code}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">Nombre</p>
              <p className="text-lg font-semibold text-gray-800">{course.name}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">Créditos</p>
              <p className="text-lg font-semibold text-gray-800">{course.credits}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">Período</p>
              <p className="text-lg font-semibold text-gray-800">Período {course.period_number}</p>
            </div>
          </CardContent>
        </Card>

        {/* Course Characteristics */}
        <Card>
          <CardHeader className="pb-3">
            <CardTitle className="text-lg flex items-center gap-2">
              <GraduationCap className="h-5 w-5 text-green-600" />
              Características
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">Tipo</p>
              <Badge variant="secondary" className="bg-blue-100 text-blue-800">
                {course.type_name}
              </Badge>
            </div>
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">Naturaleza</p>
              <Badge variant="secondary" className="bg-green-100 text-green-800">
                {course.nature_name}
              </Badge>
            </div>
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">¿Es de Ciberseguridad?</p>
              <Badge 
                variant="secondary" 
                className={course.is_cybersecurity ? "bg-red-100 text-red-800" : "bg-gray-100 text-gray-800"}
              >
                {course.is_cybersecurity ? "Sí" : "No"}
              </Badge>
            </div>
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">¿Contiene tópicos CS?</p>
              <Badge 
                variant="secondary" 
                className={course.contains_cybersecurity_topics ? "bg-purple-100 text-purple-800" : "bg-gray-100 text-gray-800"}
              >
                {course.contains_cybersecurity_topics ? "Sí" : "No"}
              </Badge>
            </div>
          </CardContent>
        </Card>

        {/* Program Information */}
        <Card>
          <CardHeader className="pb-3">
            <CardTitle className="text-lg flex items-center gap-2">
              <Calendar className="h-5 w-5 text-purple-600" />
              Información del Programa
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">Programa</p>
              <p className="text-lg font-semibold text-gray-800">{course.degree_program_name}</p>
            </div>
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">Creado por</p>
              <div className="flex items-center gap-2">
                <User className="h-4 w-4 text-gray-500" />
                <p className="text-lg font-semibold text-gray-800">{course.creator_name}</p>
              </div>
            </div>
            <div className="space-y-2">
              <p className="text-sm text-gray-500 font-medium">ID del Curso</p>
              <p className="text-sm font-mono text-gray-600">#{course.id}</p>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Course Type Description */}
      <Card>
        <CardHeader className="pb-3">
          <CardTitle className="text-lg flex items-center gap-2">
            <Icon className="h-5 w-5 text-blue-600" />
            Descripción del Tipo de Curso
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className={`p-4 rounded-lg ${courseTypeInfo.color} ${courseTypeInfo.textColor}`}>
            <div className="flex items-center gap-3 mb-2">
              <Icon className="h-6 w-6" />
              <h3 className="text-lg font-semibold">{courseTypeInfo.label}</h3>
            </div>
            <p className="text-sm opacity-90">{courseTypeInfo.description}</p>
          </div>
        </CardContent>
      </Card>

      {/* Course Topics */}
      <Card>
        <CardHeader className="pb-3">
          <div className="flex items-center justify-between">
            <CardTitle className="text-lg flex items-center gap-2">
              <BookOpen className="h-5 w-5 text-purple-600" />
              Tópicos del Curso
            </CardTitle>
            {courseTopics && courseTopics.length > 0 && (
              <div className="flex items-center gap-2">
                <span className="text-sm text-gray-600">Por página:</span>
                <select
                  value={itemsPerPage}
                  onChange={handleItemsPerPageChange}
                  className="h-8 px-2 py-1 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-white text-sm"
                >
                  <option value={3}>3</option>
                  <option value={5}>5</option>
                  <option value={10}>10</option>
                  <option value={20}>20</option>
                </select>
              </div>
            )}
          </div>
        </CardHeader>
        <CardContent>
          {topicsLoading ? (
            <div className="space-y-3">
              {Array.from({ length: itemsPerPage }).map((_, i) => (
                <div key={`skeleton-${i}`} className="flex items-center justify-between p-3 border border-gray-200 rounded-lg">
                  <div className="flex-1">
                    <div className="h-4 bg-gray-200 rounded w-3/4 mb-2 animate-pulse" />
                    <div className="h-3 bg-gray-200 rounded w-1/2 animate-pulse" />
                  </div>
                  <div className="h-6 bg-gray-200 rounded w-16 animate-pulse" />
                </div>
              ))}
            </div>
          ) : courseTopics && courseTopics.length > 0 ? (
            <>
              <div className="space-y-3">
                {currentTopics.map((topic, index) => (
                  <div key={`topic-${topic.id}-${index}`} className="flex items-center justify-between p-3 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors">
                    <div className="flex-1">
                      <h4 className="text-sm font-medium text-gray-800">{topic.topic_name}</h4>
                      <p className="text-xs text-gray-500">{topic.knowledge_area_name}</p>
                    </div>
                    <Badge variant="secondary" className="bg-purple-100 text-purple-800">
                      {topic.study_hours} horas
                    </Badge>
                  </div>
                ))}
              </div>

              {/* Pagination */}
              <div className="flex items-center justify-between pt-4 border-t border-gray-200 mt-4">
                <div className="text-sm text-gray-600">
                  Mostrando {startIndex + 1} a {Math.min(endIndex, totalTopics)} de {totalTopics} tópicos
                </div>
                <div className="flex items-center gap-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => handlePageChange(currentPage - 1)}
                    disabled={currentPage === 1}
                    className="h-8 w-8 p-0"
                  >
                    <ChevronLeft className="h-4 w-4" />
                  </Button>
                  
                  {/* Page numbers */}
                  {totalPages > 0 && (() => {
                    const pages = [];
                    const maxVisible = Math.min(5, totalPages);
                    
                    if (totalPages <= 5) {
                      for (let i = 1; i <= totalPages; i++) {
                        pages.push(i);
                      }
                    } else if (currentPage <= 3) {
                      for (let i = 1; i <= 5; i++) {
                        pages.push(i);
                      }
                    } else if (currentPage >= totalPages - 2) {
                      for (let i = totalPages - 4; i <= totalPages; i++) {
                        pages.push(i);
                      }
                    } else {
                      for (let i = currentPage - 2; i <= currentPage + 2; i++) {
                        pages.push(i);
                      }
                    }
                    
                    return pages.map((pageNumber) => (
                      <Button
                        key={`page-${pageNumber}`}
                        variant={currentPage === pageNumber ? "default" : "outline"}
                        size="sm"
                        onClick={() => handlePageChange(pageNumber)}
                        className="h-8 w-8 p-0"
                      >
                        {pageNumber}
                      </Button>
                    ));
                  })()}
                  
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => handlePageChange(currentPage + 1)}
                    disabled={currentPage >= totalPages}
                    className="h-8 w-8 p-0"
                  >
                    <ChevronRight className="h-4 w-4" />
                  </Button>
                </div>
              </div>
            </>
          ) : (
            <div className="text-center py-8 text-gray-500">
              <BookOpen className="h-8 w-8 mx-auto mb-2 text-gray-400" />
              <p>No hay tópicos asociados a este curso</p>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}

function CourseDetailSkeleton() {
  return (
    <div className="space-y-6">
      {/* Header Skeleton */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          <Skeleton className="h-10 w-24" />
          <div>
            <Skeleton className="h-8 w-64 mb-2" />
            <Skeleton className="h-4 w-48" />
          </div>
        </div>
        <Skeleton className="h-10 w-32" />
      </div>

      {/* Course Type Badge Skeleton */}
      <Card>
        <CardContent className="p-6">
          <div className="flex items-center gap-4">
            <Skeleton className="h-16 w-16 rounded-xl" />
            <div className="flex-1">
              <Skeleton className="h-6 w-3/4 mb-2" />
              <Skeleton className="h-5 w-1/2 mb-2" />
              <Skeleton className="h-6 w-24" />
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Information Grid Skeleton */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {Array.from({ length: 3 }).map((_, i) => (
          <Card key={`skeleton-card-${i}`}>
            <CardHeader className="pb-3">
              <Skeleton className="h-6 w-32" />
            </CardHeader>
            <CardContent className="space-y-4">
              {Array.from({ length: 4 }).map((_, j) => (
                <div key={`skeleton-field-${i}-${j}`} className="space-y-2">
                  <Skeleton className="h-4 w-20" />
                  <Skeleton className="h-5 w-full" />
                </div>
              ))}
            </CardContent>
          </Card>
        ))}
      </div>

      {/* Description Skeleton */}
      <Card>
        <CardHeader className="pb-3">
          <Skeleton className="h-6 w-48" />
        </CardHeader>
        <CardContent>
          <Skeleton className="h-20 w-full" />
        </CardContent>
      </Card>
    </div>
  );
}
