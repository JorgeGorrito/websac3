"use client";

import { useEffect, useRef } from "react";
import { useRouter } from "next/navigation";
import { useAppSelector } from "@/store/hooks";

export const AuthRedirector = () => {
  const router = useRouter();
  const { isAuthenticated, user, isLoading } = useAppSelector((state) => state.auth);
  const hasRedirected = useRef(false);

  useEffect(() => {
    console.log("🔄 Debug - AuthRedirector useEffect:", { isAuthenticated, user, isLoading, hasRedirected: hasRedirected.current });
    
    // Only redirect if we're not loading, have complete auth data, and haven't redirected yet
    if (!isLoading && isAuthenticated && user && !hasRedirected.current) {
      const getDashboardRoute = (role: string): string => {
        switch (role) {
          case 'admin':
            return "/admin/dashboard";
          case 'guest':
          case 'program lead':
            return "/director/dashboard";
          case 'cybersecurity_auditor':
            return "/experto/dashboard";
          default:
            return "/admin/dashboard";
        }
      };
      
      const dashboardRoute = getDashboardRoute(user.role);
      const currentPath = window.location.pathname;
      
      console.log("🔄 Debug - AuthRedirector: Current path:", currentPath, "Dashboard route:", dashboardRoute);
      
      // Only redirect if we're on a public route (login, register, etc.)
      const publicRoutes = ['/login', '/solicitar-acceso', '/registrar-usuario'];
      const isOnPublicRoute = publicRoutes.some(route => currentPath.startsWith(route));
      
      if (isOnPublicRoute) {
        console.log("🔄 Debug - AuthRedirector: On public route, redirecting to dashboard");
        hasRedirected.current = true; // Mark as redirected to prevent multiple redirections
        window.location.href = dashboardRoute;
      } else {
        console.log("🔄 Debug - AuthRedirector: Not on public route, staying put");
      }
    }
  }, [isAuthenticated, user, isLoading, router]);

  return null; // This component doesn't render anything
};
