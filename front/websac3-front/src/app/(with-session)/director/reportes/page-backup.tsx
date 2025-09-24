"use client";

import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { FileText } from "lucide-react";

export default function ReportesPage() {
  const [test, setTest] = useState(false);

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Consultar Reportes</h1>
          <p className="text-gray-600 mt-2">Test page</p>
        </div>
        <Button onClick={() => setTest(!test)}>
          <FileText className="h-4 w-4 mr-2" />
          Test
        </Button>
      </div>
    </div>
  );
}
