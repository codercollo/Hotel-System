// useItems — wired to the real API via useApi()
// Supports both:
//   - imperative: call list() / get() manually (used in components)
//   - declarative: useFeaturedItems() / useItemDetail() return useAsyncData refs (used in pages)
import type { Item } from "~/types/item.types";

// ── Imperative composable (your existing pattern, kept intact) ────────────────
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
      total.value = res.length;
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : "Failed to load items";
    } finally {
      loading.value = false;
    }
  };

  const get = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      item.value = await api.get<Item>(`/api/v1/items/${id}`);
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : "Failed to load item";
    } finally {
      loading.value = false;
    }
  };

  return { items, item, loading, error, total, list, get };
};

// ── Declarative helpers for pages (SSR-safe with useAsyncData) ────────────────

/**
 * useFeaturedItems — fetches the first `limit` active items for the home page.
 * Returns useAsyncData refs so SSR works correctly.
 */
export const useFeaturedItems = (limit = 3) => {
  const api = useApi();
  return useAsyncData<Item[]>(
    `items-featured-${limit}`,
    () =>
      api.get<Item[]>("/api/v1/items", {
        limit: String(limit),
        status: "active",
      }),
    { default: () => [] as Item[] },
  );
};

/**
 * useItemList — reactive list with optional params (for /items page).
 */
export const useItemList = (params?: MaybeRef<Record<string, string>>) => {
  const api = useApi();
  const resolvedParams = isRef(params) ? params : ref(params ?? {});

  return useAsyncData<Item[]>(
    `items-list`,
    () => api.get<Item[]>("/api/v1/items", resolvedParams.value),
    {
      default: () => [] as Item[],
      watch: [resolvedParams],
    },
  );
};

/**
 * useItemDetail — fetches a single item by id for the detail page.
 */
export const useItemDetail = (id: MaybeRef<string>) => {
  const api = useApi();
  const resolvedId = isRef(id) ? id : ref(id);

  return useAsyncData<Item | null>(
    `item-${resolvedId.value}`,
    () => api.get<Item>(`/api/v1/items/${resolvedId.value}`),
    {
      default: () => null,
      watch: [resolvedId],
    },
  );
};
