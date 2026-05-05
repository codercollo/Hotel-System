// app/server/api/demo/auth/login.post.ts
import { defineEventHandler, readBody, createError } from "h3";

const DEMO_USERS: Record<
  string,
  { password: string; user_id: string; name: string; role: string }
> = {
  "admin@hotel.com": {
    password: "Admin123!@#",
    user_id: "01JDEMOUSER001ADMINXXXXXX",
    name: "Admin User",
    role: "admin",
  },
  "guest@hotel.com": {
    password: "Guest123!",
    user_id: "01JDEMOUSER002GUESTXXXXX",
    name: "Jane Kamau",
    role: "guest",
  },
};

// Simple base64 without Buffer — works in any Nitro/edge runtime
function b64(str: string): string {
  return btoa(unescape(encodeURIComponent(str)));
}

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 400));

  const body = await readBody(event);
  const { email, password } = body as { email: string; password: string };

  const user = DEMO_USERS[email?.toLowerCase()];

  if (!user || user.password !== password) {
    throw createError({
      statusCode: 401,
      data: {
        success: false,
        error: {
          code: "INVALID_CREDENTIALS",
          message: "Invalid email or password",
        },
      },
    });
  }

  const now = new Date();
  const expiresAt = new Date(now.getTime() + 15 * 60 * 1000).toISOString();

  const payload = b64(
    JSON.stringify({ user_id: user.user_id, role: user.role, exp: expiresAt }),
  );

  return {
    success: true,
    data: {
      access_token: `demo.access.${payload}`,
      refresh_token: `demo.refresh.${user.user_id}.${Date.now()}`,
      expires_at: expiresAt,
      user: {
        user_id: user.user_id,
        email,
        name: user.name,
        role: user.role,
      },
    },
  };
});
