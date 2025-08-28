"use client";

import * as React from "react";
// icons previously used in the demo were removed from here
// to avoid unused imports and linter errors

import { NavMain } from "@/components/sidebar/nav-main";
import { NavSecondary } from "@/components/sidebar/nav-secondary";
// import { NavUser } from "@/components/sidebar/nav-user";
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

type AppSidebarProps = React.ComponentProps<typeof Sidebar> & {
  role?: RoleKey;
};

export default function AppSidebar({
  role = "EXPERTO",
  ...props
}: AppSidebarProps) {
  const data = SIDEBAR_BY_ROLE[role];
  return (
    <Sidebar variant="inset" {...props}>
      <SidebarHeader className="border-b border-primary/20">
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              size="lg"
              asChild
              className="hover:bg-primary/5 transition-colors"
            >
              <a href="#" className="group">
                <Avatar className="h-10 w-10 rounded-xl">
                  <AvatarImage
                    src="https://github.com/shadcn.png"
                    alt={data.user.name}
                  />
                  <AvatarFallback className="rounded-xl bg-primary/10 text-primary font-medium">
                    {data.user.name
                      .split(" ")
                      .map((n) => n[0])
                      .join("")
                      .toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <span className="truncate font-medium text-foreground">
                    {data.user.name}
                  </span>
                  <span className="truncate text-xs text-primary bg-primary/5 px-2 py-1 rounded-full">
                    {data.roleLabel}
                  </span>
                </div>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      {data.navSecondary ? (
        <SidebarFooter>
          <NavSecondary items={data.navSecondary} />
        </SidebarFooter>
      ) : null}
    </Sidebar>
  );
}
