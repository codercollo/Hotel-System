import type { Item, CreateItemRequest } from "~/types/item.types";

export const itemsService = {
  list: (p?: Record<string, string>) =>
    useApi().get<Item[]>("/api/v1/items", p),
  get: (id: string) => useApi().get<Item>(`/api/v1/items/${id}`),
  create: (body: CreateItemRequest) =>
    useApi().post<Item>("/api/v1/items", body),
  update: (id: string, body: Partial<CreateItemRequest>) =>
    useApi().patch<Item>(`/api/v1/items/${id}`, body),
  delete: (id: string) => useApi().del<void>(`/api/v1/items/${id}`),
};
