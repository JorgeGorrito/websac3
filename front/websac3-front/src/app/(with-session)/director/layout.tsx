"use client";

import React from "react";
import AppSidebar from "@/components/sidebar/app-sidebar";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import type { RoleKey } from "@/constants/sidebar";
import { WebSAC3Footer } from "@/components/websac3/footer/WebSAC3Footer";
import { Breadcrumbs } from "@/components/sidebar/breadcrumbs";
import { ProtectedRoute } from "@/components/auth/ProtectedRoute";

export default function DirectorLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const role: RoleKey = "DIRECTOR";

  const getPageTitle = (page: string) => {
    const titles: Record<string, string> = {
      dashboard: "Dashboard",
      "registrar-programa": "Registrar Programa",
      "cargar-datos": "Cargar Datos",
      "asesoria-experto": "Asesoría con Experto",
      reportes: "Consultar Reportes",
      "operar-modelo": "Operar Modelo",
      perfil: "Perfil",
    };
    return titles[page] || page.charAt(0).toUpperCase() + page.slice(1);
  };

  return (
    <ProtectedRoute requiredRole={["guest", "program lead"]}>
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
              <Breadcrumbs role="Director" getPageTitle={getPageTitle} />
            </div>
          </header>

          <div className="flex flex-1 flex-col h-[calc(100vh-4rem)] overflow-x-hidden">
            <div className="flex-1 overflow-y-auto overflow-x-hidden p-4">{children}</div>

            <WebSAC3Footer applyShadow={false} />
          </div>
        </SidebarInset>
      </SidebarProvider>
    </ProtectedRoute>
  );
}
