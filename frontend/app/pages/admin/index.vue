<script setup lang="ts">
definePageMeta({ layout: "admin" });
useHead({ title: "Admin Dashboard" });

const pageTitle = useState("admin-page-title");
pageTitle.value = "Dashboard";

const stats = [
  {
    label: "Total Rooms",
    value: "50",
    icon: "lucide:bed-double",
    change: "+2 this month",
    up: true,
  },
  {
    label: "Total Guests",
    value: "1,248",
    icon: "lucide:users",
    change: "+12% this month",
    up: true,
  },
  {
    label: "Orders Today",
    value: "34",
    icon: "lucide:calendar-check",
    change: "+8 vs yesterday",
    up: true,
  },
  {
    label: "Revenue MTD",
    value: "$48.2k",
    icon: "lucide:dollar-sign",
    change: "-3% vs last month",
    up: false,
  },
];

const recentOrders = [
  {
    id: "ORD-001",
    guest: "Alice Johnson",
    room: "Pearl Suite",
    checkIn: "2025-06-01",
    status: "confirmed",
  },
  {
    id: "ORD-002",
    guest: "Bob Martinez",
    room: "Deluxe Room",
    checkIn: "2025-06-03",
    status: "pending",
  },
  {
    id: "ORD-003",
    guest: "Carol White",
    room: "Family Suite",
    checkIn: "2025-06-05",
    status: "confirmed",
  },
  {
    id: "ORD-004",
    guest: "David Lee",
    room: "Honeymoon Suite",
    checkIn: "2025-06-07",
    status: "cancelled",
  },
];

const statusColor: Record<string, string> = {
  confirmed: "badge-forest",
  pending: "badge-gold",
  cancelled: "bg-red-100 text-red-700 badge",
};
</script>

<template>
  <div>
    <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4 mb-8">
      <div
        v-for="s in stats"
        :key="s.label"
        class="bg-white rounded-xl border border-surface-200 p-5"
      >
        <div class="flex items-start justify-between mb-3">
          <div
            class="w-10 h-10 rounded-lg bg-forest/8 flex items-center justify-center"
          >
            <Icon :name="s.icon" class="w-5 h-5 text-forest" />
          </div>
          <span
            :class="[
              'text-xs font-sans',
              s.up ? 'text-green-600' : 'text-red-500',
            ]"
          >
            {{ s.change }}
          </span>
        </div>
        <div class="font-display text-3xl font-semibold text-brand-900 mb-0.5">
          {{ s.value }}
        </div>
        <div class="text-xs font-sans uppercase tracking-widest text-muted">
          {{ s.label }}
        </div>
      </div>
    </div>

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
      <UiTable
        :headers="[
          { key: 'id', label: 'Order ID' },
          { key: 'guest', label: 'Guest' },
          { key: 'room', label: 'Room' },
          { key: 'checkIn', label: 'Check In' },
          { key: 'status', label: 'Status' },
        ]"
      >
        <tr
          v-for="o in recentOrders"
          :key="o.id"
          class="border-b border-surface-200 last:border-0 hover:bg-surface-100 transition-colors"
        >
          <td class="px-5 py-3.5 text-xs font-mono text-muted">{{ o.id }}</td>
          <td class="px-5 py-3.5 text-sm">{{ o.guest }}</td>
          <td class="px-5 py-3.5 text-sm text-muted">{{ o.room }}</td>
          <td class="px-5 py-3.5 text-sm text-muted">{{ o.checkIn }}</td>
          <td class="px-5 py-3.5">
            <span :class="['badge', statusColor[o.status] ?? 'badge-muted']">
              {{ o.status }}
            </span>
          </td>
        </tr>
      </UiTable>
    </div>
  </div>
</template>
