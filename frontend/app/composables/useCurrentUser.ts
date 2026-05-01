// useCurrentUser — Phase 2 wires this to the auth store
export const useCurrentUser = () => {
  const store = useAuthStore();
  return {
    user: computed(() => store.user),
    isAuthenticated: computed(() => store.isAuthenticated),
    isAdmin: computed(() => store.isAdmin),
  };
};
