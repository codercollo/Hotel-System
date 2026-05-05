// server/api/demo/items/[id].delete.ts
// Place at: server/api/demo/rooms/[id].delete.ts
import { defineEventHandler } from "h3";

export default defineEventHandler(async () => {
  await new Promise((r) => setTimeout(r, 150));
  // 204 — just return empty success
  return { success: true, data: null };
});
