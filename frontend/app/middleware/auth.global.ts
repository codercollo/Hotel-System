// middleware/auth.global.ts
const PROTECTED_PREFIXES = ["/dashboard", "/orders", "/account", "/admin"];

export default defineNuxtRouteMiddleware((to) => {
  if (import.meta.server) return;

  const store = useAuthStore();
  const isAuth = store.isAuthenticated;
  const isAuthRoute = to.path.startsWith("/auth");

  const isProtected =
    to.meta.requiresAuth === true ||
    PROTECTED_PREFIXES.some((prefix) => to.path.startsWith(prefix));

  // Authenticated user on /auth/* → send home
  if (isAuth && isAuthRoute) {
    return navigateTo("/");
  }

  // Unauthenticated user on protected route → login with redirect
  if (!isAuth && isProtected) {
    return navigateTo({
      path: "/auth/login",
      query: { redirect: to.fullPath !== "/" ? to.fullPath : undefined },
    });
  }
});
