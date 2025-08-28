import { UnillanosLogo } from "@/components/websac3/logos/UnillanosLogo";
import { LogosContainer } from "./LogosContainer";
import { FooterLine } from "./FooterLine";
import { SupportContainer } from "./SupportContainer";
import { FCBILogo } from "../logos/FCBILogo";

const WebSAC3Footer = ({applyShadow = true}: {applyShadow?: boolean}) => {
  return (
    <footer className="flex w-full justify-center items-center p-4 mt-auto">
      <SupportContainer applyShadow={applyShadow}>
        <FooterLine />
        <LogosContainer>
          <UnillanosLogo />
          <FCBILogo />
        </LogosContainer>
        <FooterLine />
      </SupportContainer>
    </footer>
  );
};

export { WebSAC3Footer };
