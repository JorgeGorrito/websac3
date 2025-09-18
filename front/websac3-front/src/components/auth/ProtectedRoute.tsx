"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAppSelector } from "@/store/hooks";

interface ProtectedRouteProps {
  children: React.ReactNode;
  requiredRole?: string;
  fallbackPath?: string;
}

export const ProtectedRoute = ({ 
  children, 
  requiredRole, 
  fallbackPath = "/login" 
}: ProtectedRouteProps) => {
  const { isAuthenticated, user, isLoading } = useAppSelector((state) => state.auth);
  const router = useRouter();

  useEffect(() => {
    if (!isLoading) {
      if (!isAuthenticated) {
        router.push(fallbackPath);
        return;
      }

      if (requiredRole && user?.role !== requiredRole) {
        // Redirect to appropriate dashboard based on user role
        const roleRoutes: Record<string, string> = {
          admin: "/admin/dashboard",
          director: "/director/dashboard",
          experto: "/experto/dashboard",
        };
        
        const userDashboard = roleRoutes[user?.role || ""] || "/admin/dashboard";
        router.push(userDashboard);
        return;
      }
    }
  }, [isAuthenticated, user, isLoading, requiredRole, fallbackPath, router]);

  // Show loading while checking authentication
  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-screen">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500 mx-auto"></div>
          <p className="mt-4 text-gray-600">Verificando autenticación...</p>
        </div>
      </div>
    );
  }

  // Don't render children if not authenticated
  if (!isAuthenticated) {
    return null;
  }

  // Don't render children if role doesn't match
  if (requiredRole && user?.role !== requiredRole) {
    return null;
  }

  return <>{children}</>;
};
