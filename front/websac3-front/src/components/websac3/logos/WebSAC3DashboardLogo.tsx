"use client";

import { useRouter } from "next/navigation";
import { WebSAC3Logo } from "./WebSAC3Logo";
import { useAuth } from "@/hooks/useAuth";

export function WebSAC3DashboardLogo() {
  const router = useRouter();
  const { user } = useAuth();

  const handleLogoClick = () => {
    const role = user?.role?.toLowerCase().trim();
    
    if (role === "admin") {
      router.push("/admin/dashboard");
    } else if (role === "guest" || role === "program lead") {
      router.push("/director/dashboard");
    } else if (role === "cybersecurity_auditor" || role === "cybersecurity auditor") {
      router.push("/experto/dashboard");
    } else {
      router.push("/login");
    }
  };

  return (
    <div
      className="h-12 w-full cursor-pointer transition-transform hover:scale-105"
      onClick={handleLogoClick}
    >
      <WebSAC3Logo />
    </div>
  );
}
