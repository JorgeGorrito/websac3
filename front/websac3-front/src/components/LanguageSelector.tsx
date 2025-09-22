"use client";

import { Button } from "@/components/ui/button";
import { setLanguage, getCurrentLanguage } from "@/services/api";
import { useState, useEffect } from "react";

export function LanguageSelector() {
  const [currentLang, setCurrentLang] = useState<string>("es");

  useEffect(() => {
    setCurrentLang(getCurrentLanguage());
  }, []);

  const handleLanguageChange = (language: string) => {
    setLanguage(language);
    setCurrentLang(language);
    // Reload the page to apply the new language
    window.location.reload();
  };

  return (
    <div className="flex gap-2 items-center">
      <span className="text-sm text-gray-600">Idioma:</span>
      <Button
        variant={currentLang === "es" ? "default" : "outline"}
        size="sm"
        onClick={() => handleLanguageChange("es")}
      >
        ES
      </Button>
      <Button
        variant={currentLang === "en" ? "default" : "outline"}
        size="sm"
        onClick={() => handleLanguageChange("en")}
      >
        EN
      </Button>
    </div>
  );
}
