<script setup lang="ts">
definePageMeta({ layout: "admin", requiresAuth: true, roles: ["admin"] });

const pageTitle = useState("admin-page-title");
pageTitle.value = "Notifications";

// ── Types ─────────────────────────────────────────────────────────────────────
type NotifKind = "booking" | "payment" | "cancellation" | "system";

interface Notification {
  id: string;
  kind: NotifKind;
  title: string;
  body: string;
  read: boolean;
  created_at: string;
}

// ── State ─────────────────────────────────────────────────────────────────────
const loading = ref(true);
const activeFilter = ref<"all" | "unread">("all");

// Replace with a real API call: const { data } = await useFetch("/api/v1/admin/notifications")
const notifications = ref<Notification[]>([
  {
    id: "1",
    kind: "booking",
    title: "New booking #BK-1042",
    body: "Room 204 booked for 3 nights from May 10.",
    read: false,
    created_at: new Date(Date.now() - 1000 * 60 * 5).toISOString(),
  },
  {
    id: "2",
    kind: "payment",
    title: "Payment confirmed",
    body: "Order #ORD-8821 paid via Stripe — $480.",
    read: false,
    created_at: new Date(Date.now() - 1000 * 60 * 32).toISOString(),
  },
  {
    id: "3",
    kind: "cancellation",
    title: "Booking cancelled",
    body: "Guest John D. cancelled reservation #BK-1039.",
    read: false,
    created_at: new Date(Date.now() - 1000 * 60 * 60 * 2).toISOString(),
  },
  {
    id: "4",
    kind: "payment",
    title: "Payment failed",
    body: "M-Pesa STK push timed out for order #ORD-8819.",
    read: true,
    created_at: new Date(Date.now() - 1000 * 60 * 60 * 5).toISOString(),
  },
  {
    id: "5",
    kind: "system",
    title: "System maintenance scheduled",
    body: "Downtime window: May 6, 02:00–03:00 UTC.",
    read: true,
    created_at: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(),
  },
  {
    id: "6",
    kind: "booking",
    title: "New booking #BK-1041",
    body: "Suite 501 booked for 1 night from May 8.",
    read: true,
    created_at: new Date(Date.now() - 1000 * 60 * 60 * 26).toISOString(),
  },
]);

setTimeout(() => (loading.value = false), 400);

// ── Derived ───────────────────────────────────────────────────────────────────
const unreadCount = computed(
  () => notifications.value.filter((n) => !n.read).length,
);

const filtered = computed(() =>
  activeFilter.value === "unread"
    ? notifications.value.filter((n) => !n.read)
    : notifications.value,
);

// ── Actions ───────────────────────────────────────────────────────────────────
function markRead(id: string) {
  const n = notifications.value.find((n) => n.id === id);
  if (n) n.read = true;
  // TODO: PATCH /api/v1/admin/notifications/:id/read
}

function markAllRead() {
  notifications.value.forEach((n) => (n.read = true));
  // TODO: POST /api/v1/admin/notifications/read-all
}

function dismiss(id: string) {
  notifications.value = notifications.value.filter((n) => n.id !== id);
  // TODO: DELETE /api/v1/admin/notifications/:id
}

// ── Helpers ───────────────────────────────────────────────────────────────────
const kindMeta: Record<NotifKind, { icon: string; color: string; bg: string }> =
  {
    booking: {
      icon: "lucide:calendar-check",
      color: "text-forest",
      bg: "bg-forest/10",
    },
    payment: {
      icon: "lucide:credit-card",
      color: "text-blue-600",
      bg: "bg-blue-50",
    },
    cancellation: {
      icon: "lucide:x-circle",
      color: "text-red-500",
      bg: "bg-red-50",
    },
    system: { icon: "lucide:info", color: "text-amber-600", bg: "bg-amber-50" },
  };

function timeAgo(iso: string) {
  const diff = Math.floor((Date.now() - new Date(iso).getTime()) / 1000);
  if (diff < 60) return `${diff}s ago`;
  if (diff < 3600) return `${Math.floor(diff / 60)}m ago`;
  if (diff < 86400) return `${Math.floor(diff / 3600)}h ago`;
  return `${Math.floor(diff / 86400)}d ago`;
}
</script>

<template>
  <div class="max-w-2xl mx-auto">
    <!-- Header row -->
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <h2 class="font-display text-2xl text-brand-900">Notifications</h2>
        <span
          v-if="unreadCount > 0"
          class="inline-flex items-center justify-center h-5 min-w-5 px-1.5 rounded-full bg-accent text-white text-[10px] font-bold"
        >
          {{ unreadCount }}
        </span>
      </div>
      <button
        v-if="unreadCount > 0"
        class="text-xs font-sans text-forest hover:underline"
        @click="markAllRead"
      >
        Mark all as read
      </button>
    </div>

    <!-- Filter tabs -->
    <div class="flex gap-1 mb-4 bg-surface-100 p-1 rounded-xl w-fit">
      <button
        v-for="tab in ['all', 'unread'] as const"
        :key="tab"
        :class="[
          'px-4 py-1.5 rounded-lg text-xs font-sans font-medium transition-all capitalize',
          activeFilter === tab
            ? 'bg-white text-brand-900 shadow-sm'
            : 'text-muted hover:text-brand-900',
        ]"
        @click="activeFilter = tab"
      >
        {{ tab }}
      </button>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="space-y-3">
      <UiSkeleton v-for="i in 4" :key="i" class="h-20 rounded-2xl" />
    </div>

    <!-- Empty -->
    <div
      v-else-if="filtered.length === 0"
      class="card p-12 text-center text-muted"
    >
      <Icon name="lucide:bell-off" class="w-8 h-8 mx-auto mb-3 opacity-40" />
      <p class="font-sans text-sm">No notifications here.</p>
    </div>

    <!-- List -->
    <div v-else class="space-y-2">
      <div
        v-for="n in filtered"
        :key="n.id"
        :class="[
          'card p-4 flex items-start gap-4 group transition-all',
          !n.read && 'ring-1 ring-forest/20 bg-forest/[0.02]',
        ]"
      >
        <!-- Icon -->
        <div
          :class="[
            'mt-0.5 w-9 h-9 rounded-xl flex items-center justify-center shrink-0',
            kindMeta[n.kind].bg,
          ]"
        >
          <Icon
            :name="kindMeta[n.kind].icon"
            class="w-4 h-4"
            :class="kindMeta[n.kind].color"
          />
        </div>

        <!-- Content -->
        <div class="flex-1 min-w-0">
          <div class="flex items-start justify-between gap-2">
            <p
              :class="[
                'text-sm font-sans leading-snug',
                n.read ? 'text-muted' : 'text-brand-900 font-medium',
              ]"
            >
              {{ n.title }}
            </p>
            <span class="text-[10px] font-sans text-muted shrink-0 mt-0.5">
              {{ timeAgo(n.created_at) }}
            </span>
          </div>
          <p class="text-xs font-sans text-muted mt-0.5 leading-relaxed">
            {{ n.body }}
          </p>
        </div>

        <!-- Actions (visible on hover) -->
        <div
          class="flex items-center gap-1 opacity-0 group-hover:opacity-100 transition-opacity shrink-0"
        >
          <button
            v-if="!n.read"
            class="p-1.5 rounded-lg hover:bg-surface-100 text-muted hover:text-forest transition-colors"
            title="Mark as read"
            @click="markRead(n.id)"
          >
            <Icon name="lucide:check" class="w-3.5 h-3.5" />
          </button>
          <button
            class="p-1.5 rounded-lg hover:bg-red-50 text-muted hover:text-red-500 transition-colors"
            title="Dismiss"
            @click="dismiss(n.id)"
          >
            <Icon name="lucide:x" class="w-3.5 h-3.5" />
          </button>
        </div>

        <!-- Unread dot -->
        <div
          v-if="!n.read"
          class="mt-2 w-1.5 h-1.5 rounded-full bg-accent shrink-0"
        />
      </div>
    </div>
  </div>
</template>
