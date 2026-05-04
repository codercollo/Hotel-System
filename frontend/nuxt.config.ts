// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  srcDir: "app/",
  modules: [
    "@nuxtjs/tailwindcss",
    "@pinia/nuxt",
    "@pinia-plugin-persistedstate/nuxt",
    "@nuxt/icon",
    "@nuxt/image",
  ],
  icon: {
    serverBundle: "local",
  },
  css: ["@/assets/css/main.css", "leaflet/dist/leaflet.css"],
  components: [
    {
      path: "~/components",
      pathPrefix: false,
    },
  ],
  runtimeConfig: {
    public: {
      apiBase: "http://localhost:8080",
      wsBase: "",
    },
  },
  compatibilityDate: "2026-05-01",
});
