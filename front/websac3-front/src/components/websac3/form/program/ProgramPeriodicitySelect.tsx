import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { useListDurationUnitsQuery } from "@/services/api";

type ProgramPeriodicitySelectProps = {
  id?: string;
  placeholder?: string;
  defaultValue?: string;
  value?: string;
  onValueChange?: (value: string) => void;
};

export function ProgramPeriodicitySelect({ id, placeholder = "Seleccionar", defaultValue, value, onValueChange }: ProgramPeriodicitySelectProps) {
  const { data: durationUnits, isLoading, error } = useListDurationUnitsQuery();

  // Determine the current value to display
  const currentValue = value || defaultValue;
  
  // Find the selected unit for display
  const selectedUnit = durationUnits?.find(unit => unit.id.toString() === currentValue);

  if (isLoading) {
    return (
      <Select disabled value={currentValue}>
        <SelectTrigger id={id} className="h-10">
          <SelectValue placeholder="Cargando..." />
        </SelectTrigger>
      </Select>
    );
  }

  if (error) {
    return (
      <Select disabled value={currentValue}>
        <SelectTrigger id={id} className="h-10">
          <SelectValue placeholder="Error al cargar" />
        </SelectTrigger>
      </Select>
    );
  }

  return (
    <Select value={currentValue} onValueChange={onValueChange}>
      <SelectTrigger id={id} className="h-10">
        <SelectValue placeholder={placeholder}>
          {selectedUnit ? selectedUnit.name : placeholder}
        </SelectValue>
      </SelectTrigger>
      <SelectContent>
        {durationUnits?.map((unit) => (
          <SelectItem key={unit.id} value={unit.id.toString()}>
            {unit.name}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
}


