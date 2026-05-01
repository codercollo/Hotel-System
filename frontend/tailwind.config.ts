import type { Config } from "tailwindcss";

export default {
  content: [
    "./app/**/*.{vue,ts,js,jsx,tsx}",
    "./app/components/**/*.{vue,ts,js}",
    "./app/layouts/**/*.{vue,ts,js}",
    "./app/pages/**/*.{vue,ts,js}",
    "./app/composables/**/*.{ts,js}",
    "./nuxt.config.{js,ts}",
  ],
  theme: {
    extend: {
      fontFamily: {
        display: ["Cormorant Garamond", "Georgia", "serif"],
        sans: ["Jost", "Helvetica Neue", "sans-serif"],
      },
      colors: {
        brand: {
          50:  "rgb(250 248 244 / <alpha-value>)",
          100: "rgb(243 238 228 / <alpha-value>)",
          200: "rgb(228 216 198 / <alpha-value>)",
          300: "rgb(207 188 163 / <alpha-value>)",
          400: "rgb(183 155 121 / <alpha-value>)",
          500: "rgb(163 130 93  / <alpha-value>)",
          600: "rgb(140 108 72  / <alpha-value>)",
          700: "rgb(73  74  46  / <alpha-value>)",
          800: "rgb(52  58  32  / <alpha-value>)",
          900: "rgb(30  34  18  / <alpha-value>)",
        },
        surface: {
          50:  "rgb(255 255 255 / <alpha-value>)",
          100: "rgb(250 249 246 / <alpha-value>)",
          200: "rgb(237 234 226 / <alpha-value>)",
          300: "rgb(73  74  80  / <alpha-value>)",
        },
        forest: {
          DEFAULT: "rgb(38  79  72  / <alpha-value>)",
          light:   "rgb(52  101 90  / <alpha-value>)",
          dark:    "rgb(25  55  48  / <alpha-value>)",
        },
        accent: {
          DEFAULT: "rgb(196 156 85  / <alpha-value>)",
          light:   "rgb(230 196 130 / <alpha-value>)",
        },
        muted: "rgb(140 133 120 / <alpha-value>)",
      },
      fontSize: {
        "display-xl": ["clamp(2.5rem, 6vw, 5rem)",   { lineHeight: "1.1" }],
        "display-lg": ["clamp(2rem, 4vw, 3.5rem)",   { lineHeight: "1.15" }],
        "display-md": ["clamp(1.5rem, 3vw, 2.5rem)", { lineHeight: "1.2" }],
        "display-sm": ["clamp(1.25rem, 2vw, 1.875rem)", { lineHeight: "1.25" }],
      },
      boxShadow: {
        card:       "0 4px 24px -4px rgba(30,34,18,0.08)",
        "card-hover":"0 12px 48px -8px rgba(30,34,18,0.15)",
        nav:        "0 2px 16px rgba(30,34,18,0.06)",
      },
      backgroundImage: {
        "gold-gradient":   "linear-gradient(135deg, rgb(196 156 85), rgb(230 196 130))",
        "forest-gradient": "linear-gradient(135deg, rgb(25 55 48), rgb(38 79 72))",
      },
      animation: {
        "fade-up": "fadeUp 0.7s cubic-bezier(0.22,1,0.36,1) both",
        "fade-in": "fadeIn 0.5s ease both",
        ticker:    "ticker 28s linear infinite",
      },
      keyframes: {
        fadeUp: {
          from: { opacity: "0", transform: "translateY(24px)" },
          to:   { opacity: "1", transform: "translateY(0)" },
        },
        fadeIn: {
          from: { opacity: "0" },
          to:   { opacity: "1" },
        },
        ticker: {
          from: { transform: "translateX(0)" },
          to:   { transform: "translateX(-50%)" },
        },
      },
      spacing: {
        18: "4.5rem",
        22: "5.5rem",
        26: "6.5rem",
      },
    },
  },
  plugins: [],
} satisfies Config;