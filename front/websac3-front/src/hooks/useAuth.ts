import { useAppSelector, useAppDispatch } from "@/store/hooks";
import { logout } from "@/store/authSlice";
import { useRouter } from "next/navigation";

export const useAuth = () => {
  const dispatch = useAppDispatch();
  const router = useRouter();
  const { user, isAuthenticated, isLoading, error } = useAppSelector((state) => state.auth);

  const handleLogout = () => {
    dispatch(logout());
    router.push("/login");
  };

  const hasRole = (role: string) => {
    return user?.role === role;
  };

  const hasPermission = (permission: string) => {
    if (!user?.permissions) return false;
    
    // Check if user has the specific permission
    return Object.values(user.permissions).some(permissions => 
      permissions.includes(permission)
    );
  };

  const canAccess = (requiredRole?: string, requiredPermission?: string) => {
    if (!isAuthenticated) return false;
    
    if (requiredRole && !hasRole(requiredRole)) return false;
    
    if (requiredPermission && !hasPermission(requiredPermission)) return false;
    
    return true;
  };

  return {
    user,
    isAuthenticated,
    isLoading,
    error,
    logout: handleLogout,
    hasRole,
    hasPermission,
    canAccess,
  };
};
