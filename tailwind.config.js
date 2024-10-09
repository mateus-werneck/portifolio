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
        inter: ["Inter", "sans-serif"],
        sourcesans: ["'Source Sans 3'", "Inter"],
        jetbrains: ["'JetBrains Mono'", "Inter"],
      },
    },
  },
  plugins: [],
};
