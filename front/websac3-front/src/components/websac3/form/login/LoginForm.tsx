"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Eye, EyeOff, User } from "lucide-react";
import { SlideToSubmit } from "../SlideToSubmit";
import { useLoginMutation } from "@/services/api";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { loginStart, loginSuccess, loginFailure, clearError } from "@/store/authSlice";
import { decodeJWT } from "@/lib/jwt";
import { getDashboardRoute } from "@/utils/roleUtils";

export const LoginForm = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [slideCompleted, setSlideCompleted] = useState(false);
  
  const router = useRouter();
  const dispatch = useAppDispatch();
  const { isLoading, error, isAuthenticated, user } = useAppSelector((state) => state.auth);
  const [loginMutation] = useLoginMutation();

  // Clear error when component mounts
  useEffect(() => {
    dispatch(clearError());
  }, [dispatch]);

  // Handle redirection after successful login
  useEffect(() => {
    if (isAuthenticated && user) {
      const dashboardRoute = getDashboardRoute(user.role);
      console.log(`🔐 LoginForm: User "${user.username}" with role "${user.role}" redirecting to: ${dashboardRoute}`);
      
      // Use window.location.href for a hard redirect
      window.location.href = dashboardRoute;
    }
  }, [isAuthenticated, user]);

  const handleSlideComplete = () => {
    setSlideCompleted(true);
  };

  const handleLogin = async () => {
    if (!slideCompleted) {
      alert("¡Debes completar el slide para iniciar sesión!");
      return;
    }

    if (!email || !password) {
      dispatch(loginFailure("Por favor completa todos los campos"));
      return;
    }

    try {
      dispatch(loginStart());
      const result = await loginMutation({ email, password }).unwrap();
      
      // Decode JWT to get user information
      const tokenPayload = decodeJWT(result.access_token);
      
      if (!tokenPayload) {
        dispatch(loginFailure("Error al procesar la respuesta del servidor"));
        return;
      }

      const user = {
        id: tokenPayload.sub,
        username: tokenPayload.username,
        email: email,
        role: tokenPayload.role,
        permissions: tokenPayload.permissions,
      };

      dispatch(loginSuccess({
        accessToken: result.access_token,
        refreshToken: result.refresh_token,
        user,
      }));

      // The useEffect will handle the redirection based on user role
    } catch (error: any) {
      const errorMessage = error?.data?.errors?.[0] || "Error al iniciar sesión";
      dispatch(loginFailure(errorMessage));
    }
  };

  const handleAccessRequest = () => {
    window.location.href = "/solicitar-acceso";
  };

  return (
    <div className="flex flex-col items-center space-y-4 px-6 py-4">
      {/* Icono de usuario */}
      <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center border-2 border-blue-300">
        <User className="w-6 h-6 text-blue-500" />
      </div>

      {/* Error message */}
      {error && (
        <div className="w-full p-3 bg-red-100 border border-red-300 rounded-md">
          <p className="text-sm text-red-600">{error}</p>
        </div>
      )}

      {/* Formulario */}
      <div className="w-full space-y-3">
        {/* Campo de email */}
        <div className="space-y-1">
          <label className="text-sm font-medium text-gray-700">Email</label>
          <Input
            type="email"
            placeholder="Ingresa tu email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full"
            disabled={isLoading}
          />
        </div>

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

        {/* Olvidaste tu contraseña */}
        <div className="text-center">
          <button className="text-sm text-blue-600 hover:text-blue-800 underline">
            ¿Olvidaste tu contraseña?
          </button>
        </div>

        <SlideToSubmit
          onComplete={handleSlideComplete}
          completedText="¡Completado!"
          slideText="Arrastra para continuar"
        />

        {/* Botones */}
        <div className="flex space-x-3 pt-2">
          <Button 
            onClick={handleLogin} 
            className="flex-1"
            disabled={isLoading || !slideCompleted}
          >
            {isLoading ? "Iniciando sesión..." : "Iniciar sesión"}
          </Button>
          <Button
            onClick={handleAccessRequest}
            variant="outline"
            className="flex-1"
            disabled={isLoading}
          >
            Solicitar acceso
          </Button>
        </div>
      </div>
    </div>
  );
};
