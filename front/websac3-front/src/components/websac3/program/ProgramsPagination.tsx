import { Button } from "@/components/ui/button";
import { ChevronLeft, ChevronRight, MoreHorizontal } from "lucide-react";

type ProgramsPaginationProps = {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
};

export function ProgramsPagination({
  currentPage,
  totalPages,
  onPageChange,
}: ProgramsPaginationProps) {
  // Always show pagination, even with 1 page
  // if (totalPages <= 1) return null;

  const getVisiblePages = () => {
    // For single page, just return [1]
    if (totalPages === 1) {
      return [1];
    }

    const delta = 2;
    const pages = [];
    
    // Always add page 1
    pages.push(1);
    
    // Add middle pages if there are more than 2 pages
    if (totalPages > 2) {
      const start = Math.max(2, currentPage - delta);
      const end = Math.min(totalPages - 1, currentPage + delta);
      
      // Add dots before middle range if needed
      if (start > 2) {
        pages.push("...");
      }
      
      // Add middle pages
      for (let i = start; i <= end; i++) {
        if (i !== 1 && i !== totalPages) { // Avoid duplicates with first and last page
          pages.push(i);
        }
      }
      
      // Add dots after middle range if needed
      if (end < totalPages - 1) {
        pages.push("...");
      }
    }
    
    // Always add last page (if more than 1 page)
    if (totalPages > 1) {
      pages.push(totalPages);
    }
    
    return pages;
  };

  const visiblePages = getVisiblePages();

  return (
    <div className="flex items-center justify-center space-x-2 mt-8">
      {/* Previous Button */}
      <Button
        variant="outline"
        size="sm"
        onClick={() => onPageChange(currentPage - 1)}
        disabled={currentPage === 1}
        className="flex items-center"
      >
        <ChevronLeft className="h-4 w-4 mr-1" />
        Anterior
      </Button>

      {/* Page Numbers */}
      <div className="flex items-center space-x-1">
        {visiblePages.map((page, index) => {
          if (page === "...") {
            return (
              <div key={`dots-${index}`} className="px-2 py-1">
                <MoreHorizontal className="h-4 w-4 text-gray-400" />
              </div>
            );
          }

          const pageNumber = page as number;
          const isCurrentPage = pageNumber === currentPage;

          return (
            <Button
              key={pageNumber}
              variant={isCurrentPage ? "default" : "outline"}
              size="sm"
              onClick={() => onPageChange(pageNumber)}
              className={`min-w-[40px] ${
                isCurrentPage
                  ? "bg-blue-600 hover:bg-blue-700 text-white"
                  : "hover:bg-gray-50"
              }`}
            >
              {pageNumber}
            </Button>
          );
        })}
      </div>

      {/* Next Button */}
      <Button
        variant="outline"
        size="sm"
        onClick={() => onPageChange(currentPage + 1)}
        disabled={currentPage === totalPages}
        className="flex items-center"
      >
        Siguiente
        <ChevronRight className="h-4 w-4 ml-1" />
      </Button>
    </div>
  );
}
