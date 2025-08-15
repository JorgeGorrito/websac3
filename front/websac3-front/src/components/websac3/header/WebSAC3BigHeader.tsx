import { WebSAC3Logo } from "../logos/WebSAC3Logo";

const WebSAC3BigHeader = () => {
  return (
    <header className="flex w-full bg-white/10 backdrop-blur-sm shadow-lg">
      <div className="h-24 ml-8 mt-6 mb-6">
        <WebSAC3Logo />
      </div>
    </header>
  );
};

export { WebSAC3BigHeader };
