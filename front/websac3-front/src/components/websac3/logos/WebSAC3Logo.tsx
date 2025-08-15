import LogoWebSAC3 from "@/../public/websac3/logo-websac3.png";
import Image from "next/image";
import "@/styles/websac3/logos/WebSAC3Logo.css";

const WebSAC3Logo = () => {
    return (
        <div className="WebSAC3Logo">
            <Image
            src={LogoWebSAC3}
            alt="Logo WebSAC3"
            width={400} 
            height={129}
            className="h-full w-auto object-contain"
            quality={100}
            draggable={false}
            />
      </div>
    );
}

export { WebSAC3Logo };