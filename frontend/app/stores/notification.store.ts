import { defineStore } from "pinia";
import type { Notification } from "~/types/notification.types";

interface NotificationState {
  items: Notification[];
}

export const useNotificationStore = defineStore("notification", {
  state: (): NotificationState => ({ items: [] }),

  getters: {
    unreadCount: (s) => s.items.filter((n) => !n.is_read).length,
    sorted: (s) =>
      [...s.items].sort(
        (a, b) =>
          new Date(b.created_at).getTime() - new Date(a.created_at).getTime(),
      ),
  },

  actions: {
    set(items: Notification[]) {
      this.items = items;
    },
    push(n: Notification) {
      this.items.unshift(n);
    },
    markRead(id: string) {
      const n = this.items.find((n) => n.id === id);
      if (n) {
        n.is_read = true;
        n.read_at = new Date().toISOString();
      }
    },
    markAllRead() {
      this.items.forEach((n) => {
        n.is_read = true;
        n.read_at = new Date().toISOString();
      });
    },
  },
});
