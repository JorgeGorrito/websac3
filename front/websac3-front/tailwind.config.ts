import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          light: '#1E4E79', // Azul principal
          dark: '#122A45',  // Azul más oscuro para modo oscuro
        },
        accent: {
          light: '#E53935', // Rojo brillante (modo claro)
          dark: '#C62828',  // Rojo oscuro (modo oscuro)
        },
        background: {
          light: '#FFFFFF', // Fondo blanco (modo claro)
          dark: '#121212',  // Fondo oscuro (modo oscuro)
        },
        secondary: {
          light: '#3B87C1', // Azul claro (modo claro)
          dark: '#246791',  // Azul claro más oscuro
        },
        neutral: {
          light: '#E5E5E5', // Gris claro (modo claro)
          dark: '#1A1A1A',  // Gris oscuro (modo oscuro)
        },
      },
    },
  },
  plugins: [],
} satisfies Config;
