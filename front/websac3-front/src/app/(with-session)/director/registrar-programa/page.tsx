"use client";

import { useRouter } from "next/navigation";
import { RegisterProgramForm } from "@/components/websac3/form/program/RegisterProgramForm";
import { Button } from "@/components/ui/button";
import { Upload, Plus } from "lucide-react";

export default function RegistrarProgramaPage() {
  const router = useRouter();

  return (
    <div className="space-y-6">
      {/* Header with Actions */}
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <div className="flex items-center justify-between mb-4">
            <div>
              <h2 className="text-2xl font-semibold">
                Registrar Programa
              </h2>
              <div className="mt-2 h-0.5 w-24 bg-gray-300 rounded" />
            </div>
            <Button
              onClick={() => router.push('/director/carga-masiva-programas')}
              className="flex items-center gap-2"
            >
              <Upload className="h-4 w-4" />
              Carga Masiva
            </Button>
          </div>
          <p className="text-gray-600 text-sm">
            Registra un programa de grado individualmente o utiliza la carga masiva para importar múltiples programas.
          </p>
        </div>
        
        <RegisterProgramForm />
      </div>

      {/* Quick Actions */}
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <h3 className="text-lg font-semibold mb-4">Opciones de Registro</h3>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors">
            <div className="flex items-center gap-3 mb-2">
              <div className="p-2 bg-blue-100 rounded-lg">
                <Plus className="h-5 w-5 text-blue-600" />
              </div>
              <h4 className="font-medium text-gray-900">Registro Individual</h4>
            </div>
            <p className="text-sm text-gray-600">
              Completa el formulario para registrar un programa de grado paso a paso.
            </p>
          </div>
          
          <div className="p-4 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors cursor-pointer" onClick={() => router.push('/director/carga-masiva-programas')}>
            <div className="flex items-center gap-3 mb-2">
              <div className="p-2 bg-green-100 rounded-lg">
                <Upload className="h-5 w-5 text-green-600" />
              </div>
              <h4 className="font-medium text-gray-900">Carga Masiva</h4>
            </div>
            <p className="text-sm text-gray-600">
              Importa múltiples programas desde un archivo Excel o CSV.
            </p>
          </div>
        </div>
      </div>

    </div>
  );
}
