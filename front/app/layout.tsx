import { Noto_Sans_JP } from "next/font/google";
import "./globals.css";
import { Provider } from "@/components/ui/provider";


const notoSansJP = Noto_Sans_JP({
  variable: "--font-noto-sans-jp",
  subsets: ["latin"],
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja" suppressHydrationWarning>
      <body
        className={`${notoSansJP.variable} antialiased bg-background`}
      >
        <Provider>
          {children}
        </Provider>
      </body>
    </html>
  );
}
