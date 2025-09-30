"use client";

import React from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { 
  User, 
  Mail, 
  Building, 
  IdCard, 
  Briefcase, 
  Shield, 
  CheckCircle, 
  XCircle
} from "lucide-react";
import { ProfileData } from "@/services/api";
import { translateRole } from "@/utils/roleTranslations";

interface ProfileViewProps {
  profile: ProfileData | undefined;
  isLoading?: boolean;
}

export function ProfileView({ profile, isLoading = false }: ProfileViewProps) {
  if (isLoading) {
    return <ProfileLoadingSkeleton />;
  }

  if (!profile) {
    return (
      <Card>
        <CardContent className="p-12">
          <div className="text-center">
            <User className="h-12 w-12 text-gray-400 mx-auto mb-4" />
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              No se pudo cargar el perfil
            </h3>
            <p className="text-gray-600">
              No se pudo obtener la información del perfil del usuario.
            </p>
          </div>
        </CardContent>
      </Card>
    );
  }

  const fullName = `${profile.name} ${profile.lastname}`.trim();

  return (
    <div className="space-y-6">
      {/* Header Card */}
      <Card className="bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200">
        <CardHeader>
          <div className="flex items-center gap-4">
            <div className="p-3 bg-blue-100 rounded-full">
              <User className="h-8 w-8 text-blue-600" />
            </div>
            <div>
              <CardTitle className="text-2xl font-bold text-gray-900">
                {fullName}
              </CardTitle>
              <CardDescription className="text-lg text-gray-600 mt-1">
                {profile.job_position}
              </CardDescription>
            </div>
            <div className="ml-auto">
              <Badge 
                className={`px-3 py-1 ${
                  profile.is_active 
                    ? 'bg-green-100 text-green-800 border-green-200' 
                    : 'bg-red-100 text-red-800 border-red-200'
                }`}
              >
                {profile.is_active ? (
                  <><CheckCircle className="h-4 w-4 mr-1" /> Activo</>
                ) : (
                  <><XCircle className="h-4 w-4 mr-1" /> Inactivo</>
                )}
              </Badge>
            </div>
          </div>
        </CardHeader>
      </Card>

      {/* Personal Information */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <User className="h-5 w-5 text-blue-600" />
            Información Personal
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="space-y-2">
              <div className="flex items-center gap-2">
                <Mail className="h-4 w-4 text-gray-500" />
                <span className="text-sm font-medium text-gray-600">Correo electrónico</span>
              </div>
              <p className="text-lg font-semibold text-gray-900">{profile.email}</p>
            </div>

            <div className="space-y-2">
              <div className="flex items-center gap-2">
                <IdCard className="h-4 w-4 text-gray-500" />
                <span className="text-sm font-medium text-gray-600">Documento de identidad</span>
              </div>
              <div className="flex items-center gap-2">
                <p className="text-lg font-semibold text-gray-900">{profile.identification_number}</p>
                <Badge variant="secondary" className="text-xs">
                  {profile.identification_type_name}
                </Badge>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Professional Information */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Briefcase className="h-5 w-5 text-green-600" />
            Información Profesional
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="space-y-2">
              <div className="flex items-center gap-2">
                <Shield className="h-4 w-4 text-gray-500" />
                <span className="text-sm font-medium text-gray-600">Rol en el sistema</span>
              </div>
              <Badge className="bg-blue-100 text-blue-800 border-blue-200 px-3 py-1">
                {translateRole(profile.role_name)}
              </Badge>
            </div>

            <div className="space-y-2">
              <div className="flex items-center gap-2">
                <Briefcase className="h-4 w-4 text-gray-500" />
                <span className="text-sm font-medium text-gray-600">Cargo</span>
              </div>
              <p className="text-lg font-semibold text-gray-900">{profile.job_position}</p>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Institution Information */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Building className="h-5 w-5 text-purple-600" />
            Institución Educativa
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="bg-white rounded-lg p-4 border border-gray-200">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div className="space-y-2">
                <div className="flex items-center gap-2">
                  <Building className="h-4 w-4 text-gray-500" />
                  <span className="text-sm font-medium text-gray-600">Nombre de la institución</span>
                </div>
                <p className="text-lg font-semibold text-gray-900">{profile.institution_name}</p>
              </div>

              <div className="space-y-2">
                <div className="flex items-center gap-2">
                  <Building className="h-4 w-4 text-gray-500" />
                  <span className="text-sm font-medium text-gray-600">Código SNIES</span>
                </div>
                <p className="text-lg font-semibold text-gray-900">{profile.institution_snies}</p>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

    </div>
  );
}

// Loading skeleton component
function ProfileLoadingSkeleton() {
  return (
    <div className="space-y-6">
      {/* Header Card Skeleton */}
      <Card className="bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200">
        <CardHeader>
          <div className="flex items-center gap-4">
            <Skeleton className="h-14 w-14 rounded-full" />
            <div className="flex-1">
              <Skeleton className="h-8 w-64 mb-2" />
              <Skeleton className="h-6 w-48" />
            </div>
            <Skeleton className="h-8 w-20" />
          </div>
        </CardHeader>
      </Card>

      {/* Cards Skeleton */}
      {Array.from({ length: 3 }).map((_, i) => (
        <Card key={i}>
          <CardHeader>
            <div className="flex items-center gap-2">
              <Skeleton className="h-5 w-5" />
              <Skeleton className="h-6 w-48" />
            </div>
          </CardHeader>
          <CardContent className="space-y-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div className="space-y-2">
                <Skeleton className="h-4 w-32" />
                <Skeleton className="h-6 w-48" />
              </div>
              <div className="space-y-2">
                <Skeleton className="h-4 w-32" />
                <Skeleton className="h-6 w-48" />
              </div>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
