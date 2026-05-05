// server/api/demo/orders/[id]/status.patch.ts
// Place at: server/api/demo/orders/[id]/status.patch.ts
import { defineEventHandler, getRouterParam, readBody } from "h3";

// In-memory status overrides for the demo session
const statusOverrides: Record<string, string> = {};

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 180));
  const id = getRouterParam(event, "id") ?? "";
  const body = (await readBody(event)) as { status: string };
  statusOverrides[id] = body.status;
  const now = new Date().toISOString();
  return {
    success: true,
    data: {
      id,
      status: body.status,
      updated_at: now,
    },
  };
});
