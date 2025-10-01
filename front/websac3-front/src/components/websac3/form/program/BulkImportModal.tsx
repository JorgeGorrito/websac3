"use client";

import React, { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { 
  useGetDegreeProgramTemplateQuery,
  useBulkImportDegreeProgramsMutation,
  type BulkImportResult
} from "@/services/api";
import { 
  Upload, 
  Download, 
  FileText, 
  CheckCircle, 
  XCircle, 
  AlertTriangle,
  X,
  FileSpreadsheet
} from "lucide-react";

interface BulkImportModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSuccess?: (result: BulkImportResult) => void;
}

export function BulkImportModal({ isOpen, onClose, onSuccess }: BulkImportModalProps) {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [isUploading, setIsUploading] = useState(false);
  
  const [bulkImport, { isLoading: isImporting }] = useBulkImportDegreeProgramsMutation();
  const { refetch: downloadTemplate, isLoading: isDownloading } = useGetDegreeProgramTemplateQuery({}, {
    skip: true, // Don't auto-fetch
  });

  const handleDownloadTemplate = async () => {
    try {
      const result = await downloadTemplate();
      if (result.data) {
        const blob = result.data as Blob;
        const url = window.URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        link.download = 'plantilla-programas-grado.xlsx';
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
        window.URL.revokeObjectURL(url);
      }
    } catch (error) {
      console.error('Error downloading template:', error);
    }
  };

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      // Validate file type
      const validTypes = [
        'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
        'application/vnd.ms-excel',
        'text/csv'
      ];
      
      if (!validTypes.includes(file.type)) {
        alert('Por favor selecciona un archivo Excel (.xlsx, .xls) o CSV válido.');
        return;
      }
      
      setSelectedFile(file);
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) return;

    setIsUploading(true);
    try {
      const result = await bulkImport({ file: selectedFile }).unwrap();
      
      if (result.result) {
        onSuccess?.(result.result);
        setSelectedFile(null);
        onClose();
      }
    } catch (error: any) {
      console.error('Error uploading file:', error);
      alert('Error al procesar el archivo. Por favor, verifica el formato y vuelve a intentar.');
    } finally {
      setIsUploading(false);
    }
  };

  const handleClose = () => {
    setSelectedFile(null);
    onClose();
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* Overlay */}
      <div 
        className="absolute inset-0 bg-black/50"
        onClick={handleClose}
      />
      
      {/* Modal Content */}
      <div className="relative bg-white rounded-lg shadow-xl max-w-2xl w-full mx-4 p-6 max-h-[90vh] overflow-y-auto">
        {/* Close Button */}
        <button
          onClick={handleClose}
          className="absolute top-4 right-4 p-1 rounded-full hover:bg-gray-100 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-300"
          aria-label="Cerrar modal"
        >
          <X className="h-5 w-5 text-gray-500 hover:text-gray-700" />
        </button>
        
        {/* Header */}
        <div className="flex items-center gap-3 pr-8 mb-6">
          <div className="p-2 bg-blue-100 rounded-lg">
            <FileSpreadsheet className="h-6 w-6 text-blue-600" />
          </div>
          <div>
            <h2 className="text-xl font-semibold text-gray-900">Carga Masiva de Programas</h2>
            <p className="text-gray-600 text-sm">Importa múltiples programas de grado desde un archivo</p>
          </div>
        </div>
        
        {/* Content */}
        <div className="space-y-6">
          {/* Step 1: Download Template */}
          <Card>
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <Download className="h-5 w-5 text-blue-600" />
                Paso 1: Descargar Plantilla
              </CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-gray-600 mb-4">
                Descarga la plantilla de Excel para conocer el formato correcto de los datos.
              </p>
              <Button 
                onClick={handleDownloadTemplate}
                disabled={isDownloading}
                className="w-full"
              >
                {isDownloading ? "Descargando..." : "Descargar Plantilla"}
              </Button>
            </CardContent>
          </Card>

          {/* Step 2: Upload File */}
          <Card>
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <Upload className="h-5 w-5 text-green-600" />
                Paso 2: Subir Archivo
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <p className="text-gray-600">
                  Selecciona el archivo Excel (.xlsx, .xls) o CSV con los datos de los programas.
                </p>
                
                <div className="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
                  <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <input
                    type="file"
                    accept=".xlsx,.xls,.csv"
                    onChange={handleFileSelect}
                    className="hidden"
                    id="file-upload"
                  />
                  <label
                    htmlFor="file-upload"
                    className="cursor-pointer inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
                  >
                    Seleccionar Archivo
                  </label>
                  {selectedFile && (
                    <div className="mt-4 p-3 bg-green-50 border border-green-200 rounded-lg">
                      <div className="flex items-center gap-2">
                        <CheckCircle className="h-4 w-4 text-green-600" />
                        <span className="text-sm font-medium text-green-800">
                          {selectedFile.name}
                        </span>
                      </div>
                      <p className="text-xs text-green-600 mt-1">
                        Tamaño: {(selectedFile.size / 1024 / 1024).toFixed(2)} MB
                      </p>
                    </div>
                  )}
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Instructions */}
          <Card className="bg-blue-50 border-blue-200">
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2 text-blue-800">
                <AlertTriangle className="h-5 w-5" />
                Instrucciones Importantes
              </CardTitle>
            </CardHeader>
            <CardContent>
              <ul className="text-sm text-blue-700 space-y-2">
                <li>• Utiliza la plantilla descargada para mantener el formato correcto</li>
                <li>• Asegúrate de que todos los campos obligatorios estén completos</li>
                <li>• Los códigos SNIES deben ser únicos</li>
                <li>• Revisa los datos antes de subir el archivo</li>
                <li>• El proceso puede tomar unos minutos dependiendo del tamaño del archivo</li>
              </ul>
            </CardContent>
          </Card>
        </div>
        
        {/* Footer */}
        <div className="flex justify-end gap-3 mt-6 pt-4 border-t border-gray-200">
          <Button
            variant="outline"
            onClick={handleClose}
            disabled={isUploading}
          >
            Cancelar
          </Button>
          <Button
            onClick={handleUpload}
            disabled={!selectedFile || isUploading}
          >
            {isUploading ? "Procesando..." : "Subir Archivo"}
          </Button>
        </div>
      </div>
    </div>
  );
}
