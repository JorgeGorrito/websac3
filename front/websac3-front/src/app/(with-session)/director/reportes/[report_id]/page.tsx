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
      console.log('Starting PDF generation with html2canvas...');
      console.log('Report detail:', reportDetail);
      
      // Capture the report component as canvas
      const canvas = await html2canvas(reportRef.current, {
        scale: 1.5, // Reduced scale for better fit
        useCORS: true,
        allowTaint: true,
        backgroundColor: '#f4f4f4',
        width: reportRef.current.scrollWidth,
        height: reportRef.current.scrollHeight,
        onclone: (clonedDoc) => {
          // Override any problematic styles in the cloned document
          const style = clonedDoc.createElement('style');
          style.textContent = `
            * {
              -webkit-print-color-adjust: exact !important;
              color-adjust: exact !important;
            }
            [style*="oklch"] {
              color: rgb(220, 53, 69) !important;
              border-color: rgb(220, 53, 69) !important;
              background-color: rgba(220, 53, 69, 0.1) !important;
            }
          `;
          clonedDoc.head.appendChild(style);
          
          // Completely replace SVG text elements to ensure no transforms
          const svgTexts = clonedDoc.querySelectorAll('svg text');
          svgTexts.forEach((textElement) => {
            const x = textElement.getAttribute('x');
            const y = textElement.getAttribute('y');
            const textContent = textElement.textContent;
            
            // Create a new text element without any transforms
            const newTextElement = clonedDoc.createElementNS('http://www.w3.org/2000/svg', 'text');
            newTextElement.setAttribute('x', x);
            newTextElement.setAttribute('text-anchor', 'middle');
            newTextElement.textContent = textContent;
            
            // Set appropriate positioning and styling
            if (x === '70') {
              // Main circle text - rotate 180 degrees for correct orientation
              newTextElement.setAttribute('y', '75');
              newTextElement.setAttribute('font-size', '20px');
              newTextElement.setAttribute('font-weight', 'bold');
              newTextElement.setAttribute('fill', '#2c3e50');
              newTextElement.setAttribute('dominant-baseline', 'middle');
              newTextElement.setAttribute('transform', 'rotate(180 70 75)');
            } else if (x === '30') {
              // Small circle text - rotate 180 degrees for correct orientation
              newTextElement.setAttribute('y', '30');
              newTextElement.setAttribute('font-size', '10px');
              newTextElement.setAttribute('font-weight', 'bold');
              newTextElement.setAttribute('fill', '#333');
              newTextElement.setAttribute('dominant-baseline', 'middle');
              newTextElement.setAttribute('transform', 'rotate(180 30 30)');
            }
            
            // Replace the old element with the new one
            textElement.parentNode.replaceChild(newTextElement, textElement);
          });
        }
      });
      
      console.log('Canvas generated successfully');
      
      // Create PDF
      const imgData = canvas.toDataURL('image/png');
      const pdf = new jsPDF({
        orientation: 'portrait',
        unit: 'mm',
        format: 'a4'
      });
      
      // Calculate dimensions to fit A4
      const pdfWidth = pdf.internal.pageSize.getWidth();
      const pdfHeight = pdf.internal.pageSize.getHeight();
      const imgWidth = pdfWidth - 10; // 5mm margin on each side
      const imgHeight = (canvas.height * imgWidth) / canvas.width;
      
      let heightLeft = imgHeight;
      let position = 5; // 5mm top margin
      
      // Add first page
      pdf.addImage(imgData, 'PNG', 5, position, imgWidth, imgHeight);
      heightLeft -= pdfHeight - 10; // Account for margins
      
      // Add additional pages if needed
      while (heightLeft >= 0) {
        position = heightLeft - imgHeight + 5; // 5mm margin
        pdf.addPage();
        pdf.addImage(imgData, 'PNG', 5, position, imgWidth, imgHeight);
        heightLeft -= pdfHeight - 10;
      }
      
      // Generate filename
      const programName = reportDetail.degree_program?.name || 'Programa';
      const date = new Date(reportDetail.created_at).toLocaleDateString('es-ES');
      const filename = `Reporte_${programName.replace(/[^a-zA-Z0-9]/g, '_')}_${date.replace(/\//g, '-')}.pdf`;
      
      // Download PDF
      pdf.save(filename);
      
      console.log('PDF download completed successfully');

    } catch (error) {
      console.error('Error generating PDF:', error);
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

        {/* Report Content */}
        <div ref={reportRef} className="bg-white rounded-lg shadow-lg overflow-hidden">
          <ReportHTMLView reportDetail={reportDetail} />
        </div>
      </div>
    </div>
  );
}