"use client";

import { useSelector, useDispatch } from "react-redux";
import { RootState } from "@/store/store";
import { hideError } from "@/store/errorSlice";
import { ErrorModal } from "@/components/ui/error-modal";
import { NotFoundModal } from "@/components/ui/not-found-modal";

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

  // Show different modals based on error type
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
