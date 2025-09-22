"use client";

import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { BookOpen, Shield, ShieldCheck } from "lucide-react";
import { useGetCourseByIdQuery } from "@/services/api";

interface CourseDetailsProps {
  courseId: number;
}

export function CourseDetails({ courseId }: CourseDetailsProps) {
  const { data: course, isLoading, error } = useGetCourseByIdQuery({ course_id: courseId });

  if (isLoading) {
    return (
      <Card>
        <CardContent className="p-6">
          <div className="text-center">Cargando curso...</div>
        </CardContent>
      </Card>
    );
  }

  if (error) {
    return (
      <Card>
        <CardContent className="p-6">
          <div className="text-center text-red-600">Error al cargar el curso</div>
        </CardContent>
      </Card>
    );
  }

  if (!course) {
    return (
      <Card>
        <CardContent className="p-6">
          <div className="text-center text-gray-600">Curso no encontrado</div>
        </CardContent>
      </Card>
    );
  }

  const getCourseTypeColor = () => {
    if (course.is_cybersecurity) return "bg-red-500";
    if (course.contains_cybersecurity_topics) return "bg-purple-500";
    return "bg-blue-500";
  };

  const getCourseTypeIcon = () => {
    if (course.is_cybersecurity) return Shield;
    if (course.contains_cybersecurity_topics) return ShieldCheck;
    return BookOpen;
  };

  const Icon = getCourseTypeIcon();

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <Icon className="h-5 w-5" />
          {course.name}
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="text-sm font-medium text-gray-600">Código:</label>
            <p className="text-lg font-semibold">{course.code}</p>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-600">Créditos:</label>
            <p className="text-lg font-semibold">{course.credits}</p>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-600">Período:</label>
            <p className="text-lg font-semibold">{course.period_number}</p>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-600">Tipo:</label>
            <p className="text-lg font-semibold">{course.type_name}</p>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-600">Naturaleza:</label>
            <p className="text-lg font-semibold">{course.nature_name}</p>
          </div>
          <div>
            <label className="text-sm font-medium text-gray-600">Programa:</label>
            <p className="text-lg font-semibold">{course.degree_program_name}</p>
          </div>
        </div>

        <div className="flex gap-2">
          {course.is_cybersecurity && (
            <Badge className="bg-red-100 text-red-800">
              <Shield className="h-3 w-3 mr-1" />
              Ciberseguridad
            </Badge>
          )}
          {course.contains_cybersecurity_topics && (
            <Badge className="bg-purple-100 text-purple-800">
              <ShieldCheck className="h-3 w-3 mr-1" />
              Con Tópicos CS
            </Badge>
          )}
        </div>

        <div>
          <label className="text-sm font-medium text-gray-600">Creado por:</label>
          <p className="text-sm text-gray-700">{course.creator_name}</p>
        </div>
      </CardContent>
    </Card>
  );
}
