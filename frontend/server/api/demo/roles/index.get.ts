// server/api/demo/roles/index.get.ts
import { defineEventHandler } from "h3";

export default defineEventHandler(async () => {
  await new Promise((r) => setTimeout(r, 120));
  return {
    success: true,
    data: [
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
        permissions: [
          "items:read",
          "orders:read",
          "orders:write",
          "payments:read",
        ],
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
    ],
  };
});
