"use client";

import { LoginForm } from "@/components/websac3/form/login/LoginForm";

export default function LoginPage() {
  return (
    <div className="absolute flex justify-center h-screen w-screen items-center">
      <div className="flex h-3/5 w-1/2 bg-white rounded-lg mt-12 shadow-2xl shadow-slate-800">
        <div className="flex flex-col w-full h-full pl-7 pr-7">
          <div className="px-8 py-6 border-b border-gray-100">
            <h1 className="text-2xl font-semibold text-gray-900 text-center">
              Iniciar Sesión
            </h1>
          </div>
          <div className="pt-5">
            <LoginForm />
          </div>
        </div>
      </div>
    </div>
  );
} 