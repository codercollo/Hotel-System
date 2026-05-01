<script setup lang="ts">
useHead({ title: "My Dashboard" });

// Phase 2: protect via auth middleware
const upcomingStays = [
  {
    id: "ORD-001",
    room: "The Pearl Suite",
    checkIn: "2025-07-15",
    checkOut: "2025-07-18",
    nights: 3,
    status: "confirmed",
    total: 135000,
  },
];

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
</script>

<template>
  <div class="max-w-4xl mx-auto section-px py-12">
    <div class="mb-10">
      <p class="eyebrow mb-2">Welcome Back</p>
      <h1 class="font-display text-display-md text-brand-900 font-light">
        Your Dashboard
      </h1>
    </div>

    <!-- Quick stats -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-10">
      <div class="card p-5 text-center">
        <div class="font-display text-3xl text-forest mb-1">3</div>
        <div class="text-xs font-sans uppercase tracking-widest text-muted">
          Total Stays
        </div>
      </div>
      <div class="card p-5 text-center">
        <div class="font-display text-3xl text-forest mb-1">1</div>
        <div class="text-xs font-sans uppercase tracking-widest text-muted">
          Upcoming
        </div>
      </div>
      <div class="card p-5 text-center">
        <div class="font-display text-3xl text-forest mb-1">2,400</div>
        <div class="text-xs font-sans uppercase tracking-widest text-muted">
          Loyalty Points
        </div>
      </div>
    </div>

    <!-- Upcoming stays -->
    <div class="card p-6 mb-6">
      <div class="flex items-center justify-between mb-5">
        <h3 class="font-display text-xl text-brand-900">Upcoming Stays</h3>
        <NuxtLink
          to="/orders"
          class="text-xs font-sans text-forest hover:underline"
          >View all →</NuxtLink
        >
      </div>
      <div
        v-for="s in upcomingStays"
        :key="s.id"
        class="flex items-center gap-5 p-4 rounded-xl bg-surface-100"
      >
        <div
          class="w-12 h-12 rounded-xl bg-forest/10 flex items-center justify-center shrink-0"
        >
          <Icon name="lucide:bed-double" class="w-6 h-6 text-forest" />
        </div>
        <div class="flex-1 min-w-0">
          <div class="font-display text-lg text-brand-900">{{ s.room }}</div>
          <div class="text-xs font-sans text-muted">
            {{ formatDate(s.checkIn) }} → {{ formatDate(s.checkOut) }} ·
            {{ s.nights }} nights
          </div>
        </div>
        <div class="text-right shrink-0">
          <UiBadge variant="forest">{{ s.status }}</UiBadge>
          <div class="text-sm font-sans font-medium mt-1">
            {{ formatPrice(s.total) }}
          </div>
        </div>
      </div>
    </div>

    <!-- Quick links -->
    <div class="grid grid-cols-2 sm:grid-cols-4 gap-3">
      <NuxtLink
        v-for="l in [
          { icon: 'lucide:search', label: 'Browse Rooms', to: '/items' },
          { icon: 'lucide:calendar', label: 'My Bookings', to: '/orders' },
          { icon: 'lucide:user', label: 'Profile', to: '/account/profile' },
          {
            icon: 'lucide:settings',
            label: 'Settings',
            to: '/account/settings',
          },
        ]"
        :key="l.to"
        :to="l.to"
        class="card p-4 flex flex-col items-center gap-2 text-center hover:border-forest transition-colors"
      >
        <Icon :name="l.icon" class="w-5 h-5 text-forest" />
        <span class="text-xs font-sans text-brand-800">{{ l.label }}</span>
      </NuxtLink>
    </div>
  </div>
</template>
