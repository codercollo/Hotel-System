// app/server/api/demo/payments/paystack.post.ts
import { defineEventHandler, readBody } from "h3";

function ulid(): string {
  const chars = "0123456789ABCDEFGHJKMNPQRSTVWXYZ";
  let id = "01J";
  for (let i = 0; i < 23; i++) id += chars[Math.floor(Math.random() * 32)];
  return id;
}

export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  // Simulate card processing latency
  await new Promise((r) => setTimeout(r, 600));

  const paymentId = ulid();
  const reference = `PAY-${Date.now()}-${Math.random().toString(36).slice(2, 7).toUpperCase()}`;
  const now = new Date().toISOString();

  // Simulate a Paystack-style checkout URL
  const checkoutUrl = `https://checkout.paystack.com/demo/${reference}`;

  return {
    success: true,
    data: {
      payment_id: paymentId,
      order_id: body.order_id,
      provider: "paystack",
      provider_ref: reference,
      checkout_url: checkoutUrl,
      checkout_code: null,
      status: "processing",
      amount: body.amount,
      currency: body.currency ?? "KES",
      message: "Redirecting to secure Paystack checkout...",
      created_at: now,
      updated_at: now,
    },
  };
});
