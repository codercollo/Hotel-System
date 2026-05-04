/**
 * admin.ts — convenience middleware that gates admin-only pages.
 *
 * Usage:
 *   definePageMeta({ middleware: ['admin'] })
 *
 * Redirects non-admins to /dashboard rather than /auth/login,
 * since they are authenticated — just not privileged enough.
 */
export default defineNuxtRouteMiddleware(() => {
  if (import.meta.server) return;

  const store = useAuthStore();

  if (store.user?.role !== "admin") {
    return navigateTo("/dashboard");
  }
});
