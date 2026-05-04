<script setup lang="ts">
definePageMeta({ requiresAuth: true });
useHead({ title: "My Bookings" });

const { orders, loading, error, list } = useOrders();

onMounted(() => list());

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

const statusMap: Record<
  string,
  { label: string; variant: "forest" | "gold" | "muted" }
> = {
  confirmed: { label: "Confirmed", variant: "forest" },
  completed: { label: "Completed", variant: "muted" },
  pending: { label: "Pending", variant: "gold" },
  processing: { label: "Processing", variant: "gold" },
  cancelled: { label: "Cancelled", variant: "muted" },
};

// Derive hotel-style display fields from the order
const roomName = (o: (typeof orders.value)[0]) =>
  (o.metadata?.room_name as string) || o.items[0]?.name || "Room Booking";

const nights = (o: (typeof orders.value)[0]) => {
  const checkIn = o.metadata?.check_in as string | undefined;
  const checkOut = o.metadata?.check_out as string | undefined;
  if (!checkIn || !checkOut) return o.items[0]?.quantity ?? 1;
  return Math.round(
    (new Date(checkOut).getTime() - new Date(checkIn).getTime()) / 86_400_000,
  );
};

const checkIn = (o: (typeof orders.value)[0]) =>
  (o.metadata?.check_in as string) || o.created_at;
const checkOut = (o: (typeof orders.value)[0]) =>
  (o.metadata?.check_out as string) || o.updated_at;
</script>

<template>
  <div class="max-w-5xl mx-auto section-px py-12">
    <div class="mb-8">
      <p class="eyebrow mb-2">Account</p>
      <h1 class="font-display text-display-md text-brand-900 font-light">
        My Bookings
      </h1>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="space-y-4">
      <UiSkeleton v-for="i in 3" :key="i" class="h-24 rounded-2xl" />
    </div>

    <!-- Error -->
    <UiAlert v-else-if="error" variant="error" :message="error" class="mb-6" />

    <!-- List -->
    <div v-else class="space-y-4">
      <NuxtLink
        v-for="o in orders"
        :key="o.id"
        :to="`/orders/${o.id}`"
        class="card p-5 flex flex-col sm:flex-row items-start sm:items-center gap-5 group"
      >
        <div
          class="w-12 h-12 rounded-xl bg-forest/8 flex items-center justify-center shrink-0"
        >
          <Icon name="lucide:bed-double" class="w-6 h-6 text-forest" />
        </div>

        <div class="flex-1 min-w-0">
          <div class="flex items-center gap-2 mb-1">
            <span class="font-display text-lg text-brand-900 leading-tight">{{
              roomName(o)
            }}</span>
            <UiBadge :variant="statusMap[o.status]?.variant ?? 'muted'">
              {{ statusMap[o.status]?.label ?? o.status }}
            </UiBadge>
          </div>
          <div class="text-xs font-sans text-muted">
            {{ formatDate(checkIn(o)) }} → {{ formatDate(checkOut(o)) }} ·
            {{ nights(o) }} night{{ nights(o) !== 1 ? "s" : "" }}
          </div>
          <div class="text-xs font-sans text-muted mt-0.5 font-mono">
            {{ o.id }}
          </div>
        </div>

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
