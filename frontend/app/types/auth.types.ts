export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface AuthUser {
  user_id: string;
  email: string;
  role: string;
  name?: string;
}

export interface AuthResponse {
  access_token: string;
  refresh_token: string;
  expires_at: string;
  user?: AuthUser;
}

export interface RefreshRequest {
  refresh_token: string;
}
