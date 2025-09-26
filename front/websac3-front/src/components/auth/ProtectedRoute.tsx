"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAppSelector } from "@/store/hooks";

interface ProtectedRouteProps {
  children: React.ReactNode;
  requiredRole?: string | string[];
  fallbackPath?: string;
}

export const ProtectedRoute = ({ 
  children, 
  requiredRole, 
  fallbackPath = "/login" 
}: ProtectedRouteProps) => {
  const { isAuthenticated, user, isLoading, isInitialized } = useAppSelector((state) => state.auth);
  const router = useRouter();

  useEffect(() => {
    // Don't do anything until auth is initialized
    if (!isInitialized) {
      return;
    }
    
    if (!isLoading) {
      if (!isAuthenticated) {
        router.push(fallbackPath);
        return;
      }

      if (requiredRole) {
        const allowedRoles = Array.isArray(requiredRole) ? requiredRole : [requiredRole];
        const userRole = user?.role?.toLowerCase().trim() || "";
        const hasRequiredRole = allowedRoles.some(role => 
          role.toLowerCase().trim() === userRole
        );
        
        if (!hasRequiredRole) {
        // Redirect to appropriate dashboard based on user role
        const getDashboardRoute = (role: string): string => {
          const normalizedRole = role?.toLowerCase().trim();
          switch (normalizedRole) {
            case 'admin':
              return "/admin/dashboard";
            case 'guest':
            case 'program lead':
              return "/director/dashboard";
            case 'cybersecurity_auditor':
            case 'cybersecurity auditor':
              return "/experto/dashboard";
            default:
              return "/admin/dashboard";
          }
        };
        
          const userDashboard = getDashboardRoute(user?.role || "");
          router.push(userDashboard);
          return;
        }
      }
    }
  }, [isAuthenticated, user, isLoading, isInitialized, requiredRole, fallbackPath, router]);

  // Show loading while checking authentication or initializing
  if (isLoading || !isInitialized) {
    return (
      <div className="flex items-center justify-center h-screen">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
          <p className="mt-4 text-gray-600">
            {!isInitialized ? "Inicializando..." : "Verificando autenticación..."}
          </p>
        </div>
      </div>
    );
  }

  // If authenticated but no user data, redirect to login to get user info
  if (isAuthenticated && !user) {
    router.push("/login");
    return null;
  }

  // Don't render children if not authenticated
  if (!isAuthenticated) {
    return null;
  }

  // Don't render children if role doesn't match
  if (requiredRole) {
    const allowedRoles = Array.isArray(requiredRole) ? requiredRole : [requiredRole];
    const userRole = user?.role?.toLowerCase().trim() || "";
    const hasRequiredRole = allowedRoles.some(role => 
      role.toLowerCase().trim() === userRole
    );
    
    if (!hasRequiredRole) {
      return null;
    }
  }

  return <>{children}</>;
};
