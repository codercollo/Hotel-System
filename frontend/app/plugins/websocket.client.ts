export default defineNuxtPlugin(() => {
  if (import.meta.server) return;

  const config = useRuntimeConfig();
  const apiBase = (config.public.apiBase as string) ?? "";
  const wsBase = apiBase.replace(/^http/, "ws");
  const authStore = useAuthStore();
  const ws = useWebSocket(wsBase);

  watch(
    () => authStore.accessToken,
    (token) => {
      if (token) {
        ws.connect(token);
      } else {
        ws.disconnect();
      }
    },
    { immediate: true },
  );
});
