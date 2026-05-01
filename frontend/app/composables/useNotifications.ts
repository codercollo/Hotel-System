// useNotifications — Phase 6 wires this to the notifications API + WebSocket
export interface Notification {
  id: string;
  type: string;
  title: string;
  body: string;
  is_read: boolean;
  created_at: string;
}

export const useNotifications = () => {
  const api = useApi();
  const notifications = ref<Notification[]>([]);
  const loading = ref(false);
  const unreadCount = computed(
    () => notifications.value.filter((n) => !n.is_read).length,
  );

  const list = async () => {
    loading.value = true;
    try {
      notifications.value = await api.get<Notification[]>(
        "/api/v1/notifications",
      );
    } catch {
      /* silently fail */
    } finally {
      loading.value = false;
    }
  };

  const markRead = async (id: string) => {
    await api.patch(`/api/v1/notifications/${id}/read`, {});
    const n = notifications.value.find((n) => n.id === id);
    if (n) n.is_read = true;
  };

  const markAllRead = async () => {
    await api.patch("/api/v1/notifications/read-all", {});
    notifications.value.forEach((n) => (n.is_read = true));
  };

  // Push new notification (called by WebSocket handler in Phase 9)
  const push = (n: Notification) => {
    notifications.value.unshift(n);
  };

  return {
    notifications,
    loading,
    unreadCount,
    list,
    markRead,
    markAllRead,
    push,
  };
};
