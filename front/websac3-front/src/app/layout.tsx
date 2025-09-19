import "@/app/globals.css";
import "@/styles/layouts/DefaultLayout.css";
import type { Metadata } from "next";
import { ReactNode } from "react";
import { ReduxProvider } from "@/store/ReduxProvider";
import { AuthInitializer } from "@/components/auth/AuthInitializer";
import { ErrorHandler } from "@/components/ErrorHandler";

export const metadata: Metadata = {
  title: "WebSAC3",
  description: "Web System for analysis of curricula cybersecurity component",
  viewport: "width=device-width, initial-scale=1.0",
};

export default function RootLayout({
  children,
}: {
  children: ReactNode;
}) {
  return (
    <html lang="es">
      <body>
        <ReduxProvider>
          <AuthInitializer />
          <ErrorHandler />
          {children}
        </ReduxProvider>
      </body>
    </html>
  );
} 