// server/api/demo/users/index.get.ts
import { defineEventHandler } from "h3";

const USERS = [
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

export default defineEventHandler(async () => {
  await new Promise((r) => setTimeout(r, 150));
  return {
    success: true,
    data: USERS,
    meta: { total: USERS.length, limit: 100, offset: 0 },
  };
});
