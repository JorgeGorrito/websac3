import { LogosContainerProps } from "@/types/websac3/footer/LogosContainer";
import React from "react";

const LogosContainer: React.FC<LogosContainerProps> = ({ children }) => {
    return (
        <div className="flex items-center space-x-8 px-6">
            {React.Children.map(children, (child) => child)}
        </div>
    );
};

export { LogosContainer };
