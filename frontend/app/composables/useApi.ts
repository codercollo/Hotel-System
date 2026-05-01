// useApi provides a typed fetch wrapper that prepends the API base URL,
// attaches the auth token from the auth store, and handles error envelopes.
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

    // Attach bearer token if present in localStorage (Phase 2 will use Pinia store)
    const token = import.meta.client
      ? localStorage.getItem("access_token")
      : null;
    if (token) {
      fetchOpts.headers = {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
        ...fetchOpts.headers,
      };
    } else {
      fetchOpts.headers = {
        "Content-Type": "application/json",
        ...fetchOpts.headers,
      };
    }

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

    const body = await res.json();
    return body.data as T;
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
