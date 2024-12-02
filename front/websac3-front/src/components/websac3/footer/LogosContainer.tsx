import { LogosContainerProps } from "@/types/websac3/footer/LogosContainer";
import React from "react";

const LogosContainer: React.FC<LogosContainerProps> = ({ children }) => {
    return (
        <div className="flex space-x-4 pr-2 pl-2">
            {React.Children.map(children, (child) => child)}
        </div>
    );
};

export { LogosContainer };
