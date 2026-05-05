// // useApi provides a typed fetch wrapper that prepends the API base URL,
// // attaches the auth token, and handles error envelopes.

// export class ApiError extends Error {
//   constructor(
//     public status: number,
//     public code: string,
//     message: string,
//     public detail?: unknown,
//   ) {
//     super(message);
//     this.name = "ApiError";
//   }
// }

// export const useApi = () => {
//   const config = useRuntimeConfig();
//   const base = config.public.apiBase as string;

//   const request = async <T>(
//     path: string,
//     opts: RequestInit & { params?: Record<string, string> } = {},
//   ): Promise<T> => {
//     const { params, ...fetchOpts } = opts;
//     let url = `${base}${path}`;
//     if (params) {
//       const q = new URLSearchParams(params);
//       url += `?${q.toString()}`;
//     }

//     // Attach bearer token if present (Phase 2 will migrate to Pinia auth store)
//     const token = import.meta.client
//       ? localStorage.getItem("access_token")
//       : null;

//     fetchOpts.headers = {
//       "Content-Type": "application/json",
//       ...(token ? { Authorization: `Bearer ${token}` } : {}),
//       ...fetchOpts.headers,
//     };

//     const res = await fetch(url, fetchOpts);

//     if (!res.ok) {
//       const body = await res.json().catch(() => ({}));
//       throw new ApiError(
//         res.status,
//         body?.error?.code ?? "UNKNOWN",
//         body?.error?.message ?? "Request failed",
//         body?.error?.detail,
//       );
//     }

//     // Handle 204 No Content
//     if (res.status === 204) return undefined as T;

//     const body = await res.json();

//     // Support both envelope { data: [...] } and bare array/object responses
//     return (body?.data !== undefined ? body.data : body) as T;
//   };

//   return {
//     get: <T>(path: string, params?: Record<string, string>) =>
//       request<T>(path, { method: "GET", params }),
//     post: <T>(path: string, body: unknown) =>
//       request<T>(path, { method: "POST", body: JSON.stringify(body) }),
//     put: <T>(path: string, body: unknown) =>
//       request<T>(path, { method: "PUT", body: JSON.stringify(body) }),
//     patch: <T>(path: string, body: unknown) =>
//       request<T>(path, { method: "PATCH", body: JSON.stringify(body) }),
//     del: <T>(path: string) => request<T>(path, { method: "DELETE" }),
//   };
// };

// app/composables/useApi.ts
export class ApiError extends Error {
  constructor(
    public status: number,
    public code: string,
    message: string,
    public detail?: unknown,
  ) {
    super(message);
    this.name = "ApiError";
  }
}

function demoPatch(path: string, method = "GET"): string {
  // ── Auth ────────────────────────────────────────────────────────────────────
  if (path === "/api/v1/auth/login") return "/api/demo/auth/login";
  if (path === "/api/v1/auth/register") return "/api/demo/auth/register";
  if (path === "/api/v1/auth/refresh") return "/api/demo/auth/login";
  if (path === "/api/v1/auth/me") return "/api/demo/auth/me";

  // ── Admin ───────────────────────────────────────────────────────────────────
  if (path === "/api/v1/admin/stats") return "/api/demo/admin/stats";

  // ── Users ───────────────────────────────────────────────────────────────────
  if (path.startsWith("/api/v1/users")) return "/api/demo/users";

  // ── Roles ───────────────────────────────────────────────────────────────────
  if (path.startsWith("/api/v1/roles")) return "/api/demo/roles";

  // ── Items → Rooms ────────────────────────────────────────────────────────────
  // PATCH/DELETE /api/v1/items/:id → /api/demo/rooms/:id
  const itemId = path.match(/^\/api\/v1\/items\/([^/]+)$/)?.[1];
  if (itemId) return `/api/demo/rooms/${itemId}`;
  // POST /api/v1/items → /api/demo/rooms (index.post if you add it, else items.post)
  if (path === "/api/v1/items" && method === "POST") return "/api/demo/items";
  if (path === "/api/v1/items") return "/api/demo/rooms";

  // ── Orders ───────────────────────────────────────────────────────────────────
  // PATCH /api/v1/orders/:id/status
  const orderStatusMatch = path.match(/^\/api\/v1\/orders\/([^/]+)\/status$/);
  if (orderStatusMatch) return `/api/demo/orders/${orderStatusMatch[1]}/status`;

  // PATCH /api/v1/orders/:id/cancel
  const orderCancelMatch = path.match(/^\/api\/v1\/orders\/([^/]+)\/cancel$/);
  if (orderCancelMatch) return `/api/demo/orders/${orderCancelMatch[1]}/cancel`;

  // GET /api/v1/orders/:id
  const orderId = path.match(/^\/api\/v1\/orders\/([^/]+)$/)?.[1];
  if (orderId) return `/api/demo/orders/${orderId}`;

  // GET/POST /api/v1/orders
  if (path === "/api/v1/orders" && method === "POST")
    return "/api/demo/bookings";
  if (path === "/api/v1/orders") return "/api/demo/orders";

  // ── Payments ─────────────────────────────────────────────────────────────────
  if (path === "/api/v1/payments/initiate") return "/api/demo/payments/mpesa";
  const payId = path.match(/^\/api\/v1\/payments\/([^/]+)$/)?.[1];
  if (payId) return `/api/demo/payments/verify?payment_id=${payId}`;

  return path;
}

export const useApi = () => {
  const config = useRuntimeConfig();
  const isDemo = config.public.mode === "demo";

  const getBase = (): string => {
    if (!isDemo) return config.public.apiBase as string;
    if (import.meta.server) {
      const reqUrl = useRequestURL();
      return `${reqUrl.protocol}//${reqUrl.host}`;
    }
    return "";
  };

  const request = async <T>(
    path: string,
    opts: {
      method?: string;
      body?: unknown;
      params?: Record<string, string>;
    } = {},
  ): Promise<T> => {
    const { params, body, method = "GET" } = opts;
    const resolvedPath = isDemo ? demoPatch(path, method) : path;
    const base = getBase();
    let url = `${base}${resolvedPath}`;

    if (params) {
      const q = new URLSearchParams(params);
      url += (url.includes("?") ? "&" : "?") + q.toString();
    }

    const token = import.meta.client
      ? localStorage.getItem("access_token")
      : null;

    try {
      const response = await $fetch<{ data?: T; success?: boolean } | T>(url, {
        method: method as "GET" | "POST" | "PUT" | "PATCH" | "DELETE",
        headers: { ...(token ? { Authorization: `Bearer ${token}` } : {}) },
        body: body ?? undefined,
      });

      if (
        response !== null &&
        typeof response === "object" &&
        "data" in (response as object)
      ) {
        return (response as { data: T }).data;
      }
      return response as T;
    } catch (err: unknown) {
      const e = err as {
        status?: number;
        data?: {
          error?: { code?: string; message?: string; detail?: unknown };
        };
        message?: string;
      };
      throw new ApiError(
        e.status ?? 500,
        e.data?.error?.code ?? "UNKNOWN",
        e.data?.error?.message ?? e.message ?? "Request failed",
        e.data?.error?.detail,
      );
    }
  };

  return {
    get: <T>(path: string, params?: Record<string, string>) =>
      request<T>(path, { method: "GET", params }),
    post: <T>(path: string, body: unknown) =>
      request<T>(path, { method: "POST", body }),
    put: <T>(path: string, body: unknown) =>
      request<T>(path, { method: "PUT", body }),
    patch: <T>(path: string, body: unknown) =>
      request<T>(path, { method: "PATCH", body }),
    del: <T>(path: string) => request<T>(path, { method: "DELETE" }),
  };
};
