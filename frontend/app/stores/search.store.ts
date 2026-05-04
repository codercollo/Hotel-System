import { defineStore } from "pinia";

export const useSearchStore = defineStore("search", {
  state: () => ({
    query: "",
    results: [] as any[],
    loading: false,
  }),
  actions: {
    setQuery(q: string) {
      this.query = q;
    },
    setResults(r: any[]) {
      this.results = r;
    },
    setLoading(v: boolean) {
      this.loading = v;
    },
    clear() {
      this.query = "";
      this.results = [];
    },
  },
});
