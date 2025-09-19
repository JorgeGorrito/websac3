import { X } from "lucide-react";
import { Button } from "@/components/ui/button";
import { getErrorInfo } from "@/utils/errorMessages";

type ErrorModalProps = {
  isOpen: boolean;
  onClose: () => void;
  statusCode?: number;
  message: string;
  errors?: string[];
};

export function ErrorModal({ 
  isOpen, 
  onClose, 
  statusCode,
  message, 
  errors = [] 
}: ErrorModalProps) {
  if (!isOpen) return null;

  const errorInfo = statusCode ? getErrorInfo(statusCode) : {
    title: "Error Inesperado",
    description: "Ha ocurrido un problema inesperado",
    icon: "❓",
    color: "gray"
  };

  const getColorClasses = (color: string) => {
    switch (color) {
      case "red":
        return {
          bg: "bg-red-100",
          text: "text-red-800",
          border: "border-red-200",
          button: "bg-red-600 hover:bg-red-700"
        };
      case "orange":
        return {
          bg: "bg-orange-100",
          text: "text-orange-800",
          border: "border-orange-200",
          button: "bg-orange-600 hover:bg-orange-700"
        };
      case "yellow":
        return {
          bg: "bg-yellow-100",
          text: "text-yellow-800",
          border: "border-yellow-200",
          button: "bg-yellow-600 hover:bg-yellow-700"
        };
      case "blue":
        return {
          bg: "bg-blue-100",
          text: "text-blue-800",
          border: "border-blue-200",
          button: "bg-blue-600 hover:bg-blue-700"
        };
      default:
        return {
          bg: "bg-gray-100",
          text: "text-gray-800",
          border: "border-gray-200",
          button: "bg-gray-600 hover:bg-gray-700"
        };
    }
  };

  const colors = getColorClasses(errorInfo.color);

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* Overlay */}
      <div 
        className="absolute inset-0 bg-black/50"
        onClick={onClose}
      />
      
      {/* Modal Content */}
      <div className="relative bg-white rounded-lg shadow-xl max-w-lg w-full mx-4 p-6">
        {/* Close Button */}
        <button
          onClick={onClose}
          className="absolute top-4 right-4 p-1 rounded-full hover:bg-gray-100 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-300"
          aria-label="Cerrar modal"
        >
          <X className="h-5 w-5 text-gray-500 hover:text-gray-700" />
        </button>
        
        {/* Header */}
        <div className="flex items-center gap-3 pr-8 mb-4">
          <div className={`flex items-center justify-center h-12 w-12 rounded-full ${colors.bg}`}>
            <span className="text-2xl">{errorInfo.icon}</span>
          </div>
          <div>
            <h2 className={`text-lg font-semibold ${colors.text}`}>
              {errorInfo.title}
            </h2>
            <p className="text-sm text-gray-600">
              {errorInfo.description}
            </p>
          </div>
        </div>
        
        {/* Backend Message */}
        <div className="mb-6 p-3 bg-gray-50 rounded-lg border-l-4 border-gray-300">
          <h3 className="text-sm font-medium text-gray-700 mb-2">
            Detalle:
          </h3>
          <p className="text-sm text-gray-700">
            {message}
          </p>
        </div>
        
        {/* Footer */}
        <div className="flex justify-end">
          <Button
            onClick={onClose}
            className={`${colors.button} text-white`}
          >
            Entendido
          </Button>
        </div>
      </div>
    </div>
  );
}
