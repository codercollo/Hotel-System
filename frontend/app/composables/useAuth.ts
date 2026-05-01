// useAuth provides reactive auth state and actions.
// The implementation is completed in Phase 2 when the auth store is wired.
export const useAuth = () => {
  const isAuthenticated = computed(() => false); // Phase 2: reads from auth store
  const user = computed(() => null);

  const login = async (_email: string, _password: string) => {
    /* Phase 2 */
  };
  const logout = async () => {
    /* Phase 2 */
  };
  const register = async (_name: string, _email: string, _password: string) => {
    /* Phase 2 */
  };

  return { isAuthenticated, user, login, logout, register };
};
