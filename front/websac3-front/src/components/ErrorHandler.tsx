"use client";

import { useSelector, useDispatch } from "react-redux";
import { RootState } from "@/store/store";
import { hideError } from "@/store/errorSlice";
import { ErrorModal } from "@/components/ui/error-modal";
import { NotFoundModal } from "@/components/ui/not-found-modal";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { CheckCircle } from "lucide-react";

export function ErrorHandler() {
  const dispatch = useDispatch();
  const { isOpen, type, title, message, errors, searchTerm, onRetry, statusCode } = useSelector((state: RootState) => state.error);

  const handleClose = () => {
    dispatch(hideError());
  };

  const handleRetry = () => {
    if (onRetry) {
      onRetry();
    }
    dispatch(hideError());
  };

  if (!isOpen) return null;

  // Show different modals based on type
  if (type === "not_found") {
    return (
      <NotFoundModal
        isOpen={isOpen}
        onClose={handleClose}
        onRetry={onRetry ? handleRetry : undefined}
        statusCode={statusCode}
        message={message}
        searchTerm={searchTerm}
      />
    );
  }

  if (type === "success") {
    return (
      <Dialog open={isOpen} onOpenChange={handleClose}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2 text-green-600">
              <CheckCircle className="h-5 w-5" />
              {title || "Éxito"}
            </DialogTitle>
            <DialogDescription className="text-gray-600">
              {message}
            </DialogDescription>
          </DialogHeader>
        </DialogContent>
      </Dialog>
    );
  }

  // Default error modal for other error types
  return (
    <ErrorModal
      isOpen={isOpen}
      onClose={handleClose}
      statusCode={statusCode}
      message={message}
      errors={errors}
    />
  );
}
