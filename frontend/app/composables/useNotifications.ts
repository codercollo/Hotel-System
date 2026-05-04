import type { Notification } from "~/types/notification.types";

export const useNotifications = () => {
  const api = useApi();
  const store = useNotificationStore();
  const loading = ref(false);
  const error = ref<string | null>(null);

  const fetch = async () => {
    loading.value = true;
    error.value = null;
    try {
      const items = await api.get<Notification[]>("/api/v1/notifications");
      store.set(items);
    } catch (e: any) {
      error.value = e.message;
    } finally {
      loading.value = false;
    }
  };

  const markRead = async (id: string) => {
    store.markRead(id);
    await api.patch(`/api/v1/notifications/${id}/read`, {});
  };

  const markAllRead = async () => {
    store.markAllRead();
    await api.patch("/api/v1/notifications/read-all", {});
  };

  return { loading, error, fetch, markRead, markAllRead };
};
