"use client";

import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
} from "@/components/ui/navigation-menu";
import Link from "next/link";
import { usePathname } from "next/navigation";

export function NavLinks() {
  const currentPathRoot = usePathname().split("/")[1];

  const links = [
    { label: "Series", href: "/series" },
    { label: "Matches", href: "/matches" },
    { label: "Players", href: "/players" },
    { label: "Stats", href: "/stats/filters" },
  ];

  return (
    <NavigationMenu>
      <NavigationMenuList>
        {links.map((link) => {
          const linkRoot = link.href.split("/")[1];

          return (
            <NavigationMenuItem key={link.href}>
              <NavigationMenuLink asChild className="text-base tracking-wide underline">
                <Link
                  href={link.href}
                  className={`${linkRoot === currentPathRoot ? "text-sky-500" : ""} hover:text-sky-500`}
                >
                  {link.label}
                </Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
          );
        })}
      </NavigationMenuList>
    </NavigationMenu>
  );
}
