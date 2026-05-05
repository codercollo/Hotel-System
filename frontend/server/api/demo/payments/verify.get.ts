// app/server/api/demo/payments/verify.get.ts
// Simulates polling for payment status.
// Returns "processing" for the first ~6 seconds, then "completed".
// This makes the UI polling loop feel realistic during the pitch demo.
import { defineEventHandler, getQuery } from "h3";

// In-memory store keyed by payment_id → first-seen timestamp
const firstSeen: Record<string, number> = {};

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 250));

  const { payment_id, order_id } = getQuery(event) as {
    payment_id?: string;
    order_id?: string;
  };

  const key = payment_id ?? order_id ?? "unknown";
  const now = Date.now();

  if (!firstSeen[key]) firstSeen[key] = now;

  const elapsedMs = now - firstSeen[key];
  // Simulate "pending → completed" transition after 6 seconds of polling
  const status = elapsedMs > 6000 ? "completed" : "processing";

  const ts = new Date().toISOString();

  return {
    success: true,
    data: {
      payment_id: key,
      order_id: order_id ?? null,
      provider: "mpesa",
      status,
      amount: 18500,
      currency: "KES",
      provider_ref: `WS${Math.floor(Math.random() * 9000000000 + 1000000000)}`,
      message:
        status === "completed"
          ? "Payment confirmed successfully."
          : "Awaiting payment confirmation...",
      created_at: ts,
      updated_at: ts,
    },
  };
});
