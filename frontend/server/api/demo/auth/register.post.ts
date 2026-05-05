// app/server/api/demo/auth/register.post.ts
import { defineEventHandler, readBody } from "h3";

function ulid(): string {
  const chars = "0123456789ABCDEFGHJKMNPQRSTVWXYZ";
  let id = "01J";
  for (let i = 0; i < 23; i++) id += chars[Math.floor(Math.random() * 32)];
  return id;
}

function b64(str: string): string {
  return btoa(unescape(encodeURIComponent(str)));
}

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 450));

  const body = await readBody(event);
  const { name, email } = body as { name: string; email: string };

  const userId = ulid();
  const now = new Date();
  const expiresAt = new Date(now.getTime() + 15 * 60 * 1000).toISOString();

  const payload = b64(
    JSON.stringify({ user_id: userId, role: "guest", exp: expiresAt }),
  );

  return {
    success: true,
    data: {
      access_token: `demo.access.${payload}`,
      refresh_token: `demo.refresh.${userId}.${Date.now()}`,
      expires_at: expiresAt,
      user: {
        user_id: userId,
        email,
        name,
        role: "guest",
      },
    },
  };
});
