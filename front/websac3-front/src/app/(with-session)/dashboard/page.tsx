"use client";
import { AppSidebar } from "@/components/sidebar/app-sidebar";
import type { RoleKey } from "@/constants/sidebar";
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
import { WebSAC3Footer } from "@/components/websac3/footer/WebSAC3Footer";
import { WebSAC3Logo } from "@/components/websac3/logos/WebSAC3Logo";
import { usePathname } from "next/navigation";

function toTitle(text: string) {
  return text
    .split(" ")
    .map((w) => (w ? w[0].toUpperCase() + w.slice(1).toLowerCase() : w))
    .join(" ");
}

export default function Page() {
  const role: RoleKey = "EXPERTO";
  const pathname = usePathname();
  const segments = pathname.split("/").filter(Boolean);
  const last = segments.at(-1) ?? "dashboard";
  const isSub = last.startsWith("dashboard-");
  const subLabel = isSub
    ? toTitle(last.replace("dashboard-", "").replaceAll("-", " "))
    : undefined;
  return (
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
                <BreadcrumbItem>
                  <BreadcrumbLink href="/dashboard">Dashboard</BreadcrumbLink>
                </BreadcrumbItem>
                {isSub && (
                  <>
                    <BreadcrumbSeparator className="hidden md:block" />
                    <BreadcrumbItem>
                      <BreadcrumbPage>{subLabel}</BreadcrumbPage>
                    </BreadcrumbItem>
                  </>
                )}
              </BreadcrumbList>
            </Breadcrumb>
          </div>
          <div className="px-4 h-10">
            <WebSAC3Logo />
          </div>
        </header>
        <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
          <div className="grid auto-rows-min gap-4 md:grid-cols-3">
            <div className="bg-muted/50 aspect-video rounded-xl" />
            <div className="bg-muted/50 aspect-video rounded-xl" />
            <div className="bg-muted/50 aspect-video rounded-xl" />
          </div>
          <div className="bg-muted/50 min-h-[100vh] flex-1 rounded-xl md:min-h-min" />
          <WebSAC3Footer applyShadow={false} />
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}
