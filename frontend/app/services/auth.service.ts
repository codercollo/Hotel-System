import type {
  LoginRequest,
  RegisterRequest,
  AuthResponse,
} from "~/types/auth.types";

export const authService = {
  login: (body: LoginRequest) =>
    useApi().post<AuthResponse>("/api/v1/auth/login", body),
  register: (body: RegisterRequest) =>
    useApi().post<AuthResponse>("/api/v1/auth/register", body),
  refresh: (refreshToken: string) =>
    useApi().post<AuthResponse>("/api/v1/auth/refresh", {
      refresh_token: refreshToken,
    }),
  logout: () => useApi().post("/api/v1/auth/logout", {}),
  me: () => useApi().get<AuthResponse["user"]>("/api/v1/auth/me"),
};
