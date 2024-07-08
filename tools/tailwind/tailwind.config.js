/** @type {import('tailwindcss').Config} */
module.exports = {
  // this path is ran relative to the root directory 
  // as per the tools.go generate commands
  content: ['ui/**/**/*.templ'],
  options: {
    defaultExtractor: content => content.match(/[\w-/:]+(?<!:)/g) || []
  },
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
