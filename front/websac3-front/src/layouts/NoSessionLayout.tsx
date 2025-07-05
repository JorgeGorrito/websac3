import { WebSAC3Footer } from "@/components/websac3/footer/WebSAC3Footer";
import { WebSAC3BigHeader } from "@/components/websac3/header/WebSAC3BigHeader";
import "@/styles/layouts/DefaultLayout.css";
import Head from "next/head";
import React from "react";

const NoSessionLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
      <>
        <Head>
            <title>WebSAC3</title>
            <meta name="description" content="Web System for analysis of curricula cybersecurity component" />
            <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        </Head>
        <main className="main-container">
          <WebSAC3BigHeader/>
            {children}
          <WebSAC3Footer/>
        </main>
      </>
  );
};

export { NoSessionLayout };
