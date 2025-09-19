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
  const [isSuccess, setIsSuccess] = useState(false)

  const [createAccessRequest, { isLoading }] = useCreateAccessRequestMutation()
  const { data: identificationTypes } = useListIdentificationTypesQuery()
  const [institutionQuery, setInstitutionQuery] = useState("")
  const { data: institutions } = useListHigherEducationInstitutionsQuery(
    institutionQuery ? { nameCont: institutionQuery, items_per_page: 20 } : undefined
  )

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setMessage("")
    setError("")
    setIsSuccess(false)

    const identification_type_id = Number(docType)
    const sniesNumber = Number(snies)
    if (!identification_type_id || !sniesNumber) {
      setError("Por favor completa tipo de documento e institución válidos.")
      return
    }

    const origin = typeof window !== "undefined" ? window.location.origin : "http://localhost:3000"
    try {
      const response = await createAccessRequest({
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
      
      // Handle success response
      if (response.result) {
        setMessage(response.result)
        setIsSuccess(true)
        // Clear form on success
        setName("")
        setLastname("")
        setDocType("")
        setIdentification("")
        setPosition("")
        setEmail("")
        setSnies("")
      }
    } catch (error: any) {
      // Handle error response
      if (error?.data?.errors && Array.isArray(error.data.errors)) {
        // Backend sends errors as array of strings
        setError(error.data.errors.join(", "))
      } else if (error?.data?.result) {
        // Sometimes errors come in result field
        setError(error.data.result)
      } else {
        setError("No se pudo enviar la solicitud. Intenta de nuevo.")
      }
    }
  }

  return (
    <form onSubmit={onSubmit} className="space-y-8 px-1">
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
          onSearch={setInstitutionQuery}
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

      {/* Success Message */}
      {message && isSuccess && (
        <div className="bg-green-50 border border-green-200 rounded-lg p-4">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <svg className="h-5 w-5 text-green-400" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
              </svg>
            </div>
            <div className="ml-3">
              <p className="text-sm font-medium text-green-800">{message}</p>
            </div>
          </div>
        </div>
      )}

      {/* Error Message */}
      {error && (
        <div className="bg-red-50 border border-red-200 rounded-lg p-4">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <svg className="h-5 w-5 text-red-400" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
              </svg>
            </div>
            <div className="ml-3">
              <p className="text-sm font-medium text-red-800">{error}</p>
            </div>
          </div>
        </div>
      )}
    </form>
  )
}

export { AccessRequestForm }
