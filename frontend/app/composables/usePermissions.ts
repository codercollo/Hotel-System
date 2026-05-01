// usePermissions — RBAC checks mirrored from backend roles
// Phase 2 expands this once roles are loaded from the API
export const usePermissions = () => {
  const store = useAuthStore();

  const can = (permission: string): boolean => {
    if (store.isAdmin) return true;
    // Expand in Phase 2 with per-user permission list from API
    const userPermissions = ["items:read", "orders:create", "orders:read"];
    return userPermissions.includes(permission);
  };

  const isAdmin = computed(() => store.isAdmin);

  return { can, isAdmin };
};
