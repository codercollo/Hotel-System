export const adminService = {
  stats: () => useApi().get("/api/v1/admin/stats"),
  listUsers: (p?: Record<string, string>) =>
    useApi().get("/api/v1/admin/users", p),
  listOrders: (p?: Record<string, string>) =>
    useApi().get("/api/v1/admin/orders", p),
  updateSetting: (key: string, value: unknown) =>
    useApi().patch("/api/v1/admin/settings", { key, value }),
};
