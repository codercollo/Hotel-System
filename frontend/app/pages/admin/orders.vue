<script setup lang="ts">
import type { Order } from "~/types/order.types";

definePageMeta({
  layout: "admin",
  middleware: ["admin"],
  requiresAuth: true,
});
useHead({ title: "Orders" });

const api = useApi();
const search = ref("");
const filterStatus = ref("all");
const updatingId = ref<string | null>(null);

const {
  data: orders,
  pending: loading,
  refresh,
} = await useAsyncData<Order[]>("admin-orders", () =>
  api.get<Order[]>("/api/v1/orders", { limit: "100", offset: "0" }),
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
  (o.metadata?.guest_name as string) || `User ${o.user_id.slice(-6)}`;

const filtered = computed(() =>
  (orders.value ?? []).filter((o) => {
    const matchSearch =
      !search.value ||
      o.id.toLowerCase().includes(search.value.toLowerCase()) ||
      roomName(o).toLowerCase().includes(search.value.toLowerCase());
    const matchStatus =
      filterStatus.value === "all" || o.status === filterStatus.value;
    return matchSearch && matchStatus;
  }),
);

const statusVariant: Record<string, "forest" | "gold" | "muted"> = {
  confirmed: "forest",
  pending: "gold",
  completed: "muted",
  cancelled: "muted",
  processing: "gold",
};

const updateStatus = async (id: string, status: string) => {
  updatingId.value = id;
  try {
    await api.patch(`/api/v1/orders/${id}/status`, { status });
    await refresh();
  } catch (e: any) {
    alert(e.message);
  } finally {
    updatingId.value = null;
  }
};

const nextStatus: Record<string, string> = {
  pending: "confirmed",
  confirmed: "processing",
  processing: "completed",
};
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="font-display text-2xl text-brand-900">Orders</h2>
        <p class="text-xs text-muted font-sans mt-0.5">
          {{ orders?.length ?? 0 }} total orders
        </p>
      </div>
    </div>

    <div class="flex flex-wrap gap-3 mb-5">
      <div class="max-w-xs flex-1">
        <UiInput
          v-model="search"
          placeholder="Search by ID or room…"
          icon="lucide:search"
        />
      </div>
      <div class="flex gap-1.5 flex-wrap">
        <button
          v-for="s in [
            'all',
            'pending',
            'confirmed',
            'processing',
            'completed',
            'cancelled',
          ]"
          :key="s"
          :class="[
            'px-3 py-1.5 rounded-md text-xs font-sans font-medium capitalize transition-colors border',
            filterStatus === s
              ? 'bg-forest text-white border-forest'
              : 'bg-white text-muted border-surface-200 hover:border-forest hover:text-forest',
          ]"
          @click="filterStatus = s"
        >
          {{ s }}
        </button>
      </div>
    </div>

    <div v-if="loading" class="space-y-3">
      <UiSkeleton v-for="i in 5" :key="i" class="h-14 rounded-lg" />
    </div>

    <UiTable
      v-else
      :headers="[
        { key: 'id', label: 'Order ID' },
        { key: 'guest', label: 'Guest' },
        { key: 'room', label: 'Room' },
        { key: 'date', label: 'Date' },
        { key: 'total', label: 'Total', align: 'right' },
        { key: 'status', label: 'Status' },
        { key: 'actions', label: '', align: 'right' },
      ]"
    >
      <tr
        v-for="o in filtered"
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
        <td class="px-5 py-3.5 text-right flex items-center justify-end gap-2">
          <button
            v-if="nextStatus[o.status]"
            class="text-xs font-sans text-forest hover:underline disabled:opacity-40"
            :disabled="updatingId === o.id"
            @click="updateStatus(o.id, nextStatus[o.status]!)"
          >
            → {{ nextStatus[o.status] }}
          </button>
          <NuxtLink
            :to="`/orders/${o.id}`"
            class="text-muted hover:text-forest transition-colors"
          >
            <Icon name="lucide:external-link" class="w-3.5 h-3.5" />
          </NuxtLink>
        </td>
      </tr>
      <tr v-if="!filtered.length">
        <td
          colspan="7"
          class="px-5 py-12 text-center text-muted font-sans text-sm"
        >
          No orders match your filters.
        </td>
      </tr>
    </UiTable>
  </div>
</template>
