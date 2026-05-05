// app/server/api/demo/payments/mpesa.post.ts
import { defineEventHandler, readBody } from "h3";

function ulid(): string {
  const chars = "0123456789ABCDEFGHJKMNPQRSTVWXYZ";
  let id = "01J";
  for (let i = 0; i < 23; i++) id += chars[Math.floor(Math.random() * 32)];
  return id;
}

export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  // Simulate STK push network latency
  await new Promise((r) => setTimeout(r, 800));

  const paymentId = ulid();
  const checkoutCode = `WS${Math.floor(Math.random() * 9000000000 + 1000000000)}`;
  const phone = body.phone ?? "254712345678";
  const now = new Date().toISOString();

  return {
    success: true,
    data: {
      payment_id: paymentId,
      order_id: body.order_id,
      provider: "mpesa",
      provider_ref: checkoutCode,
      checkout_code: checkoutCode,
      checkout_url: null,
      status: "processing",
      amount: body.amount,
      currency: "KES",
      phone,
      message: `STK Push sent to ${phone}. Enter your M-Pesa PIN to complete payment.`,
      created_at: now,
      updated_at: now,
    },
  };
});
