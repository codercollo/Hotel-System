import { defineStore } from "pinia";

interface User {
  id: string;
  name: string;
  email: string;
  role: string;
}

interface AuthState {
  accessToken: string | null;
  refreshToken: string | null;
  user: User | null;
}

export const useAuthStore = defineStore("auth", {
  state: (): AuthState => ({
    accessToken: null,
    refreshToken: null,
    user: null,
  }),
  getters: {
    isAuthenticated: (s) => !!s.accessToken,
    isAdmin: (s) => s.user?.role === "admin",
  },
  actions: {
    setTokens(access: string, refresh: string) {
      this.accessToken = access;
      this.refreshToken = refresh;
    },
    setUser(user: User) {
      this.user = user;
    },
    clear() {
      this.accessToken = null;
      this.refreshToken = null;
      this.user = null;
    },
  },
});
