// usePayments — Phase 5 wires this to the payments API
export const usePayments = () => {
  const api = useApi();
  const loading = ref(false);
  const error = ref<string | null>(null);

  const initiate = async (orderId: string, provider: string) => {
    loading.value = true;
    error.value = null;
    try {
      return await api.post("/api/v1/payments/initiate", {
        order_id: orderId,
        provider,
      });
    } catch (e: any) {
      error.value = e.message;
      throw e;
    } finally {
      loading.value = false;
    }
  };

  return { loading, error, initiate };
};
