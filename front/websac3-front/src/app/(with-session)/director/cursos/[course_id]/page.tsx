"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { useParams } from "next/navigation";
import { CourseDetailView } from "@/components/websac3/course/CourseDetailView";
import { useGetCourseByIdQuery, useListDegreeProgramsQuery } from "@/services/api";

export default function CourseDetailPage() {
  const router = useRouter();
  const params = useParams();
  const courseId = parseInt(params.course_id as string);

  // Fetch course details
  const { data: course, isLoading: courseLoading, error: courseError } = useGetCourseByIdQuery({
    course_id: courseId,
  });

  // Fetch degree programs to get program name
  const { data: programsData } = useListDegreeProgramsQuery({
    current_page: 1,
    items_per_page: 50,
  });

  const programName = programsData?.data.find(p => p.id === course?.degree_program_id)?.name || "";

  const handleEdit = () => {
    router.push(`/director/cursos/${courseId}/editar`);
  };

  const handleBack = () => {
    router.back();
  };

  if (courseError) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
          <div className="text-center text-gray-500">
            <h2 className="text-xl font-semibold mb-4">Error al cargar el curso</h2>
            <p className="mb-4">No se pudo cargar la información del curso.</p>
            <button 
              onClick={handleBack}
              className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
            >
              Volver
            </button>
          </div>
        </div>
      </div>
    );
  }

  if (!course && !courseLoading) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
          <div className="text-center text-gray-500">
            <h2 className="text-xl font-semibold mb-4">Curso no encontrado</h2>
            <p className="mb-4">El curso solicitado no existe o no tienes permisos para verlo.</p>
            <button 
              onClick={handleBack}
              className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
            >
              Volver
            </button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <CourseDetailView
      course={course}
      programName={programName}
      onEdit={handleEdit}
      onBack={handleBack}
      isLoading={courseLoading}
    />
  );
}
