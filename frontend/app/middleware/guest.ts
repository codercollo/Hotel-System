// /**
//  * guest.ts — redirects authenticated users away from guest-only pages.
//  *
//  * Usage: definePageMeta({ middleware: ['guest'] })
//  *
//  * Note: the global auth middleware already handles this for /auth/* pages.
//  * Use this named middleware for any additional guest-only routes.
//  */
// export default defineNuxtRouteMiddleware(() => {
//   const store = useAuthStore();

//   if (store.isAuthenticated) {
//     return navigateTo("/dashboard");
//   }
// });
