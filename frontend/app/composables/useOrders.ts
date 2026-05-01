// useOrders — Phase 4 wires this to the orders API
import type { Order } from "~/types/order.types";

export const useOrders = () => {
  const api = useApi();
  const orders = ref<Order[]>([]);
  const order = ref<Order | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  const list = async (params?: Record<string, string>) => {
    loading.value = true;
    error.value = null;
    try {
      orders.value = await api.get<Order[]>("/api/v1/orders", params);
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  const get = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      order.value = await api.get<Order>(`/api/v1/orders/${id}`);
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  const create = async (payload: unknown) => {
    loading.value = true;
    error.value = null;
    try {
      order.value = await api.post<Order>("/api/v1/orders", payload);
      return order.value;
    } catch (e: any) {
      error.value = e.message;
      throw e;
    } finally {
      loading.value = false;
    }
  };

  return { orders, order, loading, error, list, get, create };
};
