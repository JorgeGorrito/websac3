import { UnillanosLogo } from "@/components/websac3/logos/UnillanosLogo";
import { LogosContainer } from "./LogosContainer";
import { FooterLine } from "./FooterLine";
import { SupportContainer } from "./SupportContainer";
import { FCBILogo } from "../logos/FCBILogo";

const WebSAC3Footer = () => {
    return (
        <footer className="flex w-full h-1/2 items-end  p-2 mb-2">
            <SupportContainer>
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