"use client";

import * as React from "react";
// icons previously used in the demo were removed from here
// to avoid unused imports and linter errors

import { NavMain } from "@/components/sidebar/nav-main";
import { NavSecondary } from "@/components/sidebar/nav-secondary";
import { NavUser } from "@/components/sidebar/nav-user";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";

import { SIDEBAR_BY_ROLE, type RoleKey } from "@/constants/sidebar";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { useAuth } from "@/hooks/useAuth";
import { WebSAC3Logo } from "@/components/websac3/logos/WebSAC3Logo";

type AppSidebarProps = React.ComponentProps<typeof Sidebar> & {
  role?: RoleKey;
};

export default function AppSidebar({
  role = "EXPERTO",
  ...props
}: AppSidebarProps) {
  const data = SIDEBAR_BY_ROLE[role];
  const { user } = useAuth();
  
  // Use authenticated user data if available, otherwise fallback to role data
  const userData = user ? {
    name: user.username,
    email: user.email,
    avatar: "https://github.com/shadcn.png" // You can customize this
  } : data.user;
  return (
    <Sidebar variant="inset" {...props}>
      <SidebarHeader className="border-b border-primary/20">
        <div className="flex items-center justify-center p-2">
          <div className="h-12 w-full">
            <WebSAC3Logo />
          </div>
        </div>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      <SidebarFooter>
        {data.navSecondary ? (
          <NavSecondary items={data.navSecondary} />
        ) : null}
        <NavUser user={userData} roleLabel={data.roleLabel} />
      </SidebarFooter>
    </Sidebar>
  );
}
