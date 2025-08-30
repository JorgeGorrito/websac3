"use client";

export default function ValidateEmailPage() {
  return (
    <div className="absolute flex justify-center h-screen w-screen items-center">
      <div className="flex h-3/5 w-1/2 bg-white rounded-lg mt-12 shadow-2xl shadow-slate-800">
        <div className="flex flex-col w-full h-full pl-7 pr-7">
          <div className="px-8 py-6 border-b border-gray-100">
            <h1 className="text-2xl font-semibold text-gray-900 text-center">
              Validación de correo
            </h1>
          </div>
          <div className="p-6">
            <ValidateEmailClient />
          </div>
        </div>
      </div>
    </div>
  );
}

import { useEffect, useState } from "react";
import { useValidateAccessRequestEmailMutation } from "@/services/api";

function ValidateEmailClient() {
  const [status, setStatus] = useState<"loading" | "success" | "error">(
    "loading"
  );
  const [validateEmail] = useValidateAccessRequestEmailMutation();

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const token = params.get("token") || "";
    if (!token) {
      setStatus("error");
      return;
    }
    (async () => {
      try {
        await validateEmail({ validation_token: token }).unwrap();
        setStatus("success");
      } catch {
        setStatus("error");
      }
    })();
  }, [validateEmail]);

  if (status === "loading")
    return <p className="text-sm text-gray-600">Validando token…</p>;
  if (status === "success")
    return (
      <p className="text-green-600">
        Correo validado correctamente. Ya puedes continuar.
      </p>
    );
  return <p className="text-red-600">Token inválido o expirado.</p>;
}
