<script setup lang="ts">
const pageTitle = useState("admin-page-title", () => "Dashboard");

// Pull user from your auth composable — adjust to match your actual composable name
const { user } = useAuth();

// Derive initials from user's name or email for the avatar
const initials = computed(() => {
  if (!user.value) return "A";

  const name = user.value.name || user.value.email || "";

  return name
    .split(/[\s@.]+/)
    .filter((s): s is string => s.length > 0)
    .slice(0, 2)
    .map((s) => s.charAt(0).toUpperCase())
    .join("");
});
</script>

<template>
  <div class="flex h-screen overflow-hidden bg-surface-100">
    <AppSidebar />
    <div class="flex-1 flex flex-col min-w-0 overflow-y-auto">
      <!-- Admin topbar -->
      <header
        class="bg-white border-b border-surface-200 px-6 py-4 flex items-center justify-between shrink-0"
      >
        <h1 class="font-display text-xl text-brand-900">
          {{ pageTitle }}
        </h1>

        <div class="flex items-center gap-3">
          <!-- Bell → /admin/notifications -->
          <NuxtLink
            to="/admin/notifications"
            class="relative p-2 rounded-lg text-muted hover:text-brand-900 hover:bg-surface-100 transition-colors"
            aria-label="Notifications"
          >
            <Icon name="lucide:bell" class="w-5 h-5" />
            <!-- Unread badge — hide it when count is 0 -->
            <span
              class="absolute top-1.5 right-1.5 w-2 h-2 bg-accent rounded-full ring-2 ring-white"
            />
          </NuxtLink>

          <!-- Avatar → /admin/account -->
          <NuxtLink
            to="/admin/account"
            class="w-8 h-8 rounded-full bg-forest flex items-center justify-center text-white text-xs font-medium hover:ring-2 hover:ring-forest/40 transition-all"
            aria-label="My account"
          >
            {{ initials }}
          </NuxtLink>
        </div>
      </header>

      <main class="flex-1 p-6">
        <slot />
      </main>
    </div>
  </div>
</template>
