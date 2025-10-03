"use client";

import { HeroSection } from "./HeroSection";
import { FeaturesSection } from "./FeaturesSection";

const LandingPage = () => {
  return (
    <div className="min-h-screen">
      <HeroSection />
      <FeaturesSection />
    </div>
  );
};

export { LandingPage };
