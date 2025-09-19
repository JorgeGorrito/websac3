import { Search, X, RefreshCw } from "lucide-react";
import { Button } from "@/components/ui/button";
import { getErrorInfo } from "@/utils/errorMessages";

type NotFoundModalProps = {
  isOpen: boolean;
  onClose: () => void;
  onRetry?: () => void;
  statusCode?: number;
  message?: string;
  searchTerm?: string;
};

export function NotFoundModal({ 
  isOpen, 
  onClose, 
  onRetry,
  statusCode = 404,
  message = "No se encontraron resultados para tu búsqueda",
  searchTerm
}: NotFoundModalProps) {
  const errorInfo = getErrorInfo(statusCode);
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* Overlay */}
      <div 
        className="absolute inset-0 bg-black/50"
        onClick={onClose}
      />
      
      {/* Modal Content */}
      <div className="relative bg-white rounded-lg shadow-xl max-w-md w-full mx-4 p-6">
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
          <div className="flex items-center justify-center h-12 w-12 rounded-full bg-blue-100">
            <span className="text-2xl">{errorInfo.icon}</span>
          </div>
          <div>
            <h2 className="text-lg font-semibold text-blue-800">
              {errorInfo.title}
            </h2>
            <p className="text-sm text-gray-600">
              {errorInfo.description}
            </p>
          </div>
        </div>
        
        {/* Backend Message */}
        <div className="mb-4 p-3 bg-gray-50 rounded-lg border-l-4 border-gray-300">
          <h3 className="text-sm font-medium text-gray-700 mb-2">
            Detalle:
          </h3>
          <p className="text-sm text-gray-700">
            {message}
          </p>
        </div>
        
        {/* Search Term */}
        {searchTerm && (
          <div className="mb-4 p-3 bg-gray-50 rounded-lg">
            <p className="text-sm text-gray-500 mb-1">Búsqueda realizada:</p>
            <p className="text-sm font-medium text-gray-700">"{searchTerm}"</p>
          </div>
        )}
        
        {/* Suggestions */}
        <div className="mb-6 p-4 bg-blue-50 rounded-lg">
          <h3 className="text-sm font-medium text-blue-800 mb-2">
            Sugerencias:
          </h3>
          <ul className="text-sm text-blue-700 space-y-1">
            <li>• Verifica que el término de búsqueda esté escrito correctamente</li>
            <li>• Intenta con términos más generales</li>
            <li>• Revisa los filtros aplicados</li>
          </ul>
        </div>
        
        {/* Footer */}
        <div className="flex gap-3 justify-end">
          {onRetry && (
            <Button
              onClick={onRetry}
              variant="outline"
              className="flex items-center"
            >
              <RefreshCw className="h-4 w-4 mr-2" />
              Reintentar
            </Button>
          )}
          <Button
            onClick={onClose}
            className="bg-blue-600 hover:bg-blue-700 text-white"
          >
            Entendido
          </Button>
        </div>
      </div>
    </div>
  );
}
