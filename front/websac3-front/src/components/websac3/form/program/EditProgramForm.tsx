import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { ProgramPeriodicitySelect } from "./ProgramPeriodicitySelect";
import { ProgramFormationLevelSelect } from "./ProgramFormationLevelSelect";
import { ProgramProfessionalRolesSelect } from "./ProgramProfessionalRolesSelect";
import { useUpdateDegreeProgramMutation, useGetDegreeProgramQuery, type DegreeProgramItem } from "@/services/api";
import {
  BookOpen,
  GraduationCap,
  Clock,
  Calendar,
  Target,
  UserCheck,
  User,
  Briefcase,
  CheckCircle,
  XCircle,
  X,
  Award,
  ArrowLeft,
} from "lucide-react";

type EditProgramFormProps = {
  programId: number;
  onCancel?: () => void;
  onSubmit?: (event: React.FormEvent<HTMLFormElement>) => void;
};

export function EditProgramForm({
  programId,
  onCancel,
  onSubmit,
}: EditProgramFormProps) {
  const router = useRouter();
  
  // Fetch the program data
  const { data: programData, isLoading: isLoadingProgram, error: programError } = useGetDegreeProgramQuery({
    degree_program_id: programId,
  });

  const [formData, setFormData] = useState({
    name: "",
    total_credits: "",
    program_focus: "",
    snies: "",
    duration_unit_id: "",
    duration_value: "",
    formation_level_id: "",
    professional_role_ids: [] as number[],
    entry_profile: "",
    graduate_profile: "",
    professional_profile: "",
  });

  const [modalState, setModalState] = useState<{
    isOpen: boolean;
    type: "success" | "error";
    title: string;
    message: string;
  }>({
    isOpen: false,
    type: "success",
    title: "",
    message: "",
  });

  const [updateDegreeProgram, { isLoading, error }] = useUpdateDegreeProgramMutation();

  // Populate form data when program data is loaded
  useEffect(() => {
    if (programData) {
      setFormData({
        name: programData.name || "",
        total_credits: programData.total_credits?.toString() || "",
        program_focus: programData.program_focus || "",
        snies: programData.snies?.toString() || "",
        duration_unit_id: programData.duration_unit?.id?.toString() || "",
        duration_value: programData.duration_value?.toString() || "",
        formation_level_id: programData.formation_level?.id?.toString() || "",
        professional_role_ids: programData.professional_roles?.map(role => role.id) || [],
        entry_profile: programData.entry_profile || "",
        graduate_profile: programData.graduate_profile || "",
        professional_profile: programData.professional_profile || "",
      });
    }
  }, [programData]);

  const handleInputChange = (field: string, value: string) => {
    setFormData(prev => ({ ...prev, [field]: value }));
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    
    try {
      const response = await updateDegreeProgram({
        degree_program_id: programId,
        body: {
          name: formData.name,
          total_credits: parseInt(formData.total_credits),
          program_focus: formData.program_focus,
          snies: parseInt(formData.snies),
          duration_unit_id: parseInt(formData.duration_unit_id),
          duration_value: parseInt(formData.duration_value),
          formation_level_id: parseInt(formData.formation_level_id),
          professional_role_ids: formData.professional_role_ids,
          entry_profile: formData.entry_profile,
          graduate_profile: formData.graduate_profile,
          professional_profile: formData.professional_profile,
        }
      }).unwrap();

      // Handle success
      if (response.result) {
        setModalState({
          isOpen: true,
          type: "success",
          title: "¡Programa Actualizado!",
          message: response.result || "El programa ha sido actualizado exitosamente. Redirigiendo al dashboard...",
        });

        // Redirect to programs list after 2 seconds
        setTimeout(() => {
          router.push("/director/programas");
        }, 2000);
      }
    } catch (err: any) {
      console.error("Error updating degree program:", err);
      let errorMessage = "Error al actualizar el programa";
      
      if (err?.data?.errors && Array.isArray(err.data.errors)) {
        errorMessage = err.data.errors.join(", ");
      } else if (err?.data?.result) {
        errorMessage = err.data.result;
      }
      
      setModalState({
        isOpen: true,
        type: "error",
        title: "Error al Actualizar",
        message: errorMessage,
      });
    }
  };

  // Show loading state while fetching program data
  if (isLoadingProgram) {
    return (
      <div className="flex items-center justify-center min-h-[400px]">
        <div className="text-center">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto mb-4"></div>
          <p className="text-gray-600">Cargando datos del programa...</p>
        </div>
      </div>
    );
  }

  // Show error state if program couldn't be loaded
  if (programError || !programData) {
    return (
      <div className="text-center py-12">
        <XCircle className="h-12 w-12 text-red-500 mx-auto mb-4" />
        <h3 className="text-lg font-medium text-gray-900 mb-2">
          Error al cargar el programa
        </h3>
        <p className="text-gray-600 mb-6">
          No se pudo cargar la información del programa. Por favor, inténtalo de nuevo.
        </p>
        <Button onClick={() => router.back()} variant="outline">
          <ArrowLeft className="h-4 w-4 mr-2" />
          Volver
        </Button>
      </div>
    );
  }

  return (
    <form className="space-y-8" onSubmit={handleSubmit}>
      {/* Header with back button */}
      <div className="flex items-center gap-4 mb-6">
        <Button
          type="button"
          variant="outline"
          size="sm"
          onClick={() => router.back()}
          className="flex items-center gap-2"
        >
          <ArrowLeft className="h-4 w-4" />
          Volver
        </Button>
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Editar Programa</h1>
          <p className="text-gray-600">Modifica la información del programa de grado</p>
        </div>
      </div>

      {/* Sección de información básica */}
      <div className="relative">
        <div className="absolute -top-3 left-0 bg-primary text-primary-foreground px-4 py-1 rounded-full text-sm font-medium shadow-lg">
          Información Básica
        </div>
        <div className="border-2 border-border rounded-2xl p-6 bg-card shadow-sm">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="group space-y-3">
              <div className="flex items-center gap-2">
                <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                  <BookOpen className="w-4 h-4 text-primary" />
                </div>
                <Label
                  htmlFor="snies"
                  className="text-sm font-semibold text-foreground"
                >
                  Código SNIES
                </Label>
              </div>
              <Input
                id="snies"
                type="number"
                min="0"
                value={formData.snies}
                onChange={(e) => handleInputChange("snies", e.target.value)}
                className="h-11 border-input focus:border-ring focus:ring-ring/20 transition-all duration-200"
                placeholder="Ej. 12345"
              />
            </div>

            <div className="group space-y-3">
              <div className="flex items-center gap-2">
                <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                  <BookOpen className="w-4 h-4 text-primary" />
                </div>
                <Label
                  htmlFor="nombre-programa"
                  className="text-sm font-semibold text-foreground"
                >
                  Nombre del programa
                </Label>
              </div>
              <Input
                id="nombre-programa"
                value={formData.name}
                onChange={(e) => handleInputChange("name", e.target.value)}
                className="h-11 border-input focus:border-ring focus:ring-ring/20 transition-all duration-200"
                placeholder="Ej. Ingeniería de Sistemas"
              />
            </div>

            <div className="group space-y-3">
              <div className="flex items-center gap-2">
                <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                  <GraduationCap className="w-4 h-4 text-primary" />
                </div>
                <Label
                  htmlFor="creditos"
                  className="text-sm font-semibold text-foreground"
                >
                  Número de créditos
                </Label>
              </div>
              <Input
                id="creditos"
                type="number"
                value={formData.total_credits}
                onChange={(e) => handleInputChange("total_credits", e.target.value)}
                className="h-11 border-input focus:border-ring focus:ring-ring/20 transition-all duration-200"
                placeholder="Ej. 160"
              />
            </div>

            <div className="group space-y-3">
              <div className="flex items-center gap-2">
                <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                  <Calendar className="w-4 h-4 text-primary" />
                </div>
                <Label
                  htmlFor="periodicidad"
                  className="text-sm font-semibold text-foreground"
                >
                  Periodicidad
                </Label>
              </div>
              <ProgramPeriodicitySelect 
                id="periodicidad"
                value={formData.duration_unit_id}
                onValueChange={(value) => handleInputChange("duration_unit_id", value)}
              />
            </div>

            <div className="group space-y-3">
              <div className="flex items-center gap-2">
                <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                  <Clock className="w-4 h-4 text-primary" />
                </div>
                <Label
                  htmlFor="duracion"
                  className="text-sm font-semibold text-foreground"
                >
                  Duración
                </Label>
              </div>
              <Input
                id="duracion"
                type="number"
                min="0"
                value={formData.duration_value}
                onChange={(e) => handleInputChange("duration_value", e.target.value)}
                className="h-11 border-input focus:border-ring focus:ring-ring/20 transition-all duration-200"
                placeholder="Ej. 5"
              />
            </div>

            <div className="group space-y-3">
              <div className="flex items-center gap-2">
                <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                  <Award className="w-4 h-4 text-primary" />
                </div>
                <Label
                  htmlFor="nivel-formacion"
                  className="text-sm font-semibold text-foreground"
                >
                  Nivel de formación
                </Label>
              </div>
              <ProgramFormationLevelSelect 
                id="nivel-formacion"
                value={formData.formation_level_id}
                onValueChange={(value) => handleInputChange("formation_level_id", value)}
              />
            </div>

            <div className="group space-y-3">
              <div className="flex items-center gap-2">
                <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                  <Briefcase className="w-4 h-4 text-primary" />
                </div>
                <Label
                  htmlFor="roles-profesionales"
                  className="text-sm font-semibold text-foreground"
                >
                  Roles Profesionales de Ciberseguridad
                </Label>
              </div>
              <ProgramProfessionalRolesSelect
                id="roles-profesionales"
                value={formData.professional_role_ids}
                onValueChange={(value) => handleInputChange("professional_role_ids", value)}
              />
            </div>

            <div className="group space-y-3">
              <div className="flex items-center gap-2">
                <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                  <Target className="w-4 h-4 text-primary" />
                </div>
                <Label
                  htmlFor="enfoque"
                  className="text-sm font-semibold text-foreground"
                >
                  Enfoque del programa
                </Label>
              </div>
              <Input
                id="enfoque"
                value={formData.program_focus}
                onChange={(e) => handleInputChange("program_focus", e.target.value)}
                className="h-11 border-input focus:border-ring focus:ring-ring/20 transition-all duration-200"
                placeholder="Ej. Ingeniería de Software"
              />
            </div>
          </div>
        </div>
      </div>

      {/* Sección de perfiles */}
      <div className="relative">
        <div className="absolute -top-3 left-0 bg-secondary text-secondary-foreground px-4 py-1 rounded-full text-sm font-medium shadow-lg">
          Perfiles del Programa
        </div>
        <div className="border-2 border-border rounded-2xl p-6 bg-card shadow-sm space-y-6">
          <div className="group space-y-3">
            <div className="flex items-center gap-2">
              <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                <UserCheck className="w-4 h-4 text-secondary" />
              </div>
              <Label
                htmlFor="perfil-ingreso"
                className="text-sm font-semibold text-foreground"
              >
                Perfil de ingreso
              </Label>
            </div>
            <Textarea
              id="perfil-ingreso"
              value={formData.entry_profile}
              onChange={(e) => handleInputChange("entry_profile", e.target.value)}
              placeholder="Describa el perfil de ingreso requerido para el programa..."
              className="min-h-[100px] border-input focus:border-ring focus:ring-ring/20 transition-all duration-200 resize-none"
            />
          </div>

          <div className="group space-y-3">
            <div className="flex items-center gap-2">
              <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                <User className="w-4 h-4 text-secondary" />
              </div>
              <Label
                htmlFor="perfil-egreso"
                className="text-sm font-semibold text-foreground"
              >
                Perfil de egreso
              </Label>
            </div>
            <Textarea
              id="perfil-egreso"
              value={formData.graduate_profile}
              onChange={(e) => handleInputChange("graduate_profile", e.target.value)}
              placeholder="Describa el perfil de egreso que tendrán los estudiantes..."
              className="min-h-[100px] border-input focus:border-ring focus:ring-ring/20 transition-all duration-200 resize-none"
            />
          </div>

          <div className="group space-y-3">
            <div className="flex items-center gap-2">
              <div className="p-2 bg-muted rounded-lg group-hover:bg-secondary/20 transition-colors">
                <Briefcase className="w-4 h-4 text-secondary" />
              </div>
              <Label
                htmlFor="perfil-profesional"
                className="text-sm font-semibold text-foreground"
              >
                Perfil profesional
              </Label>
            </div>
            <Textarea
              id="perfil-profesional"
              value={formData.professional_profile}
              onChange={(e) => handleInputChange("professional_profile", e.target.value)}
              placeholder="Describa el perfil profesional que desarrollarán..."
              className="min-h-[100px] border-input focus:border-ring focus:ring-ring/20 transition-all duration-200 resize-none"
            />
          </div>
        </div>
      </div>

      {/* Botones de acción */}
      <div className="flex justify-center gap-4">
        <Button type="button" variant="outline" onClick={onCancel} disabled={isLoading}>
          Cancelar
        </Button>
        <Button type="submit" disabled={isLoading}>
          {isLoading ? "Actualizando..." : "Actualizar Programa"}
        </Button>
      </div>

      {/* Modal personalizado */}
      {modalState.isOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center">
          {/* Overlay */}
          <div 
            className="absolute inset-0 bg-black/50"
            onClick={() => setModalState(prev => ({ ...prev, isOpen: false }))}
          />
          
          {/* Modal Content */}
          <div className="relative bg-white rounded-lg shadow-xl max-w-md w-full mx-4 p-6">
            {/* Botón de cerrar */}
            <button
              onClick={() => setModalState(prev => ({ ...prev, isOpen: false }))}
              className="absolute top-4 right-4 p-1 rounded-full hover:bg-gray-100 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-300"
              aria-label="Cerrar modal"
            >
              <X className="h-5 w-5 text-gray-500 hover:text-gray-700" />
            </button>
            
            {/* Header */}
            <div className="flex items-center gap-3 pr-8 mb-4">
              {modalState.type === "success" ? (
                <div className="flex items-center justify-center h-10 w-10 rounded-full bg-green-100">
                  <CheckCircle className="h-5 w-5 text-green-600" />
                </div>
              ) : (
                <div className="flex items-center justify-center h-10 w-10 rounded-full bg-red-100">
                  <XCircle className="h-5 w-5 text-red-600" />
                </div>
              )}
              <h2 className={`text-lg font-semibold ${modalState.type === "success" ? "text-green-800" : "text-red-800"}`}>
                {modalState.title}
              </h2>
            </div>
            
            {/* Message */}
            <p className="text-gray-600 mb-6">
              {modalState.message}
            </p>
            
            {/* Footer */}
            <div className="flex justify-end">
              <Button
                onClick={() => {
                  setModalState(prev => ({ ...prev, isOpen: false }));
                  if (modalState.type === "success") {
                    router.push("/director/programas");
                  }
                }}
                className={modalState.type === "success" ? "bg-green-600 hover:bg-green-700" : "bg-red-600 hover:bg-red-700"}
              >
                {modalState.type === "success" ? "Ir a Programas" : "Entendido"}
              </Button>
            </div>
          </div>
        </div>
      )}
    </form>
  );
}
