import SelectWithSearch from "@/components/websac3/form/SelectWithSearch";

export function ProgramNameAutocomplete() {
  return (
    <SelectWithSearch
      placeHolderDefault="Seleccionar programa"
      placeHolderSearch="Buscar..."
      placeHolderNoResults="Sin resultados"
      items={[]}
    />
  );
}


