<script setup lang="ts">
const open = ref(false);
const items = ref([
  {
    id: 1,
    title: "Booking Confirmed",
    body: "Your stay at The Pearl Suite is confirmed for Jul 15.",
    time: "2m ago",
    read: false,
  },
  {
    id: 2,
    title: "Payment Received",
    body: "Payment of $1,552 has been processed successfully.",
    time: "1h ago",
    read: false,
  },
  {
    id: 3,
    title: "Check-Out Reminder",
    body: "Your check-out is tomorrow at 11:00 AM.",
    time: "1d ago",
    read: true,
  },
]);

const unread = computed(() => items.value.filter((n) => !n.read).length);
const markAll = () => items.value.forEach((n) => (n.read = true));
</script>

<template>
  <div class="relative">
    <button
      @click="open = !open"
      class="relative p-2 text-muted hover:text-brand-900 transition-colors"
    >
      <Icon name="lucide:bell" class="w-5 h-5" />
      <span
        v-if="unread > 0"
        class="absolute top-1 right-1 w-4 h-4 bg-accent text-white text-[9px] font-bold rounded-full flex items-center justify-center"
        >{{ unread }}</span
      >
    </button>

    <!-- Dropdown -->
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
            @click="markAll"
            class="text-xs text-forest font-sans hover:underline"
          >
            Mark all read
          </button>
        </div>

        <ul class="divide-y divide-surface-200 max-h-72 overflow-y-auto">
          <li
            v-for="n in items"
            :key="n.id"
            :class="[
              'px-4 py-3 hover:bg-surface-100 transition-colors',
              !n.read ? 'bg-forest/3' : '',
            ]"
          >
            <div class="flex items-start gap-2.5">
              <div
                :class="[
                  'w-2 h-2 rounded-full mt-1.5 shrink-0',
                  !n.read ? 'bg-accent' : 'bg-transparent',
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
                  {{ n.time }}
                </p>
              </div>
            </div>
          </li>
          <li
            v-if="items.length === 0"
            class="px-4 py-8 text-center text-muted font-sans text-sm"
          >
            No notifications
          </li>
        </ul>
      </div>
    </Transition>

    <!-- Backdrop -->
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
