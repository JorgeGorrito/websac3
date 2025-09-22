"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { useParams } from "next/navigation";
import { CourseEditView } from "@/components/websac3/course/CourseEditView";
import { useGetCourseByIdQuery, useListDegreeProgramsQuery } from "@/services/api";

export default function CourseEditPage() {
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

  const handleSave = () => {
    // Navigate back to course detail page
    router.push(`/director/cursos/${courseId}`);
  };

  const handleCancel = () => {
    // Navigate back to course detail page
    router.push(`/director/cursos/${courseId}`);
  };

  if (courseError) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
          <div className="text-center text-gray-500">
            <h2 className="text-xl font-semibold mb-4">Error al cargar el curso</h2>
            <p className="mb-4">No se pudo cargar la información del curso para editar.</p>
            <button 
              onClick={() => router.back()}
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
            <p className="mb-4">El curso solicitado no existe o no tienes permisos para editarlo.</p>
            <button 
              onClick={() => router.back()}
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
    <CourseEditView
      course={course}
      programName={programName}
      onSave={handleSave}
      onCancel={handleCancel}
      isLoading={courseLoading}
    />
  );
}
