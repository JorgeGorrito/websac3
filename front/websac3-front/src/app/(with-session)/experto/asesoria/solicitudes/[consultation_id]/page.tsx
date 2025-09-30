"use client";

import React, { useState, useCallback } from "react";
import { useRouter, useParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { 
  MessageSquare, 
  BookOpen, 
  Clock, 
  GraduationCap, 
  Calendar,
  User,
  FileText,
  CheckCircle,
  AlertTriangle,
  Eye,
  Mail,
  ArrowLeft,
  Send,
  X
} from "lucide-react";
import { useListPendingExpertConsultationsQuery, useAcceptExpertConsultationMutation, useRejectExpertConsultationMutation } from "@/services/api";
import { useDispatch } from "react-redux";
import { showError } from "@/store/errorSlice";

interface PendingConsultation {
  id: number;
  requester_id: number;
  requester_name: string;
  requester_email: string;
  requester_institution_snies: number;
  requester_institution_name: string;
  requester_institution_ownership: string;
  requester_job_position: string;
  degree_program_id: number;
  degree_program_name: string;
  degree_program_snies: number;
  report_id: number;
  report_score: number;
  request_message: string;
  status_id: number;
  status_name: string;
  created_at: string;
  updated_at: string;
}

export default function ConsultationDetailPage() {
  const router = useRouter();
  const params = useParams();
  const dispatch = useDispatch();
  const consultationId = parseInt(params.consultation_id as string);
  
  const [expertResponse, setExpertResponse] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  // Get the consultation data
  const { data: consultationsData, isLoading, error, refetch } = useListPendingExpertConsultationsQuery({
    current_page: 1,
    items_per_page: 100,
    lang: 'es'
  });

  const [acceptConsultation] = useAcceptExpertConsultationMutation();
  const [rejectConsultation] = useRejectExpertConsultationMutation();

  // Find the specific consultation
  const consultation = consultationsData?.data?.find(c => c.id === consultationId);

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("es-ES", {
      year: "numeric",
      month: "long",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit"
    });
  };

  const formatScore = (score: number) => {
    const percentage = Math.floor(score * 10000) / 100;
    return percentage % 1 === 0 ? percentage.toFixed(0) : percentage.toFixed(2);
  };

  const handleAccept = useCallback(async () => {
    if (!expertResponse.trim()) {
      dispatch(showError({ type: 'error', message: 'Por favor escribe una respuesta antes de aceptar la consulta.' }));
      return;
    }

    setIsSubmitting(true);
    try {
      await acceptConsultation({
        consultation_id: consultationId,
        body: { expert_response: expertResponse.trim() }
      }).unwrap();

      dispatch(showError({ type: 'success', message: 'Consulta aceptada exitosamente.' }));
      router.push('/experto/asesoria/solicitudes');
    } catch (error: any) {
      const errorMessage = error?.data?.errors?.[0] || 'Error al aceptar la consulta';
      dispatch(showError({ type: 'error', message: errorMessage }));
    } finally {
      setIsSubmitting(false);
    }
  }, [consultationId, expertResponse, acceptConsultation, dispatch, router]);

  const handleReject = useCallback(async () => {
    if (!expertResponse.trim()) {
      dispatch(showError({ type: 'error', message: 'Por favor escribe una respuesta antes de rechazar la consulta.' }));
      return;
    }

    setIsSubmitting(true);
    try {
      await rejectConsultation({
        consultation_id: consultationId,
        body: { expert_response: expertResponse.trim() }
      }).unwrap();

      dispatch(showError({ type: 'success', message: 'Consulta rechazada exitosamente.' }));
      router.push('/experto/asesoria/solicitudes');
    } catch (error: any) {
      const errorMessage = error?.data?.errors?.[0] || 'Error al rechazar la consulta';
      dispatch(showError({ type: 'error', message: errorMessage }));
    } finally {
      setIsSubmitting(false);
    }
  }, [consultationId, expertResponse, rejectConsultation, dispatch, router]);

  if (isLoading) {
    return (
      <div className="space-y-6">
        <div className="flex items-center gap-4">
          <Skeleton className="h-10 w-10" />
          <Skeleton className="h-8 w-64" />
        </div>
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div className="lg:col-span-2">
            <Card>
              <CardHeader>
                <Skeleton className="h-6 w-48" />
              </CardHeader>
              <CardContent className="space-y-4">
                <Skeleton className="h-32 w-full" />
                <Skeleton className="h-24 w-full" />
              </CardContent>
            </Card>
          </div>
          <div>
            <Card>
              <CardHeader>
                <Skeleton className="h-6 w-32" />
              </CardHeader>
              <CardContent className="space-y-4">
                <Skeleton className="h-4 w-full" />
                <Skeleton className="h-4 w-3/4" />
                <Skeleton className="h-4 w-1/2" />
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    );
  }

  if (error || !consultation) {
    return (
      <div className="space-y-6">
        <div className="flex items-center gap-4">
          <Button
            variant="outline"
            onClick={() => router.push('/experto/asesoria/solicitudes')}
            className="flex items-center gap-2"
          >
            <ArrowLeft className="h-4 w-4" />
            Volver
          </Button>
          <h1 className="text-3xl font-bold text-gray-900">Consulta no encontrada</h1>
        </div>
        
        <Card>
          <CardContent className="p-6">
            <div className="text-center">
              <AlertTriangle className="h-12 w-12 text-red-500 mx-auto mb-4" />
              <h2 className="text-xl font-semibold text-red-600 mb-2">Error al cargar la consulta</h2>
              <p className="text-gray-600 mb-4">
                {error ? 'Ocurrió un error al cargar los datos de la consulta.' : 'La consulta solicitada no fue encontrada.'}
              </p>
              <Button onClick={() => router.push('/experto/asesoria/solicitudes')}>
                Volver a Solicitudes
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center gap-4">
        <Button
          variant="outline"
          onClick={() => router.push('/experto/asesoria/solicitudes')}
          className="flex items-center gap-2"
        >
          <ArrowLeft className="h-4 w-4" />
          Volver
        </Button>
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Solicitud #{consultation.id}</h1>
          <p className="text-gray-600 mt-1">Revisa y responde a la solicitud de asesoría</p>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Main Content */}
        <div className="lg:col-span-2 space-y-6">
          {/* Request Details */}
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <CardTitle className="flex items-center gap-2">
                  <MessageSquare className="h-5 w-5 text-blue-600" />
                  Detalles de la Solicitud
                </CardTitle>
                <Badge className="bg-orange-100 text-orange-800">
                  {consultation.status_name}
                </Badge>
              </div>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div className="flex items-center space-x-2">
                  <GraduationCap className="h-4 w-4 text-gray-500" />
                  <div>
                    <p className="text-sm font-medium">{consultation.degree_program_name}</p>
                    <p className="text-xs text-gray-500">SNIES: {consultation.degree_program_snies}</p>
                  </div>
                </div>
                
                <div className="flex items-center space-x-2">
                  <FileText className="h-4 w-4 text-gray-500" />
                  <div>
                    <p className="text-sm">Reporte #{consultation.report_id}</p>
                    <p className="text-xs text-gray-500">Puntaje: {formatScore(consultation.report_score)}%</p>
                  </div>
                </div>
                
                <div className="flex items-center space-x-2">
                  <Calendar className="h-4 w-4 text-gray-500" />
                  <p className="text-sm">{formatDate(consultation.created_at)}</p>
                </div>
              </div>

              {/* Request Message */}
              <div className="bg-gray-50 rounded-lg p-4">
                <h4 className="text-sm font-medium text-gray-700 mb-2 flex items-center gap-1">
                  <MessageSquare className="h-4 w-4" />
                  Mensaje de Solicitud
                </h4>
                <p className="text-sm text-gray-600 whitespace-pre-wrap">
                  {consultation.request_message}
                </p>
              </div>
            </CardContent>
          </Card>

          {/* Expert Response */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Send className="h-5 w-5 text-green-600" />
                Tu Respuesta
              </CardTitle>
              <CardDescription>
                Escribe tu respuesta detallada a la solicitud de asesoría
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="expert-response" className="text-sm font-medium text-gray-700">
                  Respuesta del Experto *
                </Label>
                <Textarea
                  id="expert-response"
                  placeholder="Escribe tu respuesta detallada aquí..."
                  value={expertResponse}
                  onChange={(e) => setExpertResponse(e.target.value)}
                  rows={8}
                  className="w-full"
                />
              </div>
              
              <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
                <div className="flex items-start gap-3">
                  <MessageSquare className="h-5 w-5 text-blue-600 mt-0.5" />
                  <div className="text-sm text-blue-800">
                    <p className="font-medium mb-2">Recomendaciones para tu respuesta:</p>
                    <ul className="space-y-1 text-xs">
                      <li>• Sé específico y detallado en tus recomendaciones</li>
                      <li>• Incluye referencias a mejores prácticas de ciberseguridad</li>
                      <li>• Proporciona ejemplos prácticos cuando sea posible</li>
                      <li>• Sugiere recursos adicionales si es apropiado</li>
                      <li>• Mantén un tono profesional y constructivo</li>
                    </ul>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Sidebar */}
        <div className="space-y-6">
          {/* Requester Information */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <User className="h-5 w-5 text-blue-600" />
                Información del Solicitante
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-3">
                <div>
                  <p className="text-sm font-medium text-gray-700">Nombre</p>
                  <p className="text-sm text-gray-600">{consultation.requester_name}</p>
                </div>
                
                <div>
                  <p className="text-sm font-medium text-gray-700">Email</p>
                  <p className="text-sm text-gray-600 flex items-center gap-1">
                    <Mail className="h-3 w-3" />
                    {consultation.requester_email}
                  </p>
                </div>
                
                <div>
                  <p className="text-sm font-medium text-gray-700">Cargo</p>
                  <p className="text-sm text-gray-600">{consultation.requester_job_position}</p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Institution Information */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <GraduationCap className="h-5 w-5 text-green-600" />
                Institución Educativa
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-3">
                <div>
                  <p className="text-sm font-medium text-gray-700">Nombre</p>
                  <p className="text-sm text-gray-600">{consultation.requester_institution_name}</p>
                </div>
                
                <div>
                  <p className="text-sm font-medium text-gray-700">SNIES</p>
                  <p className="text-sm text-gray-600">{consultation.requester_institution_snies}</p>
                </div>
                
                <div>
                  <p className="text-sm font-medium text-gray-700">Tipo</p>
                  <p className="text-sm text-gray-600">{consultation.requester_institution_ownership}</p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Actions */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <CheckCircle className="h-5 w-5 text-purple-600" />
                Acciones
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
              <Button
                onClick={handleAccept}
                disabled={isSubmitting || !expertResponse.trim()}
                className="w-full bg-green-600 hover:bg-green-700"
              >
                {isSubmitting ? (
                  <>
                    <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></div>
                    Procesando...
                  </>
                ) : (
                  <>
                    <CheckCircle className="h-4 w-4 mr-2" />
                    Aceptar Consulta
                  </>
                )}
              </Button>
              
              <Button
                onClick={handleReject}
                disabled={isSubmitting || !expertResponse.trim()}
                variant="outline"
                className="w-full border-red-300 text-red-600 hover:bg-red-50"
              >
                {isSubmitting ? (
                  <>
                    <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-red-600 mr-2"></div>
                    Procesando...
                  </>
                ) : (
                  <>
                    <X className="h-4 w-4 mr-2" />
                    Rechazar Consulta
                  </>
                )}
              </Button>
              
              <div className="text-xs text-gray-500 text-center">
                Ambas acciones requieren una respuesta escrita
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
