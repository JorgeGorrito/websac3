"use client";

import { ChangePasswordForm } from "@/components/websac3/profile/ChangePasswordForm";

export default function CambiarContrasenaPage() {
  return (
    <div className="space-y-6">
      {/* Header */}
      <div>
        <h1 className="text-3xl font-bold text-gray-900">Cambiar Contraseña</h1>
        <p className="text-gray-600 mt-2">Actualice su contraseña para mantener su cuenta segura</p>
      </div>

      {/* Form */}
      <ChangePasswordForm />
    </div>
  );
}
