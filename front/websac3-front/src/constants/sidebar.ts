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
      { title: "Inicio", url: "/experto/dashboard", icon: Home },
      {
        title: "Reportes Retroalimentados",
        url: "/experto/reportes/retroalimentados",
        icon: FolderCheck,
      },
      {
        title: "Reportes Pendientes",
        url: "/experto/reportes/pendientes",
        icon: AlertTriangle,
      },
      {
        title: "Solicitudes De Asesoria",
        url: "/experto/asesoria/solicitudes",
        icon: Headphones,
      },
      { title: "Perfil", url: "/experto/perfil", icon: UserCog },
    ],
    navSecondary: [],
  },
  ADMIN: {
    user: defaultUser,
    roleLabel: "Administrador",
    navMain: [
      { title: "Inicio", url: "/admin/dashboard", icon: Home },
      {
        title: "Gestionar Usuarios",
        url: "/admin/usuarios",
        icon: Users,
      },
      {
        title: "Gestionar Solicitudes De Acceso",
        url: "/admin/accesos",
        icon: Inbox,
      },
    ],
    navSecondary: [],
  },
  DIRECTOR: {
    user: defaultUser,
    roleLabel: "Director de programa",
    navMain: [
      { title: "Inicio", url: "/director/dashboard", icon: Home },
      {
        title: "Programas de Grado",
        url: "/director/programas",
        icon: PlusSquare,
      },
      {
        title: "Cargar Datos",
        url: "/director/cargar-datos",
        icon: Upload,
      },
      {
        title: "Asesoria Con Experto",
        url: "/director/asesoria-experto",
        icon: Headphones,
      },
      {
        title: "Consultar Reportes",
        url: "/director/reportes",
        icon: PieChart,
      },
      {
        title: "Evaluar Componente de Ciberseguridad",
        url: "/director/operar-modelo",
        icon: Cog,
      },
      { title: "Perfil", url: "/director/perfil", icon: UserCog },
    ],
    navSecondary: [],
  },
};
