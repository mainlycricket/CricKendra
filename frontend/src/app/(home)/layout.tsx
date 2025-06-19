import type { Metadata } from "next";
import { ThemeProvider } from "@/components/theme-provider";
import "./globals.css";
import Link from "next/link";
import { ModeToggle } from "@/components/theme-toggler";

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
          <ThemeProvider
            attribute="class"
            defaultTheme="system"
            enableSystem
            disableTransitionOnChange
          >
            <header>
              <p>Menu</p>
              <ModeToggle />
              <div>
                <Link href="/matches/1">Match 1</Link>
                <Link href="/series/1">Series 1</Link>
              </div>
            </header>
            {children}
            <footer>Footer</footer>
          </ThemeProvider>
        </div>
      </body>
    </html>
  );
}
