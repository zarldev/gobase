/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ"],
  theme: {
    extend: {
      textColor: ['group-hover'],
      colors: {},
    },
    plugins: [
      "@tailwindcss/typography",
      "@tailwindcss/forms",
      "@tailwindcss/line-clamp"
    ],
  },

};
