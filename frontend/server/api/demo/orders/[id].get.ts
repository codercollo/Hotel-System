// server/api/demo/orders/[id].get.ts
import { defineEventHandler, getRouterParam } from "h3";

const now = new Date().toISOString();
const yesterday = new Date(Date.now() - 86400000).toISOString();

const DEMO_ORDERS: Record<string, object> = {
  "01JQVG9V83RYCC7G58AGPMZ7AZ": {
    id: "01JQVG9V83RYCC7G58AGPMZ7AZ",
    user_id: "01JDEMOUSER001ADMINXXXXXX",
    status: "confirmed",
    total: 45000,
    currency: "KES",
    notes: "Late check-in requested",
    metadata: { check_in: "2026-05-10", check_out: "2026-05-13", guests: 2 },
    items: [
      {
        id: "01JITEM001EXECSUITE000001",
        order_id: "01JQVG9V83RYCC7G58AGPMZ7AZ",
        item_id: "01JROOM001EXECUTIVESUITE01",
        name: "Executive Suite",
        price: 45000,
        quantity: 1,
        subtotal: 45000,
      },
    ],
    created_at: yesterday,
    updated_at: now,
  },
  "01JQVG9V83RYCC7G58AGPMZ7BB": {
    id: "01JQVG9V83RYCC7G58AGPMZ7BB",
    user_id: "01JDEMOUSER001ADMINXXXXXX",
    status: "pending",
    total: 32000,
    currency: "KES",
    notes: "",
    metadata: { check_in: "2026-05-20", check_out: "2026-05-22", guests: 4 },
    items: [
      {
        id: "01JITEM002FAMILYROOM00001",
        order_id: "01JQVG9V83RYCC7G58AGPMZ7BB",
        item_id: "01JROOM005FAMILYCONNECTED05",
        name: "Family Connected Room",
        price: 32000,
        quantity: 1,
        subtotal: 32000,
      },
    ],
    created_at: now,
    updated_at: now,
  },
};

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 100));

  const id = getRouterParam(event, "id");
  const order = id ? DEMO_ORDERS[id] : null;

  // Unknown IDs = freshly created bookings from bookings.post.ts
  // Return a generic confirmed shell so the UI never breaks
  return {
    success: true,
    data: order ?? {
      id,
      user_id: "01JDEMOUSER001ADMINXXXXXX",
      status: "confirmed",
      total: 18500,
      currency: "KES",
      notes: "",
      metadata: {},
      items: [],
      created_at: now,
      updated_at: now,
    },
  };
});
