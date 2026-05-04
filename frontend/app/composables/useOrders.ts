import type { Order, CreateOrderRequest } from "~/types/order.types";

export const useOrders = () => {
  const api = useApi();
  const orders = ref<Order[]>([]);
  const order = ref<Order | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);
  const total = ref(0);

  const list = async (params?: Record<string, string>) => {
    loading.value = true;
    error.value = null;
    try {
      // Backend returns { data: Order[], meta: { total, limit, offset } }
      // useApi unwraps data → Order[]
      orders.value = await api.get<Order[]>("/api/v1/orders", {
        limit: "50",
        offset: "0",
        ...params,
      });
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

  const create = async (payload: CreateOrderRequest) => {
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

  const cancel = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      order.value = await api.patch<Order>(`/api/v1/orders/${id}/cancel`, {});
      // Update in list if present
      const idx = orders.value.findIndex((o) => o.id === id);
      if (idx !== -1 && order.value) orders.value[idx] = order.value;
      return order.value;
    } catch (e: any) {
      error.value = e.message;
      throw e;
    } finally {
      loading.value = false;
    }
  };

  return { orders, order, loading, error, total, list, get, create, cancel };
};
