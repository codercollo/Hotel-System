// useItems — Phase 3 will wire this to the real API via useApi()
import type { Item } from "~/types/item.types";

export const useItems = () => {
  const api = useApi();
  const items = ref<Item[]>([]);
  const item = ref<Item | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);
  const total = ref(0);

  const list = async (params?: Record<string, string>) => {
    loading.value = true;
    error.value = null;
    try {
      const res = await api.get<Item[]>("/api/v1/items", params);
      items.value = res;
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
      item.value = await api.get<Item>(`/api/v1/items/${id}`);
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  return { items, item, loading, error, total, list, get };
};
