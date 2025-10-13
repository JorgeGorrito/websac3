"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Eye, EyeOff, Lock, CheckCircle } from "lucide-react";
import { useChangePasswordMutation } from "@/services/api";

export const ChangePasswordForm = () => {
  const [showCurrentPassword, setShowCurrentPassword] = useState(false);
  const [showNewPassword, setShowNewPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  
  const [changePasswordMutation, { isLoading }] = useChangePasswordMutation();

  const validatePasswords = () => {
    if (!currentPassword) {
      setError("Debes ingresar tu contraseña actual");
      return false;
    }
    
    if (newPassword.length < 8) {
      setError("La nueva contraseña debe tener al menos 8 caracteres");
      return false;
    }
    
    if (newPassword === currentPassword) {
      setError("La nueva contraseña debe ser diferente a la actual");
      return false;
    }

    // Validar que incluya letras, números y símbolos
    const hasLetter = /[a-zA-Z]/.test(newPassword);
    const hasNumber = /[0-9]/.test(newPassword);
    const hasSymbol = /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(newPassword);
    
    if (!hasLetter || !hasNumber || !hasSymbol) {
      setError("La contraseña debe incluir una combinación de letras, números y símbolos");
      return false;
    }
    
    if (newPassword !== confirmPassword) {
      setError("Las contraseñas nuevas no coinciden");
      return false;
    }
    
    return true;
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    
    if (!validatePasswords()) {
      return;
    }

    setError("");
    setSuccess("");

    try {
      await changePasswordMutation({
        current_password: currentPassword,
        new_password: newPassword,
        confirm_new_password: confirmPassword,
      }).unwrap();
      
      setSuccess("¡Contraseña cambiada exitosamente!");
      
      // Limpiar el formulario
      setCurrentPassword("");
      setNewPassword("");
      setConfirmPassword("");
      
    } catch (error: any) {
      // Manejar errores de la API
      if (error?.data?.errors && Array.isArray(error.data.errors)) {
        setError(error.data.errors[0] || "Error al cambiar la contraseña");
      } else if (error?.message) {
        setError(error.message);
      } else {
        setError("Error al cambiar la contraseña. Inténtalo de nuevo.");
      }
    }
  };

  return (
    <div className="w-full max-w-md mx-auto">
      {/* Mensaje de éxito */}
      {success && (
        <div className="mb-4 p-3 bg-green-50 border border-green-200 rounded-lg flex items-center gap-2">
          <CheckCircle className="w-5 h-5 text-green-600" />
          <p className="text-sm text-green-800">{success}</p>
        </div>
      )}

      {/* Mensaje de error */}
      {error && (
        <div className="mb-4 p-3 bg-red-50 border border-red-200 rounded-lg">
          <p className="text-sm text-red-800">{error}</p>
        </div>
      )}

      {/* Formulario */}
      <form onSubmit={handleSubmit} className="w-full space-y-4">
        {/* Contraseña actual */}
        <div className="space-y-2">
          <label className="text-sm font-medium text-gray-700">
            Contraseña Actual
          </label>
          <div className="relative">
            <Input
              type={showCurrentPassword ? "text" : "password"}
              placeholder="Ingresa tu contraseña actual"
              value={currentPassword}
              onChange={(e) => setCurrentPassword(e.target.value)}
              className="w-full pr-10"
              disabled={isLoading}
              required
            />
            <button
              type="button"
              onClick={() => setShowCurrentPassword(!showCurrentPassword)}
              className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              disabled={isLoading}
            >
              {showCurrentPassword ? (
                <EyeOff className="w-4 h-4" />
              ) : (
                <Eye className="w-4 h-4" />
              )}
            </button>
          </div>
        </div>

        {/* Nueva contraseña */}
        <div className="space-y-2">
          <label className="text-sm font-medium text-gray-700">
            Nueva Contraseña
          </label>
          <div className="relative">
            <Input
              type={showNewPassword ? "text" : "password"}
              placeholder="Ingresa tu nueva contraseña"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
              className="w-full pr-10"
              disabled={isLoading}
              required
            />
            <button
              type="button"
              onClick={() => setShowNewPassword(!showNewPassword)}
              className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              disabled={isLoading}
            >
              {showNewPassword ? (
                <EyeOff className="w-4 h-4" />
              ) : (
                <Eye className="w-4 h-4" />
              )}
            </button>
          </div>
        </div>

        {/* Confirmar nueva contraseña */}
        <div className="space-y-2">
          <label className="text-sm font-medium text-gray-700">
            Confirmar Nueva Contraseña
          </label>
          <div className="relative">
            <Input
              type={showConfirmPassword ? "text" : "password"}
              placeholder="Confirma tu nueva contraseña"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              className="w-full pr-10"
              disabled={isLoading}
              required
            />
            <button
              type="button"
              onClick={() => setShowConfirmPassword(!showConfirmPassword)}
              className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              disabled={isLoading}
            >
              {showConfirmPassword ? (
                <EyeOff className="w-4 h-4" />
              ) : (
                <Eye className="w-4 h-4" />
              )}
            </button>
          </div>
        </div>

        {/* Botones */}
        <div className="flex space-x-3 pt-4">
          <Button
            type="submit"
            disabled={isLoading}
            className="flex-1 bg-primary hover:bg-primary/90 text-white"
          >
            {isLoading ? (
              <div className="flex items-center gap-2">
                <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                <span>Cambiando...</span>
              </div>
            ) : (
              <div className="flex items-center gap-2">
                <Lock className="w-4 h-4" />
                <span>Cambiar Contraseña</span>
              </div>
            )}
          </Button>
        </div>
      </form>

      {/* Información adicional */}
      <div className="mt-6 p-4 bg-blue-50 border border-blue-200 rounded-lg">
        <h4 className="text-sm font-medium text-blue-900 mb-2">Requisitos de contraseña:</h4>
        <ul className="text-xs text-blue-800 space-y-1">
          <li>• Tener un mínimo de 8 caracteres</li>
          <li>• Ser diferente a la contraseña actual</li>
          <li>• Incluir una combinación de letras, números y símbolos</li>
        </ul>
      </div>
    </div>
  );
};
