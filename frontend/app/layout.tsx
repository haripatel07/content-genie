import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
// 1. Import the Toaster
import { Toaster } from "react-hot-toast";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Content-Genie",
  description: "AI-Powered Content Repurposing",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    // 2. Set the data-theme attribute
    <html lang="en" data-theme="dark">
      <body className={inter.className}>
        {children}
        {/* 3. Add the Toaster component */}
        <Toaster position="bottom-right" />
      </body>
    </html>
  );
}