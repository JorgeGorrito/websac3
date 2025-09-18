"use client";

import React from "react";
import AppSidebar from "@/components/sidebar/app-sidebar";
import { Breadcrumbs } from "@/components/sidebar/breadcrumbs";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import type { RoleKey } from "@/constants/sidebar";
import { WebSAC3Footer } from "@/components/websac3/footer/WebSAC3Footer";
import { WebSAC3Logo } from "@/components/websac3/logos/WebSAC3Logo";
import { ProtectedRoute } from "@/components/auth/ProtectedRoute";

export default function AdminLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const role: RoleKey = "ADMIN";

  const getPageTitle = (page: string) => {
    const titles: Record<string, string> = {
      dashboard: "Dashboard",
      usuarios: "Gestionar Usuarios",
      reportes: "Gestionar Reportes",
      asesoria: "Gestionar Solicitudes de Asesoría",
      accesos: "Gestionar Solicitudes de Acceso",
    };
    return titles[page] || page.charAt(0).toUpperCase() + page.slice(1);
  };

  return (
    <ProtectedRoute requiredRole="admin">
      <SidebarProvider>
        <AppSidebar role={role} />
        <SidebarInset>
          <header className="flex h-16 shrink-0 items-center justify-between">
            <div className="flex items-center gap-2 px-4">
              <SidebarTrigger className="-ml-1" />
              <Separator
                orientation="vertical"
                className="mr-2 data-[orientation=vertical]:h-4"
              />
              <Breadcrumbs role={role} getPageTitle={getPageTitle} />
            </div>
            <div className="px-4 h-10">
              <WebSAC3Logo />
            </div>
          </header>

          <div className="flex flex-1 flex-col h-[calc(100vh-4rem)]">
            <div className="flex-1 overflow-y-auto p-4">{children}</div>

            <WebSAC3Footer applyShadow={false} />
          </div>
        </SidebarInset>
      </SidebarProvider>
    </ProtectedRoute>
  );
}
