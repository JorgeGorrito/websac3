"use client";

import { useParams } from "next/navigation";
import { EditProgramForm } from "@/components/websac3/form/program/EditProgramForm";

export default function EditarProgramaPage() {
  const params = useParams();
  const programId = parseInt(params.program_id as string);

  // Validate program ID
  if (isNaN(programId)) {
    return (
      <div className="space-y-6">
        <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
          <div className="text-center">
            <h2 className="text-2xl font-semibold text-red-600 mb-4">
              ID de programa inválido
            </h2>
            <p className="text-gray-600">
              El ID del programa proporcionado no es válido.
            </p>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <EditProgramForm programId={programId} />
      </div>
    </div>
  );
}

