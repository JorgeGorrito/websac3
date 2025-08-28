import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";

type ProgramDurationSelectProps = {
  id?: string;
  placeholder?: string;
  defaultValue?: string;
  onValueChange?: (value: string) => void;
};

export function ProgramDurationSelect({ id, placeholder = "Seleccionar", defaultValue, onValueChange }: ProgramDurationSelectProps) {
  return (
    <Select defaultValue={defaultValue} onValueChange={onValueChange}>
      <SelectTrigger id={id} className="h-10">
        <SelectValue placeholder={placeholder} />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="6">6 semestres</SelectItem>
        <SelectItem value="8">8 semestres</SelectItem>
        <SelectItem value="10">10 semestres</SelectItem>
        <SelectItem value="otros">Otro</SelectItem>
      </SelectContent>
    </Select>
  );
}


