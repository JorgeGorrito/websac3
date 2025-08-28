"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Eye, EyeOff, User } from "lucide-react";
import { SlideToSubmit } from "../SlideToSubmit";

export const LoginForm = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [slideCompleted, setSlideCompleted] = useState(false);

  const handleSlideComplete = () => {
    setSlideCompleted(true);
  };

  const handleLogin = () => {
    if (slideCompleted) {
      alert("¡Login exitoso!");
    } else {
      alert("¡Debes completar el slide para iniciar sesión!");
    }
  };

  const handleAccessRequest = () => {
    window.location.href = "/access-request";
  };

  return (
    <div className="flex flex-col items-center space-y-4 px-6 py-4">
      {/* Icono de usuario */}
      <div className="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center border-2 border-blue-300">
        <User className="w-6 h-6 text-blue-500" />
      </div>

      {/* Formulario */}
      <div className="w-full space-y-3">
        {/* Campo de usuario */}
        <div className="space-y-1">
          <label className="text-sm font-medium text-gray-700">Usuario</label>
          <Input
            type="text"
            placeholder="Ingresa tu usuario"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="w-full"
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
            />
            <button
              type="button"
              onClick={() => setShowPassword(!showPassword)}
              className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
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
          label="Slide to Submit"
          completedText="¡Completado!"
          slideText="Arrastra para continuar"
        />

        {/* Botones */}
        <div className="flex space-x-3 pt-2">
          <Button onClick={handleLogin} className="flex-1">
            Iniciar sesión
          </Button>
          <Button
            onClick={handleAccessRequest}
            variant="outline"
            className="flex-1"
          >
            Solicitar acceso
          </Button>
        </div>
      </div>
    </div>
  );
};
