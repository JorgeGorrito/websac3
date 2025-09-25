import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { DegreeProgramItem } from "@/services/api";
import { 
  BookOpen, 
  Calendar, 
  Clock, 
  GraduationCap, 
  Target, 
  Eye,
  Edit,
  Trash2,
  Award,
  Upload,
  MessageSquare,
  FileText,
  Shield
} from "lucide-react";

export type ActionButton = {
  label: string;
  icon: React.ComponentType<{ className?: string }>;
  variant?: "default" | "outline" | "destructive" | "secondary" | "ghost" | "link";
  className?: string;
  onClick: (program: DegreeProgramItem) => void;
};

type DegreeProgramCardProps = {
  program: DegreeProgramItem;
  actions?: ActionButton[];
  // Mantener compatibilidad con la API anterior
  onView?: (program: DegreeProgramItem) => void;
  onEdit?: (program: DegreeProgramItem) => void;
  onDelete?: (program: DegreeProgramItem) => void;
};

export function DegreeProgramCard({ 
  program, 
  actions,
  onView, 
  onEdit, 
  onDelete 
}: DegreeProgramCardProps) {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('es-ES', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    });
  };

  return (
    <Card className="bg-white border border-gray-200 shadow-lg hover:shadow-xl transition-all duration-300 hover:-translate-y-1 py-0">
      <CardHeader className="pb-3 pt-4 bg-gradient-to-r from-blue-50 to-indigo-50 border-b border-gray-100">
        <div className="flex items-start justify-between">
          <div className="flex-1">
            <CardTitle className="text-xl font-bold text-gray-900 line-clamp-1 mb-2">
              {program.name}
            </CardTitle>
            <div className="flex items-center gap-2">
              <Badge className="bg-blue-100 text-blue-800 border-blue-200 font-semibold">
                SNIES: {program.snies}
              </Badge>
            </div>
          </div>
          <Badge className="bg-gradient-to-r from-green-500 to-emerald-500 text-white font-semibold px-3 py-1">
            {program.duration_unit.name}
          </Badge>
        </div>
      </CardHeader>
      
      <CardContent className="p-5 space-y-4">
        {/* Program Details */}
        <div className="space-y-3">
          {/* First Row: Credits and Duration */}
          <div className="grid grid-cols-1 sm:grid-cols-3 gap-3">
            <div className="flex items-center p-2 bg-blue-50 rounded-lg min-w-0">
              <div className="p-1.5 bg-blue-500 rounded-lg mr-2 flex-shrink-0">
                <GraduationCap className="h-3.5 w-3.5 text-white" />
              </div>
              <div className="min-w-0 flex-1">
                <p className="text-xs font-medium text-blue-600 uppercase tracking-wide">Créditos</p>
                <p className="text-base font-bold text-blue-900 truncate">{program.total_credits}</p>
              </div>
            </div>
            <div className="flex items-center p-2 bg-green-50 rounded-lg min-w-0 sm:col-span-2">
              <div className="p-1.5 bg-green-500 rounded-lg mr-2 flex-shrink-0">
                <Clock className="h-3.5 w-3.5 text-white" />
              </div>
              <div className="min-w-0 flex-1">
                <p className="text-xs font-medium text-green-600 uppercase tracking-wide mb-1">Duración</p>
                <div className="grid grid-cols-2 gap-1.5">
                  <div className="bg-green-100/50 rounded-md p-1.5">
                    <p className="text-xs font-medium text-green-700 uppercase tracking-wide mb-0.5">Periodicidad</p>
                    <p className="text-xs font-semibold text-green-800 break-words leading-tight">{program.duration_unit.name}</p>
                  </div>
                  <div className="bg-green-100/50 rounded-md p-1.5">
                    <p className="text-xs font-medium text-green-700 uppercase tracking-wide mb-0.5">Duración</p>
                    <p className="text-sm font-bold text-green-900 truncate">{program.duration_value}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
          
          {/* Second Row: Formation Level */}
          <div className="flex items-center p-3 bg-purple-50 rounded-lg min-w-0">
            <div className="p-2 bg-purple-500 rounded-lg mr-3 flex-shrink-0">
              <Award className="h-4 w-4 text-white" />
            </div>
            <div className="min-w-0 flex-1">
              <p className="text-xs font-medium text-purple-600 uppercase tracking-wide">Nivel de Formación</p>
              <p className="text-sm font-bold text-purple-900 break-words leading-tight">{program.formation_level.name}</p>
            </div>
          </div>
        </div>

        {/* Program Focus */}
        <div className="p-4 bg-purple-50 rounded-lg border border-purple-100">
          <div className="flex items-start mb-2">
            <div className="p-2 bg-purple-500 rounded-lg mr-3 flex-shrink-0">
              <Target className="h-4 w-4 text-white" />
            </div>
            <div className="min-w-0 flex-1">
              <span className="text-sm font-semibold text-purple-700 uppercase tracking-wide block mb-2">Enfoque</span>
              <p className="text-sm text-gray-700 break-words leading-relaxed">
                {program.program_focus}
              </p>
            </div>
          </div>
        </div>

        {/* Entry Profile */}
        <div className="p-4 bg-orange-50 rounded-lg border border-orange-100">
          <div className="flex items-start mb-2">
            <div className="p-2 bg-orange-500 rounded-lg mr-3 flex-shrink-0">
              <BookOpen className="h-4 w-4 text-white" />
            </div>
            <div className="min-w-0 flex-1">
              <span className="text-sm font-semibold text-orange-700 uppercase tracking-wide block mb-2">Perfil de Ingreso</span>
              <p className="text-sm text-gray-700 break-words leading-relaxed">
                {program.entry_profile}
              </p>
            </div>
          </div>
        </div>

        {/* Institution */}
        <div className="p-4 bg-indigo-50 rounded-lg border border-indigo-100">
          <div className="flex items-start mb-2">
            <div className="p-2 bg-indigo-500 rounded-lg mr-3 flex-shrink-0">
              <GraduationCap className="h-4 w-4 text-white" />
            </div>
            <div className="min-w-0 flex-1">
              <span className="text-sm font-semibold text-indigo-700 uppercase tracking-wide block mb-2">Institución</span>
              <p className="text-sm font-medium text-gray-800 break-words leading-relaxed mb-2">
                {program.higher_education_institution.name}
              </p>
              <Badge className="bg-indigo-100 text-indigo-800 border-indigo-200 font-medium">
                SNIES: {program.higher_education_institution.snies}
              </Badge>
            </div>
          </div>
        </div>

        {/* Actions */}
        <div className="flex flex-col sm:flex-row gap-3 pt-4 border-t border-gray-100">
          {actions ? (
            // Usar acciones personalizadas
            actions.map((action, index) => {
              const IconComponent = action.icon;
              return (
                <Button
                  key={index}
                  variant={action.variant || "outline"}
                  size="sm"
                  className={`flex-1 transition-all duration-200 ${action.className || ""}`}
                  onClick={() => action.onClick(program)}
                >
                  <IconComponent className="h-4 w-4 mr-2" />
                  {action.label}
                </Button>
              );
            })
          ) : (
            // Mantener compatibilidad con la API anterior
            <>
              {onView && (
                <Button
                  variant="outline"
                  size="sm"
                  className="flex-1 border-blue-300 text-blue-600 hover:bg-blue-50 hover:border-blue-400 transition-all duration-200"
                  onClick={() => onView(program)}
                >
                  <Eye className="h-4 w-4 mr-2" />
                  Ver
                </Button>
              )}
              {onEdit && (
                <Button
                  variant="outline"
                  size="sm"
                  className="flex-1 border-green-300 text-green-600 hover:bg-green-50 hover:border-green-400 transition-all duration-200"
                  onClick={() => onEdit(program)}
                >
                  <Edit className="h-4 w-4 mr-2" />
                  Editar
                </Button>
              )}
              {onDelete && (
                <Button
                  variant="outline"
                  size="sm"
                  className="border-red-300 text-red-600 hover:bg-red-50 hover:border-red-400 transition-all duration-200 sm:w-auto"
                  onClick={() => onDelete(program)}
                >
                  <Trash2 className="h-4 w-4" />
                </Button>
              )}
            </>
          )}
        </div>
      </CardContent>
    </Card>
  );
}
