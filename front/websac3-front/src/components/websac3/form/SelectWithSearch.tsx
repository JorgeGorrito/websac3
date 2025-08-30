"use client";

import { useMemo, useState } from "react";
import { Check, ChevronDown } from "lucide-react";
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
  items: { value: string; label: string }[];
  value?: string;
  onChange?: (value: string) => void;
}

export default function SelectWithSearch({
  placeHolderDefault,
  placeHolderSearch,
  placeHolderNoResults,
  items,
  value: controlledValue,
  onChange,
}: SelectWithSearchProps) {
  const [open, setOpen] = useState(false);
  const [uncontrolledValue, setUncontrolledValue] = useState("");
  const value = controlledValue ?? uncontrolledValue;
  const setValue = onChange ?? setUncontrolledValue;

  const labelByValue = useMemo(() => new Map(items.map((i) => [i.value, i.label])), [items]);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full h-11 justify-between font-normal border-gray-200 focus:border-gray-400 focus:ring-0 transition-colors bg-white hover:bg-gray-50"
        >
          <span className={cn("truncate", !value && "text-gray-500")}>
            {value ? labelByValue.get(value) : placeHolderDefault}
          </span>
          <ChevronDown className="ml-2 h-4 w-4 shrink-0 text-gray-400" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0 border-gray-200" align="start">
        <Command>
          <div className="flex items-center border-b border-gray-100 px-3">
            <CommandInput
              placeholder={placeHolderSearch}
              className="flex h-10 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-gray-500 disabled:cursor-not-allowed disabled:opacity-50"
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
                  onSelect={(currentValue) => {
                    setValue(currentValue === value ? "" : currentValue);
                    setOpen(false);
                  }}
                  className="cursor-pointer py-2 px-3 hover:bg-gray-50"
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4",
                      value === institution.value
                        ? "opacity-100 text-gray-900"
                        : "opacity-0"
                    )}
                  />
                  <span className="text-sm">{institution.label}</span>
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
