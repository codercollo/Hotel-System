// middleware/admin.ts
export default defineNuxtRouteMiddleware(() => {
  if (import.meta.server) return;

  const store = useAuthStore();

  // If store hasn't hydrated yet, let the page load —
  // the page itself should re-check or the plugin will handle it
  if (!store.isAuthenticated) {
    return navigateTo("/auth/login");
  }

  if (store.user?.role !== "admin") {
    return navigateTo("/dashboard");
  }
});
