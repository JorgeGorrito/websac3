"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useChangePasswordMutation } from "@/services/api";
import { 
  Lock, 
  Eye, 
  EyeOff, 
  CheckCircle, 
  XCircle, 
  X,
  Shield
} from "lucide-react";

export function ChangePasswordForm() {
  const router = useRouter();
  const [changePassword, { isLoading }] = useChangePasswordMutation();

  const [formData, setFormData] = useState({
    current_password: "",
    new_password: "",
    confirm_new_password: "",
  });

  const [showPasswords, setShowPasswords] = useState({
    current: false,
    new: false,
    confirm: false,
  });

  const [modalState, setModalState] = useState<{
    isOpen: boolean;
    type: "success" | "error";
    title: string;
    message: string;
  }>({
    isOpen: false,
    type: "success",
    title: "",
    message: "",
  });

  const handleInputChange = (field: string, value: string) => {
    setFormData(prev => ({ ...prev, [field]: value }));
  };

  const togglePasswordVisibility = (field: "current" | "new" | "confirm") => {
    setShowPasswords(prev => ({ ...prev, [field]: !prev[field] }));
  };

  const validateForm = (): string | null => {
    if (!formData.current_password) {
      return "Por favor, ingrese su contraseña actual";
    }
    if (!formData.new_password) {
      return "Por favor, ingrese una nueva contraseña";
    }
    if (formData.new_password.length < 8) {
      return "La nueva contraseña debe tener al menos 8 caracteres";
    }
    if (!formData.confirm_new_password) {
      return "Por favor, confirme su nueva contraseña";
    }
    if (formData.new_password !== formData.confirm_new_password) {
      return "Las contraseñas nuevas no coinciden";
    }
    if (formData.current_password === formData.new_password) {
      return "La nueva contraseña debe ser diferente a la actual";
    }
    return null;
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // Validate form
    const validationError = validateForm();
    if (validationError) {
      setModalState({
        isOpen: true,
        type: "error",
        title: "Error de validación",
        message: validationError,
      });
      return;
    }

    try {
      const response = await changePassword(formData).unwrap();

      // Handle success
      setModalState({
        isOpen: true,
        type: "success",
        title: "¡Contraseña Actualizada!",
        message: response.result || "Su contraseña ha sido actualizada exitosamente.",
      });

      // Reset form
      setFormData({
        current_password: "",
        new_password: "",
        confirm_new_password: "",
      });

      // Redirect after 2 seconds
      setTimeout(() => {
        router.back();
      }, 2000);
    } catch (err: any) {
      console.error("Error changing password:", err);
      let errorMessage = "Error al cambiar la contraseña";
      
      if (err?.data?.errors && Array.isArray(err.data.errors)) {
        errorMessage = err.data.errors.join(", ");
      } else if (err?.data?.result) {
        errorMessage = err.data.result;
      }
      
      setModalState({
        isOpen: true,
        type: "error",
        title: "Error al Cambiar Contraseña",
        message: errorMessage,
      });
    }
  };

  return (
    <div className="max-w-2xl mx-auto">
      <Card>
        <CardHeader>
          <div className="flex items-center gap-3">
            <div className="p-2 bg-blue-100 rounded-lg">
              <Shield className="h-6 w-6 text-blue-600" />
            </div>
            <div>
              <CardTitle className="text-xl">Cambiar Contraseña</CardTitle>
              <CardDescription className="mt-1">
                Actualice su contraseña para mantener su cuenta segura
              </CardDescription>
            </div>
          </div>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-6">
            {/* Current Password */}
            <div className="space-y-2">
              <Label htmlFor="current_password" className="flex items-center gap-2">
                <Lock className="h-4 w-4 text-gray-500" />
                Contraseña Actual
              </Label>
              <div className="relative">
                <Input
                  id="current_password"
                  type={showPasswords.current ? "text" : "password"}
                  value={formData.current_password}
                  onChange={(e) => handleInputChange("current_password", e.target.value)}
                  placeholder="Ingrese su contraseña actual"
                  className="pr-10"
                  required
                />
                <button
                  type="button"
                  onClick={() => togglePasswordVisibility("current")}
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700"
                >
                  {showPasswords.current ? (
                    <EyeOff className="h-4 w-4" />
                  ) : (
                    <Eye className="h-4 w-4" />
                  )}
                </button>
              </div>
            </div>

            {/* New Password */}
            <div className="space-y-2">
              <Label htmlFor="new_password" className="flex items-center gap-2">
                <Lock className="h-4 w-4 text-gray-500" />
                Nueva Contraseña
              </Label>
              <div className="relative">
                <Input
                  id="new_password"
                  type={showPasswords.new ? "text" : "password"}
                  value={formData.new_password}
                  onChange={(e) => handleInputChange("new_password", e.target.value)}
                  placeholder="Ingrese su nueva contraseña"
                  className="pr-10"
                  required
                />
                <button
                  type="button"
                  onClick={() => togglePasswordVisibility("new")}
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700"
                >
                  {showPasswords.new ? (
                    <EyeOff className="h-4 w-4" />
                  ) : (
                    <Eye className="h-4 w-4" />
                  )}
                </button>
              </div>
              <p className="text-xs text-gray-500">
                La contraseña debe tener al menos 8 caracteres
              </p>
            </div>

            {/* Confirm New Password */}
            <div className="space-y-2">
              <Label htmlFor="confirm_new_password" className="flex items-center gap-2">
                <Lock className="h-4 w-4 text-gray-500" />
                Confirmar Nueva Contraseña
              </Label>
              <div className="relative">
                <Input
                  id="confirm_new_password"
                  type={showPasswords.confirm ? "text" : "password"}
                  value={formData.confirm_new_password}
                  onChange={(e) => handleInputChange("confirm_new_password", e.target.value)}
                  placeholder="Confirme su nueva contraseña"
                  className="pr-10"
                  required
                />
                <button
                  type="button"
                  onClick={() => togglePasswordVisibility("confirm")}
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700"
                >
                  {showPasswords.confirm ? (
                    <EyeOff className="h-4 w-4" />
                  ) : (
                    <Eye className="h-4 w-4" />
                  )}
                </button>
              </div>
            </div>

            {/* Security Tips */}
            <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
              <h4 className="text-sm font-semibold text-blue-900 mb-2">Consejos de seguridad:</h4>
              <ul className="text-xs text-blue-800 space-y-1">
                <li>• Use al menos 8 caracteres</li>
                <li>• Combine letras mayúsculas y minúsculas</li>
                <li>• Incluya números y caracteres especiales</li>
                <li>• No use contraseñas fáciles de adivinar</li>
                <li>• No reutilice contraseñas de otras cuentas</li>
              </ul>
            </div>

            {/* Action Buttons */}
            <div className="flex justify-end gap-3 pt-4">
              <Button
                type="button"
                variant="outline"
                onClick={() => router.back()}
                disabled={isLoading}
              >
                Cancelar
              </Button>
              <Button type="submit" disabled={isLoading}>
                {isLoading ? "Actualizando..." : "Cambiar Contraseña"}
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>

      {/* Modal */}
      {modalState.isOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center">
          {/* Overlay */}
          <div 
            className="absolute inset-0 bg-black/50"
            onClick={() => setModalState(prev => ({ ...prev, isOpen: false }))}
          />
          
          {/* Modal Content */}
          <div className="relative bg-white rounded-lg shadow-xl max-w-md w-full mx-4 p-6">
            {/* Close Button */}
            <button
              onClick={() => setModalState(prev => ({ ...prev, isOpen: false }))}
              className="absolute top-4 right-4 p-1 rounded-full hover:bg-gray-100 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-300"
              aria-label="Cerrar modal"
            >
              <X className="h-5 w-5 text-gray-500 hover:text-gray-700" />
            </button>
            
            {/* Header */}
            <div className="flex items-center gap-3 pr-8 mb-4">
              {modalState.type === "success" ? (
                <div className="flex items-center justify-center h-10 w-10 rounded-full bg-green-100">
                  <CheckCircle className="h-5 w-5 text-green-600" />
                </div>
              ) : (
                <div className="flex items-center justify-center h-10 w-10 rounded-full bg-red-100">
                  <XCircle className="h-5 w-5 text-red-600" />
                </div>
              )}
              <h2 className={`text-lg font-semibold ${modalState.type === "success" ? "text-green-800" : "text-red-800"}`}>
                {modalState.title}
              </h2>
            </div>
            
            {/* Message */}
            <p className="text-gray-600 mb-6">
              {modalState.message}
            </p>
            
            {/* Footer */}
            <div className="flex justify-end">
              <Button
                onClick={() => {
                  setModalState(prev => ({ ...prev, isOpen: false }));
                  if (modalState.type === "success") {
                    router.back();
                  }
                }}
                className={modalState.type === "success" ? "bg-green-600 hover:bg-green-700" : "bg-red-600 hover:bg-red-700"}
              >
                {modalState.type === "success" ? "Entendido" : "Cerrar"}
              </Button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
