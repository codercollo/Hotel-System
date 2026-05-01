<script setup lang="ts">
defineProps<{ open?: boolean }>();
const route = useRoute();

const links = [
  { icon: "lucide:layout-dashboard", label: "Dashboard", to: "/admin" },
  { icon: "lucide:users", label: "Users", to: "/admin/users" },
  { icon: "lucide:shield", label: "Roles", to: "/admin/roles" },
  { icon: "lucide:bed-double", label: "Rooms", to: "/admin/items" },
  { icon: "lucide:calendar-check", label: "Orders", to: "/admin/orders" },
  { icon: "lucide:settings", label: "Settings", to: "/admin/settings" },
];
</script>

<template>
  <aside
    class="w-60 bg-brand-900 text-white flex flex-col h-full min-h-screen shrink-0"
  >
    <!-- Logo -->
    <div class="px-6 py-5 border-b border-white/10 flex items-center gap-3">
      <div
        class="w-8 h-8 rounded-full bg-accent flex items-center justify-center font-display text-base font-semibold"
      >
        B
      </div>
      <div>
        <div class="font-display text-base font-semibold leading-tight">
          Boyelle
        </div>
        <div class="text-[9px] uppercase tracking-[0.2em] text-white/40">
          Admin Panel
        </div>
      </div>
    </div>

    <!-- Nav -->
    <nav class="flex-1 px-3 py-5 space-y-0.5">
      <NuxtLink
        v-for="link in links"
        :key="link.to"
        :to="link.to"
        :class="[
          'flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm font-sans transition-all',
          route.path === link.to ||
          (link.to !== '/admin' && route.path.startsWith(link.to))
            ? 'bg-white/10 text-white font-medium'
            : 'text-white/60 hover:text-white hover:bg-white/5',
        ]"
      >
        <Icon :name="link.icon" class="w-4 h-4" />
        {{ link.label }}
      </NuxtLink>
    </nav>

    <!-- User -->
    <div class="px-4 py-4 border-t border-white/10">
      <div class="flex items-center gap-3">
        <div
          class="w-8 h-8 rounded-full bg-forest-light flex items-center justify-center text-xs font-medium"
        >
          A
        </div>
        <div class="flex-1 min-w-0">
          <div class="text-sm font-medium truncate">Admin</div>
          <div class="text-xs text-white/40 truncate">admin@hotel.com</div>
        </div>
        <button class="text-white/40 hover:text-white transition-colors">
          <Icon name="lucide:log-out" class="w-4 h-4" />
        </button>
      </div>
    </div>
  </aside>
</template>
