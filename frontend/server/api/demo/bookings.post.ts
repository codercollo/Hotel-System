// app/server/api/demo/bookings.post.ts
import { defineEventHandler, readBody } from "h3";

function ulid(): string {
  const chars = "0123456789ABCDEFGHJKMNPQRSTVWXYZ";
  const t = Date.now();
  let id = "";
  let time = t;
  for (let i = 9; i >= 0; i--) {
    id = chars[time % 32] + id;
    time = Math.floor(time / 32);
  }
  for (let i = 0; i < 16; i++) id += chars[Math.floor(Math.random() * 32)];
  return "01J" + id.slice(3);
}

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 300));

  const body = await readBody(event);

  const orderId = ulid();
  const now = new Date().toISOString();

  // Build order items from request
  const items = (body.items ?? []).map(
    (item: {
      item_id: string;
      quantity: number;
      name?: string;
      price?: number;
    }) => ({
      id: ulid(),
      order_id: orderId,
      item_id: item.item_id,
      name: item.name ?? "Room Booking",
      price: item.price ?? 18500,
      quantity: item.quantity ?? 1,
      subtotal: (item.price ?? 18500) * (item.quantity ?? 1),
    }),
  );

  const total = items.reduce(
    (sum: number, i: { subtotal: number }) => sum + i.subtotal,
    0,
  );

  const order = {
    id: orderId,
    user_id: "demo-user-01JDEMOUSER001",
    status: "pending",
    total,
    currency: "KES",
    notes: body.notes ?? "",
    metadata: body.metadata ?? {},
    items,
    created_at: now,
    updated_at: now,
  };

  return { success: true, data: order };
});
