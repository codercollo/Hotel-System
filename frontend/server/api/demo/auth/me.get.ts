// server/api/demo/auth/me.get.ts
import { defineEventHandler, getHeader } from "h3";

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 80));

  // Parse the demo token to extract user info
  const auth = getHeader(event, "authorization") ?? "";
  const token = auth.replace("Bearer ", "");

  try {
    const payload = token.split(".")[2]; // demo.access.<b64payload>
    if (payload) {
      const decoded = JSON.parse(atob(payload));
      return {
        success: true,
        data: {
          user_id: decoded.user_id,
          email: decoded.user_id?.includes("ADMIN")
            ? "admin@hotel.com"
            : "guest@hotel.com",
          name: decoded.user_id?.includes("ADMIN")
            ? "Admin User"
            : "Jane Kamau",
          role: decoded.role ?? "guest",
        },
      };
    }
  } catch {
    // fall through to default
  }

  // Fallback — return a default guest user
  return {
    success: true,
    data: {
      user_id: "01JDEMOUSER001ADMINXXXXXX",
      email: "admin@hotel.com",
      name: "Admin User",
      role: "admin",
    },
  };
});
