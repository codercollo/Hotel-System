import { authService } from "~/services/auth.service";

/**
 * useAuth — provides reactive auth state and all auth actions.
 * Wraps authStore + authService so components never touch raw tokens.
 */
export const useAuth = () => {
  const store = useAuthStore();
  const router = useRouter();

  const isAuthenticated = computed(() => store.isAuthenticated);
  const user = computed(() => store.user);
  const role = computed(() => store.user?.role ?? null);
  const isAdmin = computed(() => store.user?.role === "admin");

  /**
   * Login with email + password. Stores tokens and fetches /me profile.
   */
  async function login(email: string, password: string): Promise<void> {
    const pair = await authService.login({ email, password });
    store.setTokens(pair.access_token, pair.refresh_token);

    // Fetch full profile after login
    try {
      const me = await authService.me();
      store.setUser(me);
    } catch {
      // Non-fatal — token is valid even if /me fails
      store.setUser({ user_id: "", email, role: "guest" });
    }
  }

  /**
   * Register a new account and auto-login.
   */
  async function register(
    name: string,
    email: string,
    password: string,
  ): Promise<void> {
    const pair = await authService.register({ name, email, password });
    store.setTokens(pair.access_token, pair.refresh_token);

    try {
      const me = await authService.me();
      store.setUser(me);
    } catch {
      store.setUser({ user_id: "", email, role: "guest" });
    }
  }

  /**
   * Logout — revokes refresh token on server then clears local state.
   */
  async function logout(): Promise<void> {
    if (store.refreshToken) {
      try {
        await authService.logout(store.refreshToken);
      } catch {
        // Best-effort; clear client state regardless
      }
    }
    store.clear();
    await router.push("/auth/login");
  }

  /**
   * Silently refresh the access token using the stored refresh token.
   * Called by auth.client.ts plugin on app mount.
   */
  async function refreshSession(): Promise<boolean> {
    if (!store.refreshToken) return false;
    try {
      const pair = await authService.refresh({
        refresh_token: store.refreshToken,
      });
      store.setTokens(pair.access_token, pair.refresh_token);
      const me = await authService.me();
      store.setUser(me);
      return true;
    } catch {
      store.clear();
      return false;
    }
  }

  return {
    isAuthenticated,
    user,
    role,
    isAdmin,
    login,
    logout,
    register,
    refreshSession,
  };
};
