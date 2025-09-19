"use client";

import { useState, useRef, useEffect } from "react";
import { ArrowRight, Check } from "lucide-react";

interface SlideToSubmitProps {
  onComplete: () => void;
  label?: string;
  completedText?: string;
  slideText?: string;
  resetTrigger?: number; // Add a prop to trigger reset from parent
}

export const SlideToSubmit = ({
  onComplete,
  label = undefined, // Changed default to undefined
  completedText = "¡Completado!",
  slideText = "Arrastra para continuar",
  resetTrigger = 0,
}: SlideToSubmitProps) => {
  const [slideProgress, setSlideProgress] = useState(4);
  const [isDragging, setIsDragging] = useState(false);
  const [isCompleted, setIsCompleted] = useState(false);
  const [hasMoved, setHasMoved] = useState(false);

  const slideRef = useRef<HTMLDivElement>(null);
  const handleRef = useRef<HTMLDivElement>(null);
  const currentProgressRef = useRef(4);
  const dragStateRef = useRef({
    startX: 0,
    startProgress: 0,
    hasMoved: false
  });

  // Reset component when it mounts - but only if not already completed
  useEffect(() => {
    if (!isCompleted) {
      setSlideProgress(4);
      setHasMoved(false);
      currentProgressRef.current = 4;
      dragStateRef.current = { startX: 0, startProgress: 0, hasMoved: false };
    }
  }, []);

  // Reset when resetTrigger changes
  useEffect(() => {
    if (resetTrigger > 0) {
      setSlideProgress(4);
      setIsCompleted(false);
      setHasMoved(false);
      currentProgressRef.current = 4;
      dragStateRef.current = { startX: 0, startProgress: 0, hasMoved: false };
    }
  }, [resetTrigger]);

  const startDrag = (clientX: number) => {
    if (!slideRef.current || !handleRef.current || isCompleted) return;
    
    const rect = slideRef.current.getBoundingClientRect();
    const handleWidth = handleRef.current.offsetWidth;
    const maxDist = rect.width - handleWidth;
    
    // Initialize drag state
    dragStateRef.current = {
      startX: clientX,
      startProgress: slideProgress,
      hasMoved: false
    };
    
    setIsDragging(true);
    setHasMoved(false);

    const onMove = (moveEvt: MouseEvent | Touch) => {
      const currentX = 'clientX' in moveEvt ? moveEvt.clientX : moveEvt.clientX;
      const deltaX = Math.abs(currentX - dragStateRef.current.startX);
      
      // Require minimum movement to prevent click cheating
      if (deltaX > 15) {
        dragStateRef.current.hasMoved = true;
        setHasMoved(true);
      }

      let x = currentX - rect.left - handleWidth / 2;
      x = Math.max(0, Math.min(maxDist, x));
      const newProgress = (x / maxDist) * 100;
      
      setSlideProgress(newProgress);
      currentProgressRef.current = newProgress;
    };

    const onUp = () => {
      setIsDragging(false);
      document.removeEventListener("mousemove", onMove);
      document.removeEventListener("mouseup", onUp);
      document.removeEventListener("touchmove", onMove);
      document.removeEventListener("touchend", onUp);

      // Use the ref value which is always current
      const currentProgress = currentProgressRef.current;
      const hasActuallyMoved = dragStateRef.current.hasMoved;
      
      // Only complete if user has actually moved the slider significantly
      if (currentProgress >= 90 && hasActuallyMoved) {
        setIsCompleted(true);
        setSlideProgress(100); // Keep it at 100% when completed
        onComplete();
        // Don't auto-reset, let the parent component handle it
      } else {
        // Snap back to start if not completed
        setSlideProgress(4);
        setHasMoved(false);
        currentProgressRef.current = 4;
        dragStateRef.current = { startX: 0, startProgress: 0, hasMoved: false };
      }
    };

    document.addEventListener("mousemove", onMove);
    document.addEventListener("mouseup", onUp);
    document.addEventListener("touchmove", onMove);
    document.addEventListener("touchend", onUp);
  };

  const handleMouseDown = (e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
    startDrag(e.clientX);
  };

  const handleTouchStart = (e: React.TouchEvent) => {
    e.preventDefault();
    e.stopPropagation();
    if (e.touches.length === 1) {
      startDrag(e.touches[0].clientX);
    }
  };

  // Prevent direct clicks on the track
  const handleTrackClick = (e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
  };

  return (
    <div className="space-y-1 select-none">
      {label && <label className="text-sm font-medium text-gray-700">{label}</label>}
      <div
        ref={slideRef}
        onClick={handleTrackClick}
        className={`relative rounded-full h-12 w-full overflow-hidden border-2 transition-all duration-200 ${
          isCompleted 
            ? "bg-green-100 border-green-400" 
            : isDragging 
            ? "bg-blue-100 border-blue-400" 
            : "bg-gray-200 border-gray-300"
        }`}
      >
        {/* Progress bar */}
        <div
          className={`absolute top-0 left-0 h-full transition-all duration-200 ${
            isCompleted 
              ? "bg-green-400" 
              : "bg-blue-400"
          } ${isDragging ? "duration-0" : "ease-out"}`}
          style={{ width: `${slideProgress}%` }}
        />

        {/* Handle */}
        <div
          ref={handleRef}
          onMouseDown={handleMouseDown}
          onTouchStart={handleTouchStart}
          className={`absolute top-1 h-10 w-10 rounded-full shadow-lg flex items-center justify-center z-10 transition-all duration-200 ${
            isCompleted
              ? "bg-green-500 cursor-default"
              : isDragging
              ? "bg-white cursor-grabbing shadow-xl scale-105"
              : "bg-white cursor-grab hover:shadow-xl hover:scale-105"
          }`}
          style={{
            left: `calc(${slideProgress}% - 1.25rem)`,
            transform: isDragging ? "scale(1.05)" : "scale(1)",
          }}
        >
          {isCompleted ? (
            <Check className="w-5 h-5 text-white" />
          ) : (
            <ArrowRight className="w-4 h-4 text-blue-500" />
          )}
        </div>

        {/* Text overlay */}
        <div className="h-full flex items-center justify-center relative z-0 pointer-events-none">
          <span className={`text-xs font-medium transition-colors duration-200 ${
            isCompleted 
              ? "text-green-600" 
              : slideProgress > 50 
              ? "text-white" 
              : "text-gray-500"
          }`}>
            {isCompleted 
              ? completedText 
              : slideProgress >= 95 
              ? "¡Suelta para completar!" 
              : slideText
            }
          </span>
        </div>

        {/* Security indicator */}
        {!hasMoved && slideProgress > 10 && (
          <div className="absolute top-0 right-2 h-full flex items-center">
            <div className="w-2 h-2 bg-yellow-400 rounded-full animate-pulse" />
          </div>
        )}
      </div>
      
      {/* Instructions */}
      <div className="text-xs text-gray-500 text-center">
        {!hasMoved && slideProgress > 10 
          ? "Continúa deslizando para completar" 
          : "Desliza completamente para continuar"
        }
      </div>
    </div>
  );
};
