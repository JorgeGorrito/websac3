"use client";

import React from "react";
import AppSidebar from "@/components/sidebar/app-sidebar";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Separator } from "@/components/ui/separator";
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import type { RoleKey } from "@/constants/sidebar";
import { WebSAC3Footer } from "@/components/websac3/footer/WebSAC3Footer";
import { usePathname } from "next/navigation";
import { ProtectedRoute } from "@/components/auth/ProtectedRoute";

export default function ExpertoLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const role: RoleKey = "EXPERTO";
  const pathname = usePathname();

  const generateBreadcrumbs = () => {
    const segments = pathname.split("/").filter(Boolean);
    const breadcrumbs = [];

    if (segments.length === 1 && segments[0] === "experto") {
      breadcrumbs.push({
        title: "Experto",
        href: "/experto/dashboard",
        isCurrent: true,
      });
      return breadcrumbs;
    }

    if (segments.length > 1) {
      breadcrumbs.push({
        title: "Experto",
        href: "/experto/dashboard",
        isCurrent: false,
      });
    }

    if (segments.length > 1) {
      const currentPage = segments[1];
      const pageTitle = getPageTitle(currentPage);
      breadcrumbs.push({
        title: pageTitle,
        href: pathname,
        isCurrent: true,
      });
    }

    return breadcrumbs;
  };

  const getPageTitle = (page: string) => {
    const titles: Record<string, string> = {
      dashboard: "Dashboard",
      reportes: "Reportes",
      retroalimentados: "Reportes Retroalimentados",
      pendientes: "Reportes Pendientes",
      asesoria: "Asesoría",
      solicitudes: "Solicitudes de Asesoría",
      modelo: "Gestionar Modelo",
      perfil: "Perfil",
    };
    return titles[page] || page.charAt(0).toUpperCase() + page.slice(1);
  };

  const breadcrumbs = generateBreadcrumbs();

  return (
    <ProtectedRoute requiredRole={["cybersecurity_auditor", "cybersecurity auditor"]}>
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
              <Breadcrumb>
                <BreadcrumbList>
                  {breadcrumbs.map((breadcrumb, index) => (
                    <React.Fragment key={`${breadcrumb.title}-${index}`}>
                      <BreadcrumbItem>
                        {breadcrumb.isCurrent ? (
                          <BreadcrumbPage>{breadcrumb.title}</BreadcrumbPage>
                        ) : (
                          <BreadcrumbLink href={breadcrumb.href}>
                            {breadcrumb.title}
                          </BreadcrumbLink>
                        )}
                      </BreadcrumbItem>
                      {index < breadcrumbs.length - 1 && (
                        <BreadcrumbSeparator className="hidden md:block" />
                      )}
                    </React.Fragment>
                  ))}
                </BreadcrumbList>
              </Breadcrumb>
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
