/**
 * auth.client.ts
 *
 * Runs once on the client after the app mounts.
 * 1. Reads persisted tokens from localStorage.
 * 2. Hydrates the auth store so the UI doesn't flash as unauthenticated.
 * 3. Attempts a silent refresh to validate the stored session and
 *    rotate tokens if they are still valid.
 */
export default defineNuxtPlugin(async () => {
  if (import.meta.server) return;

  const store = useAuthStore();

  const accessToken = localStorage.getItem("access_token");
  const refreshToken = localStorage.getItem("refresh_token");
  const userRaw = localStorage.getItem("user");

  // Nothing stored — nothing to rehydrate.
  if (!accessToken || !refreshToken) return;

  // Hydrate synchronously so middleware/guards see a valid session immediately.
  store.setTokens(accessToken, refreshToken);

  if (userRaw) {
    try {
      store.setUser(JSON.parse(userRaw));
    } catch {
      // Corrupt JSON — clear and start fresh
      store.clear();
      return;
    }
  }

  // Silently validate / rotate the stored refresh token.
  // If this fails the store is cleared and the user will be redirected
  // to /auth/login by the global auth middleware on the next navigation.
  const { refreshSession } = useAuth();
  await refreshSession();
});
