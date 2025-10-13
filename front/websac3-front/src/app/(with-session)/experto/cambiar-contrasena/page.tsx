"use client";

import { ChangePasswordForm } from "@/components/websac3/form/change-password/ChangePasswordForm";

export default function ExpertoChangePasswordPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="max-w-2xl mx-auto">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            Cambiar Contraseña
          </h1>
          <p className="text-gray-600">
            Actualiza tu contraseña para mantener la seguridad de tu cuenta de experto en ciberseguridad.
          </p>
        </div>

        {/* Card del formulario */}
        <div className="bg-white rounded-lg shadow-md border border-gray-200 p-8">
          <ChangePasswordForm />
        </div>
      </div>
    </div>
  );
}


