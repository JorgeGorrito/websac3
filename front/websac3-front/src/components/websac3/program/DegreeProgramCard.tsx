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
  Trash2
} from "lucide-react";

type DegreeProgramCardProps = {
  program: DegreeProgramItem;
  onView?: (program: DegreeProgramItem) => void;
  onEdit?: (program: DegreeProgramItem) => void;
  onDelete?: (program: DegreeProgramItem) => void;
};

export function DegreeProgramCard({ 
  program, 
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
        <div className="grid grid-cols-2 gap-4">
          <div className="flex items-center p-3 bg-blue-50 rounded-lg">
            <div className="p-2 bg-blue-500 rounded-lg mr-3">
              <GraduationCap className="h-4 w-4 text-white" />
            </div>
            <div>
              <p className="text-xs font-medium text-blue-600 uppercase tracking-wide">Créditos</p>
              <p className="text-lg font-bold text-blue-900">{program.total_credits}</p>
            </div>
          </div>
          <div className="flex items-center p-3 bg-green-50 rounded-lg">
            <div className="p-2 bg-green-500 rounded-lg mr-3">
              <Clock className="h-4 w-4 text-white" />
            </div>
            <div>
              <p className="text-xs font-medium text-green-600 uppercase tracking-wide">Duración</p>
              <p className="text-lg font-bold text-green-900">{program.duration_value} {program.duration_unit.name.toLowerCase()}</p>
            </div>
          </div>
        </div>

        {/* Program Focus */}
        <div className="p-4 bg-purple-50 rounded-lg border border-purple-100">
          <div className="flex items-center mb-2">
            <div className="p-2 bg-purple-500 rounded-lg mr-3">
              <Target className="h-4 w-4 text-white" />
            </div>
            <span className="text-sm font-semibold text-purple-700 uppercase tracking-wide">Enfoque</span>
          </div>
          <p className="text-sm text-gray-700 line-clamp-2 pl-11">
            {program.program_focus}
          </p>
        </div>

        {/* Entry Profile */}
        <div className="p-4 bg-orange-50 rounded-lg border border-orange-100">
          <div className="flex items-center mb-2">
            <div className="p-2 bg-orange-500 rounded-lg mr-3">
              <BookOpen className="h-4 w-4 text-white" />
            </div>
            <span className="text-sm font-semibold text-orange-700 uppercase tracking-wide">Perfil de Ingreso</span>
          </div>
          <p className="text-sm text-gray-700 line-clamp-2 pl-11">
            {program.entry_profile}
          </p>
        </div>

        {/* Institution */}
        <div className="p-4 bg-indigo-50 rounded-lg border border-indigo-100">
          <div className="flex items-center mb-2">
            <div className="p-2 bg-indigo-500 rounded-lg mr-3">
              <GraduationCap className="h-4 w-4 text-white" />
            </div>
            <span className="text-sm font-semibold text-indigo-700 uppercase tracking-wide">Institución</span>
          </div>
          <p className="text-sm font-medium text-gray-800 pl-11 mb-1">
            {program.higher_education_institution.name}
          </p>
          <Badge className="ml-11 bg-indigo-100 text-indigo-800 border-indigo-200 font-medium">
            SNIES: {program.higher_education_institution.snies}
          </Badge>
        </div>

        {/* Actions */}
        <div className="flex gap-3 pt-4 border-t border-gray-100">
          <Button
            variant="outline"
            size="sm"
            className="flex-1 border-blue-300 text-blue-600 hover:bg-blue-50 hover:border-blue-400 transition-all duration-200"
            onClick={() => onView?.(program)}
          >
            <Eye className="h-4 w-4 mr-2" />
            Ver
          </Button>
          <Button
            variant="outline"
            size="sm"
            className="flex-1 border-green-300 text-green-600 hover:bg-green-50 hover:border-green-400 transition-all duration-200"
            onClick={() => onEdit?.(program)}
          >
            <Edit className="h-4 w-4 mr-2" />
            Editar
          </Button>
          <Button
            variant="outline"
            size="sm"
            className="border-red-300 text-red-600 hover:bg-red-50 hover:border-red-400 transition-all duration-200"
            onClick={() => onDelete?.(program)}
          >
            <Trash2 className="h-4 w-4" />
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}
