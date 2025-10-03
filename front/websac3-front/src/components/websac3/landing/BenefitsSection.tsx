"use client";

import { Card, CardContent } from "@/components/ui/card";
import { 
  Target, 
  Clock, 
  Award, 
  Globe, 
  Zap, 
  Lock,
  TrendingUp,
  Users2
} from "lucide-react";

const BenefitsSection = () => {
  const benefits = [
    {
      icon: Target,
      title: "Análisis Especializado",
      description: "Evaluación específica del componente de ciberseguridad en programas de grado con métricas precisas.",
      stat: "100%",
      statLabel: "Enfoque"
    },
    {
      icon: Clock,
      title: "Evaluación Eficiente",
      description: "Análisis automatizado del componente de ciberseguridad que reduce el tiempo de evaluación académica.",
      stat: "75%",
      statLabel: "Eficiencia"
    },
    {
      icon: Award,
      title: "Calidad Curricular",
      description: "Mejora la calidad del componente de ciberseguridad en programas académicos mediante análisis detallados.",
      stat: "90%",
      statLabel: "Mejora"
    },
    {
      icon: Globe,
      title: "Estándares Académicos",
      description: "Cumple con estándares académicos para la evaluación del componente de ciberseguridad en programas de grado.",
      stat: "100%",
      statLabel: "Cumplimiento"
    }
  ];

  const additionalBenefits = [
    {
      icon: Zap,
      title: "Análisis Rápido",
      description: "Evaluación rápida del componente de ciberseguridad en programas de grado."
    },
    {
      icon: Lock,
      title: "Seguridad Académica",
      description: "Protección de información académica sensible durante el análisis."
    },
    {
      icon: TrendingUp,
      title: "Mejora Continua",
      description: "Sistema que mejora constantemente la evaluación del componente de ciberseguridad."
    },
    {
      icon: Users2,
      title: "Colaboración Académica",
      description: "Facilita la colaboración entre diferentes roles en la evaluación académica."
    }
  ];

  return (
    <section className="py-20 bg-gradient-to-br from-gray-50 to-blue-50">
      <div className="container mx-auto px-4">
        {/* Main Benefits */}
        <div className="text-center mb-16">
          <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
            Beneficios del Análisis Especializado
          </h2>
          <p className="text-xl text-gray-600 max-w-3xl mx-auto">
            Evalúa y mejora el componente de ciberseguridad en programas de grado 
            con análisis especializado y métricas precisas.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8 mb-20">
          {benefits.map((benefit, index) => {
            const IconComponent = benefit.icon;
            return (
              <Card key={index} className="group hover:shadow-xl transition-all duration-300 border-0 shadow-lg hover:scale-105 bg-white">
                <CardContent className="p-8 text-center">
                  <div className="bg-gradient-to-br from-blue-100 to-indigo-100 w-20 h-20 rounded-full flex items-center justify-center mx-auto mb-6 group-hover:scale-110 transition-transform duration-300">
                    <IconComponent className="h-10 w-10 text-blue-600" />
                  </div>
                  
                  <div className="mb-4">
                    <div className="text-3xl font-bold text-blue-600 mb-1">
                      {benefit.stat}
                    </div>
                    <div className="text-sm text-gray-500 font-medium">
                      {benefit.statLabel}
                    </div>
                  </div>
                  
                  <h3 className="text-xl font-semibold text-gray-900 mb-3">
                    {benefit.title}
                  </h3>
                  <p className="text-gray-600 leading-relaxed">
                    {benefit.description}
                  </p>
                </CardContent>
              </Card>
            );
          })}
        </div>

        {/* Additional Benefits Grid */}
        <div className="bg-white rounded-2xl shadow-xl p-8 md:p-12">
          <div className="text-center mb-12">
            <h3 className="text-3xl font-bold text-gray-900 mb-4">
              Funcionalidades Especializadas
            </h3>
            <p className="text-lg text-gray-600">
              Características específicas para el análisis del componente de ciberseguridad
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            {additionalBenefits.map((benefit, index) => {
              const IconComponent = benefit.icon;
              return (
                <div key={index} className="text-center group">
                  <div className="bg-gradient-to-br from-green-100 to-emerald-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-4 group-hover:scale-110 transition-transform duration-300">
                    <IconComponent className="h-8 w-8 text-green-600" />
                  </div>
                  <h4 className="text-lg font-semibold text-gray-900 mb-2">
                    {benefit.title}
                  </h4>
                  <p className="text-gray-600 text-sm leading-relaxed">
                    {benefit.description}
                  </p>
                </div>
              );
            })}
          </div>
        </div>

        {/* Stats Section */}
        <div className="mt-20 bg-gradient-to-r from-blue-600 to-indigo-600 rounded-2xl p-8 md:p-12 text-white">
          <div className="text-center mb-12">
            <h3 className="text-3xl font-bold mb-4">
              Impacto Académico
            </h3>
            <p className="text-blue-100 text-lg">
              Resultados en el análisis del componente de ciberseguridad
            </p>
          </div>

          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            <div className="text-center">
              <div className="text-4xl font-bold mb-2">100%</div>
              <div className="text-blue-100">Enfoque en Ciberseguridad</div>
            </div>
            <div className="text-center">
              <div className="text-4xl font-bold mb-2">Unillanos</div>
              <div className="text-blue-100">Universidad Desarrolladora</div>
            </div>
            <div className="text-center">
              <div className="text-4xl font-bold mb-2">Especializado</div>
              <div className="text-blue-100">Análisis Académico</div>
            </div>
            <div className="text-center">
              <div className="text-4xl font-bold mb-2">Grado</div>
              <div className="text-blue-100">Programas Académicos</div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export { BenefitsSection };
