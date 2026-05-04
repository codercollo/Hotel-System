import type { Payment, InitiateResponse } from "~/types/payment.types";

export const usePayments = () => {
  const api = useApi();
  const loading = ref(false);
  const error = ref<string | null>(null);
  const payment = ref<Payment | null>(null);
  const pollInterval = ref<ReturnType<typeof setInterval> | null>(null);

  const initiate = async (
    orderId: string,
    provider: string,
    phone?: string,
  ): Promise<InitiateResponse> => {
    loading.value = true;
    error.value = null;
    try {
      const res = await api.post<InitiateResponse>(
        "/api/v1/payments/initiate",
        {
          order_id: orderId,
          provider,
          ...(phone ? { phone } : {}),
        },
      );
      return res;
    } catch (e: any) {
      error.value = e.message;
      throw e;
    } finally {
      loading.value = false;
    }
  };

  const getStatus = async (paymentId: string) => {
    try {
      payment.value = await api.get<Payment>(`/api/v1/payments/${paymentId}`);
      return payment.value;
    } catch (e: any) {
      error.value = e.message;
      return null;
    }
  };

  // Poll every 3s until terminal status or maxAttempts reached
  const pollUntilDone = (
    paymentId: string,
    onDone: (p: Payment) => void,
    maxAttempts = 20,
  ) => {
    let attempts = 0;
    pollInterval.value = setInterval(async () => {
      attempts++;
      const p = await getStatus(paymentId);
      if (!p) return;
      if (
        ["completed", "failed", "refunded"].includes(p.status) ||
        attempts >= maxAttempts
      ) {
        stopPolling();
        onDone(p);
      }
    }, 3000);
  };

  const stopPolling = () => {
    if (pollInterval.value) {
      clearInterval(pollInterval.value);
      pollInterval.value = null;
    }
  };

  onUnmounted(() => stopPolling());

  return {
    loading,
    error,
    payment,
    initiate,
    getStatus,
    pollUntilDone,
    stopPolling,
  };
};
