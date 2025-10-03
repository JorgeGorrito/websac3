"use client";

import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { 
  CheckCircle, 
  XCircle, 
  AlertTriangle,
  X,
  FileCheck,
  TrendingUp
} from "lucide-react";
import type { BulkImportCourseResult } from "@/services/api";

interface BulkImportCourseResultsProps {
  result: BulkImportCourseResult;
  onClose: () => void;
}

export function BulkImportCourseResults({ result, onClose }: BulkImportCourseResultsProps) {
  const successRate = result.total_processed > 0 
    ? Math.round((result.successful_count / result.total_processed) * 100) 
    : 0;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* Overlay */}
      <div 
        className="absolute inset-0 bg-black/50"
        onClick={onClose}
      />
      
      {/* Modal Content */}
      <div className="relative bg-white rounded-lg shadow-xl max-w-4xl w-full mx-4 p-6 max-h-[90vh] overflow-y-auto">
        {/* Close Button */}
        <button
          onClick={onClose}
          className="absolute top-4 right-4 p-1 rounded-full hover:bg-gray-100 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-300"
          aria-label="Cerrar modal"
        >
          <X className="h-5 w-5 text-gray-500 hover:text-gray-700" />
        </button>
        
        {/* Header */}
        <div className="flex items-center gap-3 pr-8 mb-6">
          <div className={`p-2 rounded-lg ${successRate === 100 ? 'bg-green-100' : successRate >= 50 ? 'bg-yellow-100' : 'bg-red-100'}`}>
            <FileCheck className={`h-6 w-6 ${successRate === 100 ? 'text-green-600' : successRate >= 50 ? 'text-yellow-600' : 'text-red-600'}`} />
          </div>
          <div>
            <h2 className="text-xl font-semibold text-gray-900">Resultados de la Importación de Cursos</h2>
            <p className="text-gray-600 text-sm">Procesamiento completado</p>
          </div>
        </div>
        
        {/* Summary Stats */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
          <Card className="bg-blue-50 border-blue-200">
            <CardContent className="p-4">
              <div className="flex items-center gap-2">
                <TrendingUp className="h-5 w-5 text-blue-600" />
                <span className="text-sm font-medium text-blue-800">Total Procesados</span>
              </div>
              <div className="text-2xl font-bold text-blue-900 mt-1">
                {result.total_processed}
              </div>
            </CardContent>
          </Card>

          <Card className="bg-green-50 border-green-200">
            <CardContent className="p-4">
              <div className="flex items-center gap-2">
                <CheckCircle className="h-5 w-5 text-green-600" />
                <span className="text-sm font-medium text-green-800">Exitosos</span>
              </div>
              <div className="text-2xl font-bold text-green-900 mt-1">
                {result.successful_count}
              </div>
            </CardContent>
          </Card>

          <Card className="bg-red-50 border-red-200">
            <CardContent className="p-4">
              <div className="flex items-center gap-2">
                <XCircle className="h-5 w-5 text-red-600" />
                <span className="text-sm font-medium text-red-800">Fallidos</span>
              </div>
              <div className="text-2xl font-bold text-red-900 mt-1">
                {result.failed_count}
              </div>
            </CardContent>
          </Card>

          <Card className={`${successRate === 100 ? 'bg-green-50 border-green-200' : successRate >= 50 ? 'bg-yellow-50 border-yellow-200' : 'bg-red-50 border-red-200'}`}>
            <CardContent className="p-4">
              <div className="flex items-center gap-2">
                <AlertTriangle className={`h-5 w-5 ${successRate === 100 ? 'text-green-600' : successRate >= 50 ? 'text-yellow-600' : 'text-red-600'}`} />
                <span className={`text-sm font-medium ${successRate === 100 ? 'text-green-800' : successRate >= 50 ? 'text-yellow-800' : 'text-red-800'}`}>
                  Tasa de Éxito
                </span>
              </div>
              <div className={`text-2xl font-bold mt-1 ${successRate === 100 ? 'text-green-900' : successRate >= 50 ? 'text-yellow-900' : 'text-red-900'}`}>
                {successRate}%
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Successful Items */}
        {result.successful_items.length > 0 && (
          <Card className="mb-6">
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2 text-green-700">
                <CheckCircle className="h-5 w-5" />
                Cursos Importados Exitosamente ({result.successful_count})
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-3 max-h-48 overflow-y-auto">
                {result.successful_items.map((item, index) => (
                  <div key={index} className="flex items-center justify-between p-3 bg-green-50 border border-green-200 rounded-lg">
                    <div>
                      <div className="font-medium text-green-900">{item.name}</div>
                      <div className="text-sm text-green-700">Código: {item.code} • Fila: {item.row_number}</div>
                      <div className="text-xs text-green-600 mt-1">{item.message}</div>
                    </div>
                    <Badge variant="secondary" className="bg-green-100 text-green-800">
                      Exitoso
                    </Badge>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        )}

        {/* Failed Items */}
        {result.failed_items.length > 0 && (
          <Card className="mb-6">
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2 text-red-700">
                <XCircle className="h-5 w-5" />
                Cursos con Errores ({result.failed_count})
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-3 max-h-48 overflow-y-auto">
                {result.failed_items.map((item, index) => (
                  <div key={index} className="p-3 bg-red-50 border border-red-200 rounded-lg">
                    <div className="flex items-start justify-between mb-2">
                      <div>
                        <div className="font-medium text-red-900">{item.name}</div>
                        <div className="text-sm text-red-700">Código: {item.code} • Fila: {item.row_number}</div>
                      </div>
                      <Badge variant="secondary" className="bg-red-100 text-red-800">
                        Error
                      </Badge>
                    </div>
                    <div className="space-y-1">
                      {item.errors.map((error, errorIndex) => (
                        <div key={errorIndex} className="text-xs text-red-600 bg-red-100 px-2 py-1 rounded">
                          • {error}
                        </div>
                      ))}
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        )}

        {/* Summary Message */}
        <div className={`p-4 rounded-lg border ${successRate === 100 ? 'bg-green-50 border-green-200' : successRate >= 50 ? 'bg-yellow-50 border-yellow-200' : 'bg-red-50 border-red-200'}`}>
          <div className="flex items-center gap-2">
            {successRate === 100 ? (
              <CheckCircle className="h-5 w-5 text-green-600" />
            ) : successRate >= 50 ? (
              <AlertTriangle className="h-5 w-5 text-yellow-600" />
            ) : (
              <XCircle className="h-5 w-5 text-red-600" />
            )}
            <span className={`font-medium ${successRate === 100 ? 'text-green-800' : successRate >= 50 ? 'text-yellow-800' : 'text-red-800'}`}>
              {successRate === 100 
                ? "¡Importación completada exitosamente!" 
                : successRate >= 50 
                ? "Importación completada con algunos errores"
                : "Importación completada con múltiples errores"
              }
            </span>
          </div>
          <p className={`text-sm mt-1 ${successRate === 100 ? 'text-green-700' : successRate >= 50 ? 'text-yellow-700' : 'text-red-700'}`}>
            {result.successful_count} de {result.total_processed} cursos fueron importados correctamente.
            {result.failed_count > 0 && " Revisa los errores para corregir los datos y volver a importar."}
          </p>
        </div>
        
        {/* Footer */}
        <div className="flex justify-end mt-6 pt-4 border-t border-gray-200">
          <Button onClick={onClose}>
            Cerrar
          </Button>
        </div>
      </div>
    </div>
  );
}


