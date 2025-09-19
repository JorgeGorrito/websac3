"use client";

import React, { createContext, useContext, useState, ReactNode } from "react";
import { ErrorModal } from "@/components/ui/error-modal";

type ErrorState = {
  isOpen: boolean;
  title?: string;
  message: string;
  errors: string[];
};

type ErrorContextType = {
  showError: (message: string, errors?: string[], title?: string) => void;
  hideError: () => void;
};

const ErrorContext = createContext<ErrorContextType | undefined>(undefined);

type ErrorProviderProps = {
  children: ReactNode;
};

export function ErrorProvider({ children }: ErrorProviderProps) {
  const [errorState, setErrorState] = useState<ErrorState>({
    isOpen: false,
    message: "",
    errors: [],
  });

  const showError = (message: string, errors: string[] = [], title?: string) => {
    setErrorState({
      isOpen: true,
      title,
      message,
      errors,
    });
  };

  const hideError = () => {
    setErrorState({
      isOpen: false,
      message: "",
      errors: [],
    });
  };

  return (
    <ErrorContext.Provider value={{ showError, hideError }}>
      {children}
      <ErrorModal
        isOpen={errorState.isOpen}
        onClose={hideError}
        title={errorState.title}
        message={errorState.message}
        errors={errorState.errors}
      />
    </ErrorContext.Provider>
  );
}

export function useError() {
  const context = useContext(ErrorContext);
  if (context === undefined) {
    throw new Error("useError must be used within an ErrorProvider");
  }
  return context;
}
