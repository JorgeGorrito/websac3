import React from "react"
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
} from "@/components/ui/command"

import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/popover"
import { Check, ChevronsUpDown } from "lucide-react"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import ItemSelectWithSearch from "@/types/websac3/Input/ItemSelectWithSearch"
import SelectWithSearchProps from "@/types/websac3/Input/SelectWithSearch"

const SelectWithSearch = ({
    placeHolderDefault = "Select An Option...",
    placeHolderSearch = "Search...",
    placeHolderNoResults = "No results found.",
} : SelectWithSearchProps) => {
    const [isOpen, setIsOpen] = React.useState(false)
    const [value, setValue] = React.useState("")
    const [data, setData] = React.useState<ItemSelectWithSearch[]>([])


    return (
       <Popover open={isOpen} onOpenChange={setIsOpen}>
            <PopoverTrigger asChild>
                <Button
                variant="outline"
                role="combobox"
                aria-expanded={isOpen}
                className="w-full justify-between px-4 py-2"
                >
                {value ? data.find((item) => item.value === value)?.label : placeHolderDefault}
                <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent 
                className="w-[var(--radix-popover-trigger-width)] p-0"
                align="start"
                sideOffset={4}
            >
                <Command>
                <CommandInput placeholder={placeHolderSearch} className="h-9" />
                <CommandList>
                    <CommandEmpty>{placeHolderNoResults}</CommandEmpty>
                    <CommandGroup>
                    {data.map((item) => (
                        <CommandItem
                        key={item.value}
                        value={item.value}
                        onSelect={(currentValue) => {
                            setValue(currentValue === value ? "" : currentValue);
                            setIsOpen(false);
                        }}
                        >
                        {item.label}
                        <Check className={cn("ml-auto h-4 w-4", value === item.value ? "opacity-100" : "opacity-0")} />
                        </CommandItem>
                    ))}
                    </CommandGroup>
                </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    )
}

export default SelectWithSearch