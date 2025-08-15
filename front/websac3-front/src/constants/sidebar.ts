import {
  Home,
  FolderCheck,
  AlertTriangle,
  Headphones,
  Settings,
  UserCog,
  Users,
  FileText,
  Inbox,
  PlusSquare,
  Upload,
  PieChart,
  Cog,
  LifeBuoy,
  Send,
} from "lucide-react";

import type { LucideIcon } from "lucide-react";

export type RoleKey = "ADMIN" | "EXPERTO" | "DIRECTOR";

export type NavItem = {
  title: string;
  url: string;
  icon: LucideIcon;
  isActive?: boolean;
  items?: { title: string; url: string }[];
};

export type SidebarData = {
  user: { name: string; email: string; avatar: string };
  roleLabel: string;
  navMain: NavItem[];
  navSecondary?: { title: string; url: string; icon: LucideIcon }[];
};

const defaultUser = {
  name: "Pepito Perez",
  email: "pepito.perez@example.com",
  avatar: "/dashboard/menu/user-circle-solid.png",
};

export const SIDEBAR_BY_ROLE: Record<RoleKey, SidebarData> = {
  EXPERTO: {
    user: defaultUser,
    roleLabel: "Experto en Ciberseguridad",
    navMain: [
      { title: "Inicio", url: "/dashboard", icon: Home, isActive: true },
      {
        title: "Reportes Retroalimentados",
        url: "/dashboard/reportes/retroalimentados",
        icon: FolderCheck,
      },
      {
        title: "Reportes Pendientes",
        url: "/dashboard/reportes/pendientes",
        icon: AlertTriangle,
      },
      {
        title: "Solicitudes De Asesoria",
        url: "/dashboard/asesoria/solicitudes",
        icon: Headphones,
      },
      { title: "Gestionar Modelo", url: "/dashboard/modelo", icon: Settings },
      { title: "Perfil", url: "/dashboard/perfil", icon: UserCog },
    ],
    navSecondary: [
      { title: "Soporte", url: "#", icon: LifeBuoy },
      { title: "Feedback", url: "#", icon: Send },
    ],
  },
  ADMIN: {
    user: defaultUser,
    roleLabel: "Administrador",
    navMain: [
      { title: "Inicio", url: "/dashboard", icon: Home, isActive: true },
      {
        title: "Gestionar Usuarios",
        url: "/dashboard/admin/usuarios",
        icon: Users,
      },
      {
        title: "Gestionar Reportes",
        url: "/dashboard/admin/reportes",
        icon: FileText,
      },
      {
        title: "Gestionar Solicitudes De Asesoria",
        url: "/dashboard/admin/asesoria",
        icon: Headphones,
      },
      {
        title: "Gestionar Solicitudes De Acceso",
        url: "/dashboard/admin/accesos",
        icon: Inbox,
      },
    ],
    navSecondary: [
      { title: "Soporte", url: "#", icon: LifeBuoy },
      { title: "Feedback", url: "#", icon: Send },
    ],
  },
  DIRECTOR: {
    user: defaultUser,
    roleLabel: "Director de programa",
    navMain: [
      { title: "Inicio", url: "/dashboard", icon: Home, isActive: true },
      {
        title: "Registrar Programa",
        url: "/dashboard/director/registrar-programa",
        icon: PlusSquare,
      },
      {
        title: "Cargar Datos",
        url: "/dashboard/director/cargar-datos",
        icon: Upload,
      },
      {
        title: "Asesoria Con Experto",
        url: "/dashboard/director/asesoria-experto",
        icon: Headphones,
      },
      {
        title: "Consultar Reportes",
        url: "/dashboard/director/reportes",
        icon: PieChart,
      },
      {
        title: "Operar Modelo",
        url: "/dashboard/director/operar-modelo",
        icon: Cog,
      },
      { title: "Perfil", url: "/dashboard/perfil", icon: UserCog },
    ],
    navSecondary: [
      { title: "Soporte", url: "#", icon: LifeBuoy },
      { title: "Feedback", url: "#", icon: Send },
    ],
  },
};
