"use client";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { 
  TrendingUp, 
  Users,
  Shield, 
  BarChart3,
  AlertTriangle,
  Target,
  Lightbulb
} from "lucide-react";

const FeaturesSection = () => {
  const features = [
    {
      icon: Shield,
      title: "Análisis de Ciberseguridad",
      description: "Evalúa específicamente el componente de ciberseguridad en programas de grado académicos.",
      color: "bg-blue-100 text-blue-600"
    },
    {
      icon: TrendingUp,
      title: "Reportes Especializados",
      description: "Genera reportes detallados sobre el componente de ciberseguridad en PDF y HTML.",
      color: "bg-purple-100 text-purple-600"
    },
    {
      icon: Users,
      title: "Asesorías con Experto",
      description: "Asesorías con experto en ciberseguridad para el análisis del componente curricular.",
      color: "bg-red-100 text-red-600"
    },
    {
      icon: BarChart3,
      title: "Métricas de Evaluación",
      description: "Visualiza métricas específicas del componente de ciberseguridad en los programas.",
      color: "bg-indigo-100 text-indigo-600"
    }
  ];

  return (
    <section className="py-20 bg-white">
      <div className="container mx-auto px-4">
        <div className="text-center mb-16">
          <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
            Análisis Especializado en Ciberseguridad
          </h2>
          <p className="text-xl text-gray-600 max-w-3xl mx-auto">
            WebSAC3 realiza la evaluación y análisis del componente de ciberseguridad 
            en programas académicos de grado con base en estándares internacionales como ACM e IEEE.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {features.map((feature, index) => {
            const IconComponent = feature.icon;
            return (
              <Card key={index} className="group hover:shadow-lg transition-all duration-300 border-0 shadow-md hover:scale-105">
                <CardHeader className="text-center pb-4">
                  <div className={`w-16 h-16 mx-auto rounded-full ${feature.color} flex items-center justify-center mb-4 group-hover:scale-110 transition-transform duration-300`}>
                    <IconComponent className="h-8 w-8" />
                  </div>
                  <CardTitle className="text-xl font-semibold text-gray-900">
                    {feature.title}
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <CardDescription className="text-gray-600 text-center leading-relaxed">
                    {feature.description}
                  </CardDescription>
                </CardContent>
              </Card>
            );
          })}
        </div>

        {/* Why WebSAC3 - Problem & Solution */}
        <div className="mt-20">
          {/* Problem Section */}
          <div className="bg-gradient-to-br from-slate-50 via-blue-50 to-slate-50 rounded-3xl p-10 md:p-14 mb-10 shadow-sm">
            <div className="text-center mb-12">
              <h3 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
                El Desafío en Ciberseguridad
              </h3>
              <p className="text-xl text-gray-600 max-w-4xl mx-auto leading-relaxed">
                La demanda global de profesionales en ciberseguridad supera la oferta educativa actual. 
                Los datos muestran una brecha significativa entre las necesidades del mercado y las capacidades formadas.
              </p>
            </div>

            {/* Key Statistics */}
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-6xl mx-auto mb-12">
              <div className="bg-white rounded-2xl p-8 shadow-md hover:shadow-xl transition-shadow duration-300">
                <div className="text-center">
                  <div className="text-5xl font-bold bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent mb-3">
                    +25%
                  </div>
                  <p className="text-gray-800 font-semibold text-lg mb-2">Brecha laboral mundial</p>
                  <p className="text-gray-600">Aumento en 2022</p>
                </div>
              </div>
              <div className="bg-white rounded-2xl p-8 shadow-md hover:shadow-xl transition-shadow duration-300">
                <div className="text-center">
                  <div className="text-5xl font-bold bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent mb-3">
                    70%
                  </div>
                  <p className="text-gray-800 font-semibold text-lg mb-2">Organizaciones</p>
                  <p className="text-gray-600">Con escasez de personal</p>
                </div>
              </div>
              <div className="bg-white rounded-2xl p-8 shadow-md hover:shadow-xl transition-shadow duration-300">
                <div className="text-center">
                  <div className="text-5xl font-bold bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent mb-3">
                    3.43M
                  </div>
                  <p className="text-gray-800 font-semibold text-lg mb-2">Déficit mundial</p>
                  <p className="text-gray-600">515K en Latinoamérica</p>
                </div>
              </div>
            </div>

            {/* Colombian Context */}
            <div className="bg-white rounded-2xl p-8 shadow-md max-w-5xl mx-auto mb-10">
              <div className="flex items-center justify-center mb-6">
                <div className="text-3xl mr-3">🇨🇴</div>
                <h4 className="text-2xl font-bold text-gray-900">Contexto Colombiano</h4>
              </div>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div className="flex items-start space-x-4">
                  <div className="flex-shrink-0 w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center">
                    <span className="text-blue-600 font-bold text-lg">3°</span>
                  </div>
                  <div>
                    <p className="text-gray-800 font-semibold mb-1">Posición en Latinoamérica</p>
                    <p className="text-gray-600">6.300 millones de intentos de ciberataques registrados</p>
                  </div>
                </div>
                <div className="flex items-start space-x-4">
                  <div className="flex-shrink-0 w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center">
                    <span className="text-blue-600 font-bold text-lg">68%</span>
                  </div>
                  <div>
                    <p className="text-gray-800 font-semibold mb-1">Empresas afectadas</p>
                    <p className="text-gray-600">Tuvieron al menos un incidente de ciberseguridad en 2022</p>
                  </div>
                </div>
              </div>
            </div>

            {/* Impact on Organizations */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-8 max-w-5xl mx-auto">
              <div className="bg-gradient-to-br from-white to-blue-50 rounded-xl p-6 border border-blue-100 shadow-sm">
                <div className="flex items-start space-x-3">
                  <div className="flex-shrink-0 w-10 h-10 bg-blue-500 rounded-lg flex items-center justify-center">
                    <BarChart3 className="h-5 w-5 text-white" />
                  </div>
                  <div>
                    <h5 className="font-bold text-gray-900 mb-2">Capacidades Insuficientes</h5>
                    <p className="text-gray-700">
                      <strong>59%</strong> de organizaciones señalan que la falta de habilidades en sus equipos 
                      dificulta responder eficazmente a incidentes.
                    </p>
                  </div>
                </div>
              </div>
              <div className="bg-gradient-to-br from-white to-blue-50 rounded-xl p-6 border border-blue-100 shadow-sm">
                <div className="flex items-start space-x-3">
                  <div className="flex-shrink-0 w-10 h-10 bg-indigo-500 rounded-lg flex items-center justify-center">
                    <TrendingUp className="h-5 w-5 text-white" />
                  </div>
                  <div>
                    <h5 className="font-bold text-gray-900 mb-2">Demanda en Aumento</h5>
                    <p className="text-gray-700">
                      Incremento constante de plazas para profesionales de ciberseguridad, 
                      requiriendo competencias actualizadas.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Solution Section */}
          <div className="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-2xl p-8 md:p-12">
            <div className="text-center mb-8">
              <div className="inline-block bg-blue-100 p-4 rounded-full mb-4">
                <Lightbulb className="h-12 w-12 text-blue-600" />
              </div>
              <h3 className="text-3xl font-bold text-gray-900 mb-4">
                WebSAC3: Fortaleciendo la Educación en Ciberseguridad
              </h3>
              <p className="text-lg text-gray-700 max-w-4xl mx-auto leading-relaxed mb-6">
                Ante esta <strong>problemática social</strong> de vulneración de sistemas y falta de capacidades, 
                WebSAC3 surge como una <strong>herramienta especializada</strong> para fortalecer las capacidades y 
                habilidades de los programas de TI en ciberseguridad.
              </p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-5xl mx-auto">
              <div className="text-center">
                <div className="bg-gradient-to-br from-blue-100 to-blue-200 w-20 h-20 rounded-2xl flex items-center justify-center mx-auto mb-4 transform hover:scale-110 transition-transform duration-300">
                  <Shield className="h-10 w-10 text-blue-700" />
                </div>
                <h4 className="text-xl font-semibold text-gray-900 mb-3">Análisis Curricular</h4>
                <p className="text-gray-700 leading-relaxed">
                  Evalúa el componente de ciberseguridad en programas académicos basándose en estándares internacionales ACM e IEEE.
                </p>
              </div>
              <div className="text-center">
                <div className="bg-gradient-to-br from-green-100 to-green-200 w-20 h-20 rounded-2xl flex items-center justify-center mx-auto mb-4 transform hover:scale-110 transition-transform duration-300">
                  <Target className="h-10 w-10 text-green-700" />
                </div>
                <h4 className="text-xl font-semibold text-gray-900 mb-3">Identificación de Brechas</h4>
                <p className="text-gray-700 leading-relaxed">
                  Detecta deficiencias en competencias y habilidades de ciberseguridad para tomar decisiones informadas de mejora.
                </p>
              </div>
              <div className="text-center">
                <div className="bg-gradient-to-br from-purple-100 to-purple-200 w-20 h-20 rounded-2xl flex items-center justify-center mx-auto mb-4 transform hover:scale-110 transition-transform duration-300">
                  <Users className="h-10 w-10 text-purple-700" />
                </div>
                <h4 className="text-xl font-semibold text-gray-900 mb-3">Asesoría Especializada</h4>
                <p className="text-gray-700 leading-relaxed">
                  Conexión con expertos en ciberseguridad para validar y enriquecer el análisis del componente curricular.
                </p>
              </div>
            </div>

            <div className="mt-10 text-center bg-blue-100/50 rounded-xl p-6 max-w-3xl mx-auto">
              <p className="text-lg text-gray-800 leading-relaxed">
                <strong>WebSAC3</strong> contribuye a cerrar la brecha de talento en ciberseguridad, 
                fortaleciendo los programas académicos para formar profesionales con las <strong>competencias 
                necesarias</strong> para enfrentar los desafíos actuales de seguridad digital.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export { FeaturesSection };
