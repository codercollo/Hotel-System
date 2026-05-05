// server/api/demo/_mockData.ts
// Shared in-memory store for demo — mutations (patch/delete) persist for the session

export const DEMO_USERS = [
  {
    id: "01JDEMOUSER001ADMINXXXXXX",
    email: "admin@hotel.com",
    name: "Admin User",
    role: "admin",
    is_active: true,
    metadata: {},
    created_at: "2025-01-10T08:00:00Z",
    updated_at: "2026-04-01T12:00:00Z",
  },
  {
    id: "01JDEMOUSER002GUESTXXXXX",
    email: "guest@hotel.com",
    name: "Jane Kamau",
    role: "guest",
    is_active: true,
    metadata: {},
    created_at: "2025-03-15T10:30:00Z",
    updated_at: "2026-04-10T09:00:00Z",
  },
  {
    id: "01JDEMOUSER003STAFFXXXXX",
    email: "staff@hotel.com",
    name: "Peter Odhiambo",
    role: "staff",
    is_active: true,
    metadata: {},
    created_at: "2025-06-01T08:00:00Z",
    updated_at: "2026-03-20T14:00:00Z",
  },
  {
    id: "01JDEMOUSER004GUESTXXXXX",
    email: "amina.wanjiku@gmail.com",
    name: "Amina Wanjiku",
    role: "guest",
    is_active: true,
    metadata: {},
    created_at: "2026-01-22T11:00:00Z",
    updated_at: "2026-04-28T16:45:00Z",
  },
  {
    id: "01JDEMOUSER005GUESTXXXXX",
    email: "tom.njoroge@outlook.com",
    name: "Tom Njoroge",
    role: "guest",
    is_active: false,
    metadata: {},
    created_at: "2026-02-14T09:15:00Z",
    updated_at: "2026-04-01T08:00:00Z",
  },
];

export const DEMO_STATS = {
  total_users: 5,
  total_items: 5,
  total_orders: 12,
  revenue_total: 487500, // KES in cents equivalent
};

export const DEMO_ROLES = [
  {
    id: "01JROLE001ADMINXXXXXXXX",
    name: "admin",
    description:
      "Full platform access — manage rooms, users, orders, and settings.",
    permissions: [
      "users:read",
      "users:write",
      "users:delete",
      "items:read",
      "items:write",
      "items:delete",
      "orders:read",
      "orders:write",
      "payments:read",
      "admin:stats",
      "admin:settings",
    ],
  },
  {
    id: "01JROLE002STAFFXXXXXXXX",
    name: "staff",
    description:
      "Front-desk staff — view and manage bookings and room availability.",
    permissions: ["items:read", "orders:read", "orders:write", "payments:read"],
  },
  {
    id: "01JROLE003GUESTXXXXXXXX",
    name: "guest",
    description:
      "Registered hotel guest — browse rooms and manage own bookings.",
    permissions: [
      "items:read",
      "orders:read:own",
      "orders:write:own",
      "payments:write:own",
    ],
  },
];
