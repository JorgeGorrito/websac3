import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { ProgramDurationSelect } from "./ProgramDurationSelect";
import { ProgramPeriodicitySelect } from "./ProgramPeriodicitySelect";
import { ProgramNameAutocomplete } from "./ProgramNameAutocomplete";
import {
  BookOpen,
  GraduationCap,
  Clock,
  Calendar,
  Target,
  UserCheck,
  User,
  Briefcase,
} from "lucide-react";

type RegisterProgramFormProps = {
  onCancel?: () => void;
  onSubmit?: (event: React.FormEvent<HTMLFormElement>) => void;
};

export function RegisterProgramForm({
  onCancel,
  onSubmit,
}: RegisterProgramFormProps) {
  return (
    <form className="space-y-8" onSubmit={onSubmit}>
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
                  htmlFor="nombre-programa"
                  className="text-sm font-semibold text-foreground"
                >
                  Nombre del programa
                </Label>
              </div>
              <ProgramNameAutocomplete />
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
              <ProgramPeriodicitySelect id="periodicidad" />
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
              <ProgramDurationSelect id="duracion" />
            </div>

            <div className="group space-y-3 md:col-start-2 md:row-start-2">
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
              placeholder="Describa el perfil profesional que desarrollarán..."
              className="min-h-[100px] border-input focus:border-ring focus:ring-ring/20 transition-all duration-200 resize-none"
            />
          </div>
        </div>
      </div>

      {/* Botones de acción */}
      <div className="flex justify-center gap-4">
        <Button type="button" variant="outline" onClick={onCancel}>
          Cancelar
        </Button>
        <Button type="submit">Guardar Programa</Button>
      </div>
    </form>
  );
}
