/**
 * usePermissions — thin RBAC helper composable.
 *
 * Reads the current user's role from the auth store and exposes
 * boolean guards used throughout the UI.
 */
export const usePermissions = () => {
  const store = useAuthStore();
  const role = computed(() => store.user?.role ?? null);

  const hasRole = (...roles: string[]) =>
    computed(() => !!role.value && roles.includes(role.value));

  return {
    role,
    isAdmin: hasRole("admin"),
    isGuest: hasRole("guest"),
    isStaff: hasRole("staff", "admin"),
    /** True if the current user owns the given resource by userId */
    isSelf: (userId: string) => computed(() => store.user?.user_id === userId),
    hasRole,
  };
};
