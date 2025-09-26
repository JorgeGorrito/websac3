"use client"

import {
  User,
  Settings,
  ChevronsUpDown,
  LogOut,
  Shield,
} from "lucide-react"

import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from "@/components/ui/sidebar"
import { useAuth } from "@/hooks/useAuth"
import { Badge } from "@/components/ui/badge"

export function NavUser({
  user,
  roleLabel,
}: {
  user: {
    name: string
    email: string
    avatar: string
  }
  roleLabel?: string
}) {
  const { isMobile } = useSidebar()
  const { logout, user: authUser } = useAuth()

  const handleLogout = () => {
    logout()
  }

  // Get role label from auth user if available
  const displayRole = roleLabel || (authUser?.role ? 
    (authUser.role === 'admin' ? 'Administrador' :
    authUser.role === 'program lead' ? 'Director de programa' :
    authUser.role === 'cybersecurity_auditor' ? 'Experto en Ciberseguridad' :
    authUser.role) : 'Usuario')

  return (
    <SidebarMenu>
      <SidebarMenuItem>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <SidebarMenuButton
              size="lg"
              className="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground hover:bg-sidebar-accent/50 transition-colors h-auto py-4"
            >
              <Avatar className="h-12 w-12 rounded-xl border-2 border-primary/20">
                <AvatarImage src={user.avatar} alt={user.name} />
                <AvatarFallback className="rounded-xl bg-gradient-to-br from-primary/20 to-primary/10 text-primary font-semibold">
                  {user.name
                    .split(" ")
                    .map((n) => n[0])
                    .join("")
                    .toUpperCase()}
                </AvatarFallback>
              </Avatar>
              <div className="grid flex-1 text-left text-sm leading-tight min-w-0 space-y-2">
                <span className="truncate font-medium text-foreground">{user.name}</span>
                <div className="flex items-center gap-1 min-w-0">
                  <Badge variant="secondary" className="text-xs px-2 py-1 bg-primary/10 text-primary border-primary/20 break-words leading-tight">
                    {displayRole}
                  </Badge>
                </div>
              </div>
              <ChevronsUpDown className="ml-auto size-4 text-muted-foreground flex-shrink-0" />
            </SidebarMenuButton>
          </DropdownMenuTrigger>
          <DropdownMenuContent
            className="w-64 rounded-xl border shadow-lg"
            side={isMobile ? "bottom" : "right"}
            align="end"
            sideOffset={4}
          >
            <DropdownMenuLabel className="p-0 font-normal">
              <div className="flex items-center gap-3 px-4 py-4 text-left">
                <Avatar className="h-12 w-12 rounded-xl border-2 border-primary/20">
                  <AvatarImage src={user.avatar} alt={user.name} />
                  <AvatarFallback className="rounded-xl bg-gradient-to-br from-primary/20 to-primary/10 text-primary font-semibold">
                    {user.name
                      .split(" ")
                      .map((n) => n[0])
                      .join("")
                      .toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div className="grid flex-1 text-left leading-tight min-w-0">
                  <span className="truncate font-medium text-foreground">{user.name}</span>
                  <span className="truncate text-sm text-muted-foreground">{user.email}</span>
                  <Badge variant="secondary" className="text-xs px-2 py-0.5 bg-primary/10 text-primary border-primary/20 w-fit mt-1">
                    {displayRole}
                  </Badge>
                </div>
              </div>
            </DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuGroup>
              <DropdownMenuItem className="gap-3 px-4 py-3">
                <User className="h-4 w-4" />
                <span>Mi Perfil</span>
              </DropdownMenuItem>
              <DropdownMenuItem className="gap-3 px-4 py-3">
                <Settings className="h-4 w-4" />
                <span>Configuración</span>
              </DropdownMenuItem>
              <DropdownMenuItem className="gap-3 px-4 py-3">
                <Shield className="h-4 w-4" />
                <span>Seguridad</span>
              </DropdownMenuItem>
            </DropdownMenuGroup>
            <DropdownMenuSeparator />
            <DropdownMenuItem 
              onClick={handleLogout}
              className="gap-3 px-4 py-3 text-red-600 hover:text-red-700 hover:bg-red-50 focus:text-red-700 focus:bg-red-50"
            >
              <LogOut className="h-4 w-4" />
              <span>Cerrar Sesión</span>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarMenuItem>
    </SidebarMenu>
  )
}
