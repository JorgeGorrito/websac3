"use client";

import React, { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { AlertTriangle, Calendar, GraduationCap, Building, User, Star } from "lucide-react";
import { useListPendingReportsQuery } from "@/services/api";

interface PendingReport {
  id: number;
  created_at: string;
  score: number;
  has_feedback: boolean;
  degree_program: {
    id: number;
    name: string;
    snies: number;
  };
  higher_education_institution: {
    id: number;
    name: string;
  };
  professional_role: {
    id: number;
    name: string;
  };
}

export default function PendingReportsPage() {
  const [currentPage, setCurrentPage] = useState(1);
  const [limit] = useState(10);

  const {
    data: reportsData,
    isLoading: loading,
    error: queryError,
    refetch
  } = useListPendingReportsQuery({ page: currentPage, limit });

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString("es-ES", {
      year: "numeric",
      month: "long",
      day: "numeric",
    });
  };

  const getScoreColor = (score: number) => {
    if (score >= 80) return "bg-green-100 text-green-800";
    if (score >= 60) return "bg-yellow-100 text-yellow-800";
    return "bg-red-100 text-red-800";
  };

  // Extract data from the query result
  const reports = reportsData?.data || [];
  const totalCount = reportsData?.total_count || 0;
  const currentPageFromAPI = reportsData?.current_page || 1;
  const itemsPerPage = reportsData?.items_per_page || 10;
  
  // Calculate total pages based on total count and items per page
  const totalPages = Math.ceil(totalCount / itemsPerPage);
  
  // Check if it's a 404 error (no reports found) vs a real error
  const is404Error = queryError && 'status' in queryError && queryError.status === 404;
  const error = queryError && !is404Error ? "Error al cargar los reportes pendientes" : null;

  const Pagination = () => {
    const pages = [];
    const maxVisiblePages = 5;
    let startPage = Math.max(1, currentPage - Math.floor(maxVisiblePages / 2));
    let endPage = Math.min(totalPages, startPage + maxVisiblePages - 1);

    if (endPage - startPage + 1 < maxVisiblePages) {
      startPage = Math.max(1, endPage - maxVisiblePages + 1);
    }

    for (let i = startPage; i <= endPage; i++) {
      pages.push(
        <Button
          key={i}
          variant={i === currentPage ? "default" : "outline"}
          size="sm"
          onClick={() => handlePageChange(i)}
          className="mx-1"
        >
          {i}
        </Button>
      );
    }

    return (
      <div className="flex items-center justify-center space-x-2 mt-6">
        <Button
          variant="outline"
          size="sm"
          onClick={() => handlePageChange(currentPage - 1)}
          disabled={currentPage === 1}
        >
          Anterior
        </Button>
        {pages}
        <Button
          variant="outline"
          size="sm"
          onClick={() => handlePageChange(currentPage + 1)}
          disabled={currentPage === totalPages}
        >
          Siguiente
        </Button>
        <div className="ml-4 text-sm text-gray-600">
          Página {currentPage} de {totalPages} ({totalCount} reportes)
        </div>
      </div>
    );
  };

  const ReportCard = ({ report }: { report: PendingReport }) => (
    <Card className="hover:shadow-md transition-shadow">
      <CardHeader className="pb-3">
        <div className="flex items-start justify-between">
          <div className="flex items-center space-x-2">
            <AlertTriangle className="h-5 w-5 text-orange-500" />
            <CardTitle className="text-lg">Reporte #{report.id}</CardTitle>
          </div>
          <Badge className={getScoreColor(report.score * 100)}>
            <Star className="h-3 w-3 mr-1" />
            {(() => {
              const percentage = Math.floor(report.score * 10000) / 100;
              return percentage % 1 === 0 ? percentage.toFixed(0) : percentage.toFixed(2);
            })()}%
          </Badge>
        </div>
      </CardHeader>
      <CardContent className="space-y-3">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="flex items-center space-x-2">
            <GraduationCap className="h-4 w-4 text-gray-500" />
            <div>
              <p className="text-sm font-medium">{report.degree_program.name}</p>
              <p className="text-xs text-gray-500">SNIES: {report.degree_program.snies}</p>
            </div>
          </div>
          
          <div className="flex items-center space-x-2">
            <Building className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{report.higher_education_institution.name}</p>
          </div>
          
          <div className="flex items-center space-x-2">
            <User className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{report.professional_role.name}</p>
          </div>
          
          <div className="flex items-center space-x-2">
            <Calendar className="h-4 w-4 text-gray-500" />
            <p className="text-sm">{formatDate(report.created_at)}</p>
          </div>
        </div>
        
        <div className="pt-3 border-t">
          <Button className="w-full" variant="outline">
            Revisar Reporte
          </Button>
        </div>
      </CardContent>
    </Card>
  );

  const LoadingSkeleton = () => (
    <div className="space-y-4">
      {Array.from({ length: 3 }).map((_, index) => (
        <Card key={index}>
          <CardHeader>
            <div className="flex items-center justify-between">
              <Skeleton className="h-6 w-32" />
              <Skeleton className="h-6 w-16" />
            </div>
          </CardHeader>
          <CardContent className="space-y-3">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <Skeleton className="h-4 w-full" />
              <Skeleton className="h-4 w-full" />
              <Skeleton className="h-4 w-full" />
              <Skeleton className="h-4 w-full" />
            </div>
            <Skeleton className="h-10 w-full" />
          </CardContent>
        </Card>
      ))}
    </div>
  );

  if (error) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6">
          <div className="text-center">
            <AlertTriangle className="h-12 w-12 text-red-500 mx-auto mb-4" />
            <h2 className="text-xl font-semibold text-red-600 mb-2">Error al cargar reportes</h2>
            <p className="text-gray-600 mb-4">{error}</p>
            <Button onClick={() => refetch()}>
              Reintentar
            </Button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6">
        <div className="mb-6">
          <h2 className="text-2xl font-semibold text-center">
            Reportes Pendientes
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
          <p className="text-center text-gray-600 mt-4">
            {totalCount > 0 ? `${totalCount} reportes pendientes de revisión` : "Estado actualizado - Sin reportes pendientes"}
          </p>
        </div>

        {loading ? (
          <LoadingSkeleton />
        ) : (reports.length === 0 || is404Error) ? (
          <div className="text-center py-12">
            <div className="bg-green-50 rounded-full w-20 h-20 flex items-center justify-center mx-auto mb-6">
              <AlertTriangle className="h-10 w-10 text-green-500" />
            </div>
            <h3 className="text-xl font-semibold text-gray-800 mb-3">
              ¡Excelente trabajo!
            </h3>
            <p className="text-gray-600 mb-2">
              No hay reportes pendientes de revisión en este momento.
            </p>
            <p className="text-sm text-gray-500">
              Todos los reportes han sido procesados o no hay nuevos reportes disponibles.
            </p>
          </div>
        ) : (
          <>
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              {reports.map((report) => (
                <ReportCard key={report.id} report={report} />
              ))}
            </div>
            
            
            {totalCount > 0 && <Pagination />}
          </>
        )}
      </div>
    </div>
  );
}
