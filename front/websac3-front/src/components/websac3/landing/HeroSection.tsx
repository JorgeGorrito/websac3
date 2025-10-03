"use client";

import { Button } from "@/components/ui/button";
import { WebSAC3Logo } from "../logos/WebSAC3Logo";
import { UnillanosLogo } from "../logos/UnillanosLogo";
import { FCBILogo } from "../logos/FCBILogo";
import { ArrowRight, Shield, Users, BarChart3, BookOpen } from "lucide-react";
import Link from "next/link";

const HeroSection = () => {
  return (
    <section className="relative min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-50 via-blue-50 to-slate-100 overflow-hidden">
      {/* Decorative Elements */}
      <div className="absolute inset-0 overflow-hidden">
        <div className="absolute -top-40 -right-40 w-80 h-80 bg-blue-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob"></div>
        <div className="absolute -bottom-40 -left-40 w-80 h-80 bg-indigo-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob animation-delay-2000"></div>
        <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-80 h-80 bg-sky-200 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-blob animation-delay-4000"></div>
      </div>
      
      <div className="container mx-auto px-6 py-16 relative z-10">
        <div className="max-w-7xl mx-auto">
          
          {/* Main Content Card */}
          <div className="bg-white/80 backdrop-blur-md rounded-3xl shadow-2xl p-10 md:p-16 mb-12">
            
            {/* Logo Section */}
            <div className="flex justify-center mb-10">
              <div className="w-full max-w-2xl">
                <WebSAC3Logo />
              </div>
            </div>

            {/* Main Title */}
            <h1 className="text-4xl md:text-5xl lg:text-6xl font-bold text-gray-900 mb-6 leading-tight text-center">
              Análisis del Componente de
              <span className="block mt-2 bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent">
                Ciberseguridad en Programas de Grado
              </span>
            </h1>

            {/* Subtitle */}
            <p className="text-lg md:text-xl text-gray-600 mb-10 max-w-4xl mx-auto leading-relaxed text-center">
              Plataforma especializada para recopilación y análisis del componente de ciberseguridad 
              en programas académicos de grado.
            </p>

            {/* CTA Buttons */}
            <div className="flex flex-col sm:flex-row gap-4 justify-center items-center mb-12">
              <Link href="/login">
                <Button size="lg" className="bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white px-10 py-6 text-lg font-semibold rounded-xl shadow-lg hover:shadow-2xl transition-all duration-300 transform hover:scale-105">
                  Iniciar Sesión
                  <ArrowRight className="ml-2 h-5 w-5" />
                </Button>
              </Link>
              <Link href="/solicitar-acceso">
                <Button variant="outline" size="lg" className="border-2 border-blue-600 text-blue-600 hover:bg-blue-600 hover:text-white px-10 py-6 text-lg font-semibold rounded-xl transition-all duration-300 transform hover:scale-105">
                  Solicitar Acceso
                </Button>
              </Link>
            </div>

            {/* Feature Cards Grid */}
            <div className="grid grid-cols-2 md:grid-cols-4 gap-6 mb-12">
              <div className="flex flex-col items-center p-6 bg-gradient-to-br from-blue-50 to-blue-100 rounded-2xl hover:shadow-lg transition-all duration-300 transform hover:-translate-y-1">
                <div className="bg-white p-4 rounded-xl mb-4 shadow-md">
                  <Shield className="h-8 w-8 text-blue-600" />
                </div>
                <h3 className="font-bold text-gray-800 text-center mb-1">Análisis Especializado</h3>
                <p className="text-xs text-gray-600 text-center">Componente ciberseguridad</p>
              </div>
              <div className="flex flex-col items-center p-6 bg-gradient-to-br from-green-50 to-green-100 rounded-2xl hover:shadow-lg transition-all duration-300 transform hover:-translate-y-1">
                <div className="bg-white p-4 rounded-xl mb-4 shadow-md">
                  <BookOpen className="h-8 w-8 text-green-600" />
                </div>
                <h3 className="font-bold text-gray-800 text-center mb-1">Programas de Grado</h3>
                <p className="text-xs text-gray-600 text-center">Evaluación académica</p>
              </div>
              <div className="flex flex-col items-center p-6 bg-gradient-to-br from-purple-50 to-purple-100 rounded-2xl hover:shadow-lg transition-all duration-300 transform hover:-translate-y-1">
                <div className="bg-white p-4 rounded-xl mb-4 shadow-md">
                  <BarChart3 className="h-8 w-8 text-purple-600" />
                </div>
                <h3 className="font-bold text-gray-800 text-center mb-1">Evaluación</h3>
                <p className="text-xs text-gray-600 text-center">Métricas detalladas</p>
              </div>
              <div className="flex flex-col items-center p-6 bg-gradient-to-br from-orange-50 to-orange-100 rounded-2xl hover:shadow-lg transition-all duration-300 transform hover:-translate-y-1">
                <div className="bg-white p-4 rounded-xl mb-4 shadow-md">
                  <Users className="h-8 w-8 text-orange-600" />
                </div>
                <h3 className="font-bold text-gray-800 text-center mb-1">Comunidad Académica</h3>
                <p className="text-xs text-gray-600 text-center">Universidad de los Llanos</p>
              </div>
            </div>

            {/* Sponsors Section */}
            <div className="border-t border-gray-200 pt-10">
              <p className="text-sm uppercase tracking-wider text-gray-500 font-semibold mb-8 text-center">
                Proyecto de proyección social
              </p>
              <div className="flex flex-col sm:flex-row items-center justify-center gap-16">
                <div className="w-80 h-36 flex items-center justify-center opacity-80 hover:opacity-100 transition-opacity duration-300 transform hover:scale-105">
                  <UnillanosLogo />
                </div>
                <div className="w-80 h-36 flex items-center justify-center opacity-80 hover:opacity-100 transition-opacity duration-300 transform hover:scale-105">
                  <FCBILogo />
                </div>
              </div>
            </div>
          </div>

        </div>
      </div>

      <style jsx>{`
        @keyframes blob {
          0%, 100% { transform: translate(0, 0) scale(1); }
          33% { transform: translate(30px, -50px) scale(1.1); }
          66% { transform: translate(-20px, 20px) scale(0.9); }
        }
        .animate-blob {
          animation: blob 7s infinite;
        }
        .animation-delay-2000 {
          animation-delay: 2s;
        }
        .animation-delay-4000 {
          animation-delay: 4s;
        }
      `}</style>
    </section>
  );
};

export { HeroSection };
