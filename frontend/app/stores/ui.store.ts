import { defineStore } from "pinia";

export const useUIStore = defineStore("ui", {
  state: () => ({
    sidebarOpen: true,
    theme: "light" as "light" | "dark",
  }),
  actions: {
    toggleSidebar() {
      this.sidebarOpen = !this.sidebarOpen;
    },
    toggleTheme() {
      this.theme = this.theme === "light" ? "dark" : "light";
    },
  },
});
