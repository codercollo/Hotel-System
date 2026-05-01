// useSearch — Phase 8 wires this to the search API
export const useSearch = () => {
  const api = useApi();
  const query = ref("");
  const results = ref<any[]>([]);
  const loading = ref(false);

  const search = useDebounceFn(async (q: string) => {
    if (!q.trim()) {
      results.value = [];
      return;
    }
    loading.value = true;
    try {
      results.value = await api.get("/api/v1/search", { q });
    } catch {
      results.value = [];
    } finally {
      loading.value = false;
    }
  }, 300);

  watch(query, search);

  return { query, results, loading, search };
};

// Minimal debounce helper (replaces VueUse dependency)
function useDebounceFn<T extends (...args: any[]) => any>(
  fn: T,
  delay: number,
) {
  let timer: ReturnType<typeof setTimeout>;
  return (...args: Parameters<T>) => {
    clearTimeout(timer);
    timer = setTimeout(() => fn(...args), delay);
  };
}
