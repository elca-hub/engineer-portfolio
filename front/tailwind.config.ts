import type { Config } from "tailwindcss";

export default {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./constants/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "#F7F7F7",
        foreground: "#020826",
        primary: "#A8C8D6",
        secondary: "#6DA6C3",
        subtext: "#5B5F71",
        lightblue: "#EDF8FD",
      },
    },
  },
  plugins: [],
} satisfies Config;
