"use client";

import React, { useState, useMemo } from "react";
import { Check, ChevronsUpDown, X, Search } from "lucide-react";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Badge } from "@/components/ui/badge";
import { useListProfessionalRolesQuery } from "@/services/api";
import { Skeleton } from "@/components/ui/skeleton";

type ProgramProfessionalRolesSelectProps = {
  id: string;
  placeholder?: string;
  defaultValue?: number[];
  value?: number[];
  onValueChange: (value: number[]) => void;
};

export function ProgramProfessionalRolesSelect({
  id,
  placeholder = "Seleccionar roles profesionales",
  defaultValue = [],
  value = [],
  onValueChange,
}: ProgramProfessionalRolesSelectProps) {
  const [open, setOpen] = useState(false);
  const [searchValue, setSearchValue] = useState("");

  // Usar filtros dinámicos para la búsqueda
  const { data: professionalRolesResponse, isLoading, error } = useListProfessionalRolesQuery({
    current_page: 1,
    items_per_page: 50, // Obtener más roles para la selección
    filters: searchValue ? { 'name[cont]': searchValue } : undefined
  });

  // Extraer los roles de la respuesta paginada
  const professionalRoles = professionalRolesResponse?.data || [];

  // Los roles ya vienen filtrados del servidor
  const filteredRoles = professionalRoles;

  // Get selected roles for display
  const selectedRoles = useMemo(() => {
    if (!professionalRoles || !Array.isArray(professionalRoles)) return [];
    return professionalRoles.filter((role) => value.includes(role.id));
  }, [professionalRoles, value]);

  const handleSelect = (roleId: number) => {
    const newValue = value.includes(roleId)
      ? value.filter((id) => id !== roleId)
      : [...value, roleId];
    onValueChange(newValue);
  };

  const handleRemove = (roleId: number) => {
    const newValue = value.filter((id) => id !== roleId);
    onValueChange(newValue);
  };

  if (isLoading) {
    return <Skeleton className="h-11 w-full" />;
  }

  if (error) {
    return (
      <div className="h-11 w-full border border-red-300 rounded-md flex items-center px-3 text-red-600 text-sm">
        Error al cargar roles profesionales
      </div>
    );
  }

  return (
    <div className="space-y-2">
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            className="w-full justify-between h-11"
          >
            <span className="truncate">
              {selectedRoles.length === 0
                ? placeholder
                : `${selectedRoles.length} rol${selectedRoles.length !== 1 ? 'es' : ''} seleccionado${selectedRoles.length !== 1 ? 's' : ''}`}
            </span>
            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-full p-0" align="start">
          <Command>
            <div className="flex items-center border-b px-3">
              <Search className="mr-2 h-4 w-4 shrink-0 opacity-50" />
              <input
                className="flex h-10 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50"
                placeholder="Buscar roles profesionales..."
                value={searchValue}
                onChange={(e) => setSearchValue(e.target.value)}
              />
            </div>
            <CommandList>
              <CommandEmpty>No se encontraron roles profesionales.</CommandEmpty>
              <CommandGroup>
                {filteredRoles.map((role) => (
                  <CommandItem
                    key={role.id}
                    value={role.name}
                    onSelect={() => handleSelect(role.id)}
                    className="flex items-center justify-between"
                  >
                    <div className="flex items-center">
                      <Check
                        className={cn(
                          "mr-2 h-4 w-4",
                          value.includes(role.id) ? "opacity-100" : "opacity-0"
                        )}
                      />
                      <span>{role.name}</span>
                    </div>
                  </CommandItem>
                ))}
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>

      {/* Selected roles display */}
      {selectedRoles.length > 0 && (
        <div className="flex flex-wrap gap-2">
          {selectedRoles.map((role) => (
            <Badge
              key={role.id}
              variant="secondary"
              className="flex items-center gap-1 px-2 py-1"
            >
              <span className="text-xs">{role.name}</span>
              <button
                type="button"
                onClick={() => handleRemove(role.id)}
                className="ml-1 hover:bg-gray-300 rounded-full p-0.5"
              >
                <X className="h-3 w-3" />
              </button>
            </Badge>
          ))}
        </div>
      )}
    </div>
  );
}
