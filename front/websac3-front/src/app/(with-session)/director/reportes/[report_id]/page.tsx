"use client";

import React, { useRef } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { useGetReportDetailQuery } from '@/services/api';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Skeleton } from '@/components/ui/skeleton';
import { ArrowLeft, Download, FileText } from 'lucide-react';
import { useDispatch } from 'react-redux';
import { showError } from '@/store/errorSlice';
import html2canvas from 'html2canvas';
import jsPDF from 'jspdf';
import ReportHTMLView from '@/components/websac3/report/ReportHTMLView';
import ReportFeedbackView from '@/components/websac3/report/ReportFeedbackView';

export default function ReportDetailPage() {
  const params = useParams();
  const router = useRouter();
  const dispatch = useDispatch();
  const reportId = parseInt(params.report_id as string);
  const reportRef = useRef<HTMLDivElement>(null);

  // Fetch report detail
  const { data: reportDetail, isLoading, error } = useGetReportDetailQuery(reportId);

  const downloadReportAsPDF = async () => {
    if (!reportDetail || !reportRef.current) {
      dispatch(showError('No hay datos del reporte disponibles.'));
      return;
    }

    try {
      // Capturar el componente como canvas
      const canvas = await html2canvas(reportRef.current, {
        scale: 1.5,
        useCORS: true,
        allowTaint: true,
        backgroundColor: '#ffffff', // <- blanco para evitar el tono rosado
        width: reportRef.current.scrollWidth,
        height: reportRef.current.scrollHeight,
        onclone: (clonedDoc) => {
          // Reglas mínimas para impresión; sin el selector [style*="oklch"]
          const style = clonedDoc.createElement('style');
          style.textContent = `
            * {
              -webkit-print-color-adjust: exact !important;
              color-adjust: exact !important;
              print-color-adjust: exact !important;
            }
          `;
          clonedDoc.head.appendChild(style);

          // Normalización opcional de <svg> <text> si la necesitas
          const svgTexts = clonedDoc.querySelectorAll('svg text');
          svgTexts.forEach((textElement) => {
            const x = textElement.getAttribute('x');
            const textContent = textElement.textContent || '';
            const newTextElement = clonedDoc.createElementNS('http://www.w3.org/2000/svg', 'text');
            if (x) newTextElement.setAttribute('x', x);
            newTextElement.setAttribute('text-anchor', 'middle');
            newTextElement.textContent = textContent;

            if (x === '70') {
              newTextElement.setAttribute('y', '75');
              newTextElement.setAttribute('font-size', '20px');
              newTextElement.setAttribute('font-weight', 'bold');
              newTextElement.setAttribute('fill', '#2c3e50');
              newTextElement.setAttribute('dominant-baseline', 'middle');
              newTextElement.setAttribute('transform', 'rotate(180 70 75)');
            } else if (x === '30') {
              newTextElement.setAttribute('y', '30');
              newTextElement.setAttribute('font-size', '10px');
              newTextElement.setAttribute('font-weight', 'bold');
              newTextElement.setAttribute('fill', '#333');
              newTextElement.setAttribute('dominant-baseline', 'middle');
              newTextElement.setAttribute('transform', 'rotate(180 30 30)');
            }

            textElement.parentNode?.replaceChild(newTextElement, textElement);
          });
        }
      });

      // Crear PDF
      const imgData = canvas.toDataURL('image/png');
      const pdf = new jsPDF({
        orientation: 'portrait',
        unit: 'mm',
        format: 'a4'
      });

      // Dimensiones para A4 con márgenes de 5mm
      const pdfWidth = pdf.internal.pageSize.getWidth();
      const pdfHeight = pdf.internal.pageSize.getHeight();
      const margin = 5;
      const imgWidth = pdfWidth - margin * 2;
      const imgHeight = (canvas.height * imgWidth) / canvas.width;

      let heightLeft = imgHeight;
      let position = margin;

      // Primera página
      pdf.addImage(imgData, 'PNG', margin, position, imgWidth, imgHeight);
      heightLeft -= (pdfHeight - margin * 2);

      // Páginas adicionales (si hace falta)
      while (heightLeft > 0) {
        pdf.addPage();
        position = heightLeft - imgHeight + margin;
        pdf.addImage(imgData, 'PNG', margin, position, imgWidth, imgHeight);
        heightLeft -= (pdfHeight - margin * 2);
      }

      // Nombre del archivo
      const programName = reportDetail.degree_program?.name || 'Programa';
      const date = new Date(reportDetail.created_at).toLocaleDateString('es-ES');
      const filename = `Reporte_${programName.replace(/[^a-zA-Z0-9]/g, '_')}_${date.replace(/\//g, '-')}.pdf`;

      pdf.save(filename);
    } catch (err: any) {
      console.error('Error generating PDF:', err);
      dispatch(showError(`Error al generar el PDF: ${err?.message || 'desconocido'}`));
    }
  };

  if (isLoading) {
    return (
      <div className="min-h-screen bg-gray-100 p-10">
        <div className="max-w-4xl mx-auto">
          <div className="mb-6">
            <Button onClick={() => router.back()} variant="outline" className="mb-4">
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
            <Button onClick={() => router.back()} variant="outline" className="mb-4">
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
                <Button onClick={() => router.back()}>Volver a la lista</Button>
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
          <Button onClick={() => router.back()} variant="outline" className="mb-4">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Volver a la lista de reportes
          </Button>
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-2xl font-bold text-gray-900">Reporte de Evaluación</h1>
              <p className="text-gray-600">
                {reportDetail.degree_program?.name} - {reportDetail.professional_role?.name}
              </p>
            </div>
            <Button onClick={downloadReportAsPDF} className="bg-blue-600 hover:bg-blue-700">
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

          {/* Expert Feedback */}
          <div className="lg:col-span-1">
            <ReportFeedbackView reportId={reportId} />
          </div>
        </div>
      </div>
    </div>
  );
}

