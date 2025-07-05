import { Input } from "@/components/ui/input"
import {Label} from "@/components/ui/label"
import { InputContainer } from "../InputContainer"
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
  } from "@/components/ui/select"
import SelectWithSearch from "../SelectWithSearch"


const AccessRequestForm = () => {
    return (
        <form action="#">
            <div className="relative flex flex-wrap">
                <InputContainer>
                    <Label htmlFor="name"> Nombre(s) <span className="text-red-600 text-lg">*</span></Label>
                    <Input id="name"/>
                </InputContainer>
                <InputContainer>
                    <Label htmlFor="lastname"> Apellido(s) <span className="text-red-600 text-lg">*</span></Label>
                    <Input id="lastname"/>
                </InputContainer>
                <InputContainer>
                    <Label htmlFor="documentType"> Tipo de Documento <span className="text-red-600 text-lg">*</span></Label>
                    <Select>
                    <SelectTrigger id="documentType">
                        <SelectValue placeholder="Tipos de documento" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectGroup>
                        <SelectItem value="cc">Cedula de ciudadania</SelectItem>
                        <SelectItem value="ce">Cedula extranjera</SelectItem>
                        <SelectItem value="ci">Codigo de la institucion</SelectItem>
                        </SelectGroup>
                    </SelectContent>
                    </Select>
                </InputContainer>
                <InputContainer>
                    <Label htmlFor="name"> Nº de Identificación <span className="text-red-600 text-lg">*</span></Label>
                    <Input id="name"/>
                </InputContainer>
                <InputContainer>
                    <Label htmlFor="name"> Cargo <span className="text-red-600 text-lg">*</span></Label>
                    <Input id="name"/>
                </InputContainer>
                <InputContainer>
                    <Label htmlFor="name"> Email <span className="text-red-600 text-lg">*</span></Label>
                    <Input id="name"/>
                </InputContainer>
                <InputContainer>
                    <Label htmlFor="name"> Institución de Educación Superior<span className="text-red-600 text-lg">*</span></Label>
                    <SelectWithSearch 
                        placeHolderDefault="Seleccione una Institución"
                        placeHolderSearch="Buscar..."
                        placeHolderNoResults="No hay resultados."
                    />
                </InputContainer>
            </div>
        </form>
    );
}

export { AccessRequestForm };