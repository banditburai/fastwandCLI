/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
      "./**/*.py",    
    ],
    theme: {
      extend: {},
    },
    daisyui: {
      themes: ["retro", "dim"],
    },
    plugins: [require("@tailwindcss/typography"), require("daisyui")],
  }