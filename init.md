# Init Project

## Init React
Buat folder react
```
npm create vite client -- --template react
```

Install dependencies
```
cd client
npm install
```

Coba run`client`
```
npm run dev
```

## Install Tailwind

Install konfigurasi utama tailwind `tailwind.config.js`

```
npm install -D tailwindcss@3 postcss autoprefixer
```

Init `postcss.config.js`
```
npx tailwindcss init -p
```

Ubah `tailwind.config.js` jadi kayak gini
```
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

Tambah tailwind directive di `index.css`
```
@tailwind base;
@tailwind components;
@tailwind utilities;
```
