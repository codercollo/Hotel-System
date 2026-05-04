import type { AdminStats } from "~/types/admin.types";
import type { User } from "~/types/user.types";
import type { Order } from "~/types/order.types";
import type { Item } from "~/types/item.types";

export const adminService = {
  stats: () => useApi().get<AdminStats>("/api/v1/admin/stats"),
  listUsers: (p?: Record<string, string>) =>
    useApi().get<User[]>("/api/v1/users", p),
  listItems: (p?: Record<string, string>) =>
    useApi().get<Item[]>("/api/v1/items", p),
  listOrders: (p?: Record<string, string>) =>
    useApi().get<Order[]>("/api/v1/orders", p),
  createItem: (body: unknown) => useApi().post<Item>("/api/v1/items", body),
  updateItem: (id: string, body: unknown) =>
    useApi().patch<Item>(`/api/v1/items/${id}`, body),
  deleteItem: (id: string) => useApi().del(`/api/v1/items/${id}`),
  updateOrderStatus: (id: string, status: string) =>
    useApi().patch<Order>(`/api/v1/orders/${id}/status`, { status }),
};
