terminal:
npm init
npm install -D tailwindcss
npx tailwindcss init

tailwind.config.js:
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/html/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [],
}

create folder src and css file named input.css:
@tailwind base;
@tailwind components;
@tailwind utilities;

terminal:
npx tailwindcss -i ./src/input.css -o ./views/css/style_tailwind.css --watch

install plugin daisyui:
terminal:
npm i -D daisyui@latest

tailwind.config.js:
module.exports = {
  //...
  plugins: [require("daisyui")],
}