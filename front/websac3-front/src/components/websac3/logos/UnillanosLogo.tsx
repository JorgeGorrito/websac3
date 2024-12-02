import Image from "next/image";
import LogoUnillanos from "@/../public/unillanos/logo-unillanos.png"

const UnillanosLogo : React.FC = () => {
    return (
        <Image
            src={LogoUnillanos}
            alt="Logo de la Universidad de los Llanos"
            height={0}
            width={150}
            draggable={false}
        />
    );
};

export { UnillanosLogo };