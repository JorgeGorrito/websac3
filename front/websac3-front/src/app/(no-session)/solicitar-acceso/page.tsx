"use client";

import { AccessRequestForm } from "@/components/websac3/form/access-request/AccessRequestForm";

export default function AccessRequestPage() {
  return (
    <div className=" w-full flex items-center justify-center px-4 py-8">
      <div className="w-full max-w-3xl bg-white rounded-lg shadow-2xl shadow-slate-800">
        <div className="flex flex-col w-full h-full px-6 sm:px-8">
          <div className="px-2 sm:px-4 py-6 border-b border-gray-100">
            <h1 className="text-2xl font-semibold text-gray-900 text-center">
              Solicitud de Acceso
            </h1>
          </div>
          <div className="py-6 overflow-auto">
            <AccessRequestForm />
          </div>
        </div>
      </div>
    </div>
  );
}
