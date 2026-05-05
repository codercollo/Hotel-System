// server/api/demo/items/[id].patch.ts
// Place at: server/api/demo/items/[id].patch.ts  (inside rooms folder or items folder)
// Since items map to rooms in demo, we handle patch on the rooms route
import { defineEventHandler, getRouterParam, readBody } from "h3";

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 200));
  const id = getRouterParam(event, "id") ?? "";
  const body = await readBody(event);
  const now = new Date().toISOString();
  return {
    success: true,
    data: {
      id,
      ...body,
      updated_at: now,
    },
  };
});
