import type {
  LoginRequest,
  RegisterRequest,
  AuthResponse,
  AuthUser,
  RefreshRequest,
} from "~/types/auth.types";

export const authService = {
  login: (body: LoginRequest) =>
    useApi().post<AuthResponse>("/api/v1/auth/login", body),

  register: (body: RegisterRequest) =>
    useApi().post<AuthResponse>("/api/v1/auth/register", body),

  refresh: (body: RefreshRequest) =>
    useApi().post<AuthResponse>("/api/v1/auth/refresh", body),

  logout: (refreshToken: string) =>
    useApi().post<void>("/api/v1/auth/logout", { refresh_token: refreshToken }),

  me: () => useApi().get<AuthUser>("/api/v1/auth/me"),
};
