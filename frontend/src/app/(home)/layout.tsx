import type { Metadata } from "next";
import { ThemeProvider } from "@/components/theme-provider";
import "./globals.css";
import Link from "next/link";
import { ModeToggle } from "@/components/theme-toggler";
import { NavLinks } from "@/components/common/nav-links.component";

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
                <Link className="text-3xl font-bold" href={`/`}>
                  CricKendra
                </Link>
                <div className="hidden md:block">
                  <NavLinks />
                </div>
                <ModeToggle />
              </div>
              <div className="w-full md:hidden">
                <NavLinks />
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
