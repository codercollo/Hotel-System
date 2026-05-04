import { defineStore } from "pinia";
import type { User } from "~/types/user.types";

export const useUserStore = defineStore("user", {
  state: () => ({
    profile: null as User | null,
    loading: false,
  }),
  actions: {
    setProfile(u: User) {
      this.profile = u;
    },
    clear() {
      this.profile = null;
    },
  },
});
