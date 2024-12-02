import { HeaderContainerProps } from "@/types/websac3/header/HeaderContainer";
import "@/styles/websac3/header/HeaderContainer.css";

const HeaderContainer: React.FC<HeaderContainerProps> = ({ children }) => {
    return (
        <header className="HeaderContainer">
            {children}
        </header>
    );
};

export { HeaderContainer };
