<script setup lang="ts">
useHead({ title: "My Bookings" });

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

// Placeholder data — replaced by useOrders() in Phase 4
const orders = [
  {
    id: "ORD-001",
    room: "The Pearl Suite",
    checkIn: "2025-07-15",
    checkOut: "2025-07-18",
    nights: 3,
    total: 135000,
    status: "confirmed",
  },
  {
    id: "ORD-002",
    room: "Deluxe Room",
    checkIn: "2025-05-10",
    checkOut: "2025-05-12",
    nights: 2,
    total: 50000,
    status: "completed",
  },
  {
    id: "ORD-003",
    room: "Honeymoon Suite",
    checkIn: "2025-04-01",
    checkOut: "2025-04-04",
    nights: 3,
    total: 195000,
    status: "completed",
  },
  {
    id: "ORD-004",
    room: "Classic Room",
    checkIn: "2025-06-20",
    checkOut: "2025-06-21",
    nights: 1,
    total: 18000,
    status: "cancelled",
  },
];

const statusMap: Record<
  string,
  { label: string; variant: "forest" | "gold" | "muted" }
> = {
  confirmed: { label: "Confirmed", variant: "forest" },
  completed: { label: "Completed", variant: "muted" },
  pending: { label: "Pending", variant: "gold" },
  cancelled: { label: "Cancelled", variant: "muted" },
};
</script>

<template>
  <div class="max-w-5xl mx-auto section-px py-12">
    <div class="mb-8">
      <p class="eyebrow mb-2">Account</p>
      <h1 class="font-display text-display-md text-brand-900 font-light">
        My Bookings
      </h1>
    </div>

    <div class="space-y-4">
      <NuxtLink
        v-for="o in orders"
        :key="o.id"
        :to="`/orders/${o.id}`"
        class="card p-5 flex flex-col sm:flex-row items-start sm:items-center gap-5 group"
      >
        <!-- Icon -->
        <div
          class="w-12 h-12 rounded-xl bg-forest/8 flex items-center justify-center shrink-0"
        >
          <Icon name="lucide:bed-double" class="w-6 h-6 text-forest" />
        </div>

        <!-- Details -->
        <div class="flex-1 min-w-0">
          <div class="flex items-center gap-2 mb-1">
            <span class="font-display text-lg text-brand-900 leading-tight">{{
              o.room
            }}</span>
            <UiBadge :variant="statusMap[o.status]?.variant ?? 'muted'">
              {{ statusMap[o.status]?.label ?? o.status }}
            </UiBadge>
          </div>
          <div class="text-xs font-sans text-muted">
            {{ formatDate(o.checkIn) }} → {{ formatDate(o.checkOut) }} ·
            {{ o.nights }} night{{ o.nights > 1 ? "s" : "" }}
          </div>
          <div class="text-xs font-sans text-muted mt-0.5 font-mono">
            {{ o.id }}
          </div>
        </div>

        <!-- Total + arrow -->
        <div class="flex items-center gap-4 shrink-0">
          <div class="text-right">
            <div class="font-display text-xl text-brand-900">
              {{ formatPrice(o.total) }}
            </div>
            <div class="text-xs text-muted font-sans">total charged</div>
          </div>
          <Icon
            name="lucide:chevron-right"
            class="w-4 h-4 text-muted group-hover:text-forest transition-colors"
          />
        </div>
      </NuxtLink>

      <div
        v-if="orders.length === 0"
        class="py-20 text-center text-muted font-sans"
      >
        <Icon
          name="lucide:calendar-x"
          class="w-12 h-12 mx-auto mb-4 opacity-30"
        />
        <p class="mb-4">You have no bookings yet.</p>
        <NuxtLink to="/items"
          ><UiButton size="sm">Browse Rooms</UiButton></NuxtLink
        >
      </div>
    </div>
  </div>
</template>
