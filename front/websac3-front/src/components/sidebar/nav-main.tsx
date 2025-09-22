"use client"

import { ChevronRight, type LucideIcon } from "lucide-react"
import { usePathname } from "next/navigation"

import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible"
import {
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuAction,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "@/components/ui/sidebar"

export function NavMain({
  items,
}: {
  items: {
    title: string
    url: string
    icon: LucideIcon
    items?: {
      title: string
      url: string
    }[]
  }[]
}) {
  const pathname = usePathname();

  // Función para determinar si un elemento está activo
  const isItemActive = (url: string) => {
    if (url === pathname) return true;
    
    // Para URLs que terminan en dashboard, también considerar activas si estamos en la raíz del rol
    if (url.endsWith('/dashboard')) {
      const rolePath = url.replace('/dashboard', '');
      return pathname === rolePath;
    }
    
    return false;
  };

  return (
    <SidebarGroup>
      <SidebarGroupLabel className="text-primary/70 font-medium text-xs uppercase tracking-wider px-2 py-1">
        Navegación
      </SidebarGroupLabel>
      <SidebarMenu>
        {items.map((item) => {
          const isActive = isItemActive(item.url);
          
          return (
            <Collapsible key={item.title} asChild defaultOpen={isActive}>
              <SidebarMenuItem>
                <SidebarMenuButton 
                  asChild 
                  tooltip={item.title}
                  isActive={isActive}
                  className="group hover:bg-primary/5 transition-colors data-[active=true]:bg-primary/10 data-[active=true]:text-primary"
                >
                  <a href={item.url}>
                    <div className={`p-1.5 rounded-md transition-colors ${
                      isActive 
                        ? "text-primary" 
                        : "text-foreground group-hover:text-foreground"
                    }`}>
                      <item.icon className="w-4 h-4" />
                    </div>
                    <span className="text-sm font-medium">{item.title}</span>
                  </a>
                </SidebarMenuButton>
                {item.items?.length ? (
                  <>
                    <CollapsibleTrigger asChild>
                      <SidebarMenuAction className="data-[state=open]:rotate-90 text-foreground hover:text-foreground transition-colors">
                        <ChevronRight className="w-4 h-4" />
                        <span className="sr-only">Toggle</span>
                      </SidebarMenuAction>
                    </CollapsibleTrigger>
                    <CollapsibleContent>
                      <SidebarMenuSub>
                        {item.items?.map((subItem) => (
                          <SidebarMenuSubItem key={subItem.title}>
                            <SidebarMenuSubButton asChild className="hover:bg-muted/30 hover:text-foreground transition-colors">
                              <a href={subItem.url}>
                                <span className="text-sm">{subItem.title}</span>
                              </a>
                            </SidebarMenuSubButton>
                          </SidebarMenuSubItem>
                        ))}
                      </SidebarMenuSub>
                    </CollapsibleContent>
                  </>
                ) : null}
              </SidebarMenuItem>
            </Collapsible>
          );
        })}
      </SidebarMenu>
    </SidebarGroup>
  )
}
