import React, { useRef } from 'react';
import { usePDF } from 'react-to-pdf';
import { ReportDetail } from '@/services/api';

interface ReportPDFAlternativeProps {
  reportDetail: ReportDetail;
}

const ReportPDFAlternative: React.FC<ReportPDFAlternativeProps> = ({ reportDetail }) => {
  const { toPDF, targetRef } = usePDF({
    filename: `Reporte_${reportDetail.degree_program?.name?.replace(/[^a-zA-Z0-9]/g, '_')}_${new Date(reportDetail.created_at).toLocaleDateString('es-ES').replace(/\//g, '-')}.pdf`,
    page: {
      margin: 20,
      format: 'a4',
      orientation: 'portrait',
    },
    canvas: {
      mimeType: 'image/png',
      qualityRatio: 1,
    },
  });

  const getScoreColor = (score: number) => {
    if (score >= 0.9) return 'text-green-600';
    if (score >= 0.8) return 'text-yellow-600';
    return 'text-red-600';
  };

  const getScoreStatus = (score: number) => {
    if (score >= 0.9) return 'Excelente';
    if (score >= 0.8) return 'Aceptable';
    return 'Por mejorar';
  };

  const percentage = Math.round(reportDetail.score * 100);

  return (
    <div className="space-y-6">
      {/* Download Button */}
      <div className="flex justify-end">
        <button
          onClick={() => toPDF()}
          className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg flex items-center gap-2"
        >
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          Descargar PDF
        </button>
      </div>

      {/* Report Content */}
      <div ref={targetRef} className="bg-white p-8 rounded-lg shadow-lg">
        {/* Header */}
        <div className="text-center mb-8">
          <img 
            src="/websac3/logo-websac3.png" 
            alt="WebSAC3" 
            className="mx-auto mb-4 h-16"
          />
          <h1 className="text-2xl font-bold text-gray-800 mb-2">Reporte de Evaluación</h1>
          <p className="text-gray-600">Resultado de la evaluación del componente de ciberseguridad</p>
        </div>

        {/* Institution and Program Info */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          {/* Institution Card */}
          <div className="bg-gray-50 p-6 rounded-lg border">
            <h3 className="text-lg font-semibold text-gray-800 mb-4">Información Institucional</h3>
            <div className="space-y-3">
              <div>
                <span className="text-sm text-gray-500">SNIES Institución:</span>
                <p className="font-semibold text-gray-800">
                  {reportDetail.degree_program?.higher_education_institution?.snies || 'N/A'}
                </p>
              </div>
              <div>
                <span className="text-sm text-gray-500">Institución:</span>
                <p className="font-semibold text-gray-800">
                  {reportDetail.degree_program?.higher_education_institution?.name || 'No disponible'}
                </p>
              </div>
              <div>
                <span className="text-sm text-gray-500">Fecha del reporte:</span>
                <p className="font-semibold text-gray-800">
                  {new Date(reportDetail.created_at).toLocaleDateString('es-ES')}
                </p>
              </div>
            </div>
          </div>

          {/* Program Card */}
          <div className="bg-gray-50 p-6 rounded-lg border">
            <h3 className="text-lg font-semibold text-gray-800 mb-4">Información del Programa</h3>
            <div className="space-y-3">
              <div>
                <span className="text-sm text-gray-500">SNIES Programa:</span>
                <p className="font-semibold text-gray-800">
                  {reportDetail.degree_program?.snies || 'N/A'}
                </p>
              </div>
              <div>
                <span className="text-sm text-gray-500">Programa:</span>
                <p className="font-semibold text-gray-800">
                  {reportDetail.degree_program?.name || 'N/A'}
                </p>
              </div>
              <div>
                <span className="text-sm text-gray-500">Rol profesional:</span>
                <p className="font-semibold text-gray-800">
                  {reportDetail.professional_role?.name || 'N/A'}
                </p>
              </div>
              <div>
                <span className="text-sm text-gray-500">Puntaje total:</span>
                <p className="font-semibold text-gray-800">
                  {reportDetail.score.toFixed(2)}
                </p>
              </div>
            </div>
          </div>
        </div>

        {/* Overall Score */}
        <div className="text-center mb-8">
          <div className="inline-flex flex-col items-center">
            <div className={`w-32 h-32 rounded-full border-8 border-gray-200 flex items-center justify-center mb-4 ${getScoreColor(reportDetail.score)}`}>
              <span className="text-2xl font-bold">{percentage}%</span>
            </div>
            <span className={`px-4 py-2 rounded-full text-sm font-semibold ${getScoreColor(reportDetail.score)} bg-opacity-10`}>
              {getScoreStatus(reportDetail.score)}
            </span>
          </div>
        </div>

        {/* Knowledge Areas */}
        <div className="space-y-6">
          {reportDetail.knowledge_area_reports?.map((area, index) => {
            const areaPercentage = Math.round((area.score_got / area.score_expected) * 100);
            
            return (
              <div key={index} className="border rounded-lg p-6">
                <div className="flex justify-between items-start mb-4">
                  <div className="flex-1">
                    <h3 className="text-lg font-semibold text-purple-700 mb-2">
                      Área de conocimiento: {area.name}
                    </h3>
                    <p className="text-sm text-gray-600">
                      <span className="font-semibold">Horas esperadas:</span> {area.total_learn_hours_expected.toFixed(2)} · 
                      <span className="font-semibold"> Horas alcanzadas:</span> {area.total_learn_hours_actual.toFixed(2)} · 
                      <span className="font-semibold"> Peso (0–1):</span> {area.score_expected.toFixed(2)} · 
                      <span className="font-semibold"> Ponderado (obtenido):</span> {area.score_got.toFixed(2)}
                    </p>
                  </div>
                  <div className="flex flex-col items-center ml-4">
                    <div className={`w-16 h-16 rounded-full border-4 border-gray-200 flex items-center justify-center mb-2 ${getScoreColor(area.score_got / area.score_expected)}`}>
                      <span className="text-sm font-bold">{areaPercentage}%</span>
                    </div>
                    <span className={`px-2 py-1 rounded text-xs font-semibold ${getScoreColor(area.score_got / area.score_expected)} bg-opacity-10`}>
                      {getScoreStatus(area.score_got / area.score_expected)}
                    </span>
                  </div>
                </div>

                {/* Topics Table */}
                <div className="overflow-x-auto">
                  <table className="w-full border-collapse border border-gray-300">
                    <thead>
                      <tr className="bg-gray-100">
                        <th className="border border-gray-300 px-4 py-2 text-left text-sm font-semibold text-gray-700">Temática</th>
                        <th className="border border-gray-300 px-4 py-2 text-left text-sm font-semibold text-gray-700">Horas esperadas</th>
                        <th className="border border-gray-300 px-4 py-2 text-left text-sm font-semibold text-gray-700">Horas alcanzadas</th>
                      </tr>
                    </thead>
                    <tbody>
                      {area.topic_reports?.map((topic, topicIndex) => (
                        <tr key={topicIndex}>
                          <td className="border border-gray-300 px-4 py-2 text-sm text-gray-800">{topic.name}</td>
                          <td className="border border-gray-300 px-4 py-2 text-sm text-gray-800">{topic.learn_hours_expected.toFixed(2)}</td>
                          <td className="border border-gray-300 px-4 py-2 text-sm text-gray-800">{topic.learn_hours_actual.toFixed(2)}</td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            );
          })}
        </div>

        {/* Footer */}
        <div className="mt-12 pt-8 border-t border-gray-200 text-center">
          <div className="flex justify-center items-center gap-8 mb-4">
            <img 
              src="/unillanos/logo-unillanos.png" 
              alt="Universidad de los Llanos" 
              className="h-12"
            />
            <img 
              src="/unillanos/logo-fcbi.png" 
              alt="Facultad de Ciencias Básicas e Ingeniería" 
              className="h-12"
            />
          </div>
          <p className="text-sm text-gray-500">
            Universidad de los Llanos • Facultad de Ciencias Básicas e Ingeniería
          </p>
          <p className="text-xs text-gray-400 mt-1">
            Sistema WebSAC3 – Web System for Analysis of Curricula Cybersecurity Component
          </p>
        </div>
      </div>
    </div>
  );
};

export default ReportPDFAlternative;
