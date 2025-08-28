"use client";

import React from "react";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { usePathname } from "next/navigation";

type BreadcrumbItem = {
  title: string;
  href: string;
  isCurrent: boolean;
};

interface BreadcrumbsProps {
  role: string;
  getPageTitle: (page: string) => string;
}

export function Breadcrumbs({ role, getPageTitle }: BreadcrumbsProps) {
  const pathname = usePathname();

  const generateBreadcrumbs = (): BreadcrumbItem[] => {
    const segments = pathname.split("/").filter(Boolean);
    const breadcrumbs: BreadcrumbItem[] = [];

    if (segments.length === 1 && segments[0] === role.toLowerCase()) {
      breadcrumbs.push({
        title: role,
        href: `/${role.toLowerCase()}/dashboard`,
        isCurrent: true,
      });
      return breadcrumbs;
    }

    if (segments.length > 1) {
      breadcrumbs.push({
        title: role,
        href: `/${role.toLowerCase()}/dashboard`,
        isCurrent: false,
      });
    }

    if (segments.length > 1) {
      const currentPage = segments[1];
      const pageTitle = getPageTitle(currentPage);
      breadcrumbs.push({
        title: pageTitle,
        href: pathname,
        isCurrent: true,
      });
    }

    return breadcrumbs;
  };

  const breadcrumbs = generateBreadcrumbs();

  return (
    <Breadcrumb>
      <BreadcrumbList>
        {breadcrumbs.map((breadcrumb, index) => (
          <React.Fragment key={`${breadcrumb.title}-${index}`}>
            <BreadcrumbItem>
              {breadcrumb.isCurrent ? (
                <BreadcrumbPage>{breadcrumb.title}</BreadcrumbPage>
              ) : (
                <BreadcrumbLink href={breadcrumb.href}>
                  {breadcrumb.title}
                </BreadcrumbLink>
              )}
            </BreadcrumbItem>
            {index < breadcrumbs.length - 1 && (
              <BreadcrumbSeparator className="hidden md:block" />
            )}
          </React.Fragment>
        ))}
      </BreadcrumbList>
    </Breadcrumb>
  );
}
