import SelectWithSearch from "@/components/websac3/form/SelectWithSearch";

type ProgramNameAutocompleteProps = {
  label?: string; // reserved for future form libs
};

export function ProgramNameAutocomplete(_props: ProgramNameAutocompleteProps) {
  return (
    <SelectWithSearch
      placeHolderDefault="Seleccionar programa"
      placeHolderSearch="Buscar..."
      placeHolderNoResults="Sin resultados"
    />
  );
}


