/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ}", "./**/*.templ"],
  safelist: [],
  plugins: [require('@tailwindcss/aspect-ratio'), require("daisyui")],
  daisyui: {
    themes: ["lemonade", "dark"]
  }
}
