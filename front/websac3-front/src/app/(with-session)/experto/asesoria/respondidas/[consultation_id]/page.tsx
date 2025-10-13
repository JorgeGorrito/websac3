"use client";

import React from "react";
import { useRouter, useParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { 
  MessageSquare, 
  GraduationCap, 
  Calendar,
  User,
  FileText,
  CheckCircle,
  AlertTriangle,
  ArrowLeft,
  Mail
} from "lucide-react";
import { useGetExpertConsultationDetailQuery } from "@/services/api";

export default function AnsweredConsultationDetailPage() {
  const router = useRouter();
  const params = useParams();
  const consultationId = parseInt(params.consultation_id as string);

  // Get the consultation data
  const { data: consultation, isLoading, error } = useGetExpertConsultationDetailQuery({
    consultation_id: consultationId,
    lang: 'es'
  });

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

  const getStatusBadgeColor = (statusName: string) => {
    const status = statusName.toLowerCase();
    if (status.includes('aceptada') || status.includes('accepted')) {
      return 'bg-green-100 text-green-800';
    } else if (status.includes('rechazada') || status.includes('rejected')) {
      return 'bg-red-100 text-red-800';
    } else if (status.includes('cerrada') || status.includes('closed')) {
      return 'bg-gray-100 text-gray-800';
    }
    return 'bg-blue-100 text-blue-800';
  };

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
            onClick={() => router.push('/experto/asesoria/respondidas')}
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
              <Button onClick={() => router.push('/experto/asesoria/respondidas')}>
                Volver a Solicitudes Respondidas
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
          onClick={() => router.push('/experto/asesoria/respondidas')}
          className="flex items-center gap-2"
        >
          <ArrowLeft className="h-4 w-4" />
          Volver
        </Button>
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Solicitud Respondida #{consultation.id}</h1>
          <p className="text-gray-600 mt-1">Detalles de la consulta respondida</p>
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
                <Badge className={getStatusBadgeColor(consultation.status_name)}>
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
                  <div>
                    <p className="text-xs text-gray-500">Fecha de solicitud</p>
                    <p className="text-sm">{formatDate(consultation.created_at)}</p>
                  </div>
                </div>

                {consultation.answered_at && (
                  <div className="flex items-center space-x-2">
                    <CheckCircle className="h-4 w-4 text-green-500" />
                    <div>
                      <p className="text-xs text-gray-500">Respondida el</p>
                      <p className="text-sm">{formatDate(consultation.answered_at)}</p>
                    </div>
                  </div>
                )}
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

          {/* Expert Response (Read Only) */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <CheckCircle className="h-5 w-5 text-green-600" />
                Tu Respuesta
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="bg-green-50 border border-green-200 rounded-lg p-4">
                <p className="text-sm text-green-900 whitespace-pre-wrap">
                  {consultation.expert_response}
                </p>
              </div>
              
              <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
                <div className="flex items-start gap-3">
                  <CheckCircle className="h-5 w-5 text-blue-600 mt-0.5" />
                  <div className="text-sm text-blue-800">
                    <p className="font-medium mb-1">Consulta respondida</p>
                    <p className="text-xs">
                      Esta consulta ya fue respondida el {consultation.answered_at && formatDate(consultation.answered_at)}. 
                      La información se muestra en modo de solo lectura.
                    </p>
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
              </div>
            </CardContent>
          </Card>

          {/* Status Summary */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <CheckCircle className="h-5 w-5 text-green-600" />
                Estado de la Consulta
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-3">
                <div className="flex items-center justify-between">
                  <span className="text-sm text-gray-600">Estado:</span>
                  <Badge className={getStatusBadgeColor(consultation.status_name)}>
                    {consultation.status_name}
                  </Badge>
                </div>
                
                <div className="text-xs text-green-700 bg-green-50 p-3 rounded-lg">
                  <p className="font-medium mb-1">✓ Consulta completada</p>
                  <p>Has proporcionado una respuesta a esta solicitud de asesoría.</p>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}


