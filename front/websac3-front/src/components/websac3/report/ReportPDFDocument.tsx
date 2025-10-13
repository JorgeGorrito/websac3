import React from 'react';
import { Document, Page, Text, View, StyleSheet, Image, Font } from '@react-pdf/renderer';
import { ReportDetail } from '@/services/api';

// Use default fonts that are already available in @react-pdf/renderer

// Create styles based on the Go HTML template
const styles = StyleSheet.create({
  page: {
    flexDirection: 'column',
    backgroundColor: '#f4f4f4',
    padding: 20,
    fontFamily: 'Helvetica',
  },
  container: {
    width: '100%',
    backgroundColor: '#ffffff',
    borderRadius: 8,
    padding: 20,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.1,
    shadowRadius: 12,
  },
  header: {
    alignItems: 'center',
    marginBottom: 16,
  },
  logo: {
    width: 120,
    height: 48,
  },
  mainSection: {
    flexDirection: 'row',
    marginBottom: 4,
  },
  leftColumn: {
    flex: 0.65,
    marginRight: 20,
  },
  rightColumn: {
    flex: 0.35,
    alignItems: 'center',
    justifyContent: 'center',
    paddingLeft: 15,
  },
  title: {
    fontSize: 16,
    fontWeight: 'bold',
    color: '#2c3e50',
    marginBottom: 6,
  },
  subtitle: {
    fontSize: 12,
    color: '#666',
    marginBottom: 10,
    lineHeight: 1.4,
  },
  card: {
    backgroundColor: '#f8f9fa',
    borderWidth: 1,
    borderColor: '#e9ecef',
    borderStyle: 'solid',
    borderRadius: 6,
    padding: 8,
    marginBottom: 6,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.05,
    shadowRadius: 4,
  },
  cardRow: {
    flexDirection: 'row',
  },
  cardColumn: {
    flex: 1,
    marginRight: 16,
  },
  cardColumnLast: {
    flex: 1,
  },
  cardLabel: {
    fontSize: 9,
    color: '#6c757d',
    fontWeight: 500,
    marginBottom: 2,
  },
  cardValue: {
    fontSize: 12,
    fontWeight: 700,
    color: '#2c3e50',
    lineHeight: 1.3,
  },
  scoreContainer: {
    alignItems: 'center',
    justifyContent: 'center',
    height: '100%',
  },
  scoreCircle: {
    width: 80,
    height: 80,
    borderRadius: 40,
    borderWidth: 4,
    borderColor: '#e9ecef',
    borderStyle: 'solid',
    alignItems: 'center',
    justifyContent: 'center',
    marginBottom: 8,
  },
  scoreText: {
    fontSize: 14,
    fontWeight: 'bold',
    color: '#2c3e50',
    textAlign: 'center',
    lineHeight: 1.2,
    fontFamily: 'Helvetica',
  },
  scoreStatus: {
    fontSize: 10,
    fontWeight: 600,
    paddingTop: 4,
    paddingBottom: 4,
    paddingLeft: 8,
    paddingRight: 8,
    borderRadius: 15,
    borderWidth: 1,
    borderStyle: 'solid',
    textAlign: 'center',
  },
  separator: {
    height: 1,
    backgroundColor: '#e0e0e0',
    marginVertical: 8,
  },
  knowledgeArea: {
    marginVertical: 8,
  },
  knowledgeAreaHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    marginBottom: 6,
  },
  knowledgeAreaInfo: {
    flex: 1,
  },
  knowledgeAreaTitle: {
    fontSize: 14,
    fontWeight: 'bold',
    color: '#5e35b1',
    marginBottom: 6,
  },
  knowledgeAreaDetails: {
    fontSize: 11,
    color: '#555',
    marginTop: 4,
  },
  smallScoreContainer: {
    alignItems: 'center',
    width: 80,
  },
  smallScoreCircle: {
    width: 50,
    height: 50,
    borderRadius: 25,
    borderWidth: 3,
    borderColor: '#e0e0e0',
    borderStyle: 'solid',
    alignItems: 'center',
    justifyContent: 'center',
  },
  smallScoreText: {
    fontSize: 9,
    fontWeight: 'bold',
    color: '#333',
    textAlign: 'center',
    lineHeight: 1.2,
    fontFamily: 'Helvetica',
    transform: 'none',
  },
  smallScoreStatus: {
    fontSize: 8,
    fontWeight: 600,
    marginTop: 3,
    textAlign: 'center',
  },
  topicsTable: {
    marginBottom: 8,
  },
  tableHeader: {
    flexDirection: 'row',
    backgroundColor: '#f0f1f6',
  },
  tableHeaderCell: {
    flex: 1,
    fontSize: 10,
    color: '#333',
    padding: 6,
    borderWidth: 1,
    borderColor: '#e6e8f0',
    borderStyle: 'solid',
    textAlign: 'left',
  },
  tableRow: {
    flexDirection: 'row',
  },
  tableCell: {
    flex: 1,
    fontSize: 10,
    color: '#333',
    padding: 6,
    borderWidth: 1,
    borderColor: '#ececf5',
    borderStyle: 'solid',
    textAlign: 'left',
  },
  footer: {
    alignItems: 'center',
    marginTop: 12,
    paddingTop: 12,
    borderTopWidth: 1,
    borderTopColor: '#e0e0e0',
    borderTopStyle: 'solid',
  },
  footerLogos: {
    flexDirection: 'row',
    justifyContent: 'center',
    marginBottom: 8,
  },
  footerLogo: {
    width: 80,
    height: 40,
    marginHorizontal: 8,
  },
  footerText: {
    fontSize: 10,
    color: '#aaa',
    textAlign: 'center',
    marginBottom: 6,
  },
  footerSubtext: {
    fontSize: 9,
    color: '#bbb',
    textAlign: 'center',
  },
});

interface ReportPDFDocumentProps {
  reportDetail: ReportDetail;
}

const ReportPDFDocument: React.FC<ReportPDFDocumentProps> = ({ reportDetail }) => {
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

  const percentage = Math.round(reportDetail.score * 100);
  const scoreColor = getScoreColor(reportDetail.score);
  const scoreStatus = getScoreStatus(reportDetail.score);

  return (
    <Document>
      <Page size="A4" style={styles.page}>
        <View style={styles.container}>
          {/* Header with Logo */}
          <View style={styles.header}>
            <Image 
              style={styles.logo}
              src="/websac3/logo-websac3.png"
            />
          </View>

          {/* Main Section with Cards and Score */}
          <View style={styles.mainSection}>
            {/* Left Column: Information */}
            <View style={styles.leftColumn}>
              <Text style={styles.title}>Reporte de Evaluación</Text>
              <Text style={styles.subtitle}>
                Resultado de la evaluación del componente de ciberseguridad
              </Text>
              
              {/* Institution Card */}
              <View style={styles.card}>
                <View style={styles.cardRow}>
                  <View style={styles.cardColumn}>
                    <Text style={styles.cardLabel}>SNIES Institución</Text>
                    <Text style={styles.cardValue}>
                      {reportDetail.degree_program?.higher_education_institution?.snies || 'N/A'}
                    </Text>
                  </View>
                  <View style={styles.cardColumn}>
                    <Text style={styles.cardLabel}>Institución</Text>
                    <Text style={styles.cardValue}>
                      {reportDetail.degree_program?.higher_education_institution?.name || 'No disponible'}
                    </Text>
                  </View>
                  <View style={styles.cardColumnLast}>
                    <Text style={styles.cardLabel}>Fecha del reporte</Text>
                    <Text style={styles.cardValue}>
                      {new Date(reportDetail.created_at).toLocaleDateString('es-ES')}
                    </Text>
                  </View>
                </View>
              </View>

              {/* Program Card */}
              <View style={styles.card}>
                <View style={styles.cardRow}>
                  <View style={styles.cardColumn}>
                    <Text style={styles.cardLabel}>SNIES Programa</Text>
                    <Text style={styles.cardValue}>
                      {reportDetail.degree_program?.snies || 'N/A'}
                    </Text>
                  </View>
                  <View style={styles.cardColumn}>
                    <Text style={styles.cardLabel}>Programa</Text>
                    <Text style={styles.cardValue}>
                      {reportDetail.degree_program?.name || 'N/A'}
                    </Text>
                  </View>
                  <View style={styles.cardColumn}>
                    <Text style={styles.cardLabel}>Rol profesional</Text>
                    <Text style={styles.cardValue}>
                      {reportDetail.professional_role?.name || 'N/A'}
                    </Text>
                  </View>
                  <View style={styles.cardColumnLast}>
                    <Text style={styles.cardLabel}>Puntaje total</Text>
                    <Text style={styles.cardValue}>
                      {reportDetail.score.toFixed(2)}
                    </Text>
                  </View>
                </View>
              </View>
            </View>

            {/* Right Column: Score Chart */}
            <View style={styles.rightColumn}>
              <View style={styles.scoreContainer}>
                <View style={styles.scoreCircle}>
                  <Text style={styles.scoreText}>{percentage}%</Text>
                </View>
                <View style={[styles.scoreStatus, { borderColor: scoreColor }]}>
                  <Text style={{ color: scoreColor }}>{scoreStatus}</Text>
                </View>
              </View>
            </View>
          </View>

          {/* Separator */}
          <View style={styles.separator} />

          {/* Knowledge Areas */}
          {reportDetail.knowledge_area_reports?.map((area, index) => {
            const areaPercentage = Math.round((area.score_got / area.score_expected) * 100);
            const areaScoreColor = getScoreColor(area.score_got / area.score_expected);
            const areaScoreStatus = getScoreStatus(area.score_got / area.score_expected);

            return (
              <View key={index} style={styles.knowledgeArea}>
                {/* Knowledge Area Header */}
                <View style={styles.knowledgeAreaHeader}>
                  <View style={styles.knowledgeAreaInfo}>
                    <Text style={styles.knowledgeAreaTitle}>
                      Área de conocimiento: {area.name}
                    </Text>
                    <Text style={styles.knowledgeAreaDetails}>
                      <Text style={{ fontWeight: 'bold' }}>Horas esperadas:</Text> {area.total_learn_hours_expected.toFixed(2)} · 
                      <Text style={{ fontWeight: 'bold' }}> Horas alcanzadas:</Text> {area.total_learn_hours_actual.toFixed(2)} · 
                      <Text style={{ fontWeight: 'bold' }}> Peso (0–1):</Text> {area.score_expected.toFixed(2)} · 
                      <Text style={{ fontWeight: 'bold' }}> Ponderado (obtenido):</Text> {area.score_got.toFixed(2)}
                    </Text>
                  </View>
                  
                  <View style={styles.smallScoreContainer}>
                    <View style={[styles.smallScoreCircle, { borderColor: areaScoreColor }]}>
                      <Text style={styles.smallScoreText}>{areaPercentage}%</Text>
                    </View>
                    <Text style={[styles.smallScoreStatus, { color: areaScoreColor }]}>
                      {areaScoreStatus}
                    </Text>
                  </View>
                </View>

                {/* Topics Table */}
                <View style={styles.topicsTable}>
                  <View style={styles.tableHeader}>
                    <Text style={styles.tableHeaderCell}>Temática</Text>
                    <Text style={styles.tableHeaderCell}>Horas esperadas</Text>
                    <Text style={styles.tableHeaderCell}>Horas alcanzadas</Text>
                  </View>
                  {area.topic_reports?.map((topic, topicIndex) => (
                    <View key={topicIndex} style={styles.tableRow}>
                      <Text style={styles.tableCell}>{topic.name}</Text>
                      <Text style={styles.tableCell}>{topic.learn_hours_expected.toFixed(2)}</Text>
                      <Text style={styles.tableCell}>{topic.learn_hours_actual.toFixed(2)}</Text>
                    </View>
                  ))}
                </View>
              </View>
            );
          })}

          {/* Footer */}
          <View style={styles.footer}>
            <View style={styles.footerLogos}>
              <Image 
                style={styles.footerLogo}
                src="/unillanos/logo-unillanos.png"
              />
              <Image 
                style={styles.footerLogo}
                src="/unillanos/logo-fcbi.png"
              />
            </View>
            <Text style={styles.footerText}>
              Universidad de los Llanos • Facultad de Ciencias Básicas e Ingeniería
            </Text>
            <Text style={styles.footerSubtext}>
              Sistema WebSAC3 – Web System for Analysis of Curricula Cybersecurity Component
            </Text>
          </View>
        </View>
      </Page>
    </Document>
  );
};

export default ReportPDFDocument;