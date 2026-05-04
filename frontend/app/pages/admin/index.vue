<script setup lang="ts">
import type { AdminStats } from "~/types/admin.types";
import type { Order } from "~/types/order.types";

definePageMeta({
  layout: "admin",
  middleware: ["admin"],
  requiresAuth: true,
});
useHead({ title: "Admin Dashboard" });

const api = useApi();

const { data: stats, pending: statsLoading } = await useAsyncData<AdminStats>(
  "admin-stats",
  () => api.get<AdminStats>("/api/v1/admin/stats"),
);

const { data: recentOrders, pending: ordersLoading } = await useAsyncData<
  Order[]
>("admin-recent-orders", () =>
  api.get<Order[]>("/api/v1/orders", { limit: "5", offset: "0" }),
);

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });

const formatDate = (d: string) =>
  new Date(d).toLocaleDateString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
  });

const roomName = (o: Order) =>
  (o.metadata?.room_name as string) || o.items[0]?.name || "Room Booking";

const guestName = (o: Order) =>
  (o.metadata?.guest_name as string) || `Guest #${o.user_id.slice(-4)}`;

const statCards = computed(() => [
  {
    label: "Total Rooms",
    value: stats.value?.total_items ?? "—",
    icon: "lucide:bed-double",
    change: null,
    up: true,
  },
  {
    label: "Total Users",
    value: stats.value?.total_users ?? "—",
    icon: "lucide:users",
    change: null,
    up: true,
  },
  {
    label: "Total Orders",
    value: stats.value?.total_orders ?? "—",
    icon: "lucide:calendar-check",
    change: null,
    up: true,
  },
  {
    label: "Revenue (completed)",
    value: stats.value ? formatPrice(stats.value.revenue_total) : "—",
    icon: "lucide:dollar-sign",
    change: null,
    up: true,
  },
]);

const statusVariant: Record<string, "forest" | "gold" | "muted"> = {
  confirmed: "forest",
  pending: "gold",
  completed: "muted",
  cancelled: "muted",
  processing: "gold",
};
</script>

<template>
  <div>
    <!-- Stat cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4 mb-8">
      <AdminStatCard
        v-for="s in statCards"
        :key="s.label"
        :label="s.label"
        :value="statsLoading ? '…' : String(s.value)"
        :icon="s.icon"
      />
    </div>

    <!-- Recent orders -->
    <div class="bg-white rounded-xl border border-surface-200 p-5">
      <div class="flex items-center justify-between mb-5">
        <h3 class="font-display text-xl text-brand-900">Recent Orders</h3>
        <NuxtLink
          to="/admin/orders"
          class="text-xs font-sans text-forest hover:underline"
        >
          View all →
        </NuxtLink>
      </div>

      <div v-if="ordersLoading" class="space-y-3">
        <UiSkeleton v-for="i in 4" :key="i" class="h-12 rounded-lg" />
      </div>

      <UiTable
        v-else
        :headers="[
          { key: 'id', label: 'Order ID' },
          { key: 'guest', label: 'Guest' },
          { key: 'room', label: 'Room' },
          { key: 'date', label: 'Placed' },
          { key: 'total', label: 'Total', align: 'right' },
          { key: 'status', label: 'Status' },
        ]"
      >
        <tr
          v-for="o in recentOrders ?? []"
          :key="o.id"
          class="border-b border-surface-200 last:border-0 hover:bg-surface-100 transition-colors"
        >
          <td class="px-5 py-3.5 font-mono text-xs text-muted">
            {{ o.id.slice(-8) }}
          </td>
          <td class="px-5 py-3.5 text-sm">{{ guestName(o) }}</td>
          <td class="px-5 py-3.5 text-sm text-muted">{{ roomName(o) }}</td>
          <td class="px-5 py-3.5 text-sm text-muted">
            {{ formatDate(o.created_at) }}
          </td>
          <td class="px-5 py-3.5 text-sm font-medium text-right">
            {{ formatPrice(o.total) }}
          </td>
          <td class="px-5 py-3.5">
            <UiBadge :variant="statusVariant[o.status] ?? 'muted'">{{
              o.status
            }}</UiBadge>
          </td>
        </tr>
        <tr v-if="!recentOrders?.length">
          <td
            colspan="6"
            class="px-5 py-10 text-center text-muted font-sans text-sm"
          >
            No orders yet.
          </td>
        </tr>
      </UiTable>
    </div>
  </div>
</template>
