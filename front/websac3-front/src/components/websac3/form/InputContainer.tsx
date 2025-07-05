import { InputContainerProps } from "@/types/websac3/form/InputContainer";
import React from "react";

const InputContainer : React.FC<InputContainerProps> = ({children}) => {
    return (
        <div className="w-full md:w-1/2 mt-1 mb-1 pl-3 pr-3">
            {children}
        </div>
    );
};

export { InputContainer };