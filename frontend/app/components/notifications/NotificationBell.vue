<script setup lang="ts">
const store = useNotificationStore();
const { markRead, markAllRead } = useNotifications();

const open = ref(false);

const items = computed(() => store.sorted);
const unread = computed(() => store.unreadCount);

const formatTime = (d: string) => {
  const diff = Date.now() - new Date(d).getTime();
  const m = Math.floor(diff / 60_000);
  if (m < 1) return "just now";
  if (m < 60) return `${m}m ago`;
  const h = Math.floor(m / 60);
  if (h < 24) return `${h}h ago`;
  return `${Math.floor(h / 24)}d ago`;
};

const handleMarkRead = async (id: string) => {
  await markRead(id);
};

const handleMarkAll = async () => {
  await markAllRead();
};
</script>

<template>
  <div class="relative">
    <button
      class="relative p-2 text-muted hover:text-brand-900 transition-colors"
      @click="open = !open"
    >
      <Icon name="lucide:bell" class="w-5 h-5" />
      <span
        v-if="unread > 0"
        class="absolute top-1 right-1 w-4 h-4 bg-accent text-white text-[9px] font-bold rounded-full flex items-center justify-center"
        >{{ unread > 9 ? "9+" : unread }}</span
      >
    </button>

    <Transition name="dropdown">
      <div
        v-if="open"
        class="absolute right-0 mt-2 w-80 bg-white rounded-xl shadow-card-hover border border-surface-200 z-50 overflow-hidden"
      >
        <div
          class="flex items-center justify-between px-4 py-3 border-b border-surface-200"
        >
          <span class="font-display text-base text-brand-900"
            >Notifications</span
          >
          <button
            class="text-xs text-forest font-sans hover:underline"
            @click="handleMarkAll"
          >
            Mark all read
          </button>
        </div>

        <ul class="divide-y divide-surface-200 max-h-72 overflow-y-auto">
          <li
            v-for="n in items"
            :key="n.id"
            :class="[
              'px-4 py-3 hover:bg-surface-100 transition-colors cursor-pointer',
              !n.is_read ? 'bg-forest/[0.03]' : '',
            ]"
            @click="handleMarkRead(n.id)"
          >
            <div class="flex items-start gap-2.5">
              <div
                :class="[
                  'w-2 h-2 rounded-full mt-1.5 shrink-0',
                  !n.is_read ? 'bg-accent' : 'bg-transparent',
                ]"
              />
              <div class="flex-1 min-w-0">
                <p class="text-sm font-sans font-medium text-brand-900">
                  {{ n.title }}
                </p>
                <p class="text-xs text-muted font-sans leading-relaxed mt-0.5">
                  {{ n.body }}
                </p>
                <p class="text-[10px] text-muted font-sans mt-1">
                  {{ formatTime(n.created_at) }}
                </p>
              </div>
            </div>
          </li>
          <li
            v-if="items.length === 0"
            class="px-4 py-8 text-center text-muted font-sans text-sm"
          >
            No notifications yet
          </li>
        </ul>
      </div>
    </Transition>

    <div v-if="open" class="fixed inset-0 z-40" @click="open = false" />
  </div>
</template>

<style scoped>
.dropdown-enter-active,
.dropdown-leave-active {
  transition:
    opacity 0.15s ease,
    transform 0.15s ease;
}
.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>
