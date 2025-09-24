# 📄 Generación de PDFs desde Componentes React

Este directorio contiene diferentes implementaciones para generar PDFs desde componentes React en el proyecto WebSAC3.

## 🎯 Opciones Disponibles

### 1. **@react-pdf/renderer** (Recomendada)
**Archivo:** `ReportPDFDocument.tsx`

**Ventajas:**
- ✅ Renderiza componentes React nativos
- ✅ Control total sobre el diseño
- ✅ Componentes específicos para PDF (Document, Page, Text, View, Image)
- ✅ Muy estable y bien mantenida
- ✅ Soporte completo para estilos CSS
- ✅ Genera PDFs nativos (no imágenes)

**Desventajas:**
- ❌ Requiere reescribir el componente con componentes específicos de PDF
- ❌ Limitaciones en algunos estilos CSS avanzados

**Uso:**
```tsx
import { pdf } from '@react-pdf/renderer';
import ReportPDFDocument from '@/components/websac3/report/ReportPDFDocument';

const downloadPDF = async () => {
  const blob = await pdf(<ReportPDFDocument reportDetail={reportDetail} />).toBlob();
  // Descargar el PDF
};
```

### 2. **react-to-pdf** (Alternativa Simple)
**Archivo:** `ReportPDFAlternative.tsx`

**Ventajas:**
- ✅ Usa componentes React existentes
- ✅ Conversión automática
- ✅ Menos configuración
- ✅ Mantiene estilos CSS

**Desventajas:**
- ❌ Genera PDFs como imágenes (menor calidad de texto)
- ❌ Menos control sobre el layout
- ❌ Puede tener problemas con estilos complejos

**Uso:**
```tsx
import { usePDF } from 'react-to-pdf';

const { toPDF, targetRef } = usePDF({
  filename: 'reporte.pdf',
  page: { margin: 20, format: 'a4' }
});

// En el JSX
<div ref={targetRef}>
  {/* Tu componente existente */}
</div>
<button onClick={() => toPDF()}>Descargar PDF</button>
```

### 3. **html2canvas + jsPDF** (Captura Visual)
**Implementación anterior en el archivo principal**

**Ventajas:**
- ✅ Captura exactamente lo que se ve en pantalla
- ✅ Incluye todos los estilos y gráficos
- ✅ Funciona con cualquier componente React

**Desventajas:**
- ❌ Genera PDFs como imágenes
- ❌ Problemas con colores `oklch` de Tailwind CSS
- ❌ Más complejo de implementar
- ❌ Archivos más grandes

## 🚀 Implementación Actual

### **ReportPDFDocument.tsx**
La implementación actual usa `@react-pdf/renderer` y incluye:

- **Header completo** con logo WebSAC3
- **Información institucional** en cards
- **Gráficos circulares** con porcentajes
- **Tablas** con bordes y formato
- **Colores dinámicos** según rendimiento
- **Footer** con logos institucionales
- **Múltiples páginas** si es necesario

### **Características del PDF Generado:**
- ✅ **Formato A4** estándar
- ✅ **Márgenes** apropiados (30px)
- ✅ **Tipografía** Roboto
- ✅ **Colores** dinámicos (verde/amarillo/rojo)
- ✅ **Gráficos** circulares con porcentajes
- ✅ **Tablas** formateadas
- ✅ **Imágenes** incluidas
- ✅ **Layout** profesional

## 🔧 Instalación

```bash
# Para @react-pdf/renderer (implementación actual)
npm install @react-pdf/renderer

# Para react-to-pdf (alternativa)
npm install react-to-pdf

# Para html2canvas + jsPDF (captura visual)
npm install html2canvas jspdf
```

## 📝 Uso en el Proyecto

### **Función de Descarga Actual:**
```tsx
const downloadReportAsPDF = async () => {
  if (!reportDetail) return;
  
  try {
    const blob = await pdf(<ReportPDFDocument reportDetail={reportDetail} />).toBlob();
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = `Reporte_${programName}_${date}.pdf`;
    link.click();
    URL.revokeObjectURL(url);
  } catch (error) {
    console.error('Error generating PDF:', error);
  }
};
```

## 🎨 Personalización

### **Estilos CSS:**
Los estilos se definen usando `StyleSheet.create()` de `@react-pdf/renderer`:

```tsx
const styles = StyleSheet.create({
  page: {
    flexDirection: 'column',
    backgroundColor: '#ffffff',
    padding: 30,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#2c3e50',
  },
  // ... más estilos
});
```

### **Colores Dinámicos:**
```tsx
const getScoreColor = (score: number) => {
  if (score >= 0.9) return '#28a745'; // Verde
  if (score >= 0.8) return '#ffc107'; // Amarillo
  return '#dc3545'; // Rojo
};
```

## 🔄 Migración entre Opciones

### **De html2canvas a @react-pdf/renderer:**
1. Instalar `@react-pdf/renderer`
2. Crear componente PDF específico
3. Reemplazar función de descarga
4. Probar y ajustar estilos

### **De @react-pdf/renderer a react-to-pdf:**
1. Instalar `react-to-pdf`
2. Usar componente React existente
3. Agregar `usePDF` hook
4. Configurar opciones de PDF

## 🐛 Solución de Problemas

### **Problemas Comunes:**

1. **Imágenes no se cargan:**
   - Verificar rutas de imágenes
   - Usar URLs absolutas
   - Configurar CORS si es necesario

2. **Estilos no se aplican:**
   - Usar `StyleSheet.create()` en lugar de CSS
   - Verificar propiedades soportadas
   - Usar colores hexadecimales

3. **Texto se corta:**
   - Ajustar márgenes de página
   - Usar `break` para saltos de página
   - Verificar altura de contenido

4. **PDF muy grande:**
   - Optimizar imágenes
   - Reducir calidad si es necesario
   - Usar compresión

## 📚 Recursos Adicionales

- [Documentación @react-pdf/renderer](https://react-pdf.org/)
- [Ejemplos de componentes](https://react-pdf.org/components)
- [Guía de estilos](https://react-pdf.org/styling)
- [react-to-pdf GitHub](https://github.com/ivmarcos/react-to-pdf)
