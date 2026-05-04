// useApi provides a typed fetch wrapper that prepends the API base URL,
// attaches the auth token, and handles error envelopes.

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

export const useApi = () => {
  const config = useRuntimeConfig();
  const base = config.public.apiBase as string;

  const request = async <T>(
    path: string,
    opts: RequestInit & { params?: Record<string, string> } = {},
  ): Promise<T> => {
    const { params, ...fetchOpts } = opts;
    let url = `${base}${path}`;
    if (params) {
      const q = new URLSearchParams(params);
      url += `?${q.toString()}`;
    }

    // Attach bearer token if present (Phase 2 will migrate to Pinia auth store)
    const token = import.meta.client
      ? localStorage.getItem("access_token")
      : null;

    fetchOpts.headers = {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...fetchOpts.headers,
    };

    const res = await fetch(url, fetchOpts);

    if (!res.ok) {
      const body = await res.json().catch(() => ({}));
      throw new ApiError(
        res.status,
        body?.error?.code ?? "UNKNOWN",
        body?.error?.message ?? "Request failed",
        body?.error?.detail,
      );
    }

    // Handle 204 No Content
    if (res.status === 204) return undefined as T;

    const body = await res.json();

    // Support both envelope { data: [...] } and bare array/object responses
    return (body?.data !== undefined ? body.data : body) as T;
  };

  return {
    get: <T>(path: string, params?: Record<string, string>) =>
      request<T>(path, { method: "GET", params }),
    post: <T>(path: string, body: unknown) =>
      request<T>(path, { method: "POST", body: JSON.stringify(body) }),
    put: <T>(path: string, body: unknown) =>
      request<T>(path, { method: "PUT", body: JSON.stringify(body) }),
    patch: <T>(path: string, body: unknown) =>
      request<T>(path, { method: "PATCH", body: JSON.stringify(body) }),
    del: <T>(path: string) => request<T>(path, { method: "DELETE" }),
  };
};
