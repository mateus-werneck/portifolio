/** @type {import('tailwindcss').Config} */
export default {
  content: ["./view/**/*.html"],
  theme: {
    extend: {
      flex: {
        2: "2 2 0%",
        3: "3 3 0%",
        4: "4 4 0%",
      },
      fontFamily: {
        sans: ["Montserrat", "Roboto"],
        inter: ["Inter", "sans-serif"],
        sourcesans: ["'Source Sans 3'", "Inter"],
        jetbrains: ["'JetBrains Mono'", "Inter"],
        lumios: ["'Lumios Typewriter New'", "Inter"] 
      },
     colors: {
         "primary-color": "#2f4f4f",
         "secondary-color": "#A1452D",
         "title-color": "#9CB3A3",
         "sub-title-color": "#897B30",
         "button-text-color": "#9CB3A3",
         "button-border-color": "#897B30"
     },
    backgroundColor: {
        "primary-color": "#2f4f4f",
        "button-hover": "#A1452D"
    },
    },
  },
  plugins: [],
};
