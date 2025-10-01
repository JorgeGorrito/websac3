"use client";

import { DirectorStats } from "@/components/websac3/dashboard/DirectorStats";

export default function DirectorDashboardPage() {
  return (
    <div className="space-y-6">
      {/* Statistics Section */}
      <DirectorStats />

      {/* Welcome Section */}
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="mb-6">
          <h2 className="text-2xl font-semibold text-center">
            Dashboard Director
          </h2>
          <div className="mx-auto mt-2 h-0.5 w-24 bg-gray-300 rounded" />
        </div>
        
        <div className="mt-8 text-center">
          <p className="text-gray-600">
            Bienvenido al dashboard de Director de Programa. Utiliza el menú lateral para acceder a todas las funcionalidades disponibles.
          </p>
        </div>
      </div>
    </div>
  );
}
