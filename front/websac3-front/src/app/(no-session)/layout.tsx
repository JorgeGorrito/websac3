"use client";

import { WebSAC3Footer } from "@/components/websac3/footer/WebSAC3Footer";
import { WebSAC3BigHeader } from "@/components/websac3/header/WebSAC3BigHeader";
import "@/styles/layouts/DefaultLayout.css";
import { ReactNode } from "react";
import { usePathname } from "next/navigation";

export default function NoSessionLayout({
  children,
}: {
  children: ReactNode;
}) {
  const pathname = usePathname();
  const isLandingPage = pathname === "/";

  if (isLandingPage) {
    return <>{children}</>;
  }

  return (
    <main className="main-container">
      <WebSAC3BigHeader />
      {children}
      <WebSAC3Footer />
    </main>
  );
}