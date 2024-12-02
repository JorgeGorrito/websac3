import Image from "next/image";
import LogoUnillanos from "@/../public/unillanos/logo-fcbi.png"

const FCBILogo : React.FC = () => {
    return (
        <Image
            src={LogoUnillanos}
            alt="Logo de la Facultad de Ciencias Basicas e Ingenieria de la Universidad de los Llanos"
            height={0}
            width={150}
            draggable={false}
        />
    );
};

export { FCBILogo };