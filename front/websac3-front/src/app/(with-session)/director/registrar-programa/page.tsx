"use client";

import { RegisterProgramForm } from "@/components/websac3/form/program/RegisterProgramForm";

export default function RegistrarProgramaPage() {
  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <h2 className="text-2xl font-semibold text-center">
            Registrar Programa
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>
        <RegisterProgramForm />
      </div>
    </div>
  );
}
