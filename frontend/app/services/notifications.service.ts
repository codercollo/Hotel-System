import type { Notification } from "~/composables/useNotifications";

export const notificationsService = {
  list: (p?: Record<string, string>) =>
    useApi().get<Notification[]>("/api/v1/notifications", p),
  markRead: (id: string) =>
    useApi().patch(`/api/v1/notifications/${id}/read`, {}),
  markAllRead: () => useApi().patch("/api/v1/notifications/read-all", {}),
};
