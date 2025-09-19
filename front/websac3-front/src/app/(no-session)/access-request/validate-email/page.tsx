"use client";

import { useEffect, useState } from "react";
import { useValidateAccessRequestEmailMutation } from "@/services/api";
import { Button } from "@/components/ui/button";
import { CheckCircle, XCircle, Mail } from "lucide-react";

export default function ValidateEmailPage() {
  const [status, setStatus] = useState<"idle" | "loading" | "success" | "error">("idle");
  const [message, setMessage] = useState<string>("");
  const [token, setToken] = useState<string>("");
  const [validateEmail] = useValidateAccessRequestEmailMutation();

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const urlToken = params.get("token") || "";
    if (!urlToken) {
      setStatus("error");
      setMessage("Token no encontrado en la URL");
      return;
    }
    setToken(urlToken);
  }, []);

  const handleValidateEmail = async () => {
    if (!token) return;
    
    setStatus("loading");
    setMessage("");
    
    try {
      const response = await validateEmail({ validation_token: token }).unwrap();
      setStatus("success");
      setMessage(response || "Correo validado correctamente. Ya puedes continuar.");
    } catch (error: any) {
      setStatus("error");
      if (error?.data?.errors && Array.isArray(error.data.errors)) {
        setMessage(error.data.errors.join(", "));
      } else if (error?.data?.result) {
        setMessage(error.data.result);
      } else {
        setMessage("Token inválido o expirado.");
      }
    }
  };

  return (
    <div className="w-full flex items-center justify-center px-4 py-12">
      <div className="w-full max-w-sm bg-white rounded-lg shadow-2xl shadow-slate-800 overflow-hidden">
        {/* Header */}
        <div className="bg-blue-50 px-6 py-4 border-b border-blue-100">
          <div className="flex items-center justify-center">
            <div className="flex items-center justify-center h-10 w-10 rounded-full bg-blue-100 mr-3">
              <Mail className="h-5 w-5 text-blue-600" />
            </div>
            <div>
              <h1 className="text-lg font-semibold text-gray-900">
                Validación de Correo
              </h1>
              <p className="text-xs text-gray-600">
                Confirma tu email para completar la solicitud
              </p>
            </div>
          </div>
        </div>

        {/* Content */}
        <div className="p-6">
          {status === "idle" && (
            <div className="text-center">
              <p className="text-gray-600 mb-4 text-sm">
                Haz clic en el botón para validar tu correo electrónico
              </p>
              <Button
                onClick={handleValidateEmail}
                className="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2.5 px-4 rounded-lg transition-colors"
              >
                Validar Email
              </Button>
            </div>
          )}

          {status === "loading" && (
            <div className="text-center py-4">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto mb-3"></div>
              <p className="text-gray-600 text-sm">Validando tu correo electrónico...</p>
            </div>
          )}

          {status === "success" && (
            <div className="text-center">
              <div className="flex items-center justify-center h-12 w-12 rounded-full bg-green-100 mx-auto mb-4">
                <CheckCircle className="h-6 w-6 text-green-600" />
              </div>
              <h3 className="text-lg font-medium text-gray-900 mb-2">
                ¡Email Validado Exitosamente!
              </h3>
              <p className="text-sm text-gray-600 mb-4">
                {message}
              </p>
              <div className="bg-blue-50 border border-blue-200 rounded-lg p-4 mb-6">
                <p className="text-sm text-blue-800">
                  <strong>Próximos pasos:</strong> Tu solicitud de acceso será revisada por un administrador. 
                  Recibirás una notificación por correo electrónico una vez que se apruebe o rechace tu solicitud.
                </p>
              </div>
              <Button
                onClick={() => window.location.href = "/login"}
                className="w-full bg-green-600 hover:bg-green-700 text-white font-medium py-2.5 px-4 rounded-lg transition-colors"
              >
                Ir al Login
              </Button>
            </div>
          )}

          {status === "error" && (
            <div className="text-center">
              <div className="flex items-center justify-center h-12 w-12 rounded-full bg-red-100 mx-auto mb-4">
                <XCircle className="h-6 w-6 text-red-600" />
              </div>
              <h3 className="text-lg font-medium text-gray-900 mb-2">
                Error de Validación
              </h3>
              <p className="text-sm text-gray-600 mb-6">
                {message}
              </p>
              <div className="space-y-3">
                <Button
                  onClick={handleValidateEmail}
                  className="w-full bg-red-600 hover:bg-red-700 text-white font-medium py-2.5 px-4 rounded-lg transition-colors"
                >
                  Intentar de Nuevo
                </Button>
                <Button
                  onClick={() => window.location.href = "/access-request"}
                  variant="outline"
                  className="w-full border-gray-300 text-gray-700 hover:bg-gray-50 font-medium py-2.5 px-4 rounded-lg transition-colors"
                >
                  Nueva Solicitud
                </Button>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
