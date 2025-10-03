import { SupportContainerProps } from "@/types/websac3/footer/SupportContainer";
import { cn } from "@/lib/utils";
import React from "react";

const SupportContainer : React.FC<SupportContainerProps> = ({children, applyShadow = true}) =>  {
    return (
        <div className={cn("flex items-center justify-center h-16 w-full bg-white rounded-lg border border-gray-200", applyShadow && "shadow-lg")}>
            {children}
        </div>
    );
};

export { SupportContainer };