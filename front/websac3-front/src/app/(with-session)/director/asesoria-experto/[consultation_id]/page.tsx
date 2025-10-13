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

export default function ConsultationDetailPage() {
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
    if (status.includes('pendiente') || status.includes('pending')) {
      return 'bg-orange-100 text-orange-800';
    } else if (status.includes('aceptada') || status.includes('accepted')) {
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
            onClick={() => router.push('/director/asesoria-experto')}
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
              <Button onClick={() => router.push('/director/asesoria-experto')}>
                Volver a Asesorías
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
          onClick={() => router.push('/director/asesoria-experto')}
          className="flex items-center gap-2"
        >
          <ArrowLeft className="h-4 w-4" />
          Volver
        </Button>
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Solicitud de Asesoría #{consultation.id}</h1>
          <p className="text-gray-600 mt-1">Detalles de la consulta con el experto en ciberseguridad</p>
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

                {consultation.updated_at && consultation.updated_at !== consultation.created_at && (
                  <div className="flex items-center space-x-2">
                    <Calendar className="h-4 w-4 text-gray-500" />
                    <div>
                      <p className="text-xs text-gray-500">Última actualización</p>
                      <p className="text-sm">{formatDate(consultation.updated_at)}</p>
                    </div>
                  </div>
                )}
              </div>

              {/* Request Message */}
              <div className="bg-gray-50 rounded-lg p-4">
                <h4 className="text-sm font-medium text-gray-700 mb-2 flex items-center gap-1">
                  <MessageSquare className="h-4 w-4" />
                  Tu Mensaje de Solicitud
                </h4>
                <p className="text-sm text-gray-600 whitespace-pre-wrap">
                  {consultation.request_message}
                </p>
              </div>

              {/* Expert Response */}
              {consultation.expert_response && (
                <div className="bg-green-50 border border-green-200 rounded-lg p-4">
                  <h4 className="text-sm font-medium text-green-800 mb-2 flex items-center gap-1">
                    <CheckCircle className="h-4 w-4" />
                    Respuesta del Experto
                  </h4>
                  <p className="text-sm text-green-900 whitespace-pre-wrap">
                    {consultation.expert_response}
                  </p>
                </div>
              )}

              {!consultation.expert_response && consultation.status_name.toLowerCase().includes('pendiente') && (
                <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
                  <div className="flex items-start gap-3">
                    <MessageSquare className="h-5 w-5 text-blue-600 mt-0.5" />
                    <div className="text-sm text-blue-800">
                      <p className="font-medium mb-1">En espera de respuesta</p>
                      <p className="text-xs">El experto en ciberseguridad revisará tu solicitud y te responderá pronto.</p>
                    </div>
                  </div>
                </div>
              )}
            </CardContent>
          </Card>

          {/* View Report Button */}
          <Card>
            <CardContent className="p-4">
              <Button
                onClick={() => router.push(`/director/reportes/${consultation.report_id}`)}
                className="w-full"
                variant="outline"
              >
                <FileText className="h-4 w-4 mr-2" />
                Ver Reporte Completo
              </Button>
            </CardContent>
          </Card>
        </div>

        {/* Sidebar */}
        <div className="space-y-6">
          {/* Requester Information (Tu información) */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <User className="h-5 w-5 text-blue-600" />
                Tu Información
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

          {/* Expert Information */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <User className="h-5 w-5 text-purple-600" />
                Experto Asignado
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-3">
                <div>
                  <p className="text-sm font-medium text-gray-700">Nombre</p>
                  <p className="text-sm text-gray-600">{consultation.expert_name}</p>
                </div>
                
                <div>
                  <p className="text-sm font-medium text-gray-700">Email</p>
                  <p className="text-sm text-gray-600 flex items-center gap-1">
                    <Mail className="h-3 w-3" />
                    {consultation.expert_email}
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Status Information */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <CheckCircle className="h-5 w-5 text-purple-600" />
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
                
                {consultation.status_name.toLowerCase().includes('pendiente') && (
                  <p className="text-xs text-gray-500">
                    Tu consulta está siendo revisada por el experto en ciberseguridad.
                  </p>
                )}
                
                {consultation.status_name.toLowerCase().includes('aceptada') && (
                  <p className="text-xs text-green-700 bg-green-50 p-2 rounded">
                    ✓ Tu consulta ha sido respondida. Revisa la respuesta del experto arriba.
                  </p>
                )}
                
                {consultation.status_name.toLowerCase().includes('rechazada') && (
                  <p className="text-xs text-red-700 bg-red-50 p-2 rounded">
                    El experto ha proporcionado una respuesta. Revisa los detalles arriba.
                  </p>
                )}
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
