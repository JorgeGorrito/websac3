import ItemSelectWithSearch from "./ItemSelectWithSearch"

interface SelectWithSearchProps {
    placeHolderDefault: string
    placeHolderSearch: string 
    placeHolderNoResults: string
    onSearch: (query: string) => Promise<ItemSelectWithSearch[]>
}

export default SelectWithSearchProps