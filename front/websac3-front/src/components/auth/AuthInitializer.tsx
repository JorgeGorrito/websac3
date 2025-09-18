"use client";

import { useEffect, useRef } from "react";
import { useAppDispatch } from "@/store/hooks";
import { loginSuccess, setInitialized } from "@/store/authSlice";
import { decodeJWT } from "@/lib/jwt";

export const AuthInitializer = () => {
  const dispatch = useAppDispatch();
  const hasInitialized = useRef(false);

  useEffect(() => {
    if (hasInitialized.current) {
      console.log("🔧 Debug - AuthInitializer: Already initialized, skipping");
      return;
    }

    console.log("🔧 Debug - AuthInitializer: Initializing auth state");
    hasInitialized.current = true;
    
    if (typeof window !== "undefined") {
      const token = localStorage.getItem("access_token");
      const refreshToken = localStorage.getItem("refresh_token");
      
      if (token) {
        console.log("🔧 Debug - AuthInitializer: Found token in localStorage");
        
        // Decode JWT to get user information
        const tokenPayload = decodeJWT(token);
        
        if (tokenPayload) {
          console.log("🔧 Debug - AuthInitializer: Decoded JWT payload:", tokenPayload);
          
          const user = {
            id: tokenPayload.sub,
            username: tokenPayload.username,
            email: tokenPayload.email || "",
            role: tokenPayload.role,
            permissions: tokenPayload.permissions,
          };

          console.log("🔧 Debug - AuthInitializer: User object created:", user);

          // Dispatch loginSuccess to set both authentication and user data
          dispatch(loginSuccess({
            accessToken: token,
            refreshToken: refreshToken || "",
            user,
          }));
          
          console.log("🔧 Debug - AuthInitializer: Auth state initialized successfully");
          
          // Handle redirection here to avoid multiple components competing
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
          
          console.log("🔧 Debug - AuthInitializer: Current path:", currentPath, "Dashboard route:", dashboardRoute);
          
          // Only redirect if we're on a public route
          const publicRoutes = ['/login', '/access-request', '/register-user', '/'];
          const isOnPublicRoute = publicRoutes.some(route => currentPath === route || currentPath.startsWith(route + '/'));
          
          if (isOnPublicRoute) {
            console.log("🔧 Debug - AuthInitializer: On public route, redirecting to dashboard");
            // Use window.location.href for a hard redirect to prevent any client-side routing issues
            window.location.href = dashboardRoute;
          } else {
            console.log("🔧 Debug - AuthInitializer: Not on public route, staying put");
          }
        } else {
          console.log("🔧 Debug - AuthInitializer: Failed to decode JWT, clearing tokens");
          localStorage.removeItem("access_token");
          localStorage.removeItem("refresh_token");
          dispatch(setInitialized()); // Mark as initialized even if no valid token
        }
      } else {
        console.log("🔧 Debug - AuthInitializer: No token found in localStorage");
        dispatch(setInitialized()); // Mark as initialized even if no token
      }
    }
  }, [dispatch]);

  return null; // This component doesn't render anything
};
