"use client"

import { useState } from "react"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import SelectWithSearch from "../SelectWithSearch"
import { useCreateAccessRequestMutation, useListHigherEducationInstitutionsQuery, useListIdentificationTypesQuery } from "@/services/api"

const AccessRequestForm = () => {
  const [name, setName] = useState("")
  const [lastname, setLastname] = useState("")
  const [docType, setDocType] = useState<string>("")
  const [identification, setIdentification] = useState("")
  const [position, setPosition] = useState("")
  const [email, setEmail] = useState("")
  const [snies, setSnies] = useState<string>("")
  const [message, setMessage] = useState<string>("")
  const [error, setError] = useState<string>("")

  const [createAccessRequest, { isLoading }] = useCreateAccessRequestMutation()
  const { data: identificationTypes } = useListIdentificationTypesQuery()
  const { data: institutions } = useListHigherEducationInstitutionsQuery()

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setMessage("")
    setError("")

    const identification_type_id = Number(docType)
    const sniesNumber = Number(snies)
    if (!identification_type_id || !sniesNumber) {
      setError("Por favor completa tipo de documento y SNIES válidos.")
      return
    }

    const origin = typeof window !== "undefined" ? window.location.origin : "http://localhost:3000"
    try {
      await createAccessRequest({
        person: {
          email,
          higher_education_institution_snies: sniesNumber,
          identification_number: identification,
          identification_type_id,
          job_position: position,
          lastname,
          name,
        },
        register_user_url: `${origin}/register-user/token=`,
        validation_email_url: `${origin}/access-request/validate-email?token=`,
      }).unwrap()
      setMessage("Solicitud enviada. Revisa tu correo para validar.")
      setName("")
      setLastname("")
      setDocType("")
      setIdentification("")
      setPosition("")
      setEmail("")
      setSnies("")
    } catch (err) {
      setError("No se pudo enviar la solicitud. Intenta de nuevo.")
    }
  }

  return (
    <form onSubmit={onSubmit} className="space-y-8">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="space-y-2">
          <Label htmlFor="name" className="text-sm font-medium text-gray-700">
            Nombre(s) <span className="text-red-500">*</span>
          </Label>
          <Input
            id="name"
            className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
            placeholder="Tu nombre"
            value={name}
            onChange={(e) => setName(e.target.value)}
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
            value={lastname}
            onChange={(e) => setLastname(e.target.value)}
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="documentType" className="text-sm font-medium text-gray-700">
            Tipo de Documento <span className="text-red-500">*</span>
          </Label>
          <Select value={docType} onValueChange={setDocType}>
            <SelectTrigger
              id="documentType"
              className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
            >
              <SelectValue placeholder="Seleccionar" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                {identificationTypes?.map((t) => (
                  <SelectItem key={t.id} value={String(t.id)}>{t.name}</SelectItem>
                ))}
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
            value={identification}
            onChange={(e) => setIdentification(e.target.value)}
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
            value={position}
            onChange={(e) => setPosition(e.target.value)}
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
            value={email}
            onChange={(e) => setEmail(e.target.value)}
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
          items={(institutions ?? []).map((i) => ({ value: String(i.snies), label: i.name }))}
          value={snies}
          onChange={setSnies}
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="snies" className="text-sm font-medium text-gray-700">
          SNIES de la institución <span className="text-red-500">*</span>
        </Label>
        <Input
          id="snies"
          className="h-10 border-gray-200 focus:border-gray-400 focus:ring-1 focus:ring-gray-400 transition-all"
          placeholder="Ej: 1119"
          value={snies}
          onChange={(e) => setSnies(e.target.value)}
        />
      </div>

      <div className="pt-6">
        <Button
          type="submit"
          className="w-full h-11 bg-gray-900 hover:bg-gray-800 text-white font-medium transition-colors rounded-lg"
          disabled={isLoading}
        >
          {isLoading ? "Enviando..." : "Enviar Solicitud"}
        </Button>
      </div>

      {message && <p className="text-green-600 text-sm">{message}</p>}
      {error && <p className="text-red-600 text-sm">{error}</p>}
    </form>
  )
}

export { AccessRequestForm }
