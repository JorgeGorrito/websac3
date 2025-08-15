"use client";

import { useState, useRef } from "react";
import { ArrowRight } from "lucide-react";

interface SlideToSubmitProps {
  onComplete: () => void;
  label?: string;
  completedText?: string;
  slideText?: string;
}

export const SlideToSubmit = ({
  onComplete,
  label = "Slide to Submit",
  completedText = "¡Completado!",
  slideText = "Arrastra para continuar",
}: SlideToSubmitProps) => {
  const [slideProgress, setSlideProgress] = useState(4);
  const [isDragging, setIsDragging] = useState(false);

  const slideRef = useRef<HTMLDivElement>(null);
  const handleRef = useRef<HTMLDivElement>(null);

  const startDrag = () => {
    if (!slideRef.current || !handleRef.current) return;
    setIsDragging(true);

    const rect = slideRef.current.getBoundingClientRect();
    const handleWidth = handleRef.current.offsetWidth;
    const maxDist = rect.width - handleWidth;

    const onMove = (moveEvt: MouseEvent) => {
      let x = moveEvt.clientX - rect.left - handleWidth / 2;
      x = Math.max(0, Math.min(maxDist, x));
      setSlideProgress((x / maxDist) * 100);
    };

    const onUp = () => {
      setIsDragging(false);
      document.removeEventListener("mousemove", onMove);
      document.removeEventListener("mouseup", onUp);

      if (slideProgress >= 95) {
        onComplete();
        setSlideProgress(0);
      }
    };

    document.addEventListener("mousemove", onMove);
    document.addEventListener("mouseup", onUp);
  };

  const handleMouseDown = (e: React.MouseEvent) => {
    e.preventDefault();
    startDrag();
  };

  const handleTouchStart = (e: React.TouchEvent) => {
    e.preventDefault();
    startDrag();
  };

  return (
    <div className="space-y-1 select-none">
      <label className="text-sm font-medium text-gray-700">{label}</label>
      <div
        ref={slideRef}
        className="relative bg-gray-200 rounded-full h-10 w-full overflow-hidden"
      >
        <div
          className={`absolute top-0 left-0 h-full bg-blue-400 ${
            isDragging ? "duration-0" : "transition-all duration-300 ease-out"
          }`}
          style={{ width: `${slideProgress}%` }}
        />

        <div
          ref={handleRef}
          onMouseDown={handleMouseDown}
          onTouchStart={handleTouchStart}
          className={`absolute top-1 h-8 w-8 bg-white rounded-full shadow-md flex items-center justify-center cursor-grab active:cursor-grabbing z-10 ${
            isDragging ? "" : "hover:shadow-lg transition-shadow"
          }`}
          style={{
            left: `calc(${slideProgress}% - 1rem)`,
            transform: "none",
          }}
        >
          <ArrowRight className="w-4 h-4 text-blue-500" />
        </div>

        <div className="h-full flex items-center justify-center relative z-0 pointer-events-none">
          <span className="text-gray-500 text-xs">
            {slideProgress >= 95 ? completedText : slideText}
          </span>
        </div>
      </div>
    </div>
  );
};
