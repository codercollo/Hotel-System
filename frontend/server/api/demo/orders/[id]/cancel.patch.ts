// server/api/demo/orders/[id]/cancel.patch.ts
// Place at: server/api/demo/orders/[id]/cancel.patch.ts
import { defineEventHandler, getRouterParam } from "h3";

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 180));
  const id = getRouterParam(event, "id") ?? "";
  const now = new Date().toISOString();
  return {
    success: true,
    data: {
      id,
      status: "cancelled",
      updated_at: now,
    },
  };
});
