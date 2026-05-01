// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  srcDir: "app/",
  modules: ["@nuxtjs/tailwindcss", "@pinia/nuxt", "@nuxt/icon"],
  icon: {
    serverBundle: "local",
  },
  css: ["@/assets/css/main.css"],
  components: [
    {
      path: "~/components",
      pathPrefix: false,
    },
  ],
  compatibilityDate: "2026-05-01",
});
