/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['../ui/**/**/*.templ'],
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
