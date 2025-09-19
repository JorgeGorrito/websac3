import SelectWithSearch from "@/components/websac3/form/SelectWithSearch";

type ProgramNameAutocompleteProps = {
  value?: string;
  onChange?: (value: string) => void;
};

export function ProgramNameAutocomplete({ value, onChange }: ProgramNameAutocompleteProps) {
  return (
    <SelectWithSearch
      placeHolderDefault="Seleccionar programa"
      placeHolderSearch="Buscar..."
      placeHolderNoResults="Sin resultados"
      items={[]}
      value={value}
      onChange={onChange}
    />
  );
}


