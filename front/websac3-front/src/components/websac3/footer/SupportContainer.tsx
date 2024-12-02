import { SupportContainerProps } from "@/types/websac3/footer/SupportContainer";
import React from "react";

const SupportContainer : React.FC<SupportContainerProps> = ({children}) =>  {
    return (
        <div className="flex items-center h-10 w-full">
            {children}
        </div>
    );
};

export { SupportContainer };