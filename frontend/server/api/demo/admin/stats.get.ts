// server/api/demo/admin/stats.get.ts
import { defineEventHandler } from "h3";

export default defineEventHandler(async () => {
  await new Promise((r) => setTimeout(r, 200));
  return {
    success: true,
    data: {
      total_users: 5,
      total_items: 5,
      total_orders: 12,
      revenue_total: 487500,
    },
  };
});
