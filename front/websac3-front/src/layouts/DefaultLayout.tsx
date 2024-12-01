import WebSAC3Header from "@/components/websac3-header/websac3Header";
import React from "react";

const DefaultLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <>
      <WebSAC3Header />
      {children}
    </>
  );
};

export { DefaultLayout };
