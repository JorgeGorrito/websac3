"use client";

import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { useListFormationLevelsQuery } from "@/services/api";

type ProgramFormationLevelSelectProps = {
  id: string;
  placeholder?: string;
  defaultValue?: string;
  value?: string;
  onValueChange?: (value: string) => void;
};

export function ProgramFormationLevelSelect({ 
  id, 
  placeholder = "Seleccionar", 
  defaultValue, 
  value, 
  onValueChange 
}: ProgramFormationLevelSelectProps) {
  const { data: formationLevels, isLoading, error } = useListFormationLevelsQuery();

  // Determine the current value to display
  const currentValue = value || defaultValue;
  
  // Find the selected level for display
  const selectedLevel = formationLevels?.find(level => level.id.toString() === currentValue);

  if (isLoading) {
    return (
      <Select disabled value={currentValue}>
        <SelectTrigger id={id}>
          <SelectValue placeholder="Cargando..." />
        </SelectTrigger>
      </Select>
    );
  }

  if (error) {
    return (
      <Select disabled value={currentValue}>
        <SelectTrigger id={id}>
          <SelectValue placeholder="Error al cargar" />
        </SelectTrigger>
      </Select>
    );
  }

  return (
    <Select 
      value={currentValue} 
      onValueChange={onValueChange}
    >
      <SelectTrigger id={id}>
        <SelectValue placeholder={placeholder}>
          {selectedLevel ? selectedLevel.name : placeholder}
        </SelectValue>
      </SelectTrigger>
      <SelectContent>
        {formationLevels?.map((level) => (
          <SelectItem key={level.id} value={level.id.toString()}>
            {level.name}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
}
