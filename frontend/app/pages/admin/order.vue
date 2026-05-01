<script setup lang="ts">
definePageMeta({ layout: "admin" });
useHead({ title: "Orders" });

const search = ref("");
const filterStatus = ref("all");

const orders = [
  {
    id: "ORD-001",
    guest: "Alice Johnson",
    room: "Pearl Suite",
    checkIn: "2025-07-15",
    total: 155250,
    status: "confirmed",
  },
  {
    id: "ORD-002",
    guest: "Bob Martinez",
    room: "Deluxe Room",
    checkIn: "2025-06-20",
    total: 52500,
    status: "pending",
  },
  {
    id: "ORD-003",
    guest: "Carol White",
    room: "Family Suite",
    checkIn: "2025-06-28",
    total: 114000,
    status: "confirmed",
  },
  {
    id: "ORD-004",
    guest: "David Lee",
    room: "Honeymoon Suite",
    checkIn: "2025-05-10",
    total: 195000,
    status: "completed",
  },
  {
    id: "ORD-005",
    guest: "Eve Turner",
    room: "Classic Room",
    checkIn: "2025-04-22",
    total: 18000,
    status: "cancelled",
  },
];

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });

const filtered = computed(() =>
  orders.filter((o) => {
    const matchSearch =
      !search.value ||
      o.guest.toLowerCase().includes(search.value.toLowerCase()) ||
      o.id.toLowerCase().includes(search.value.toLowerCase());
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
};
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="font-display text-2xl text-brand-900">Orders</h2>
        <p class="text-xs text-muted font-sans mt-0.5">
          {{ orders.length }} total orders
        </p>
      </div>
    </div>

    <!-- Filters -->
    <div class="flex flex-wrap gap-3 mb-5">
      <div class="max-w-xs flex-1">
        <UiInput
          v-model="search"
          placeholder="Search by guest or order ID…"
          icon="lucide:search"
        />
      </div>
      <div class="flex gap-1.5">
        <button
          v-for="s in ['all', 'pending', 'confirmed', 'completed', 'cancelled']"
          :key="s"
          @click="filterStatus = s"
          :class="[
            'px-3 py-1.5 rounded-md text-xs font-sans font-medium capitalize transition-colors border',
            filterStatus === s
              ? 'bg-forest text-white border-forest'
              : 'bg-white text-muted border-surface-200 hover:border-forest hover:text-forest',
          ]"
        >
          {{ s }}
        </button>
      </div>
    </div>

    <UiTable
      :headers="[
        { key: 'id', label: 'Order ID' },
        { key: 'guest', label: 'Guest' },
        { key: 'room', label: 'Room' },
        { key: 'checkIn', label: 'Check In' },
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
        <td class="px-5 py-3.5 font-mono text-xs text-muted">{{ o.id }}</td>
        <td class="px-5 py-3.5 text-sm font-sans">{{ o.guest }}</td>
        <td class="px-5 py-3.5 text-sm text-muted">{{ o.room }}</td>
        <td class="px-5 py-3.5 text-sm text-muted">{{ o.checkIn }}</td>
        <td class="px-5 py-3.5 text-sm font-sans font-medium text-right">
          {{ formatPrice(o.total) }}
        </td>
        <td class="px-5 py-3.5">
          <UiBadge :variant="statusVariant[o.status] ?? 'muted'">{{
            o.status
          }}</UiBadge>
        </td>
        <td class="px-5 py-3.5 text-right">
          <NuxtLink
            :to="`/orders/${o.id}`"
            class="text-xs font-sans text-forest hover:underline"
            >View</NuxtLink
          >
        </td>
      </tr>
      <tr v-if="filtered.length === 0">
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
