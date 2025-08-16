import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";

type ProgramPeriodicitySelectProps = {
  id?: string;
  placeholder?: string;
  defaultValue?: string;
  onValueChange?: (value: string) => void;
};

export function ProgramPeriodicitySelect({ id, placeholder = "Seleccionar", defaultValue, onValueChange }: ProgramPeriodicitySelectProps) {
  return (
    <Select defaultValue={defaultValue} onValueChange={onValueChange}>
      <SelectTrigger id={id} className="h-10">
        <SelectValue placeholder={placeholder} />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="semestral">Semestral</SelectItem>
        <SelectItem value="anual">Anual</SelectItem>
        <SelectItem value="trimestral">Trimestral</SelectItem>
      </SelectContent>
    </Select>
  );
}


