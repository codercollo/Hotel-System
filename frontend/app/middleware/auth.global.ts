// // server/api/demo/auth/login.post.ts
// import { defineEventHandler, readBody, readRawBody, createError } from "h3";

// interface DemoUser {
//   password: string;
//   user_id: string;
//   name: string;
//   role: string;
// }

// const DEMO_USERS: Record<string, DemoUser> = {
//   "admin@hotel.com": {
//     password: "Admin123!@#",
//     user_id: "01JDEMOUSER001ADMINXXXXXX",
//     name: "Admin User",
//     role: "admin",
//   },
//   "guest@hotel.com": {
//     password: "Guest123!",
//     user_id: "01JDEMOUSER002GUESTXXXXX",
//     name: "Jane Kamau",
//     role: "guest",
//   },
// };

// function b64(str: string): string {
//   return btoa(unescape(encodeURIComponent(str)));
// }

// export default defineEventHandler(async (event) => {
//   await new Promise((r) => setTimeout(r, 400));

//   let email: string | undefined;
//   let password: string | undefined;

//   try {
//     const body = await readBody(event);
//     email = body?.email;
//     password = body?.password;
//   } catch {
//     // fallback below
//   }

//   if (!email || !password) {
//     try {
//       const raw = await readRawBody(event, "utf-8");
//       if (raw) {
//         const parsed = JSON.parse(raw) as { email?: string; password?: string };
//         email = parsed?.email;
//         password = parsed?.password;
//       }
//     } catch {
//       // ignore
//     }
//   }

//   if (!email || !password) {
//     throw createError({
//       statusCode: 400,
//       data: {
//         success: false,
//         error: { code: "BAD_REQUEST", message: "email and password required" },
//       },
//     });
//   }

//   const user = DEMO_USERS[email.trim().toLowerCase()];

//   if (!user || user.password !== password) {
//     throw createError({
//       statusCode: 401,
//       data: {
//         success: false,
//         error: {
//           code: "INVALID_CREDENTIALS",
//           message: "Invalid email or password",
//         },
//       },
//     });
//   }

//   const now = new Date();
//   const expiresAt = new Date(now.getTime() + 15 * 60 * 1000).toISOString();
//   const payload = b64(
//     JSON.stringify({ user_id: user.user_id, role: user.role, exp: expiresAt }),
//   );

//   return {
//     success: true,
//     data: {
//       access_token: `demo.access.${payload}`,
//       refresh_token: `demo.refresh.${user.user_id}.${Date.now()}`,
//       expires_at: expiresAt,
//       user: {
//         user_id: user.user_id,
//         email: email.trim().toLowerCase(),
//         name: user.name,
//         role: user.role,
//       },
//     },
//   };
// });

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
    // Only keep the redirect if it's a safe internal path (not /auth/*)
    const redirect =
      to.fullPath !== "/" && !to.path.startsWith("/auth")
        ? to.fullPath
        : undefined;

    return navigateTo({ path: "/auth/login", query: { redirect } });
  }
});
