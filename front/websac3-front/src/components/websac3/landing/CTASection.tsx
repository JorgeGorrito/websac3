"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { 
  ArrowRight, 
  LogIn, 
  Mail, 
  Shield,
  CheckCircle,
  Star
} from "lucide-react";
import Link from "next/link";

const CTASection = () => {
  const ctaOptions = [
    {
      icon: LogIn,
      title: "Acceso Directo",
      description: "Si ya tienes una cuenta, inicia sesión para acceder al sistema.",
      buttonText: "Iniciar Sesión",
      buttonHref: "/login",
      variant: "default" as const,
      color: "bg-blue-600 hover:bg-blue-700"
    },
    {
      icon: Mail,
      title: "Solicitar Acceso",
      description: "Solicita acceso al sistema si eres parte de una institución educativa.",
      buttonText: "Solicitar Acceso",
      buttonHref: "/solicitar-acceso",
      variant: "outline" as const,
      color: "border-green-600 text-green-600 hover:bg-green-600 hover:text-white"
    }
  ];

  const features = [
    "Análisis del componente de ciberseguridad",
    "Reportes especializados en PDF y HTML",
    "Evaluación de programas de grado",
    "Roles académicos diferenciados",
    "Seguridad académica garantizada",
    "Desarrollado por Universidad de los Llanos"
  ];

  return (
    <section className="py-20 bg-gradient-to-br from-blue-600 via-indigo-600 to-purple-600 relative overflow-hidden">
      {/* Background Pattern */}
      <div className="absolute inset-0 opacity-30">
        <div className="w-full h-full" style={{
          backgroundImage: `url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23ffffff' fill-opacity='0.05'%3E%3Ccircle cx='30' cy='30' r='1'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E")`
        }}></div>
      </div>
      
      <div className="container mx-auto px-4 relative z-10">
        <div className="text-center mb-16">
          <div className="flex justify-center mb-6">
            <div className="bg-white/10 backdrop-blur-sm p-4 rounded-full">
              <Shield className="h-12 w-12 text-white" />
            </div>
          </div>
          <h2 className="text-4xl md:text-5xl font-bold text-white mb-6">
            ¿Listo para Analizar Ciberseguridad?
          </h2>
          <p className="text-xl text-blue-100 max-w-3xl mx-auto leading-relaxed">
            Únete a la comunidad académica que evalúa el componente de ciberseguridad 
            en programas de grado con WebSAC3.
          </p>
        </div>

        {/* CTA Cards */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8 mb-16 max-w-4xl mx-auto">
          {ctaOptions.map((option, index) => {
            const IconComponent = option.icon;
            return (
              <Card key={index} className="group hover:shadow-2xl transition-all duration-300 border-0 shadow-xl hover:scale-105 bg-white/10 backdrop-blur-sm">
                <CardContent className="p-8 text-center">
                  <div className="bg-white/20 w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-6 group-hover:scale-110 transition-transform duration-300">
                    <IconComponent className="h-8 w-8 text-white" />
                  </div>
                  
                  <h3 className="text-2xl font-semibold text-white mb-4">
                    {option.title}
                  </h3>
                  
                  <p className="text-blue-100 mb-6 leading-relaxed">
                    {option.description}
                  </p>
                  
                  <Link href={option.buttonHref}>
                    <Button 
                      variant={option.variant}
                      size="lg" 
                      className={`w-full ${option.color} transition-all duration-300 font-semibold`}
                    >
                      {option.buttonText}
                      <ArrowRight className="ml-2 h-5 w-5" />
                    </Button>
                  </Link>
                </CardContent>
              </Card>
            );
          })}
        </div>

        {/* Features List */}
        <div className="bg-white/10 backdrop-blur-sm rounded-2xl p-8 md:p-12">
          <div className="text-center mb-8">
            <h3 className="text-3xl font-bold text-white mb-4">
              Especializado en Análisis de Ciberseguridad
            </h3>
            <p className="text-blue-100 text-lg">
              Funcionalidades diseñadas para evaluar el componente de ciberseguridad en programas de grado
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {features.map((feature, index) => (
              <div key={index} className="flex items-center space-x-3">
                <div className="bg-green-500 p-1 rounded-full">
                  <CheckCircle className="h-5 w-5 text-white" />
                </div>
                <span className="text-white font-medium">{feature}</span>
              </div>
            ))}
          </div>
        </div>

        {/* Final CTA */}
        <div className="text-center mt-16">
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link href="/login">
              <Button size="lg" className="bg-white text-blue-600 hover:bg-blue-50 px-8 py-4 text-lg font-semibold rounded-lg shadow-lg hover:shadow-xl transition-all duration-300">
                Acceder al Sistema
                <ArrowRight className="ml-2 h-5 w-5" />
              </Button>
            </Link>
            <Link href="/solicitar-acceso">
              <Button variant="outline" size="lg" className="border-2 border-white text-white hover:bg-white hover:text-blue-600 px-8 py-4 text-lg font-semibold rounded-lg transition-all duration-300">
                Más Información
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
};

export { CTASection };
