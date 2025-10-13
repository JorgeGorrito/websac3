"use client";

import { useEffect, useRef } from "react";
import { useAppDispatch } from "@/store/hooks";
import { loginSuccess, setInitialized } from "@/store/authSlice";
import { decodeJWT } from "@/lib/jwt";
import { getDashboardRoute } from "@/utils/roleUtils";

export const AuthInitializer = () => {
  const dispatch = useAppDispatch();
  const hasInitialized = useRef(false);

  useEffect(() => {
    if (hasInitialized.current) {
      return;
    }

    hasInitialized.current = true;
    
    if (typeof window !== "undefined") {
      const token = localStorage.getItem("access_token");
      const refreshToken = localStorage.getItem("refresh_token");
      
      if (token) {
        // Decode JWT to get user information
        const tokenPayload = decodeJWT(token);
        
        if (tokenPayload) {
          const user = {
            id: tokenPayload.sub,
            username: tokenPayload.username,
            email: tokenPayload.email || "",
            role: tokenPayload.role,
            permissions: tokenPayload.permissions,
          };

          // Dispatch loginSuccess to set both authentication and user data
          dispatch(loginSuccess({
            accessToken: token,
            refreshToken: refreshToken || "",
            user,
          }));
          
          // Handle redirection here to avoid multiple components competing
          const dashboardRoute = getDashboardRoute(user.role);
          const currentPath = window.location.pathname;
          
          // Only redirect if we're on a public route
          const publicRoutes = ['/login', '/solicitar-acceso', '/registrar-usuario', '/'];
          const isOnPublicRoute = publicRoutes.some(route => currentPath === route || currentPath.startsWith(route + '/'));
          
          if (isOnPublicRoute) {
            // Use window.location.href for a hard redirect to prevent any client-side routing issues
            window.location.href = dashboardRoute;
          }
        } else {
          localStorage.removeItem("access_token");
          localStorage.removeItem("refresh_token");
          dispatch(setInitialized()); // Mark as initialized even if no valid token
        }
      } else {
        dispatch(setInitialized()); // Mark as initialized even if no token
      }
    }
  }, [dispatch]);

  return null; // This component doesn't render anything
};
