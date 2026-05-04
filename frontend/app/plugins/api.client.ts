// Injects $api into the Nuxt app instance.
// Components can use const { $api } = useNuxtApp() for direct access,
// but prefer the useApi() composable in most cases.
// plugins/api.ts
export default defineNuxtPlugin(() => {
  const api = useApi();
  return {
    provide: { api },
  };
});
