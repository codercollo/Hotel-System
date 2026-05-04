import type { SearchResult } from "~/types/search.types";

export const useSearch = () => {
  const api = useApi();
  const query = ref("");
  const results = ref<SearchResult[]>([]);
  const loading = ref(false);

  let timer: ReturnType<typeof setTimeout>;

  const search = async (q: string) => {
    if (!q.trim()) {
      results.value = [];
      return;
    }
    loading.value = true;
    try {
      results.value = await api.get<SearchResult[]>("/api/v1/search", { q });
    } catch {
      results.value = [];
    } finally {
      loading.value = false;
    }
  };

  const debouncedSearch = (q: string) => {
    clearTimeout(timer);
    timer = setTimeout(() => search(q), 300);
  };

  watch(query, debouncedSearch);

  const clear = () => {
    query.value = "";
    results.value = [];
  };

  return { query, results, loading, search, debouncedSearch, clear };
};
