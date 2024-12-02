import { UnillanosLogo } from "@/components/websac3/logos/UnillanosLogo";
import { LogosContainer } from "./LogosContainer";
import { FooterLine } from "./FooterLine";
import { SupportContainer } from "./SupportContainer";
import { FCBILogo } from "../logos/FCBILogo";

const WebSAC3Footer = () => {
    return (
        <footer className="flex w-full h-20 bg-secondary-light p-2">
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