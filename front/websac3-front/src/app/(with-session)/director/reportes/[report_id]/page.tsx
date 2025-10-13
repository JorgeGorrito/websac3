"use client";

import React from 'react';
import { useParams, useRouter } from 'next/navigation';
import { useGetReportDetailQuery } from '@/services/api';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';
import { ArrowLeft, Download, FileText } from 'lucide-react';
import { useDispatch } from 'react-redux';
import { showError } from '@/store/errorSlice';
import jsPDF from 'jspdf';
import ReportHTMLView from '@/components/websac3/report/ReportHTMLView';
import ReportFeedbackView from '@/components/websac3/report/ReportFeedbackView';

export default function ReportDetailPage() {
  const params = useParams();
  const router = useRouter();
  const dispatch = useDispatch();
  const reportId = parseInt(params.report_id as string);

  // Fetch report detail
  const { data: reportDetail, isLoading, error } = useGetReportDetailQuery(reportId);

  const downloadReportAsPDF = () => {
    if (!reportDetail) {
      dispatch(showError('No hay datos del reporte disponibles.'));
      return;
    }

    try {
      const doc = new jsPDF();
      const pageWidth = doc.internal.pageSize.getWidth();
      const pageHeight = doc.internal.pageSize.getHeight();
      let yPos = 20;

      // Título
      doc.setFontSize(18);
      doc.setFont('helvetica', 'bold');
      doc.text('Reporte de Evaluación', pageWidth / 2, yPos, { align: 'center' });
      yPos += 10;

      // Subtítulo
      doc.setFontSize(10);
      doc.setFont('helvetica', 'normal');
      doc.text('Resultado de la evaluación del componente de ciberseguridad', pageWidth / 2, yPos, { align: 'center' });
      yPos += 15;

      // Información institucional
      doc.setFontSize(12);
      doc.setFont('helvetica', 'bold');
      doc.text('Información Institucional', 20, yPos);
      yPos += 7;
      
      doc.setFontSize(10);
      doc.setFont('helvetica', 'normal');
      doc.text(`Institución: ${reportDetail.degree_program?.higher_education_institution?.name || 'N/A'}`, 20, yPos);
      yPos += 6;
      doc.text(`SNIES Institución: ${reportDetail.degree_program?.higher_education_institution?.snies || 'N/A'}`, 20, yPos);
      yPos += 6;
      doc.text(`Fecha: ${new Date(reportDetail.created_at).toLocaleDateString('es-ES')}`, 20, yPos);
      yPos += 10;

      // Información del programa
      doc.setFont('helvetica', 'bold');
      doc.text('Información del Programa', 20, yPos);
      yPos += 7;
      
      doc.setFont('helvetica', 'normal');
      doc.text(`Programa: ${reportDetail.degree_program?.name || 'N/A'}`, 20, yPos);
      yPos += 6;
      doc.text(`SNIES Programa: ${reportDetail.degree_program?.snies || 'N/A'}`, 20, yPos);
      yPos += 6;
      doc.text(`Rol Profesional: ${reportDetail.professional_role?.name || 'N/A'}`, 20, yPos);
      yPos += 6;
      
      // Puntaje total - TEXTO HORIZONTAL
      const percentage = Math.round(reportDetail.score * 100);
      doc.setFontSize(14);
      doc.setFont('helvetica', 'bold');
      doc.text(`Puntaje Total: ${percentage}%`, 20, yPos);
      yPos += 10;

      // Áreas de conocimiento
      reportDetail.knowledge_area_reports?.forEach((area, index) => {
        if (yPos > pageHeight - 40) {
          doc.addPage();
          yPos = 20;
        }

        doc.setFontSize(12);
        doc.setFont('helvetica', 'bold');
        doc.text(`Área ${index + 1}: ${area.name}`, 20, yPos);
        yPos += 7;

        const areaPercentage = Math.round((area.score_got / area.score_expected) * 100);
        doc.setFontSize(10);
        doc.setFont('helvetica', 'normal');
        doc.text(`Porcentaje: ${areaPercentage}%`, 25, yPos);
        yPos += 6;
        doc.text(`Horas esperadas: ${area.total_learn_hours_expected.toFixed(2)}`, 25, yPos);
        yPos += 6;
        doc.text(`Horas alcanzadas: ${area.total_learn_hours_actual.toFixed(2)}`, 25, yPos);
        yPos += 10;

        // Tópicos
        if (area.topic_reports && area.topic_reports.length > 0) {
          doc.setFont('helvetica', 'bold');
          doc.text('Temáticas:', 25, yPos);
          yPos += 6;

          area.topic_reports.forEach((topic) => {
            if (yPos > pageHeight - 20) {
              doc.addPage();
              yPos = 20;
            }

            doc.setFont('helvetica', 'normal');
            const topicText = `• ${topic.name}: ${topic.learn_hours_expected.toFixed(2)}h esperadas, ${topic.learn_hours_actual.toFixed(2)}h alcanzadas`;
            const splitText = doc.splitTextToSize(topicText, pageWidth - 50);
            doc.text(splitText, 30, yPos);
            yPos += 5 * splitText.length;
          });
        }

        yPos += 5;
      });

      // Generar nombre de archivo
      const programName = reportDetail.degree_program?.name || 'Programa';
      const date = new Date(reportDetail.created_at).toLocaleDateString('es-ES');
      const filename = `Reporte_${programName.replace(/[^a-zA-Z0-9]/g, '_')}_${date.replace(/\//g, '-')}.pdf`;

      // Descargar
      doc.save(filename);
      console.log('PDF generado exitosamente');

    } catch (error) {
      console.error('Error generando PDF:', error);
      dispatch(showError(`Error al generar el PDF: ${error.message}`));
    }
  };

  if (isLoading) {
    return (
      <div className="min-h-screen bg-gray-100 p-10">
        <div className="max-w-4xl mx-auto">
          <div className="mb-6">
            <Button
              onClick={() => router.back()}
              variant="outline"
              className="mb-4"
            >
              <ArrowLeft className="h-4 w-4 mr-2" />
              Volver
            </Button>
          </div>
          <Card>
            <CardContent className="p-12">
              <div className="space-y-4">
                <Skeleton className="h-8 w-3/4" />
                <Skeleton className="h-4 w-1/2" />
                <div className="space-y-2">
                  <Skeleton className="h-4 w-full" />
                  <Skeleton className="h-4 w-full" />
                  <Skeleton className="h-4 w-3/4" />
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    );
  }

  if (error || !reportDetail) {
    return (
      <div className="min-h-screen bg-gray-100 p-10">
        <div className="max-w-4xl mx-auto">
          <div className="mb-6">
            <Button
              onClick={() => router.back()}
              variant="outline"
              className="mb-4"
            >
              <ArrowLeft className="h-4 w-4 mr-2" />
              Volver
            </Button>
          </div>
          <Card>
            <CardContent className="p-12">
              <div className="text-center text-gray-500">
                <FileText className="h-16 w-16 text-gray-300 mx-auto mb-4" />
                <h3 className="text-lg font-medium text-gray-900 mb-2">
                  Error al cargar el reporte
                </h3>
                <p className="text-gray-600 mb-4">
                  No se pudo cargar el reporte solicitado. Inténtalo de nuevo más tarde.
                </p>
                <Button onClick={() => router.back()}>
                  Volver a la lista
                </Button>
              </div>
            </CardContent>
          </Card>
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
            Volver a la lista de reportes
          </Button>
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-2xl font-bold text-gray-900">
                Reporte de Evaluación
              </h1>
              <p className="text-gray-600">
                {reportDetail.degree_program?.name} - {reportDetail.professional_role?.name}
              </p>
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
            <div className="bg-white rounded-lg shadow-lg overflow-hidden">
              <ReportHTMLView reportDetail={reportDetail} />
            </div>
          </div>

          {/* Expert Feedback */}
          <div className="lg:col-span-1">
            <ReportFeedbackView reportId={reportId} />
          </div>
        </div>
      </div>
    </div>
  );
}