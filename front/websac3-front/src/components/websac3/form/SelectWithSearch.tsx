"use client";

import { useMemo, useState } from "react";
import { Check, ChevronDown, MapPin, Building2 } from "lucide-react";
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

interface SelectWithSearchProps {
  placeHolderDefault: string;
  placeHolderSearch: string;
  placeHolderNoResults: string;
  items: { value: string; label: string; subtitle?: string }[];
  value?: string;
  onChange?: (value: string) => void;
  onSearch?: (query: string) => void;
}

export default function SelectWithSearch({
  placeHolderDefault,
  placeHolderSearch,
  placeHolderNoResults,
  items,
  value: controlledValue,
  onChange,
  onSearch,
}: SelectWithSearchProps) {
  const [open, setOpen] = useState(false);
  const [uncontrolledValue, setUncontrolledValue] = useState("");
  const value = controlledValue ?? uncontrolledValue;
  const setValue = onChange ?? setUncontrolledValue;

  const labelByValue = useMemo(
    () => new Map(items.map((i) => [i.value, i.label])),
    [items]
  );

  const selectedItem = useMemo(
    () => items.find((item) => item.value === value),
    [items, value]
  );

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full h-12 justify-between border-gray-300 hover:border-gray-400 focus:border-blue-500 focus:ring-2 focus:ring-blue-100 transition-all"
        >
          <div className="flex items-center text-left min-w-0 flex-1">
            {value && selectedItem ? (
              <div className="flex items-center gap-3 min-w-0 flex-1">
                <div className="flex-shrink-0 p-1.5 bg-blue-50 rounded-lg">
                  <Building2 className="h-4 w-4 text-blue-600" />
                </div>
                <div className="flex flex-col min-w-0 flex-1">
                  <span className="text-sm font-semibold text-gray-900 truncate">
                    {selectedItem.label}
                  </span>
                  {selectedItem.subtitle && (
                    <div className="flex items-center gap-1 mt-0.5">
                      <MapPin className="h-3 w-3 text-gray-400" />
                      <span className="text-xs text-gray-600 truncate">
                        {selectedItem.subtitle}
                      </span>
                    </div>
                  )}
                </div>
              </div>
            ) : (
              <div className="flex items-center gap-3">
                <div className="flex-shrink-0 p-1.5 bg-gray-100 rounded-lg">
                  <Building2 className="h-4 w-4 text-gray-400" />
                </div>
                <span className="text-sm text-gray-500">
                  {placeHolderDefault}
                </span>
              </div>
            )}
          </div>
          <ChevronDown className="ml-2 h-4 w-4 shrink-0 text-gray-400" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0 border-gray-200 shadow-lg" align="start">
        <Command shouldFilter={false}>
          <div className="flex items-center border-b border-gray-100 px-4 py-2">
            <CommandInput
              placeholder={placeHolderSearch}
              className="flex h-10 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-gray-500 disabled:cursor-not-allowed disabled:opacity-50"
              onValueChange={(q) => onSearch?.(q)}
            />
          </div>
          <CommandList className="max-h-60">
            <CommandEmpty className="py-6 text-center text-sm text-gray-500">
              {placeHolderNoResults}
            </CommandEmpty>
            <CommandGroup>
              {items.map((institution) => (
                <CommandItem
                  key={institution.value}
                  value={institution.value}
                  keywords={[institution.label, institution.subtitle || ""]}
                  onSelect={(currentValue) => {
                    setValue(currentValue === value ? "" : currentValue);
                    setOpen(false);
                  }}
                  className="cursor-pointer py-3 px-4 hover:bg-blue-50 transition-colors"
                >
                  <div className="flex items-center gap-3 min-w-0 flex-1">
                    <div className={cn(
                      "flex-shrink-0 p-1.5 rounded-lg",
                      value === institution.value 
                        ? "bg-blue-100" 
                        : "bg-gray-100"
                    )}>
                      <Building2 className={cn(
                        "h-4 w-4",
                        value === institution.value 
                          ? "text-blue-600" 
                          : "text-gray-600"
                      )} />
                    </div>
                    <div className="flex flex-col min-w-0 flex-1">
                      <span className={cn(
                        "text-sm font-semibold truncate",
                        value === institution.value 
                          ? "text-blue-900" 
                          : "text-gray-900"
                      )}>
                        {institution.label}
                      </span>
                      {institution.subtitle && (
                        <div className="flex items-center gap-1 mt-1">
                          <MapPin className={cn(
                            "h-3 w-3",
                            value === institution.value 
                              ? "text-blue-500" 
                              : "text-gray-400"
                          )} />
                          <span className={cn(
                            "text-xs truncate",
                            value === institution.value 
                              ? "text-blue-700" 
                              : "text-gray-600"
                          )}>
                            {institution.subtitle}
                          </span>
                        </div>
                      )}
                    </div>
                    <Check
                      className={cn(
                        "h-4 w-4 flex-shrink-0",
                        value === institution.value
                          ? "opacity-100 text-blue-600"
                          : "opacity-0"
                      )}
                    />
                  </div>
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
