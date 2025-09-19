export type ErrorCode = 400 | 401 | 403 | 404 | 422 | 429 | 500 | 502 | 503 | 504;

export interface ErrorInfo {
  title: string;
  description: string;
  icon: string;
  color: string;
}

export const getErrorInfo = (statusCode: number): ErrorInfo => {
  const errorMap: Record<ErrorCode, ErrorInfo> = {
    400: {
      title: "Solicitud Incorrecta",
      description: "La información enviada no es válida o está incompleta",
      icon: "⚠️",
      color: "orange"
    },
    401: {
      title: "Sesión Expirada",
      description: "Tu sesión ha expirado, por favor inicia sesión nuevamente",
      icon: "🔐",
      color: "blue"
    },
    403: {
      title: "Acceso Denegado",
      description: "No cuentas con los privilegios requeridos para realizar esta acción",
      icon: "🚫",
      color: "red"
    },
    404: {
      title: "No Encontrado",
      description: "El recurso solicitado no existe o no está disponible",
      icon: "🔍",
      color: "blue"
    },
    422: {
      title: "Datos Inválidos",
      description: "Los datos proporcionados no cumplen con los requisitos",
      icon: "❌",
      color: "orange"
    },
    429: {
      title: "Demasiadas Solicitudes",
      description: "Has realizado demasiadas solicitudes, intenta nuevamente más tarde",
      icon: "⏰",
      color: "yellow"
    },
    500: {
      title: "Error del Servidor",
      description: "Ha ocurrido un problema interno, nuestro equipo está trabajando para solucionarlo",
      icon: "🔧",
      color: "red"
    },
    502: {
      title: "Servicio No Disponible",
      description: "El servicio está temporalmente fuera de línea",
      icon: "🌐",
      color: "red"
    },
    503: {
      title: "Servicio en Mantenimiento",
      description: "El sistema está en mantenimiento, intenta nuevamente más tarde",
      icon: "🔨",
      color: "yellow"
    },
    504: {
      title: "Tiempo de Espera Agotado",
      description: "La solicitud tardó demasiado en procesarse",
      icon: "⏱️",
      color: "orange"
    }
  };

  return errorMap[statusCode as ErrorCode] || {
    title: "Error Inesperado",
    description: "Ha ocurrido un problema inesperado",
    icon: "❓",
    color: "gray"
  };
};
