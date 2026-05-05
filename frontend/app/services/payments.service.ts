// export const paymentsService = {
//   initiate: (orderId: string, provider: string) =>
//     useApi().post("/api/v1/payments/initiate", { order_id: orderId, provider }),
//   getStatus: (paymentId: string) =>
//     useApi().get(`/api/v1/payments/${paymentId}`),
// };

// app/services/payments.service.ts
import type { Payment, InitiateResponse } from "~/types/payment.types";

export const paymentsService = {
  /**
   * Initiate a payment.
   * In demo mode, routes to the correct demo endpoint per provider.
   * In real mode, calls the Go backend.
   */
  initiate: (
    orderId: string,
    provider: "mpesa" | "paystack" | string,
    options?: { phone?: string; amount?: number; currency?: string },
  ): Promise<InitiateResponse> => {
    const config = useRuntimeConfig();
    const isDemo = config.public.mode === "demo";

    if (isDemo) {
      const endpoint =
        provider === "paystack"
          ? "/api/demo/payments/paystack"
          : "/api/demo/payments/mpesa";

      return useApi().post<InitiateResponse>(endpoint, {
        order_id: orderId,
        provider,
        phone: options?.phone,
        amount: options?.amount,
        currency: options?.currency ?? "KES",
      });
    }

    // Real mode — Go backend handles provider routing
    return useApi().post<InitiateResponse>("/api/v1/payments/initiate", {
      order_id: orderId,
      provider,
      phone: options?.phone,
    });
  },

  /**
   * Poll payment status.
   * In demo mode, hits the verify stub which auto-completes after ~6s.
   */
  getStatus: (paymentId: string): Promise<Payment> => {
    const config = useRuntimeConfig();
    const isDemo = config.public.mode === "demo";

    if (isDemo) {
      return useApi().get<Payment>(
        `/api/demo/payments/verify?payment_id=${paymentId}`,
      );
    }

    return useApi().get<Payment>(`/api/v1/payments/${paymentId}`);
  },

  /**
   * List all payments for an order.
   */
  listForOrder: (orderId: string): Promise<Payment[]> =>
    useApi().get<Payment[]>(`/api/v1/orders/${orderId}/payments`),
};
