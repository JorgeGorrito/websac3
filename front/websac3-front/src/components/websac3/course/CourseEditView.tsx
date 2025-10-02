"use client";

import React, { useState, useEffect } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { 
  ArrowLeft, 
  CheckCircle, 
  BookOpen, 
  Clock, 
  GraduationCap, 
  Target, 
  Search, 
  ChevronLeft, 
  ChevronRight,
  Save,
  X,
  Upload
} from "lucide-react";
import { 
  useListTopicsQuery, 
  useListCourseTypesQuery, 
  useListCourseNaturesQuery,
  useGetCourseTopicsQuery,
  useUpdateCourseMutation,
  CourseItem,
  TopicItem,
  CourseTopicItem
} from "@/services/api";
import { useRouter } from "next/navigation";

interface CourseEditViewProps {
  course: CourseItem | undefined;
  programName: string;
  onSave?: () => void;
  onCancel?: () => void;
  isLoading?: boolean;
}

export function CourseEditView({ 
  course, 
  programName, 
  onSave, 
  onCancel, 
  isLoading = false 
}: CourseEditViewProps) {
  const router = useRouter();
  // Topics pagination and search state
  const [topicsCurrentPage, setTopicsCurrentPage] = useState(1);
  const [topicsItemsPerPage, setTopicsItemsPerPage] = useState(10);
  const [topicsSearchTerm, setTopicsSearchTerm] = useState("");

  // Fetch topics with pagination and search
  const { data: topicsData, isLoading: topicsLoading } = useListTopicsQuery({
    current_page: topicsCurrentPage,
    items_per_page: topicsItemsPerPage,
    filters: topicsSearchTerm ? { "name[cont]": topicsSearchTerm } : undefined,
  });

  // Fetch course types and natures
  const { data: courseTypesData, isLoading: courseTypesLoading } = useListCourseTypesQuery({
    current_page: 1,
    items_per_page: 50,
  });

  const { data: courseNaturesData, isLoading: courseNaturesLoading } = useListCourseNaturesQuery({
    current_page: 1,
    items_per_page: 50,
  });

  const [currentStep, setCurrentStep] = useState(1);
  const [formData, setFormData] = useState({
    code: course?.code || "",
    name: course?.name || "",
    credits: course?.credits?.toString() || "",
    period_number: course?.period_number?.toString() || "",
    type_id: course?.type_id?.toString() || "",
    nature_id: course?.nature_id?.toString() || "",
    is_cybersecurity: course?.is_cybersecurity?.toString() || "",
    contains_cybersecurity_topics: course?.contains_cybersecurity_topics?.toString() || "",
    course_topics: [] as Array<{ topic_id: number; study_hours: number }>,
  });

  const [selectedTopics, setSelectedTopics] = useState<Array<{ id: number; name: string; study_hours: number }>>([]);
  const [showSuccessModal, setShowSuccessModal] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);

  // Fetch course topics
  const { data: courseTopics, isLoading: courseTopicsLoading } = useGetCourseTopicsQuery(
    { course_id: course?.id || 0 },
    { skip: !course?.id }
  );

  // Mutation for updating course
  const [updateCourse, { isLoading: isUpdatingCourse }] = useUpdateCourseMutation();

  // Update form data when course is loaded
  useEffect(() => {
    if (course) {
      setFormData({
        code: course.code || "",
        name: course.name || "",
        credits: course.credits?.toString() || "",
        period_number: course.period_number?.toString() || "",
        type_id: course.type_id?.toString() || "",
        nature_id: course.nature_id?.toString() || "",
        is_cybersecurity: course.is_cybersecurity?.toString() || "",
        contains_cybersecurity_topics: course.contains_cybersecurity_topics?.toString() || "",
        course_topics: [],
      });
    }
  }, [course]);

  // Update selected topics when course topics are loaded
  useEffect(() => {
    if (courseTopics && courseTopics.length > 0) {
      const topicsWithHours = courseTopics.map(topic => ({
        id: topic.topic_id,
        name: topic.topic_name,
        study_hours: topic.study_hours
      }));
      setSelectedTopics(topicsWithHours);
    }
  }, [courseTopics]);

  const steps = [
    {
      id: 1,
      title: "Información Básica",
      description: "Datos generales del curso",
    },
    {
      id: 2,
      title: "Características",
      description: "Naturaleza y tipo del curso",
    },
    {
      id: 3,
      title: "Temáticas",
      description: "Tópicos y horas de estudio",
    },
    { 
      id: 4, 
      title: "Revisión", 
      description: "Revisa y confirma los cambios",
    },
  ];

  const handleInputChange = (field: string, value: string | number) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
    
    // If user changes contains_cybersecurity_topics to "false", clear selected topics
    if (field === "contains_cybersecurity_topics" && value === "false") {
      setSelectedTopics([]);
    }
  };

  const handleTopicToggle = (topic: { id: number; name: string }) => {
    setSelectedTopics(prev => {
      const exists = prev.find(t => t.id === topic.id);
      if (exists) {
        return prev.filter(t => t.id !== topic.id);
      } else {
        return [...prev, { ...topic, study_hours: 0 }];
      }
    });
  };

  const handleStudyHoursChange = (topicId: number, hours: number) => {
    setSelectedTopics(prev => 
      prev.map(topic => 
        topic.id === topicId ? { ...topic, study_hours: hours } : topic
      )
    );
  };

  const removeTopic = (topicId: number) => {
    setSelectedTopics(prev => prev.filter(topic => topic.id !== topicId));
  };

  // Topics pagination handlers
  const handleTopicsPageChange = (page: number) => {
    setTopicsCurrentPage(page);
  };

  const handleTopicsItemsPerPageChange = (itemsPerPage: number) => {
    setTopicsItemsPerPage(itemsPerPage);
    setTopicsCurrentPage(1);
  };

  const handleTopicsSearch = (searchTerm: string) => {
    setTopicsSearchTerm(searchTerm);
    setTopicsCurrentPage(1);
  };

  const nextStep = () => {
    if (currentStep < steps.length) {
      if (currentStep === 2 && formData.contains_cybersecurity_topics === "false") {
        setCurrentStep(4);
      } else {
        setCurrentStep(currentStep + 1);
      }
    }
  };

  const prevStep = () => {
    if (currentStep > 1) {
      if (currentStep === 4 && formData.contains_cybersecurity_topics === "false") {
        setCurrentStep(2);
      } else {
        setCurrentStep(currentStep - 1);
      }
    }
  };

  const handleSubmit = async () => {
    if (isSubmitting) return;
    
    setIsSubmitting(true);
    
    try {
      const courseData = {
        id: course.id,
        code: formData.code,
        name: formData.name,
        credits: parseInt(formData.credits),
        period_number: parseInt(formData.period_number),
        type_id: parseInt(formData.type_id),
        nature_id: parseInt(formData.nature_id),
        is_cybersecurity: formData.is_cybersecurity === "true",
        contains_cybersecurity_topics: formData.contains_cybersecurity_topics === "true",
        degree_program_id: course.degree_program_id,
        course_topics: selectedTopics.map(topic => ({
          topic_id: topic.id,
          study_hours: topic.study_hours
        }))
      };

      await updateCourse({ course_id: course.id, data: courseData }).unwrap();
      setShowSuccessModal(true);
    } catch (error) {
      console.error("Error al actualizar el curso:", error);
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleSuccessModalClose = () => {
    setShowSuccessModal(false);
    onSave?.();
  };

  if (isLoading || courseTypesLoading || courseNaturesLoading || courseTopicsLoading || !course) {
    return <CourseEditSkeleton />;
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          {onCancel && (
            <Button 
              variant="ghost" 
              onClick={onCancel}
              className="text-gray-600 hover:text-gray-800"
            >
              <ArrowLeft className="w-4 h-4 mr-2" />
              Cancelar
            </Button>
          )}
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Editar Curso</h1>
            <p className="text-gray-600 mt-1">{programName}</p>
          </div>
        </div>
        <Button
          variant="outline"
          onClick={() => router.push(`/director/carga-masiva-topicos?course_id=${course?.id}`)}
          className="flex items-center gap-2 bg-blue-50 hover:bg-blue-100 text-blue-700 border-blue-300"
        >
          <Upload className="w-4 h-4" />
          Carga Masiva de Tópicos
        </Button>
      </div>

      {/* Course Info Card */}
      <Card className="bg-gradient-to-r from-blue-50 via-indigo-50 to-purple-50 border-blue-200 shadow-lg">
        <CardContent className="p-6">
          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <div className="p-4 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl shadow-lg">
                <GraduationCap className="h-10 w-10 text-white" />
              </div>
              <div className="ml-6">
                <p className="text-sm font-semibold text-blue-700 uppercase tracking-wide mb-1">Editando Curso</p>
                <p className="text-2xl font-bold text-gray-900 mb-1">{course.name}</p>
                <p className="text-sm text-gray-600 mb-3">{course.code}</p>
                <div className="flex items-center gap-3">
                  <Badge className="bg-blue-100 text-blue-800 border-blue-200 font-semibold px-3 py-1">
                    {course.credits} Créditos
                  </Badge>
                  <Badge className="bg-indigo-100 text-indigo-800 border-indigo-200 font-semibold px-3 py-1">
                    Período {course.period_number}
                  </Badge>
                </div>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Steps indicator */}
      <div className="mb-8">
        <div className="flex items-center justify-between">
          {steps.map((step, index) => (
            <div key={step.id} className="flex items-center">
              <div className="flex flex-col items-center">
                <div
                  className={`flex items-center justify-center w-12 h-12 rounded-xl border-2 transition-all duration-200 ${
                    currentStep >= step.id
                      ? "bg-blue-500 border-blue-500 text-white"
                      : "bg-white border-gray-300 text-gray-500"
                  }`}
                >
                  <span className="text-sm font-semibold">{step.id}</span>
                </div>
                <div className="mt-2 text-center">
                  <h3
                    className={`text-sm font-semibold ${
                      currentStep >= step.id
                        ? "text-blue-600"
                        : "text-gray-500"
                    }`}
                  >
                    {step.title}
                  </h3>
                  <p className="text-xs text-gray-400 mt-1">{step.description}</p>
                </div>
              </div>
              {index < steps.length - 1 && (
                <div
                  className={`w-16 h-1 mx-4 rounded-full transition-all duration-200 ${
                    currentStep > step.id ? "bg-blue-500" : "bg-gray-200"
                  }`}
                />
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Form content */}
      <Card className="mb-6 border border-gray-200">
        <CardContent className="p-6">
          {currentStep === 1 && (
            <div className="space-y-4">
              <h3 className="text-lg font-semibold text-gray-800 mb-4">
                Información Básica del Curso
              </h3>

              <div className="grid lg:grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="code" className="text-sm font-medium text-gray-700">
                    Código del Curso *
                  </Label>
                  <Input
                    id="code"
                    value={formData.code}
                    onChange={(e) => handleInputChange("code", e.target.value)}
                    placeholder="Ej: PROG301"
                    className="h-10"
                  />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="credits" className="text-sm font-medium text-gray-700">
                    Número de Créditos *
                  </Label>
                  <Input
                    id="credits"
                    type="number"
                    value={formData.credits}
                    onChange={(e) => handleInputChange("credits", e.target.value)}
                    placeholder="Ej: 3"
                    min="1"
                    max="6"
                    className="h-10"
                  />
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="name" className="text-sm font-medium text-gray-700">
                  Nombre del Curso *
                </Label>
                <Input
                  id="name"
                  value={formData.name}
                  onChange={(e) => handleInputChange("name", e.target.value)}
                  placeholder="Ej: Programación Avanzada"
                  className="h-10"
                />
              </div>

              <div className="space-y-2">
                <Label htmlFor="period_number" className="text-sm font-medium text-gray-700">
                  Número de Período *
                </Label>
                <Input
                  id="period_number"
                  type="number"
                  value={formData.period_number}
                  onChange={(e) => handleInputChange("period_number", e.target.value)}
                  placeholder="Ej: 3"
                  min="1"
                  className="h-10"
                />
              </div>
            </div>
          )}

          {currentStep === 2 && (
            <div className="space-y-4">
              <h3 className="text-lg font-semibold text-gray-800 mb-4">
                Características del Curso
              </h3>

              <div className="grid lg:grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="type_id" className="text-sm font-medium text-gray-700">
                    Tipo de Curso *
                  </Label>
                  <select
                    id="type_id"
                    value={formData.type_id}
                    onChange={(e) => handleInputChange("type_id", e.target.value)}
                    className="w-full h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-white"
                  >
                    <option value="">Selecciona el tipo</option>
                    {courseTypesData?.data.map((courseType) => (
                      <option key={courseType.id} value={courseType.id}>
                        {courseType.name}
                      </option>
                    ))}
                  </select>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="nature_id" className="text-sm font-medium text-gray-700">
                    Naturaleza *
                  </Label>
                  <select
                    id="nature_id"
                    value={formData.nature_id}
                    onChange={(e) => handleInputChange("nature_id", e.target.value)}
                    className="w-full h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-white"
                  >
                    <option value="">Selecciona la naturaleza</option>
                    {courseNaturesData?.data.map((courseNature) => (
                      <option key={courseNature.id} value={courseNature.id}>
                        {courseNature.name}
                      </option>
                    ))}
                  </select>
                </div>
              </div>

              <div className="grid lg:grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="is_cybersecurity" className="text-sm font-medium text-gray-700">
                    ¿Es un curso de ciberseguridad? *
                  </Label>
                  <select
                    id="is_cybersecurity"
                    value={formData.is_cybersecurity}
                    onChange={(e) => handleInputChange("is_cybersecurity", e.target.value)}
                    className="w-full h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-white"
                  >
                    <option value="">Selecciona una opción</option>
                    <option value="true">Sí</option>
                    <option value="false">No</option>
                  </select>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="contains_cybersecurity_topics" className="text-sm font-medium text-gray-700">
                    ¿Contiene tópicos de ciberseguridad? *
                  </Label>
                  <select
                    id="contains_cybersecurity_topics"
                    value={formData.contains_cybersecurity_topics}
                    onChange={(e) => handleInputChange("contains_cybersecurity_topics", e.target.value)}
                    className="w-full h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-white"
                  >
                    <option value="">Selecciona una opción</option>
                    <option value="true">Sí</option>
                    <option value="false">No</option>
                  </select>
                </div>
              </div>
            </div>
          )}

          {currentStep === 3 && (
            <div className="space-y-4">
              <h3 className="text-lg font-semibold text-gray-800 mb-4">
                Temáticas del Curso
              </h3>

              <div className="space-y-4">
                <p className="text-sm text-gray-600">
                  Selecciona los tópicos que cubre este curso y especifica las horas de estudio para cada uno:
                </p>
                
                {/* Search and Filters */}
                <div className="flex flex-col sm:flex-row gap-4 mb-4">
                  <div className="flex-1">
                    <div className="relative">
                      <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" />
                      <Input
                        placeholder="Buscar tópicos..."
                        value={topicsSearchTerm}
                        onChange={(e) => handleTopicsSearch(e.target.value)}
                        className="pl-10 h-10"
                      />
                    </div>
                  </div>
                  <div className="flex items-center gap-2">
                    <Label htmlFor="topics-per-page" className="text-sm text-gray-600 whitespace-nowrap">
                      Por página:
                    </Label>
                    <select
                      id="topics-per-page"
                      value={topicsItemsPerPage}
                      onChange={(e) => handleTopicsItemsPerPageChange(parseInt(e.target.value))}
                      className="h-10 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-white"
                    >
                      <option value={5}>5</option>
                      <option value={10}>10</option>
                      <option value={20}>20</option>
                      <option value={50}>50</option>
                    </select>
                  </div>
                </div>

                {/* Topics List */}
                <div className="space-y-3 max-h-96 overflow-y-auto">
                  {topicsLoading ? (
                    <div className="space-y-3">
                      {Array.from({ length: topicsItemsPerPage }).map((_, i) => (
                        <div key={i} className="flex items-center space-x-3 p-3 border border-gray-200 rounded-lg">
                          <Skeleton className="h-4 w-4" />
                          <div className="flex-1">
                            <Skeleton className="h-4 w-3/4 mb-1" />
                            <Skeleton className="h-3 w-1/2" />
                          </div>
                          <Skeleton className="h-8 w-16" />
                        </div>
                      ))}
                    </div>
                  ) : (
                    <>
                      {topicsData?.data.map((topic) => (
                        <div key={topic.id} className="flex items-center space-x-3 p-3 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors">
                          <input
                            type="checkbox"
                            id={`topic-${topic.id}`}
                            checked={selectedTopics.some(t => t.id === topic.id)}
                            onChange={() => handleTopicToggle(topic)}
                            className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                          />
                          <div className="flex-1">
                            <label htmlFor={`topic-${topic.id}`} className="text-sm font-medium text-gray-700 cursor-pointer">
                              {topic.name}
                            </label>
                            <p className="text-xs text-gray-500">{topic.knowledge_area_name}</p>
                          </div>
                          {selectedTopics.some(t => t.id === topic.id) && (
                            <div className="flex items-center space-x-2">
                              <Label htmlFor={`hours-${topic.id}`} className="text-xs text-gray-600">
                                Horas:
                              </Label>
                              <Input
                                id={`hours-${topic.id}`}
                                type="number"
                                min="0"
                                value={selectedTopics.find(t => t.id === topic.id)?.study_hours || 0}
                                onChange={(e) => handleStudyHoursChange(topic.id, parseInt(e.target.value) || 0)}
                                className="w-16 h-8 text-xs"
                              />
                            </div>
                          )}
                        </div>
                      ))}
                      
                      {topicsData?.data.length === 0 && (
                        <div className="text-center py-8 text-gray-500">
                          <BookOpen className="h-8 w-8 mx-auto mb-2 text-gray-400" />
                          <p>No se encontraron tópicos</p>
                          {topicsSearchTerm && (
                            <p className="text-sm">Intenta con otros términos de búsqueda</p>
                          )}
                        </div>
                      )}
                    </>
                  )}
                </div>

                {/* Pagination */}
                {topicsData && topicsData.total_count > 0 && (
                  <div className="flex items-center justify-between pt-4 border-t border-gray-200">
                    <div className="text-sm text-gray-600">
                      Mostrando {((topicsCurrentPage - 1) * topicsItemsPerPage) + 1} a {Math.min(topicsCurrentPage * topicsItemsPerPage, topicsData.total_count)} de {topicsData.total_count} tópicos
                    </div>
                    <div className="flex items-center gap-2">
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => handleTopicsPageChange(topicsCurrentPage - 1)}
                        disabled={topicsCurrentPage === 1}
                        className="h-8 w-8 p-0"
                      >
                        <ChevronLeft className="h-4 w-4" />
                      </Button>
                      
                      {/* Page numbers */}
                      {Array.from({ length: Math.min(5, Math.ceil(topicsData.total_count / topicsItemsPerPage)) }, (_, i) => {
                        const totalPages = Math.ceil(topicsData.total_count / topicsItemsPerPage);
                        let pageNumber;
                        
                        if (totalPages <= 5) {
                          pageNumber = i + 1;
                        } else if (topicsCurrentPage <= 3) {
                          pageNumber = i + 1;
                        } else if (topicsCurrentPage >= totalPages - 2) {
                          pageNumber = totalPages - 4 + i;
                        } else {
                          pageNumber = topicsCurrentPage - 2 + i;
                        }
                        
                        return (
                          <Button
                            key={pageNumber}
                            variant={topicsCurrentPage === pageNumber ? "default" : "outline"}
                            size="sm"
                            onClick={() => handleTopicsPageChange(pageNumber)}
                            className="h-8 w-8 p-0"
                          >
                            {pageNumber}
                          </Button>
                        );
                      })}
                      
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => handleTopicsPageChange(topicsCurrentPage + 1)}
                        disabled={topicsCurrentPage >= Math.ceil(topicsData.total_count / topicsItemsPerPage)}
                        className="h-8 w-8 p-0"
                      >
                        <ChevronRight className="h-4 w-4" />
                      </Button>
                    </div>
                  </div>
                )}

                {/* Selected Topics Summary */}
                {selectedTopics.length > 0 && (
                  <div className="mt-4 p-4 bg-blue-50 rounded-lg border border-blue-200">
                    <h4 className="text-sm font-semibold text-blue-800 mb-2">
                      Tópicos seleccionados ({selectedTopics.length}):
                    </h4>
                    <div className="space-y-1">
                      {selectedTopics.map((topic) => (
                        <div key={topic.id} className="flex justify-between items-center text-sm">
                          <span className="text-blue-700">{topic.name}</span>
                          <div className="flex items-center gap-2">
                            <div className="flex items-center gap-1">
                              <input
                                type="number"
                                min="1"
                                max="999"
                                value={topic.study_hours}
                                onChange={(e) => handleStudyHoursChange(topic.id, parseInt(e.target.value) || 1)}
                                className="w-16 h-6 px-2 text-xs border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-white"
                              />
                              <span className="text-xs text-gray-500">horas</span>
                            </div>
                            <button
                              onClick={() => removeTopic(topic.id)}
                              className="h-6 w-6 flex items-center justify-center rounded-full hover:bg-red-100 text-gray-400 hover:text-red-600 transition-colors"
                              title="Quitar tópico"
                            >
                              <X className="h-3 w-3" />
                            </button>
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>
                )}
              </div>
            </div>
          )}

          {currentStep === 4 && (
            <div className="space-y-4">
              <h3 className="text-lg font-semibold text-gray-800 mb-4">
                Revisión de Cambios
              </h3>

              <div className="bg-gray-50 rounded-lg p-4 border border-gray-200">
                <div className="grid lg:grid-cols-2 gap-4">
                  <div className="space-y-3">
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium">Código</p>
                      <p className="font-semibold text-gray-800">{formData.code || "No especificado"}</p>
                    </div>
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium">Nombre</p>
                      <p className="font-semibold text-gray-800">{formData.name || "No especificado"}</p>
                    </div>
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium">Créditos</p>
                      <p className="font-semibold text-gray-800">{formData.credits || "No especificado"}</p>
                    </div>
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium">Período</p>
                      <p className="font-semibold text-gray-800">{formData.period_number || "No especificado"}</p>
                    </div>
                  </div>
                  <div className="space-y-3">
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium">Tipo</p>
                      <p className="font-semibold text-gray-800">
                        {courseTypesData?.data.find(type => type.id === parseInt(formData.type_id))?.name || "No especificado"}
                      </p>
                    </div>
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium">Naturaleza</p>
                      <p className="font-semibold text-gray-800">
                        {courseNaturesData?.data.find(nature => nature.id === parseInt(formData.nature_id))?.name || "No especificado"}
                      </p>
                    </div>
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium">¿Es de Ciberseguridad?</p>
                      <p className="font-semibold text-gray-800">
                        {formData.is_cybersecurity === "true" ? "Sí" : formData.is_cybersecurity === "false" ? "No" : "No especificado"}
                      </p>
                    </div>
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium">¿Contiene tópicos de ciberseguridad?</p>
                      <p className="font-semibold text-gray-800">
                        {formData.contains_cybersecurity_topics === "true" ? "Sí" : formData.contains_cybersecurity_topics === "false" ? "No" : "No especificado"}
                      </p>
                    </div>
                  </div>
                </div>
                {selectedTopics.length > 0 && (
                  <div className="mt-4">
                    <div className="bg-white p-3 rounded-lg border border-gray-200">
                      <p className="text-sm text-gray-500 font-medium mb-2">Tópicos seleccionados:</p>
                      <div className="space-y-1">
                        {selectedTopics.map((topic) => (
                          <div key={topic.id} className="flex justify-between items-center text-sm">
                            <span>{topic.name}</span>
                            <div className="flex items-center gap-2">
                              <div className="flex items-center gap-1">
                                <input
                                  type="number"
                                  min="1"
                                  max="999"
                                  value={topic.study_hours}
                                  onChange={(e) => handleStudyHoursChange(topic.id, parseInt(e.target.value) || 1)}
                                  className="w-16 h-6 px-2 text-xs border border-gray-300 rounded focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 bg-white"
                                />
                                <span className="text-xs text-gray-500">horas</span>
                              </div>
                              <button
                                onClick={() => removeTopic(topic.id)}
                                className="h-6 w-6 flex items-center justify-center rounded-full hover:bg-red-100 text-gray-400 hover:text-red-600 transition-colors"
                                title="Quitar tópico"
                              >
                                <X className="h-3 w-3" />
                              </button>
                            </div>
                          </div>
                        ))}
                      </div>
                    </div>
                  </div>
                )}
              </div>
            </div>
          )}
        </CardContent>
      </Card>

      {/* Navigation buttons */}
      <div className="flex justify-between items-center">
        <Button
          variant="outline"
          onClick={prevStep}
          disabled={currentStep === 1}
          className="px-6 py-2 text-gray-700 border-gray-300 hover:bg-gray-50 disabled:opacity-50"
        >
          Anterior
        </Button>

        <div className="flex gap-3">
          {currentStep < steps.length ? (
            <Button 
              onClick={nextStep}
              disabled={
                (currentStep === 1 && (!formData.code || !formData.name || !formData.credits || !formData.period_number)) ||
                (currentStep === 2 && (!formData.type_id || !formData.nature_id || !formData.is_cybersecurity || !formData.contains_cybersecurity_topics)) ||
                (currentStep === 3 && formData.contains_cybersecurity_topics === "true" && selectedTopics.length === 0)
              }
              className="px-6 py-2 bg-blue-600 hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Siguiente
            </Button>
          ) : (
            <Button
              onClick={handleSubmit}
              disabled={isSubmitting || isUpdatingCourse}
              className="px-6 py-2 bg-green-600 hover:bg-green-700 flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {isSubmitting || isUpdatingCourse ? (
                <>
                  <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin" />
                  Guardando...
                </>
              ) : (
                <>
                  <Save className="w-4 h-4" />
                  Guardar Cambios
                </>
              )}
            </Button>
          )}
        </div>
      </div>

      {/* Success Modal */}
      {showSuccessModal && (
        <div 
          className="fixed inset-0 bg-black/50 flex items-center justify-center z-[9999] p-4"
          onClick={handleSuccessModalClose}
        >
          <div 
            className="bg-white rounded-xl shadow-2xl max-w-md w-full mx-4 transform transition-all"
            onClick={(e) => e.stopPropagation()}
          >
            <div className="p-6 text-center">
              <div className="mx-auto flex items-center justify-center h-16 w-16 rounded-full bg-green-100 mb-4">
                <CheckCircle className="h-8 w-8 text-green-600" />
              </div>
              <h3 className="text-lg font-semibold text-gray-900 mb-2">
                ¡Curso Actualizado Exitosamente!
              </h3>
              <p className="text-gray-600 mb-6">
                El curso <strong>{formData.name}</strong> ha sido actualizado correctamente.
              </p>
              <Button
                onClick={handleSuccessModalClose}
                className="px-6 py-2 bg-blue-600 hover:bg-blue-700 text-white"
              >
                Continuar
              </Button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

function CourseEditSkeleton() {
  return (
    <div className="space-y-6">
      {/* Header Skeleton */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          <Skeleton className="h-10 w-24" />
          <div>
            <Skeleton className="h-8 w-48 mb-2" />
            <Skeleton className="h-4 w-32" />
          </div>
        </div>
      </div>

      {/* Course Info Card Skeleton */}
      <Card>
        <CardContent className="p-6">
          <div className="flex items-center">
            <Skeleton className="h-16 w-16 rounded-xl" />
            <div className="ml-6 flex-1">
              <Skeleton className="h-4 w-32 mb-2" />
              <Skeleton className="h-6 w-3/4 mb-2" />
              <Skeleton className="h-4 w-1/2 mb-3" />
              <div className="flex gap-3">
                <Skeleton className="h-6 w-20" />
                <Skeleton className="h-6 w-24" />
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Steps Skeleton */}
      <div className="flex items-center justify-between">
        {Array.from({ length: 4 }).map((_, i) => (
          <div key={i} className="flex flex-col items-center">
            <Skeleton className="h-12 w-12 rounded-xl mb-2" />
            <Skeleton className="h-4 w-20 mb-1" />
            <Skeleton className="h-3 w-16" />
          </div>
        ))}
      </div>

      {/* Form Skeleton */}
      <Card>
        <CardContent className="p-6">
          <Skeleton className="h-6 w-48 mb-4" />
          <div className="space-y-4">
            {Array.from({ length: 4 }).map((_, i) => (
              <div key={i} className="space-y-2">
                <Skeleton className="h-4 w-24" />
                <Skeleton className="h-10 w-full" />
              </div>
            ))}
          </div>
        </CardContent>
      </Card>

      {/* Navigation Skeleton */}
      <div className="flex justify-between">
        <Skeleton className="h-10 w-24" />
        <Skeleton className="h-10 w-32" />
      </div>
    </div>
  );
}
