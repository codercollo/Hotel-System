// /**
//  * role.ts — RBAC gate for routes that require a specific role.
//  *
//  * Usage in page:
//  *   definePageMeta({ middleware: ['role'], roles: ['admin'] })
//  *
//  * Routes without `meta.roles` are skipped — this middleware is opt-in,
//  * not applied globally. Authentication is handled separately by auth.global.ts.
//  */
// export default defineNuxtRouteMiddleware((to) => {
//   if (import.meta.server) return;

//   const allowedRoles = to.meta.roles as string[] | undefined;
//   if (!allowedRoles || allowedRoles.length === 0) return;

//   const store = useAuthStore();
//   const userRole = store.user?.role;

//   if (!userRole || !allowedRoles.includes(userRole)) {
//     // Authenticated but insufficient role → redirect to their own dashboard
//     return navigateTo("/dashboard");
//   }
// });
