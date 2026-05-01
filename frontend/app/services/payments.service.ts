export const paymentsService = {
  initiate: (orderId: string, provider: string) =>
    useApi().post("/api/v1/payments/initiate", { order_id: orderId, provider }),
  getStatus: (paymentId: string) =>
    useApi().get(`/api/v1/payments/${paymentId}`),
};
