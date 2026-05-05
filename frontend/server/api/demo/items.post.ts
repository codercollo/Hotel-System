// server/api/demo/items.post.ts
// Handles POST /api/demo/items (create room)
import { defineEventHandler, readBody } from "h3";

function ulid(): string {
  const chars = "0123456789ABCDEFGHJKMNPQRSTVWXYZ";
  let id = "01J";
  for (let i = 0; i < 23; i++) id += chars[Math.floor(Math.random() * 32)];
  return id;
}

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 250));
  const body = await readBody(event);
  const now = new Date().toISOString();
  return {
    success: true,
    data: {
      id: ulid(),
      name: body.name ?? "New Room",
      description: body.description ?? "",
      price: body.price ?? 0,
      currency: body.currency ?? "KES",
      status: body.status ?? "active",
      stock: body.stock ?? 1,
      images: body.images ?? [],
      metadata: body.metadata ?? {},
      created_at: now,
      updated_at: now,
    },
  };
});
