"use client";

import { useState, useEffect, useMemo, useCallback } from "react";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { 
  useGetDegreeProgramTemplateQuery,
  useBulkImportDegreeProgramsMutation,
  useGetDurationUnitsQuery,
  useGetFormationLevelsQuery,
  useGetProfessionalRolesQuery,
  type BulkImportResult
} from "@/services/api";
import { 
  Upload, 
  Download, 
  FileText, 
  CheckCircle, 
  XCircle, 
  AlertTriangle,
  ArrowLeft,
  FileSpreadsheet,
  Clock,
  GraduationCap,
  Users,
  Copy,
  Check,
  Search,
  X
} from "lucide-react";
import { BulkImportResults } from "@/components/websac3/form/program/BulkImportResults";

// Separate component to prevent re-renders
const ReferenceDataCard = ({ 
  title, 
  data, 
  isLoading, 
  error,
  icon: Icon, 
  color,
  searchValue,
  onSearchChange,
  searchPlaceholder,
  onCopy
}: {
  title: string;
  data: any[] | undefined;
  isLoading: boolean;
  error: any;
  icon: React.ElementType;
  color: string;
  searchValue: string;
  onSearchChange: (value: string) => void;
  searchPlaceholder: string;
  onCopy: (id: string, field: string) => void;
}) => {
  // Ensure data is an array and handle error states
  const safeData = Array.isArray(data) ? data : [];
  // Only show error state if it's a real error object with isError flag
  const hasError = error && !isLoading && error.isError;
  const isEmpty = !hasError && !isLoading && safeData.length === 0;
  
  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-lg flex items-center gap-2">
          <Icon className={`h-5 w-5 ${color}`} />
          {title}
        </CardTitle>
      </CardHeader>
      <CardContent>
        {/* Search Input */}
        <div className="relative mb-4">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-gray-400" />
          <input
            type="text"
            placeholder={searchPlaceholder}
            value={searchValue}
            onChange={(e) => onSearchChange(e.target.value)}
            className="w-full pl-10 pr-10 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
          {searchValue && (
            <button
              onClick={() => onSearchChange("")}
              className="absolute right-3 top-1/2 transform -translate-y-1/2 p-1 hover:bg-gray-200 rounded transition-colors"
              title="Limpiar búsqueda"
            >
              <X className="h-4 w-4 text-gray-400" />
            </button>
          )}
        </div>

        {isLoading ? (
          <div className="space-y-2">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-3/4" />
            <Skeleton className="h-4 w-1/2" />
          </div>
        ) : hasError ? (
          <div className="text-center py-4">
            <div className="flex flex-col items-center gap-2">
              <AlertTriangle className="h-8 w-8 text-orange-500" />
              <p className="text-sm text-gray-600">
                Error al cargar los datos
              </p>
              <p className="text-xs text-gray-500">
                Intenta recargar la página
              </p>
            </div>
          </div>
        ) : isEmpty ? (
          <div className="text-center py-4">
            <p className="text-sm text-gray-500">
              {searchValue ? "No se encontraron resultados" : "No hay datos disponibles"}
            </p>
          </div>
        ) : (
          <div className="space-y-2 max-h-48 overflow-y-auto">
            {safeData.map((item) => (
              <div key={item.id} className="flex items-center justify-between p-2 bg-gray-50 rounded border">
                <div className="flex-1">
                  <span className="font-mono text-sm bg-blue-100 text-blue-800 px-2 py-1 rounded">
                    {item.id}
                  </span>
                  <span className="ml-2 text-sm text-gray-700">{item.name}</span>
                </div>
                <button
                  onClick={() => onCopy(item.id.toString(), `${title}-${item.id}`)}
                  className="ml-2 p-1 hover:bg-gray-200 rounded transition-colors"
                  title="Copiar ID"
                >
                  <Copy className="h-4 w-4 text-gray-500" />
                </button>
              </div>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  );
};

export default function CargaMasivaProgramasPage() {
  const router = useRouter();
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [isUploading, setIsUploading] = useState(false);
  const [importResult, setImportResult] = useState<BulkImportResult | null>(null);
  const [copiedField, setCopiedField] = useState<string | null>(null);
  
  // Search states (for input)
  const [durationUnitSearchInput, setDurationUnitSearchInput] = useState("");
  const [formationLevelSearchInput, setFormationLevelSearchInput] = useState("");
  const [professionalRoleSearchInput, setProfessionalRoleSearchInput] = useState("");
  
  // Debounced search states (for API calls)
  const [durationUnitSearch, setDurationUnitSearch] = useState("");
  const [formationLevelSearch, setFormationLevelSearch] = useState("");
  const [professionalRoleSearch, setProfessionalRoleSearch] = useState("");

  const [bulkImport, { isLoading: isImporting }] = useBulkImportDegreeProgramsMutation();
  const { refetch: downloadTemplate, isLoading: isDownloading } = useGetDegreeProgramTemplateQuery({}, {
    skip: true,
  });

  // Memoize query parameters to prevent unnecessary re-renders
  const durationUnitQueryParams = useMemo(() => ({
    name: durationUnitSearch || undefined,
  }), [durationUnitSearch]);

  const formationLevelQueryParams = useMemo(() => ({
    name: formationLevelSearch || undefined,
  }), [formationLevelSearch]);

  const professionalRoleQueryParams = useMemo(() => ({
    name: professionalRoleSearch || undefined,
  }), [professionalRoleSearch]);

  // Reference data queries with search filters
  const { data: durationUnits, isLoading: isLoadingDurationUnits, error: durationUnitsError } = useGetDurationUnitsQuery(durationUnitQueryParams, {
    skip: false,
  });
  const { data: formationLevels, isLoading: isLoadingFormationLevels, error: formationLevelsError } = useGetFormationLevelsQuery(formationLevelQueryParams, {
    skip: false,
  });
  const { data: professionalRoles, isLoading: isLoadingProfessionalRoles, error: professionalRolesError } = useGetProfessionalRolesQuery(professionalRoleQueryParams, {
    skip: false,
  });

  // Debounce search inputs
  useEffect(() => {
    const timer = setTimeout(() => {
      setDurationUnitSearch(durationUnitSearchInput);
    }, 300);
    return () => clearTimeout(timer);
  }, [durationUnitSearchInput]);

  useEffect(() => {
    const timer = setTimeout(() => {
      setFormationLevelSearch(formationLevelSearchInput);
    }, 300);
    return () => clearTimeout(timer);
  }, [formationLevelSearchInput]);

  useEffect(() => {
    const timer = setTimeout(() => {
      setProfessionalRoleSearch(professionalRoleSearchInput);
    }, 300);
    return () => clearTimeout(timer);
  }, [professionalRoleSearchInput]);

  const handleDownloadTemplate = async () => {
    try {
      const result = await downloadTemplate();
      if (result.data) {
        const blob = result.data as Blob;
        const url = window.URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        link.download = 'plantilla-programas-grado.xlsx';
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
        window.URL.revokeObjectURL(url);
      }
    } catch (error) {
      console.error('Error downloading template:', error);
    }
  };

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      const validTypes = [
        'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
        'application/vnd.ms-excel',
        'text/csv'
      ];
      
      if (!validTypes.includes(file.type)) {
        alert('Por favor selecciona un archivo Excel (.xlsx, .xls) o CSV válido.');
        return;
      }
      
      setSelectedFile(file);
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) return;

    setIsUploading(true);
    try {
      const result = await bulkImport({ file: selectedFile }).unwrap();
      
      if (result.result) {
        setImportResult(result.result);
        setSelectedFile(null);
      }
    } catch (error: any) {
      console.error('Error uploading file:', error);
      alert('Error al procesar el archivo. Por favor, verifica el formato y vuelve a intentar.');
    } finally {
      setIsUploading(false);
    }
  };

  const copyToClipboard = useCallback(async (text: string, field: string) => {
    try {
      await navigator.clipboard.writeText(text);
      setCopiedField(field);
      setTimeout(() => setCopiedField(null), 2000);
    } catch (err) {
      console.error('Failed to copy: ', err);
    }
  }, []);


  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="bg-white rounded-xl shadow-sm p-6 md:p-8">
        <div className="flex items-center gap-4 mb-6">
          <Button
            variant="outline"
            size="sm"
            onClick={() => router.back()}
            className="flex items-center gap-2"
          >
            <ArrowLeft className="h-4 w-4" />
            Volver
          </Button>
          <div>
            <h1 className="text-2xl font-semibold text-gray-900">Carga Masiva de Programas</h1>
            <p className="text-gray-600 text-sm">Importa múltiples programas de grado desde un archivo</p>
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Reference Data */}
        <div className="lg:col-span-1 space-y-4">
          <Card className="bg-blue-50 border-blue-200">
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2 text-blue-800">
                <AlertTriangle className="h-5 w-5" />
                Guía de IDs
              </CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-sm text-blue-700 mb-4">
                Utiliza estos IDs en tu archivo CSV/Excel. Haz clic en el ícono de copiar para copiar el ID.
              </p>
            </CardContent>
          </Card>

          <ReferenceDataCard
            title="Unidades de Duración (duration_unit_id)"
            data={durationUnits}
            isLoading={isLoadingDurationUnits}
            error={durationUnitsError}
            icon={Clock}
            color="text-blue-600"
            searchValue={durationUnitSearchInput}
            onSearchChange={setDurationUnitSearchInput}
            searchPlaceholder="Buscar unidad de duración..."
            onCopy={copyToClipboard}
          />

          <ReferenceDataCard
            title="Niveles de Formación (formation_level_id)"
            data={formationLevels}
            isLoading={isLoadingFormationLevels}
            error={formationLevelsError}
            icon={GraduationCap}
            color="text-green-600"
            searchValue={formationLevelSearchInput}
            onSearchChange={setFormationLevelSearchInput}
            searchPlaceholder="Buscar nivel de formación..."
            onCopy={copyToClipboard}
          />

          <ReferenceDataCard
            title="Roles Profesionales (professional_role_ids)"
            data={professionalRoles}
            isLoading={isLoadingProfessionalRoles}
            error={professionalRolesError}
            icon={Users}
            color="text-purple-600"
            searchValue={professionalRoleSearchInput}
            onSearchChange={setProfessionalRoleSearchInput}
            searchPlaceholder="Buscar rol profesional..."
            onCopy={copyToClipboard}
          />
        </div>

        {/* Upload Section */}
        <div className="lg:col-span-2 space-y-6">
          {/* Step 1: Download Template */}
          <Card>
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <Download className="h-5 w-5 text-blue-600" />
                Paso 1: Descargar Plantilla
              </CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-gray-600 mb-4">
                Descarga la plantilla de Excel para conocer el formato correcto de los datos.
              </p>
              <Button 
                onClick={handleDownloadTemplate}
                disabled={isDownloading}
                className="w-full"
              >
                {isDownloading ? "Descargando..." : "Descargar Plantilla"}
              </Button>
            </CardContent>
          </Card>

          {/* Step 2: Upload File */}
          <Card>
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2">
                <Upload className="h-5 w-5 text-green-600" />
                Paso 2: Subir Archivo
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <p className="text-gray-600">
                  Selecciona el archivo Excel (.xlsx, .xls) o CSV con los datos de los programas.
                </p>
                
                <div className="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center">
                  <FileText className="h-12 w-12 text-gray-400 mx-auto mb-4" />
                  <input
                    type="file"
                    accept=".xlsx,.xls,.csv"
                    onChange={handleFileSelect}
                    className="hidden"
                    id="file-upload"
                  />
                  <label
                    htmlFor="file-upload"
                    className="cursor-pointer inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50"
                  >
                    Seleccionar Archivo
                  </label>
                  {selectedFile && (
                    <div className="mt-4 p-3 bg-green-50 border border-green-200 rounded-lg">
                      <div className="flex items-center gap-2">
                        <CheckCircle className="h-4 w-4 text-green-600" />
                        <span className="text-sm font-medium text-green-800">
                          {selectedFile.name}
                        </span>
                      </div>
                      <p className="text-xs text-green-600 mt-1">
                        Tamaño: {(selectedFile.size / 1024 / 1024).toFixed(2)} MB
                      </p>
                    </div>
                  )}
                </div>

                <div className="flex gap-3">
                  <Button
                    onClick={handleUpload}
                    disabled={!selectedFile || isUploading}
                    className="flex-1"
                  >
                    {isUploading ? "Procesando..." : "Subir Archivo"}
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Instructions */}
          <Card className="bg-yellow-50 border-yellow-200">
            <CardHeader>
              <CardTitle className="text-lg flex items-center gap-2 text-yellow-800">
                <AlertTriangle className="h-5 w-5" />
                Instrucciones Importantes
              </CardTitle>
            </CardHeader>
            <CardContent>
              <ul className="text-sm text-yellow-700 space-y-2">
                <li>• Utiliza la plantilla descargada para mantener el formato correcto</li>
                <li>• Usa los IDs mostrados en la guía lateral para los campos requeridos</li>
                <li>• Asegúrate de que todos los campos obligatorios estén completos</li>
                <li>• Los códigos SNIES deben ser únicos</li>
                <li>• Los roles profesionales se pueden separar por comas si hay múltiples</li>
                <li>• Revisa los datos antes de subir el archivo</li>
              </ul>
            </CardContent>
          </Card>
        </div>
      </div>

      {/* Results Modal */}
      {importResult && (
        <BulkImportResults
          result={importResult}
          onClose={() => setImportResult(null)}
        />
      )}
    </div>
  );
}
