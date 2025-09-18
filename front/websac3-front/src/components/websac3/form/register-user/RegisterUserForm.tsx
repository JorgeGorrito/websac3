"use client";

import { useState, useEffect } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Eye, EyeOff, Lock, CheckCircle } from "lucide-react";
import { useCreateUserFromTokenMutation } from "@/services/api";

export const RegisterUserForm = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  
  const router = useRouter();
  const searchParams = useSearchParams();
  const [createUserMutation] = useCreateUserFromTokenMutation();

  // Extract token from URL
  const token = searchParams.get("token");

  useEffect(() => {
    if (!token) {
      setError("Token de registro no válido o faltante");
    }
  }, [token]);

  const validatePasswords = () => {
    if (password.length < 6) {
      setError("La contraseña debe tener al menos 6 caracteres");
      return false;
    }
    
    if (password !== confirmPassword) {
      setError("Las contraseñas no coinciden");
      return false;
    }
    
    return true;
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");
    setSuccess("");

    if (!token) {
      setError("Token de registro no válido");
      return;
    }

    if (!validatePasswords()) {
      return;
    }

    try {
      setIsLoading(true);
      const result = await createUserMutation({
        password,
        confirm_password: confirmPassword,
        create_user_token: token,
      }).unwrap();

      setSuccess("¡Cuenta creada exitosamente! Redirigiendo al login...");
      
      // Redirect to login after 2 seconds
      setTimeout(() => {
        router.push("/login");
      }, 2000);
      
    } catch (error: any) {
      const errorMessage = error?.data?.errors?.[0] || "Error al crear la cuenta";
      setError(errorMessage);
    } finally {
      setIsLoading(false);
    }
  };

  const handleAccessRequest = () => {
    router.push("/access-request");
  };

  return (
    <div className="flex flex-col items-center space-y-4 px-6 py-4">
      {/* Icono de candado */}
      <div className="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center border-2 border-green-300">
        <Lock className="w-6 h-6 text-green-500" />
      </div>

      {/* Error message */}
      {error && (
        <div className="w-full p-3 bg-red-100 border border-red-300 rounded-md">
          <p className="text-sm text-red-600">{error}</p>
        </div>
      )}

      {/* Success message */}
      {success && (
        <div className="w-full p-3 bg-green-100 border border-green-300 rounded-md flex items-center space-x-2">
          <CheckCircle className="w-4 h-4 text-green-600" />
          <p className="text-sm text-green-600">{success}</p>
        </div>
      )}

      {/* Formulario */}
      <form onSubmit={handleSubmit} className="w-full space-y-3">
        {/* Campo de contraseña */}
        <div className="space-y-1">
          <label className="text-sm font-medium text-gray-700">
            Contraseña
          </label>
          <div className="relative">
            <Input
              type={showPassword ? "text" : "password"}
              placeholder="Ingresa tu contraseña"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full pr-10"
              disabled={isLoading}
              required
            />
            <button
              type="button"
              onClick={() => setShowPassword(!showPassword)}
              className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              disabled={isLoading}
            >
              {showPassword ? (
                <EyeOff className="w-4 h-4" />
              ) : (
                <Eye className="w-4 h-4" />
              )}
            </button>
          </div>
        </div>

        {/* Campo de confirmar contraseña */}
        <div className="space-y-1">
          <label className="text-sm font-medium text-gray-700">
            Confirmar Contraseña
          </label>
          <div className="relative">
            <Input
              type={showConfirmPassword ? "text" : "password"}
              placeholder="Confirma tu contraseña"
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

        {/* Información de contraseña */}
        <div className="text-xs text-gray-500">
          <p>• La contraseña debe tener al menos 6 caracteres</p>
          <p>• Asegúrate de que ambas contraseñas coincidan</p>
        </div>

        {/* Botones */}
        <div className="flex space-x-3 pt-2">
          <Button 
            type="submit"
            className="flex-1"
            disabled={isLoading || !token}
          >
            {isLoading ? "Creando cuenta..." : "Crear Cuenta"}
          </Button>
          <Button
            type="button"
            onClick={handleAccessRequest}
            variant="outline"
            className="flex-1"
            disabled={isLoading}
          >
            Solicitar acceso
          </Button>
        </div>
      </form>
    </div>
  );
};
