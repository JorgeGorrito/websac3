"use client";

import React, { useRef, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { useGetReportDetailQuery, useSubmitReportFeedbackMutation, useGetReportFeedbackQuery } from '@/services/api';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';
import { Textarea } from '@/components/ui/textarea';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { ArrowLeft, Download, FileText, MessageSquare, Send, Star, AlertTriangle, CheckCircle, Target, Lightbulb } from 'lucide-react';
import { useDispatch } from 'react-redux';
import { showError, showSuccess } from '@/store/errorSlice';
import html2canvas from 'html2canvas';
import jsPDF from 'jspdf';
import ReportHTMLView from '@/components/websac3/report/ReportHTMLView';
import ReportFeedbackView from '@/components/websac3/report/ReportFeedbackView';

export default function ExpertReportReviewPage() {
  const params = useParams();
  const router = useRouter();
  const dispatch = useDispatch();
  const reportId = parseInt(params.report_id as string);
  const reportRef = useRef<HTMLDivElement>(null);
  
  const [feedback, setFeedback] = useState({
    general_comments: '',
    recommendations: '',
    knowledge_area_feedbacks: [] as Array<{
      knowledge_area_report_id: number;
      comments: string;
    }>
  });

  // Fetch report detail
  const { data: reportDetail, isLoading, error } = useGetReportDetailQuery(reportId);
  
  // Check if feedback already exists
  const { data: existingFeedback, isLoading: isLoadingFeedback } = useGetReportFeedbackQuery(reportId);
  
  // Submit feedback mutation
  const [submitFeedback, { isLoading: isSubmittingFeedback }] = useSubmitReportFeedbackMutation();

  const downloadReportAsPDF = async () => {
    if (!reportRef.current) return;

    try {
      const canvas = await html2canvas(reportRef.current, {
        scale: 2,
        useCORS: true,
        allowTaint: true,
        backgroundColor: '#ffffff'
      });

      const imgData = canvas.toDataURL('image/png');
      const pdf = new jsPDF('p', 'mm', 'a4');
      
      const imgWidth = 210;
      const pageHeight = 295;
      const imgHeight = (canvas.height * imgWidth) / canvas.width;
      let heightLeft = imgHeight;

      let position = 0;

      pdf.addImage(imgData, 'PNG', 0, position, imgWidth, imgHeight);
      heightLeft -= pageHeight;

      while (heightLeft >= 0) {
        position = heightLeft - imgHeight;
        pdf.addPage();
        pdf.addImage(imgData, 'PNG', 0, position, imgWidth, imgHeight);
        heightLeft -= pageHeight;
      }

      pdf.save(`reporte-evaluacion-${reportId}.pdf`);
    } catch (error) {
      console.error('Error generating PDF:', error);
      dispatch(showError({
        type: 'error',
        message: 'Error al generar el PDF del reporte',
        errors: ['No se pudo descargar el reporte. Inténtalo de nuevo.']
      }));
    }
  };

  const handleSubmitFeedback = async () => {
    // Validar campos requeridos
    if (!feedback.general_comments.trim()) {
      dispatch(showError({
        type: 'error',
        message: 'Comentarios generales requeridos',
        errors: ['Debes escribir comentarios generales antes de enviar la retroalimentación.']
      }));
      return;
    }

    if (!feedback.recommendations.trim()) {
      dispatch(showError({
        type: 'error',
        message: 'Recomendaciones requeridas',
        errors: ['Debes escribir recomendaciones antes de enviar la retroalimentación.']
      }));
      return;
    }

    // Validar que al menos un área de conocimiento tenga feedback
    const hasKnowledgeAreaFeedback = feedback.knowledge_area_feedbacks.some(
      area => area.comments.trim()
    );

    if (!hasKnowledgeAreaFeedback) {
      dispatch(showError({
        type: 'error',
        message: 'Feedback de áreas de conocimiento requerido',
        errors: ['Debes proporcionar retroalimentación para al menos una área de conocimiento.']
      }));
      return;
    }
    
    try {
      const payload = {
        report_id: reportId,
        general_comments: feedback.general_comments,
        recommendations: feedback.recommendations,
        knowledge_area_feedbacks: feedback.knowledge_area_feedbacks.filter(
          area => area.comments.trim()
        )
      };

      const result = await submitFeedback(payload).unwrap();
      
      // Mostrar mensaje de éxito
      dispatch(showSuccess({
        message: result.message || 'Tu retroalimentación ha sido enviada correctamente.',
        title: 'Retroalimentación enviada exitosamente'
      }));
      
      // Limpiar el formulario
      setFeedback({
        general_comments: '',
        recommendations: '',
        knowledge_area_feedbacks: []
      });
      
      // Opcional: redirigir de vuelta a la lista
      // router.push('/experto/reportes/pendientes');
      
    } catch (error: any) {
      console.error('Error submitting feedback:', error);
      
      // Manejar errores de la API
      const errorMessage = error?.data?.errors?.[0] || 
                          error?.data?.message || 
                          'No se pudo enviar tu retroalimentación. Inténtalo de nuevo.';
      
      dispatch(showError({
        type: 'error',
        message: 'Error al enviar retroalimentación',
        errors: [errorMessage]
      }));
    }
  };

  const updateKnowledgeAreaFeedback = (index: number, value: string) => {
    const updatedFeedbacks = [...feedback.knowledge_area_feedbacks];
    if (!updatedFeedbacks[index]) {
      updatedFeedbacks[index] = {
        knowledge_area_report_id: reportDetail.knowledge_area_reports[index]?.id || 0,
        comments: ''
      };
    }
    updatedFeedbacks[index].comments = value;
    setFeedback({ ...feedback, knowledge_area_feedbacks: updatedFeedbacks });
  };

  const getScoreColor = (score: number) => {
    const percentage = score * 100;
    if (percentage >= 80) return 'text-green-600';
    if (percentage >= 60) return 'text-yellow-600';
    return 'text-red-600';
  };

  const getScoreLabel = (score: number) => {
    const percentage = score * 100;
    if (percentage >= 80) return 'Excelente';
    if (percentage >= 60) return 'Bueno';
    return 'Necesita Mejora';
  };

  const formatScore = (score: number) => {
    const percentage = Math.floor(score * 10000) / 100;
    return percentage.toFixed(2);
  };

  if (isLoading) {
    return (
      <div className="min-h-screen bg-gray-100 p-10">
        <div className="max-w-6xl mx-auto">
          <div className="mb-6">
            <Skeleton className="h-10 w-48 mb-4" />
            <Skeleton className="h-8 w-96" />
          </div>
          <div className="bg-white rounded-lg shadow-lg p-6">
            <Skeleton className="h-96 w-full" />
          </div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-100 p-10">
        <div className="max-w-6xl mx-auto">
          <div className="mb-6">
            <Button
              onClick={() => router.back()}
              variant="outline"
              className="mb-4"
            >
              <ArrowLeft className="h-4 w-4 mr-2" />
              Volver a la lista de reportes
            </Button>
          </div>
          <div className="bg-white rounded-lg shadow-lg p-12">
            <div className="text-center">
              <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
              <h3 className="text-lg font-medium text-gray-900 mb-2">
                Error al cargar el reporte
              </h3>
              <p className="text-gray-600">
                No se pudo cargar el reporte solicitado. Inténtalo de nuevo más tarde.
              </p>
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 p-10">
      <div className="max-w-6xl mx-auto">
        {/* Header with Navigation */}
        <div className="mb-6">
          <Button
            onClick={() => router.back()}
            variant="outline"
            className="mb-4"
          >
            <ArrowLeft className="h-4 w-4 mr-2" />
            Volver a reportes pendientes
          </Button>
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-2xl font-bold text-gray-900">
                Revisar Reporte de Evaluación
              </h1>
              <p className="text-gray-600">
                {reportDetail.degree_program?.name} - {reportDetail.professional_role?.name}
              </p>
              <div className="flex items-center gap-4 mt-2">
                <div className="flex items-center gap-2">
                  <Star className="h-4 w-4 text-yellow-500" />
                  <span className={`font-medium ${getScoreColor(reportDetail.score)}`}>
                    Puntaje: {formatScore(reportDetail.score)}%
                  </span>
                </div>
                <span className={`px-2 py-1 rounded-full text-xs font-medium ${
                  getScoreColor(reportDetail.score) === 'text-green-600' ? 'bg-green-100 text-green-800' :
                  getScoreColor(reportDetail.score) === 'text-yellow-600' ? 'bg-yellow-100 text-yellow-800' :
                  'bg-red-100 text-red-800'
                }`}>
                  {getScoreLabel(reportDetail.score)}
                </span>
              </div>
            </div>
            <Button
              onClick={downloadReportAsPDF}
              className="bg-blue-600 hover:bg-blue-700"
            >
              <Download className="h-4 w-4 mr-2" />
              Descargar PDF
            </Button>
          </div>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Report Content */}
          <div className="lg:col-span-2">
            <div ref={reportRef} className="bg-white rounded-lg shadow-lg overflow-hidden">
              <ReportHTMLView reportDetail={reportDetail} />
            </div>
          </div>

          {/* Feedback Panel */}
          <div className="lg:col-span-1">
            {isLoadingFeedback ? (
              <Card className="sticky top-6">
                <CardContent className="p-6">
                  <div className="space-y-4">
                    <Skeleton className="h-6 w-48" />
                    <Skeleton className="h-4 w-full" />
                    <Skeleton className="h-4 w-3/4" />
                    <Skeleton className="h-20 w-full" />
                  </div>
                </CardContent>
              </Card>
            ) : existingFeedback ? (
              // Show existing feedback
              <ReportFeedbackView reportId={reportId} />
            ) : (
              // Show feedback creation form
              <Card className="sticky top-6">
                <CardHeader>
                  <CardTitle className="flex items-center gap-2">
                    <MessageSquare className="h-5 w-5 text-blue-600" />
                    Retroalimentación del Experto
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <Tabs defaultValue="general" className="w-full">
                    <TabsList className="grid w-full grid-cols-3">
                      <TabsTrigger value="general" className="text-xs">General</TabsTrigger>
                      <TabsTrigger value="areas" className="text-xs">Áreas</TabsTrigger>
                      <TabsTrigger value="info" className="text-xs">Info</TabsTrigger>
                    </TabsList>
                    
                    <TabsContent value="general" className="space-y-4 mt-4">
                      <div>
                        <Label htmlFor="general-comments" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                          <MessageSquare className="h-4 w-4 text-blue-500" />
                          Comentarios Generales *
                        </Label>
                        <Textarea
                          id="general-comments"
                          placeholder="Escribe tus comentarios generales sobre el reporte de evaluación..."
                          value={feedback.general_comments}
                          onChange={(e) => setFeedback({ ...feedback, general_comments: e.target.value })}
                          className="mt-2 min-h-[120px] resize-none"
                          disabled={isSubmittingFeedback}
                        />
                        <p className="text-xs text-gray-500 mt-1">
                          {feedback.general_comments.length}/1000 caracteres
                        </p>
                      </div>
                      
                      <div>
                        <Label htmlFor="recommendations" className="text-sm font-medium text-gray-700 flex items-center gap-2">
                          <Lightbulb className="h-4 w-4 text-yellow-500" />
                          Recomendaciones *
                        </Label>
                        <Textarea
                          id="recommendations"
                          placeholder="Proporciona recomendaciones específicas para mejorar el programa de ciberseguridad..."
                          value={feedback.recommendations}
                          onChange={(e) => setFeedback({ ...feedback, recommendations: e.target.value })}
                          className="mt-2 min-h-[120px] resize-none"
                          disabled={isSubmittingFeedback}
                        />
                        <p className="text-xs text-gray-500 mt-1">
                          {feedback.recommendations.length}/1000 caracteres
                        </p>
                      </div>
                    </TabsContent>
                    
                    <TabsContent value="areas" className="space-y-4 mt-4">
                      <div className="text-sm text-gray-600 mb-4">
                        Proporciona comentarios específicos para cada área de conocimiento evaluada.
                      </div>
                      
                      {reportDetail.knowledge_area_reports?.map((area, index) => (
                        <Card key={area.id} className="border border-gray-200">
                          <CardHeader className="pb-3">
                            <CardTitle className="text-sm font-medium text-gray-800 flex items-center gap-2">
                              <Target className="h-4 w-4 text-green-500" />
                              {area.name}
                            </CardTitle>
                            <div className="flex items-center gap-2 text-xs text-gray-500">
                              <span>Puntaje: {formatScore(area.score_got)}%</span>
                              <span>•</span>
                              <span>Esperado: {formatScore(area.score_expected)}%</span>
                            </div>
                          </CardHeader>
                          <CardContent>
                            <div>
                              <Label className="text-xs font-medium text-gray-600">
                                Comentarios sobre esta área
                              </Label>
                              <Textarea
                                placeholder="Escribe tus comentarios específicos sobre esta área de conocimiento, incluyendo observaciones, fortalezas, debilidades y sugerencias de mejora..."
                                value={feedback.knowledge_area_feedbacks[index]?.comments || ''}
                                onChange={(e) => updateKnowledgeAreaFeedback(index, e.target.value)}
                                className="mt-2 min-h-[100px] text-xs resize-none"
                                disabled={isSubmittingFeedback}
                              />
                              <p className="text-xs text-gray-500 mt-1">
                                {(feedback.knowledge_area_feedbacks[index]?.comments || '').length}/500 caracteres
                              </p>
                            </div>
                          </CardContent>
                        </Card>
                      ))}
                    </TabsContent>
                    
                    <TabsContent value="info" className="space-y-4 mt-4">
                      <div className="space-y-3">
                        <h4 className="text-sm font-medium text-gray-700 flex items-center gap-2">
                          <FileText className="h-4 w-4 text-blue-500" />
                          Información del Reporte
                        </h4>
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
                        </div>
                      </div>
                    </TabsContent>
                  </Tabs>
                  
                  <div className="pt-4 border-t border-gray-200 mt-4">
                    <Button
                      onClick={handleSubmitFeedback}
                      disabled={isSubmittingFeedback}
                      className="w-full bg-green-600 hover:bg-green-700"
                    >
                      {isSubmittingFeedback ? (
                        <>
                          <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                          Enviando...
                        </>
                      ) : (
                        <>
                          <Send className="h-4 w-4 mr-2" />
                          Enviar Retroalimentación
                        </>
                      )}
                    </Button>
                  </div>
                </CardContent>
              </Card>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
