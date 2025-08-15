import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import SelectWithSearch from "../SelectWithSearch"

const AccessRequestForm = () => {
  return (
    <form action="#" className="space-y-8">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="space-y-2">
          <Label htmlFor="name" className="text-sm font-medium text-gray-700">
            Nombre(s) <span className="text-red-500">*</span>
          </Label>
          <Input
            id="name"
            className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
            placeholder="Tu nombre"
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="lastname" className="text-sm font-medium text-gray-700">
            Apellido(s) <span className="text-red-500">*</span>
          </Label>
          <Input
            id="lastname"
            className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
            placeholder="Tus apellidos"
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="documentType" className="text-sm font-medium text-gray-700">
            Tipo de Documento <span className="text-red-500">*</span>
          </Label>
          <Select>
            <SelectTrigger
              id="documentType"
              className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
            >
              <SelectValue placeholder="Seleccionar" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem value="cc">Cédula de ciudadanía</SelectItem>
                <SelectItem value="ce">Cédula extranjera</SelectItem>
                <SelectItem value="ci">Código de la institución</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>

        <div className="space-y-2">
          <Label htmlFor="identification" className="text-sm font-medium text-gray-700">
            Número de Identificación <span className="text-red-500">*</span>
          </Label>
          <Input
            id="identification"
            className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
            placeholder="123456789"
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="position" className="text-sm font-medium text-gray-700">
            Cargo <span className="text-red-500">*</span>
          </Label>
          <Input
            id="position"
            className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
            placeholder="Tu cargo"
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="email" className="text-sm font-medium text-gray-700">
            Email <span className="text-red-500">*</span>
          </Label>
          <Input
            id="email"
            type="email"
            className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
            placeholder="tu@email.com"
          />
        </div>
      </div>

      <div className="space-y-2">
        <Label htmlFor="institution" className="text-sm font-medium text-gray-700">
          Institución de Educación Superior <span className="text-red-500">*</span>
        </Label>
        <SelectWithSearch
          placeHolderDefault="Seleccionar institución"
          placeHolderSearch="Buscar..."
          placeHolderNoResults="Sin resultados"
        />
      </div>

      <div className="pt-6">
        <Button
          type="submit"
          className="w-full h-11 bg-gray-900 hover:bg-gray-800 text-white font-medium transition-colors rounded-lg"
        >
          Enviar Solicitud
        </Button>
      </div>
    </form>
  )
}

export { AccessRequestForm }
