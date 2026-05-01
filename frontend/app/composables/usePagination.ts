export const usePagination = (initialLimit = 20) => {
  const page = ref(1);
  const limit = ref(initialLimit);

  const reset = () => {
    page.value = 1;
  };
  const next = () => {
    page.value++;
  };
  const prev = () => {
    if (page.value > 1) page.value--;
  };
  const goto = (p: number) => {
    page.value = p;
  };

  const offset = computed(() => (page.value - 1) * limit.value);

  return { page, limit, offset, reset, next, prev, goto };
};
