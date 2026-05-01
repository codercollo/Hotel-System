import type { Order, CreateOrderRequest } from "~/types/order.types";

export const ordersService = {
  list: (p?: Record<string, string>) =>
    useApi().get<Order[]>("/api/v1/orders", p),
  get: (id: string) => useApi().get<Order>(`/api/v1/orders/${id}`),
  create: (body: CreateOrderRequest) =>
    useApi().post<Order>("/api/v1/orders", body),
  updateStatus: (id: string, status: string) =>
    useApi().patch<Order>(`/api/v1/orders/${id}/status`, { status }),
  cancel: (id: string) =>
    useApi().patch<Order>(`/api/v1/orders/${id}/cancel`, {}),
};
