import type { Metadata } from "next";
import { ThemeProvider } from "@/components/theme-provider";
import "./globals.css";
import Link from "next/link";
import { ModeToggle } from "@/components/theme-toggler";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
} from "@/components/ui/navigation-menu";

export const metadata: Metadata = {
  title: "CricKendra",
  description: "Centre of Cricket",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body>
        <div className="container mx-auto p-2">
          <ThemeProvider attribute="class" defaultTheme="system" enableSystem disableTransitionOnChange>
            <header className="py-4 px-4 md:px-16 flex flex-col gap-4">
              <div className="flex justify-between items-center">
                <h2 className="text-3xl font-bold">CricKendra</h2>
                <div className="hidden md:block">
                  <Menu />
                </div>
                <ModeToggle />
              </div>
              <div className="w-full md:hidden">
                <Menu />
              </div>
            </header>
            <main className="md:px-16 pb-4 md:pt-4">{children}</main>
            <footer></footer>
          </ThemeProvider>
        </div>
      </body>
    </html>
  );
}

function Menu() {
  const links = [
    { label: "Series 1", href: "/series/1" },
    { label: "Match 1", href: "/matches/1" },
    { label: "Player 1", href: "/players/1" },
  ];

  return (
    <NavigationMenu>
      <NavigationMenuList>
        {links.map((link) => {
          return (
            <NavigationMenuItem key={link.href}>
              <NavigationMenuLink asChild className="text-base tracking-wide underline">
                <Link href={link.href} className="hover:text-sky-500">
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
