"use client";

import React from 'react';
import { ReportDetail } from '@/services/api';

interface ReportHTMLViewProps {
  reportDetail: ReportDetail;
}

const ReportHTMLView: React.FC<ReportHTMLViewProps> = ({ reportDetail }) => {
  const getScoreColor = (score: number) => {
    if (score >= 0.9) return '#28a745';
    if (score >= 0.8) return '#ffc107';
    return '#dc3545';
  };

  const getScoreStatus = (score: number) => {
    if (score >= 0.9) return 'Excelente';
    if (score >= 0.8) return 'Aceptable';
    return 'Por mejorar';
  };

  const percentage = (Math.floor(reportDetail.score * 10000) / 100).toFixed(2);
  const scoreColor = getScoreColor(reportDetail.score);
  const scoreStatus = getScoreStatus(reportDetail.score);
  
  // Calcular stroke-dasharray para el círculo principal (r=60, circunferencia = 2πr = 376.99)
  const mainCircumference = 376.99;
  const mainStrokeDasharray = `${(reportDetail.score * mainCircumference).toFixed(2)} ${mainCircumference}`;

  return (
    <div style={{ 
      fontFamily: 'Arial, sans-serif', 
      backgroundColor: '#f4f4f4', 
      margin: 0, 
      padding: '20px 10px',
      minHeight: '100vh'
    }}>
      <div style={{ 
        width: '100%',
        maxWidth: '900px',
        backgroundColor: '#ffffff', 
        borderRadius: '8px', 
        boxShadow: '0 4px 12px rgba(0,0,0,0.1)', 
        padding: '20px',
        margin: '0 auto'
      }}>
        {/* Header with Logo */}
        <div style={{ 
          textAlign: 'center', 
          paddingBottom: '16px',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center'
        }}>
          <img 
            src="/websac3/logo-websac3.png" 
            alt="WebSAC3" 
            width="180"
            style={{ display: 'block', margin: '0 auto' }}
          />
        </div>

        {/* Main Section */}
        <div style={{ paddingBottom: '8px' }}>
          <div style={{ display: 'flex', gap: '30px', alignItems: 'flex-start' }}>
            {/* Left Column - Information */}
            <div style={{ flex: '0 0 70%' }}>
              <h2 style={{ color: '#2c3e50', margin: '0 0 8px', fontSize: '24px', fontWeight: 'bold' }}>
                Reporte de Evaluación
              </h2>
              <p style={{ fontSize: '14px', color: '#666', margin: '0 0 20px', lineHeight: 1.4 }}>
                Resultado de la evaluación del componente de ciberseguridad
              </p>
              
              {/* Institution Card */}
              <div style={{ 
                padding: '16px', 
                background: '#f8f9fa', 
                border: '1px solid #e9ecef', 
                borderRadius: '8px', 
                boxShadow: '0 2px 4px rgba(0,0,0,0.05)',
                marginBottom: '16px'
              }}>
                <div style={{ display: 'flex', gap: '16px' }}>
                  <div style={{ flex: '0 0 20%' }}>
                    <div style={{ fontSize: '11px', color: '#6c757d', marginBottom: '4px', fontWeight: 500 }}>
                      SNIES Institución
                    </div>
                    <div style={{ fontSize: '15px', fontWeight: 700, color: '#2c3e50' }}>
                      {reportDetail.degree_program?.higher_education_institution?.snies || 'N/A'}
                    </div>
                  </div>
                  <div style={{ flex: '0 0 50%' }}>
                    <div style={{ fontSize: '11px', color: '#6c757d', marginBottom: '4px', fontWeight: 500 }}>
                      Institución
                    </div>
                    <div style={{ fontSize: '15px', fontWeight: 700, color: '#2c3e50', lineHeight: 1.3 }}>
                      {reportDetail.degree_program?.higher_education_institution?.name || 'No disponible'}
                    </div>
                  </div>
                  <div style={{ flex: '0 0 30%' }}>
                    <div style={{ fontSize: '11px', color: '#6c757d', marginBottom: '4px', fontWeight: 500 }}>
                      Fecha del reporte
                    </div>
                    <div style={{ fontSize: '15px', fontWeight: 700, color: '#2c3e50' }}>
                      {new Date(reportDetail.created_at).toLocaleDateString('es-ES')}
                    </div>
                  </div>
                </div>
              </div>
              
              {/* Program Card */}
              <div style={{ 
                padding: '16px', 
                background: '#f8f9fa', 
                border: '1px solid #e9ecef', 
                borderRadius: '8px', 
                boxShadow: '0 2px 4px rgba(0,0,0,0.05)'
              }}>
                <div style={{ display: 'flex', gap: '16px' }}>
                  <div style={{ flex: '0 0 20%' }}>
                    <div style={{ fontSize: '11px', color: '#6c757d', marginBottom: '4px', fontWeight: 500 }}>
                      SNIES Programa
                    </div>
                    <div style={{ fontSize: '15px', fontWeight: 700, color: '#2c3e50' }}>
                      {reportDetail.degree_program?.snies || 'N/A'}
                    </div>
                  </div>
                  <div style={{ flex: '0 0 30%' }}>
                    <div style={{ fontSize: '11px', color: '#6c757d', marginBottom: '4px', fontWeight: 500 }}>
                      Programa
                    </div>
                    <div style={{ fontSize: '15px', fontWeight: 700, color: '#2c3e50', lineHeight: 1.3 }}>
                      {reportDetail.degree_program?.name || 'N/A'}
                    </div>
                  </div>
                  <div style={{ flex: '0 0 30%' }}>
                    <div style={{ fontSize: '11px', color: '#6c757d', marginBottom: '4px', fontWeight: 500 }}>
                      Rol profesional
                    </div>
                    <div style={{ fontSize: '15px', fontWeight: 700, color: '#2c3e50' }}>
                      {reportDetail.professional_role?.name || 'N/A'}
                    </div>
                  </div>
                  <div style={{ flex: '0 0 20%' }}>
                    <div style={{ fontSize: '11px', color: '#6c757d', marginBottom: '4px', fontWeight: 500 }}>
                      Puntaje total
                    </div>
                    <div style={{ fontSize: '15px', fontWeight: 700, color: '#2c3e50' }}>
                      {reportDetail.score.toFixed(2)}
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            {/* Right Column - Score Chart */}
            <div style={{ 
              flex: '0 0 30%', 
              display: 'flex', 
              flexDirection: 'column', 
              alignItems: 'center', 
              justifyContent: 'center',
              paddingLeft: '20px'
            }}>
              <svg width="140" height="140" style={{ transform: 'rotate(-90deg)', marginBottom: '12px' }}>
                <circle cx="70" cy="70" r="60" fill="none" stroke="#e9ecef" strokeWidth="6"/>
                <circle 
                  cx="70" 
                  cy="70" 
                  r="60" 
                  fill="none" 
                  stroke={scoreColor} 
                  strokeWidth="6" 
                  strokeDasharray={mainStrokeDasharray}
                  strokeLinecap="round"
                />
                <text 
                  x="70" 
                  y="75" 
                  textAnchor="middle" 
                  transform="rotate(90 70 70)"
                  style={{ 
                    fontSize: '20px', 
                    fontWeight: 'bold', 
                    fill: '#2c3e50' 
                  }}
                >
                  {percentage}%
                </text>
              </svg>
              <div style={{ 
                textAlign: 'center', 
                fontSize: '13px', 
                color: scoreColor, 
                fontWeight: 600, 
                background: `${scoreColor}1a`, 
                padding: '6px 12px', 
                borderRadius: '20px', 
                border: `1px solid ${scoreColor}` 
              }}>
                {scoreStatus}
              </div>
            </div>
          </div>
          
          <hr style={{ border: 'none', borderTop: '1px solid #e0e0e0', margin: '16px 0' }} />
        </div>

        {/* Knowledge Areas */}
        {reportDetail.knowledge_area_reports?.map((area, index) => {
          const areaProgress = area.score_got / area.score_expected;
          const areaPercentage = (Math.floor(areaProgress * 10000) / 100).toFixed(2);
          const areaScoreColor = getScoreColor(areaProgress);
          const areaScoreStatus = getScoreStatus(areaProgress);
          
          // Calcular stroke-dasharray para círculos pequeños (r=22, circunferencia = 138.23)
          const smallCircumference = 138.23;
          const areaStrokeDasharray = `${(areaProgress * smallCircumference).toFixed(2)} ${smallCircumference}`;

          return (
            <div key={index} style={{ margin: '14px 0 10px' }}>
              {/* Knowledge Area Header */}
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: '10px' }}>
                <div style={{ flex: 1 }}>
                  <h3 style={{ color: '#5e35b1', margin: '0 0 8px', fontSize: '16px', fontWeight: 'bold' }}>
                    Área de conocimiento: {area.name}
                  </h3>
                  <p style={{ fontSize: '13px', color: '#555', margin: '6px 0 0' }}>
                    <strong>Horas esperadas:</strong> {area.total_learn_hours_expected.toFixed(2)} ·{' '}
                    <strong>Horas alcanzadas:</strong> {area.total_learn_hours_actual.toFixed(2)} ·{' '}
                    <strong>Peso (0–1):</strong> {area.score_expected.toFixed(2)} ·{' '}
                    <strong>Ponderado (obtenido):</strong> {area.score_got.toFixed(2)}
                  </p>
                </div>
                <div style={{ 
                  display: 'flex', 
                  flexDirection: 'column', 
                  alignItems: 'center',
                  width: '100px',
                  marginLeft: '20px'
                }}>
                  <svg width="60" height="60" style={{ transform: 'rotate(-90deg)' }}>
                    <circle cx="30" cy="30" r="22" fill="none" stroke="#e0e0e0" strokeWidth="4"/>
                    <circle 
                      cx="30" 
                      cy="30" 
                      r="22" 
                      fill="none" 
                      stroke={areaScoreColor} 
                      strokeWidth="4" 
                      strokeDasharray={areaStrokeDasharray}
                      strokeLinecap="round"
                    />
                    <text 
                      x="30" 
                      y="35" 
                      textAnchor="middle" 
                      transform="rotate(90 30 30)"
                      style={{ 
                        fontSize: '10px', 
                        fontWeight: 'bold', 
                        fill: '#333' 
                      }}
                    >
                      {areaPercentage}%
                    </text>
                  </svg>
                  <div style={{ 
                    textAlign: 'center', 
                    fontSize: '9px', 
                    color: areaScoreColor, 
                    fontWeight: 600, 
                    marginTop: '4px' 
                  }}>
                    {areaScoreStatus}
                  </div>
                </div>
              </div>

              {/* Topics Table */}
              <div style={{ 
                border: '1px solid #e6e8f0',
                borderRadius: '4px',
                overflow: 'hidden',
                marginBottom: '14px'
              }}>
                {/* Table Header */}
                <div style={{ 
                  display: 'flex',
                  background: '#f0f1f6'
                }}>
                  <div style={{ 
                    flex: 1,
                    fontSize: '12px', 
                    color: '#333', 
                    padding: '8px', 
                    borderRight: '1px solid #e6e8f0',
                    fontWeight: 'bold'
                  }}>
                    Temática
                  </div>
                  <div style={{ 
                    flex: 1,
                    fontSize: '12px', 
                    color: '#333', 
                    padding: '8px', 
                    borderRight: '1px solid #e6e8f0',
                    fontWeight: 'bold'
                  }}>
                    Horas esperadas
                  </div>
                  <div style={{ 
                    flex: 1,
                    fontSize: '12px', 
                    color: '#333', 
                    padding: '8px',
                    fontWeight: 'bold'
                  }}>
                    Horas alcanzadas
                  </div>
                </div>
                
                {/* Table Rows */}
                {area.topic_reports?.map((topic, topicIndex) => (
                  <div key={topicIndex} style={{ 
                    display: 'flex',
                    borderTop: '1px solid #ececf5'
                  }}>
                    <div style={{ 
                      flex: 1,
                      fontSize: '12px', 
                      color: '#333', 
                      padding: '8px', 
                      borderRight: '1px solid #ececf5'
                    }}>
                      {topic.name}
                    </div>
                    <div style={{ 
                      flex: 1,
                      fontSize: '12px', 
                      color: '#333', 
                      padding: '8px', 
                      borderRight: '1px solid #ececf5'
                    }}>
                      {topic.learn_hours_expected.toFixed(2)}
                    </div>
                    <div style={{ 
                      flex: 1,
                      fontSize: '12px', 
                      color: '#333', 
                      padding: '8px'
                    }}>
                      {topic.learn_hours_actual.toFixed(2)}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          );
        })}

        {/* Unexpected Knowledge Areas */}
        {reportDetail.unexpected_knowledge_area_reports && reportDetail.unexpected_knowledge_area_reports.length > 0 && (
          <>
            <hr style={{ border: 'none', borderTop: '2px solid #28a745', margin: '20px 0' }} />
            <h2 style={{ 
              color: '#28a745', 
              margin: '0 0 16px', 
              fontSize: '20px', 
              textAlign: 'center',
              fontWeight: 'bold'
            }}>
              📚 Áreas de Conocimiento Adicionales
            </h2>
            <p style={{ 
              fontSize: '13px', 
              color: '#666', 
              margin: '0 0 20px', 
              textAlign: 'center', 
              lineHeight: 1.4 
            }}>
              Temáticas cubiertas en el programa que no se esperaban en el rol profesional
            </p>

            {reportDetail.unexpected_knowledge_area_reports.map((area, index) => (
              <div key={index} style={{ 
                margin: '14px 0 10px', 
                border: '1px solid #d4edda', 
                borderRadius: '8px', 
                backgroundColor: '#f8fff9',
                padding: '16px'
              }}>
                <h3 style={{ 
                  color: '#28a745', 
                  margin: '0 0 8px', 
                  fontSize: '16px',
                  fontWeight: 'bold'
                }}>
                  🌱 {area.name}
                </h3>
                <p style={{ 
                  fontSize: '13px', 
                  color: '#555', 
                  margin: '6px 0 12px' 
                }}>
                  <strong>Total de horas cubiertas:</strong> {area.total_learn_hours.toFixed(2)}
                </p>
                
                {/* Unexpected Topics Table */}
                <div style={{ 
                  border: '1px solid #c3e6cb',
                  borderRadius: '4px',
                  overflow: 'hidden'
                }}>
                  {/* Table Header */}
                  <div style={{ 
                    display: 'flex',
                    background: '#d4edda'
                  }}>
                    <div style={{ 
                      flex: 1,
                      fontSize: '12px', 
                      color: '#155724', 
                      padding: '8px', 
                      borderRight: '1px solid #c3e6cb',
                      fontWeight: 'bold'
                    }}>
                      Temática
                    </div>
                    <div style={{ 
                      flex: 1,
                      fontSize: '12px', 
                      color: '#155724', 
                      padding: '8px',
                      fontWeight: 'bold'
                    }}>
                      Horas cubiertas
                    </div>
                  </div>
                  
                  {/* Table Rows */}
                  {area.topic_reports?.map((topic, topicIndex) => (
                    <div key={topicIndex} style={{ 
                      display: 'flex',
                      borderTop: '1px solid #d4edda'
                    }}>
                      <div style={{ 
                        flex: 1,
                        fontSize: '12px', 
                        color: '#333', 
                        padding: '8px', 
                        borderRight: '1px solid #d4edda'
                      }}>
                        {topic.topic?.name || topic.name}
                      </div>
                      <div style={{ 
                        flex: 1,
                        fontSize: '12px', 
                        color: '#333', 
                        padding: '8px'
                      }}>
                        {topic.learn_hours_actual.toFixed(2)}
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            ))}
          </>
        )}

        {/* Footer */}
        <hr style={{ border: 'none', borderTop: '1px solid #e0e0e0', margin: '20px 0' }} />
        <div style={{ 
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          gap: '20px',
          margin: '12px 0'
        }}>
          <img 
            src="/unillanos/logo-unillanos.png" 
            alt="Universidad de los Llanos" 
            style={{ width: '100px', height: 'auto' }} 
          />
          <img 
            src="/unillanos/logo-fcbi.png" 
            alt="Facultad de Ciencias Básicas e Ingeniería" 
            style={{ width: '100px', height: 'auto' }} 
          />
        </div>
        <p style={{ fontSize: '12px', color: '#aaa', textAlign: 'center', margin: 0 }}>
          Universidad de los Llanos • Facultad de Ciencias Básicas e Ingeniería
        </p>
        <p style={{ fontSize: '11px', color: '#bbb', textAlign: 'center', margin: '8px 0 0' }}>
          Sistema WebSAC3 – Web System for Analysis of Curricula Cybersecurity Component
        </p>
      </div>
    </div>
  );
};

export default ReportHTMLView;