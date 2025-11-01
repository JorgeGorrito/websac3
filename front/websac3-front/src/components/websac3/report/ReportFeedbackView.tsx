"use client";

import React from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Skeleton } from '@/components/ui/skeleton';
import { 
  MessageSquare, 
  User, 
  Calendar, 
  Target, 
  Lightbulb, 
  AlertTriangle,
  Star,
  CheckCircle,
  FileText
} from 'lucide-react';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { useGetReportFeedbackQuery, useGetReportDetailQuery } from '@/services/api';

interface ReportFeedbackViewProps {
  reportId: number;
}

export default function ReportFeedbackView({ reportId }: ReportFeedbackViewProps) {
  const { data: feedback, isLoading, error } = useGetReportFeedbackQuery(reportId);
  const { data: reportDetail } = useGetReportDetailQuery(reportId);

  const formatScore = (score: number) => {
    const percentage = Math.floor(score * 10000) / 100;
    return percentage % 1 === 0 ? percentage.toFixed(0) : percentage.toFixed(2);
  };

  const getScoreColor = (score: number) => {
    if (score >= 0.9) return 'text-green-600';
    if (score >= 0.8) return 'text-yellow-600';
    return 'text-red-600';
  };

  const getScoreLabel = (score: number) => {
    if (score >= 0.9) return 'Excelente';
    if (score >= 0.8) return 'Aceptable';
    return 'Por mejorar';
  };

  if (isLoading) {
    return (
      <Card className="mt-6">
        <CardHeader>
          <Skeleton className="h-6 w-48" />
        </CardHeader>
        <CardContent className="space-y-4">
          <Skeleton className="h-4 w-full" />
          <Skeleton className="h-4 w-3/4" />
          <Skeleton className="h-20 w-full" />
          <Skeleton className="h-20 w-full" />
        </CardContent>
      </Card>
    );
  }

  if (error) {
    return (
      <Card className="mt-6 border-orange-200 bg-orange-50">
        <CardContent className="p-6">
          <div className="flex items-center gap-3">
            <AlertTriangle className="h-5 w-5 text-orange-500" />
            <div>
              <h3 className="text-sm font-medium text-orange-800">
                Sin retroalimentación disponible
              </h3>
              <p className="text-sm text-orange-600 mt-1">
                Este reporte aún no ha sido revisado por un experto en ciberseguridad.
              </p>
            </div>
          </div>
        </CardContent>
      </Card>
    );
  }

  if (!feedback) {
    return null;
  }

  return (
    <Card className="sticky top-6 border-green-200 bg-green-50">
      <CardHeader className="pb-4">
        <CardTitle className="flex items-center gap-2 text-green-800 text-lg">
          <CheckCircle className="h-5 w-5" />
          Retroalimentación del Experto
        </CardTitle>
        <div className="space-y-2 text-sm text-green-700">
          {feedback.auditor && (
            <div className="flex items-center gap-2">
              <User className="h-4 w-4" />
              <span>Experto: {feedback.auditor.name || 'No disponible'}</span>
            </div>
          )}
          <div className="flex items-center gap-2">
            <Calendar className="h-4 w-4" />
            <span>
              {feedback.created_at ? new Date(feedback.created_at).toLocaleDateString('es-ES', {
                year: 'numeric',
                month: 'short',
                day: 'numeric'
              }) : 'Fecha no disponible'}
            </span>
          </div>
        </div>
      </CardHeader>
      
      <CardContent>
        <Tabs defaultValue="general" className="w-full">
          <TabsList className="grid w-full grid-cols-3">
            <TabsTrigger value="general" className="text-xs">General</TabsTrigger>
            <TabsTrigger value="areas" className="text-xs">Áreas</TabsTrigger>
            <TabsTrigger value="info" className="text-xs">Info</TabsTrigger>
          </TabsList>
          
          <TabsContent value="general" className="space-y-4 mt-4">
            {/* Comentarios Generales */}
            {feedback.general_comments && (
              <div>
                <h4 className="text-sm font-semibold text-gray-800 mb-2 flex items-center gap-2">
                  <MessageSquare className="h-4 w-4 text-blue-500" />
                  Comentarios Generales
                </h4>
                <div className="bg-white rounded-lg p-3 border border-gray-200">
                  <p className="text-sm text-gray-700 leading-relaxed">
                    {feedback.general_comments}
                  </p>
                </div>
              </div>
            )}

            {/* Recomendaciones */}
            {feedback.recommendations && (
              <div>
                <h4 className="text-sm font-semibold text-gray-800 mb-2 flex items-center gap-2">
                  <Lightbulb className="h-4 w-4 text-yellow-500" />
                  Recomendaciones
                </h4>
                <div className="bg-white rounded-lg p-3 border border-gray-200">
                  <p className="text-sm text-gray-700 leading-relaxed">
                    {feedback.recommendations}
                  </p>
                </div>
              </div>
            )}
          </TabsContent>
          
          <TabsContent value="areas" className="space-y-4 mt-4">
            {/* Feedback por Áreas de Conocimiento */}
            {feedback.knowledge_area_feedbacks && feedback.knowledge_area_feedbacks.length > 0 ? (
              <div className="space-y-3">
                {feedback.knowledge_area_feedbacks.map((areaFeedback) => (
                  <Card key={areaFeedback.id} className="border border-gray-200">
                    <CardHeader className="pb-2">
                      <div className="flex items-center justify-between">
                        <CardTitle className="text-xs font-medium text-gray-800 flex items-center gap-1">
                          <Target className="h-3 w-3 text-green-500" />
                          {areaFeedback.knowledge_area_name || 'Área de conocimiento'}
                        </CardTitle>
                      </div>
                    </CardHeader>
                    <CardContent className="pt-0">
                      {areaFeedback.comments && (
                        <div>
                          <h5 className="text-xs font-medium text-gray-600 mb-1">
                            Comentarios
                          </h5>
                          <div className="bg-gray-50 rounded p-2">
                            <p className="text-xs text-gray-700 leading-relaxed">
                              {areaFeedback.comments}
                            </p>
                          </div>
                        </div>
                      )}
                    </CardContent>
                  </Card>
                ))}
              </div>
            ) : (
              <div className="text-center py-8 text-gray-500">
                <Target className="h-8 w-8 mx-auto mb-2 text-gray-400" />
                <p className="text-sm">No hay retroalimentación específica por áreas</p>
              </div>
            )}
          </TabsContent>
          
          <TabsContent value="info" className="space-y-4 mt-4">
            <div className="space-y-3">
              <h4 className="text-sm font-medium text-gray-700 flex items-center gap-2">
                <FileText className="h-4 w-4 text-blue-500" />
                Información del Reporte
              </h4>
              {reportDetail && (
                <div className="space-y-2 text-sm text-gray-600">
                  <div className="flex justify-between">
                    <span>ID del Reporte:</span>
                    <span className="font-medium">#{reportDetail.id}</span>
                  </div>
                  <div className="flex justify-between">
                    <span>Fecha de creación:</span>
                    <span className="font-medium">
                      {new Date(reportDetail.created_at).toLocaleDateString('es-ES')}
                    </span>
                  </div>
                  <div className="flex justify-between">
                    <span>Institución:</span>
                    <span className="font-medium text-right text-xs">
                      {reportDetail.degree_program?.higher_education_institution?.name}
                    </span>
                  </div>
                  <div className="flex justify-between">
                    <span>SNIES:</span>
                    <span className="font-medium">
                      {reportDetail.degree_program?.higher_education_institution?.snies}
                    </span>
                  </div>
                  <div className="flex justify-between">
                    <span>Rol Profesional:</span>
                    <span className="font-medium">
                      {reportDetail.professional_role?.name}
                    </span>
                  </div>
                  <div className="flex justify-between">
                    <span>Puntaje Total:</span>
                    <span className={`font-medium ${getScoreColor(reportDetail.score)}`}>
                      {formatScore(reportDetail.score)}%
                    </span>
                  </div>
                </div>
              )}
            </div>
          </TabsContent>
        </Tabs>
      </CardContent>
    </Card>
  );
}
