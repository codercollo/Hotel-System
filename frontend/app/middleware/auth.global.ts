// Global middleware: routes with meta.requiresAuth = true redirect to /auth/login
// when the user is not authenticated. Phase 2 wires this to the auth store.
export default defineNuxtRouteMiddleware((to) => {
  const protectedPrefixes = ["/dashboard", "/orders", "/account", "/admin"];
  const needsAuth = protectedPrefixes.some((p) => to.path.startsWith(p));

  if (!needsAuth) return;

  // Phase 2: replace with real auth store check
  const isAuthenticated = false;

  if (!isAuthenticated) {
    return navigateTo("/auth/login");
  }
});
