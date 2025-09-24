"use client";

import { RegisterUserForm } from "@/components/websac3/form/register-user/RegisterUserForm";

export default function RegisterUserPage() {
  return (
    <div className="absolute flex justify-center h-screen w-screen items-center">
      <div className="flex h-3/5 w-1/2 bg-white rounded-lg mt-12 shadow-2xl shadow-slate-800">
        <div className="flex flex-col w-full h-full pl-7 pr-7">
          <div className="px-8 py-6 border-b border-gray-100">
            <h1 className="text-2xl font-semibold text-gray-900 text-center">
              Crear Cuenta
            </h1>
            <p className="text-sm text-gray-600 text-center mt-2">
              Completa tu registro estableciendo una contraseña
            </p>
          </div>
          <div className="pt-5">
            <RegisterUserForm />
          </div>
        </div>
      </div>
    </div>
  );
}
