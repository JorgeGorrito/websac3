import { WebSAC3Logo } from "../logos/WebSAC3Logo";

const WebSAC3BigHeader = () => {
    return (
        <header  className="flex w-full h-1/2 bg-secondary-light shadow-md shadow-slate-600">
            <div className="h-1/3 ml-12 mt-8 ">
                <WebSAC3Logo />
            </div>
        </header>
    );
};

export { WebSAC3BigHeader };