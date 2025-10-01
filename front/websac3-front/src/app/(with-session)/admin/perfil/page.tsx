"use client";

import React from "react";
import { useGetProfileQuery } from "@/services/api";
import { ProfileView } from "@/components/websac3/profile/ProfileView";

export default function PerfilPage() {
  const { data: profile, isLoading, error } = useGetProfileQuery();

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Mi Perfil</h1>
          <p className="text-gray-600 mt-2">Información personal y profesional</p>
        </div>
      </div>

      {/* Profile Content */}
      <ProfileView profile={profile} isLoading={isLoading} />
    </div>
  );
}

