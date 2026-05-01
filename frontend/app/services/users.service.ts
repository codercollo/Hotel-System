import type { User } from "~/types/user.types";

export const usersService = {
  list: (p?: Record<string, string>) =>
    useApi().get<User[]>("/api/v1/users", p),
  get: (id: string) => useApi().get<User>(`/api/v1/users/${id}`),
  update: (id: string, body: Partial<User>) =>
    useApi().patch<User>(`/api/v1/users/${id}`, body),
  delete: (id: string) => useApi().del<void>(`/api/v1/users/${id}`),
};
