"use client";

import React, { useState } from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Upload, FormInput, FileSpreadsheet, ArrowLeft, CheckCircle } from "lucide-react";

export default function CargarDatosPage() {
  const [selectedOption, setSelectedOption] = useState<"form" | "csv" | null>(
    null
  );

  const handleOptionSelect = (option: "form" | "csv") => {
    setSelectedOption(option);
  };

  const handleBack = () => {
    setSelectedOption(null);
  };

  if (selectedOption === "form") {
    return <FormularioCursos />;
  }

  if (selectedOption === "csv") {
    return <PlantillaCSV onBack={handleBack} />;
  }

  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <h2 className="text-2xl font-semibold text-center">
            Cargar Datos
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>
        
        <p className="text-gray-600 text-center mb-8 max-w-2xl mx-auto">
          Selecciona cómo deseas cargar los datos de los cursos. Puedes usar
          un formulario para entrada manual o una plantilla CSV para carga
          masiva.
        </p>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          <Card 
            className="cursor-pointer hover:shadow-md transition-shadow duration-200 border border-gray-200"
            onClick={() => handleOptionSelect("form")}
          >
            <CardHeader className="text-center pb-4">
              <div className="mx-auto w-16 h-16 bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl flex items-center justify-center mb-4">
                <FormInput className="w-8 h-8 text-blue-600" />
              </div>
              <CardTitle className="text-lg font-semibold text-gray-800">Formulario</CardTitle>
              <CardDescription className="text-gray-600">
                Ingresa los datos de los cursos manualmente paso a paso
              </CardDescription>
            </CardHeader>
            <CardContent className="pb-6">
              <p className="text-gray-600 text-center text-sm">
                Ideal para cargar pocos cursos o cuando necesitas revisar cada
                entrada individualmente
              </p>
            </CardContent>
          </Card>

          <Card 
            className="cursor-pointer hover:shadow-md transition-shadow duration-200 border border-gray-200"
            onClick={() => handleOptionSelect("csv")}
          >
            <CardHeader className="text-center pb-4">
              <div className="mx-auto w-16 h-16 bg-gradient-to-br from-green-50 to-green-100 rounded-xl flex items-center justify-center mb-4">
                <FileSpreadsheet className="w-8 h-8 text-green-600" />
              </div>
              <CardTitle className="text-lg font-semibold text-gray-800">Plantilla CSV</CardTitle>
              <CardDescription className="text-gray-600">
                Descarga la plantilla y sube el archivo con todos los cursos
              </CardDescription>
            </CardHeader>
            <CardContent className="pb-6">
              <p className="text-gray-600 text-center text-sm">
                Perfecto para cargar muchos cursos de una vez de manera eficiente
              </p>
            </CardContent>
          </Card>
        </div>

        <div className="bg-gray-50 rounded-xl p-6 border border-gray-200">
          <h3 className="text-lg font-semibold text-gray-800 mb-4">
            Datos que se pueden cargar:
          </h3>
          <div className="grid md:grid-cols-2 gap-4 text-gray-700">
            <ul className="space-y-2">
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                Periodicidad (ej: 1 Semestre)
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                Nombre del curso
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                Número de créditos
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                Naturaleza (virtual, presencial)
              </li>
            </ul>
            <ul className="space-y-2">
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                Tipo (teoría, laboratorio)
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                ¿Es curso de ciberseguridad? (s/n)
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                Temáticas del curso
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}

// Componente para el formulario de cursos
function FormularioCursos() {
  const [currentStep, setCurrentStep] = useState(1);
  const [formData, setFormData] = useState({
    periodicidad: "",
    nombreCurso: "",
    creditos: "",
    naturaleza: "",
    tipo: "",
    esCiberseguridad: "",
    tematicas: "",
  });

  const steps = [
    {
      id: 1,
      title: "Información Básica",
      description: "Datos generales del curso",
    },
    {
      id: 2,
      title: "Características",
      description: "Naturaleza y tipo del curso",
    },
    {
      id: 3,
      title: "Clasificación",
      description: "Ciberseguridad y temáticas",
    },
    { 
      id: 4, 
      title: "Revisión", 
      description: "Revisa y confirma los datos",
    },
  ];

  const handleInputChange = (field: string, value: string) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
  };

  const nextStep = () => {
    if (currentStep < steps.length) {
      setCurrentStep(currentStep + 1);
    }
  };

  const prevStep = () => {
    if (currentStep > 1) {
      setCurrentStep(currentStep - 1);
    }
  };

  const handleSubmit = () => {
    console.log("Datos del curso:", formData);
    // Aquí iría la lógica para enviar los datos
    alert("Curso registrado exitosamente");
  };

  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <div className="flex items-center gap-4 mb-4">
            <Button 
              variant="ghost" 
              onClick={() => window.history.back()}
              className="text-gray-600 hover:text-gray-800"
            >
              <ArrowLeft className="w-4 h-4 mr-2" />
              Volver
            </Button>
          </div>
          <h2 className="text-2xl font-semibold text-center">
            Formulario de Cursos
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>

        {/* Steps indicator */}
        <div className="mb-8">
          <div className="flex items-center justify-between">
            {steps.map((step, index) => (
              <div key={step.id} className="flex items-center">
                <div className="flex flex-col items-center">
                  <div
                    className={`flex items-center justify-center w-12 h-12 rounded-xl border-2 transition-all duration-200 ${
                      currentStep >= step.id
                        ? "bg-blue-500 border-blue-500 text-white"
                        : "bg-white border-gray-300 text-gray-500"
                    }`}
                  >
                    <span className="text-sm font-semibold">{step.id}</span>
                  </div>
                  <div className="mt-2 text-center">
                    <h3
                      className={`text-sm font-semibold ${
                        currentStep >= step.id
                          ? "text-blue-600"
                          : "text-gray-500"
                      }`}
                    >
                      {step.title}
                    </h3>
                    <p className="text-xs text-gray-400 mt-1">{step.description}</p>
                  </div>
                </div>
                {index < steps.length - 1 && (
                  <div
                    className={`w-16 h-1 mx-4 rounded-full transition-all duration-200 ${
                      currentStep > step.id ? "bg-blue-500" : "bg-gray-200"
                    }`}
                  />
                )}
              </div>
            ))}
          </div>
        </div>

        {/* Form content */}
        <Card className="mb-6 border border-gray-200">
          <CardContent className="p-6">
            {currentStep === 1 && (
              <div className="space-y-4">
                <h3 className="text-lg font-semibold text-gray-800 mb-4">
                  Información Básica del Curso
                </h3>

                <div className="grid lg:grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="block text-sm font-medium text-gray-700">
                      Periodicidad *
                    </label>
                    <select
                      value={formData.periodicidad}
                      onChange={(e) =>
                        handleInputChange("periodicidad", e.target.value)
                      }
                      className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white"
                    >
                      <option value="">Selecciona la periodicidad</option>
                      <option value="1-semestre">1 Semestre</option>
                      <option value="2-semestres">2 Semestres</option>
                      <option value="3-semestres">3 Semestres</option>
                      <option value="4-semestres">4 Semestres</option>
                    </select>
                  </div>

                  <div className="space-y-2">
                    <label className="block text-sm font-medium text-gray-700">
                      Número de Créditos *
                    </label>
                    <input
                      type="number"
                      value={formData.creditos}
                      onChange={(e) =>
                        handleInputChange("creditos", e.target.value)
                      }
                      className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200"
                      placeholder="Ej: 3"
                      min="1"
                      max="6"
                    />
                  </div>
                </div>

                <div className="space-y-2">
                  <label className="block text-sm font-medium text-gray-700">
                    Nombre del Curso *
                  </label>
                  <input
                    type="text"
                    value={formData.nombreCurso}
                    onChange={(e) =>
                      handleInputChange("nombreCurso", e.target.value)
                    }
                    className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200"
                    placeholder="Ej: Introducción a la Programación"
                  />
                </div>
              </div>
            )}

            {currentStep === 2 && (
              <div className="space-y-4">
                <h3 className="text-lg font-semibold text-gray-800 mb-4">
                  Características del Curso
                </h3>

                <div className="grid lg:grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="block text-sm font-medium text-gray-700">
                      Naturaleza *
                    </label>
                    <select
                      value={formData.naturaleza}
                      onChange={(e) =>
                        handleInputChange("naturaleza", e.target.value)
                      }
                      className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white"
                    >
                      <option value="">Selecciona la naturaleza</option>
                      <option value="virtual">Virtual</option>
                      <option value="presencial">Presencial</option>
                      <option value="hibrido">Híbrido</option>
                    </select>
                  </div>

                  <div className="space-y-2">
                    <label className="block text-sm font-medium text-gray-700">
                      Tipo *
                    </label>
                    <select
                      value={formData.tipo}
                      onChange={(e) =>
                        handleInputChange("tipo", e.target.value)
                      }
                      className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white"
                    >
                      <option value="">Selecciona el tipo</option>
                      <option value="teoria">Teoría</option>
                      <option value="laboratorio">Laboratorio</option>
                      <option value="teoria-laboratorio">
                        Teoría + Laboratorio
                      </option>
                    </select>
                  </div>
                </div>
              </div>
            )}

            {currentStep === 3 && (
              <div className="space-y-4">
                <h3 className="text-lg font-semibold text-gray-800 mb-4">
                  Clasificación del Curso
                </h3>

                <div className="space-y-4">
                  <div className="space-y-2">
                    <label className="block text-sm font-medium text-gray-700">
                      ¿Es un curso de ciberseguridad? *
                    </label>
                    <select
                      value={formData.esCiberseguridad}
                      onChange={(e) =>
                        handleInputChange("esCiberseguridad", e.target.value)
                      }
                      className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 bg-white"
                    >
                      <option value="">Selecciona una opción</option>
                      <option value="si">Sí</option>
                      <option value="no">No</option>
                    </select>
                  </div>

                  <div className="space-y-2">
                    <label className="block text-sm font-medium text-gray-700">
                      Temáticas del Curso
                    </label>
                    <textarea
                      value={formData.tematicas}
                      onChange={(e) =>
                        handleInputChange("tematicas", e.target.value)
                      }
                      className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200 resize-none"
                      rows={4}
                      placeholder="Describe las temáticas principales del curso..."
                    />
                  </div>
                </div>
              </div>
            )}

            {currentStep === 4 && (
              <div className="space-y-4">
                <h3 className="text-lg font-semibold text-gray-800 mb-4">
                  Revisión de Datos
                </h3>

                <div className="bg-gray-50 rounded-lg p-4 border border-gray-200">
                  <div className="grid lg:grid-cols-2 gap-4">
                    <div className="space-y-3">
                      <div className="bg-white p-3 rounded-lg border border-gray-200">
                        <p className="text-sm text-gray-500 font-medium">Periodicidad</p>
                        <p className="font-semibold text-gray-800">
                          {formData.periodicidad || "No especificado"}
                        </p>
                      </div>
                      <div className="bg-white p-3 rounded-lg border border-gray-200">
                        <p className="text-sm text-gray-500 font-medium">Créditos</p>
                        <p className="font-semibold text-gray-800">
                          {formData.creditos || "No especificado"}
                        </p>
                      </div>
                      <div className="bg-white p-3 rounded-lg border border-gray-200">
                        <p className="text-sm text-gray-500 font-medium">
                          Nombre del Curso
                        </p>
                        <p className="font-semibold text-gray-800">
                          {formData.nombreCurso || "No especificado"}
                        </p>
                      </div>
                    </div>
                    <div className="space-y-3">
                      <div className="bg-white p-3 rounded-lg border border-gray-200">
                        <p className="text-sm text-gray-500 font-medium">Naturaleza</p>
                        <p className="font-semibold text-gray-800">
                          {formData.naturaleza || "No especificado"}
                        </p>
                      </div>
                      <div className="bg-white p-3 rounded-lg border border-gray-200">
                        <p className="text-sm text-gray-500 font-medium">Tipo</p>
                        <p className="font-semibold text-gray-800">
                          {formData.tipo || "No especificado"}
                        </p>
                      </div>
                      <div className="bg-white p-3 rounded-lg border border-gray-200">
                        <p className="text-sm text-gray-500 font-medium">
                          ¿Es de Ciberseguridad?
                        </p>
                        <p className="font-semibold text-gray-800">
                          {formData.esCiberseguridad || "No especificado"}
                        </p>
                      </div>
                    </div>
                  </div>
                  {formData.tematicas && (
                    <div className="mt-4">
                      <div className="bg-white p-3 rounded-lg border border-gray-200">
                        <p className="text-sm text-gray-500 font-medium">Temáticas</p>
                        <p className="font-semibold text-gray-800">{formData.tematicas}</p>
                      </div>
                    </div>
                  )}
                </div>
              </div>
            )}
          </CardContent>
        </Card>

        {/* Navigation buttons */}
        <div className="flex justify-between items-center">
          <Button
            variant="outline"
            onClick={prevStep}
            disabled={currentStep === 1}
            className="px-6 py-2 text-gray-700 border-gray-300 hover:bg-gray-50 disabled:opacity-50"
          >
            Anterior
          </Button>

          <div className="flex gap-3">
            {currentStep < steps.length ? (
              <Button 
                onClick={nextStep}
                className="px-6 py-2 bg-blue-600 hover:bg-blue-700"
              >
                Siguiente
              </Button>
            ) : (
              <Button
                onClick={handleSubmit}
                className="px-6 py-2 bg-green-600 hover:bg-green-700 flex items-center gap-2"
              >
                <CheckCircle className="w-4 h-4" />
                Subir Curso
              </Button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

// Componente para la plantilla CSV
function PlantillaCSV({ onBack }: { onBack: () => void }) {
  const handleFileUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      console.log("Archivo seleccionado:", file.name);
      // Aquí iría la lógica para procesar el archivo CSV
      alert("Archivo CSV cargado exitosamente");
    }
  };

  const downloadTemplate = () => {
    // Crear y descargar la plantilla CSV
    const csvContent =
      "Periodicidad,Nombre Curso,Créditos,Naturaleza,Tipo,Es Ciberseguridad,Temáticas\n1 Semestre,Introducción a la Programación,3,Virtual,Teoría,No,Programación básica, algoritmos\n";
    const blob = new Blob([csvContent], { type: "text/csv" });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = "plantilla_cursos.csv";
    a.click();
    window.URL.revokeObjectURL(url);
  };

  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <div className="flex items-center gap-4 mb-4">
            <Button 
              variant="ghost" 
              onClick={onBack}
              className="text-gray-600 hover:text-gray-800"
            >
              <ArrowLeft className="w-4 h-4 mr-2" />
              Volver
            </Button>
          </div>
          <h2 className="text-2xl font-semibold text-center">
            Plantilla CSV
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>

        <p className="text-gray-600 text-center mb-8 max-w-2xl mx-auto">
          Descarga la plantilla CSV, complétala con los datos de tus cursos
          y súbela aquí para procesamiento automático
        </p>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          <Card className="border border-gray-200">
            <CardHeader className="pb-4">
              <CardTitle className="flex items-center gap-3 text-lg font-semibold text-gray-800">
                <div className="w-10 h-10 bg-gradient-to-br from-green-50 to-green-100 rounded-lg flex items-center justify-center">
                  <FileSpreadsheet className="w-5 h-5 text-green-600" />
                </div>
                Descargar Plantilla
              </CardTitle>
              <CardDescription className="text-gray-600">
                Obtén la plantilla CSV con el formato correcto y ejemplos
              </CardDescription>
            </CardHeader>
            <CardContent className="pb-6">
              <Button 
                onClick={downloadTemplate} 
                className="w-full py-3 bg-green-600 hover:bg-green-700"
              >
                Descargar Plantilla CSV
              </Button>
              <p className="text-sm text-gray-600 mt-3 text-center">
                La plantilla incluye ejemplos de cómo llenar cada campo
                correctamente
              </p>
            </CardContent>
          </Card>

          <Card className="border border-gray-200">
            <CardHeader className="pb-4">
              <CardTitle className="flex items-center gap-3 text-lg font-semibold text-gray-800">
                <div className="w-10 h-10 bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg flex items-center justify-center">
                  <Upload className="w-5 h-5 text-blue-600" />
                </div>
                Subir Archivo
              </CardTitle>
              <CardDescription className="text-gray-600">
                Sube tu archivo CSV completado para procesamiento
              </CardDescription>
            </CardHeader>
            <CardContent className="pb-6">
              <div className="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center hover:border-green-400 transition-colors duration-200">
                <input
                  type="file"
                  accept=".csv"
                  onChange={handleFileUpload}
                  className="hidden"
                  id="csv-upload"
                />
                <label htmlFor="csv-upload" className="cursor-pointer">
                  <Upload className="w-12 h-12 text-gray-400 mx-auto mb-3" />
                  <p className="text-sm text-gray-600 font-medium">
                    Haz clic para seleccionar tu archivo CSV
                  </p>
                  <p className="text-xs text-gray-500 mt-1">
                    Solo archivos .csv
                  </p>
                </label>
              </div>
            </CardContent>
          </Card>
        </div>

        <div className="bg-gray-50 rounded-xl p-6 border border-gray-200">
          <h3 className="text-lg font-semibold text-gray-800 mb-4">
            Instrucciones para la plantilla CSV:
          </h3>
          <div className="grid md:grid-cols-2 gap-4 text-gray-700">
            <ul className="space-y-2">
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                Usa comas (,) para separar campos
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                Encierra texto en comillas si contiene comas
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                No dejes filas vacías
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                Respeta el formato de la primera fila
              </li>
            </ul>
            <ul className="space-y-2">
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                Periodicidad: 1 Semestre, 2 Semestres, etc.
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                Naturaleza: Virtual, Presencial, Híbrido
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                Tipo: Teoría, Laboratorio, Teoría + Laboratorio
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                Ciberseguridad: Sí o No
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}