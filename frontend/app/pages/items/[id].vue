<script setup lang="ts">
const route = useRoute();
useHead({ title: "Room Detail" });

// Placeholder — replaced by useItems in Phase 3
const room = {
  id: route.params.id,
  name: "The Pearl Suite",
  description:
    "Our most requested suite, The Pearl offers sweeping city panoramas from its floor-to-ceiling windows. Featuring a king-size bed dressed in 600-thread-count Egyptian cotton, a separate living area, marble en-suite, and personalised butler service around the clock.",
  price: 45000,
  rating: 5.0,
  beds: 2,
  baths: 2,
  sqft: 400,
  amenities: [
    "King Bed",
    "City View",
    "Butler Service",
    "Mini Bar",
    "Marble Bathroom",
    "High-Speed WiFi",
    "Smart TV",
    "In-Room Dining",
  ],
};

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });

const checkIn = ref("");
const checkOut = ref("");
const guests = ref("1");
const loading = ref(false);

const book = async () => {
  loading.value = true;
  await new Promise((r) => setTimeout(r, 1000));
  loading.value = false;
  navigateTo("/orders");
};
</script>

<template>
  <div class="max-w-7xl mx-auto section-px py-10">
    <!-- Breadcrumb -->
    <nav class="flex items-center gap-2 text-xs font-sans text-muted mb-8">
      <NuxtLink to="/" class="hover:text-forest transition-colors"
        >Home</NuxtLink
      >
      <Icon name="lucide:chevron-right" class="w-3 h-3" />
      <NuxtLink to="/items" class="hover:text-forest transition-colors"
        >Rooms & Suites</NuxtLink
      >
      <Icon name="lucide:chevron-right" class="w-3 h-3" />
      <span class="text-brand-900">{{ room.name }}</span>
    </nav>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
      <!-- Left: details -->
      <div class="lg:col-span-2">
        <!-- Image placeholder -->
        <div
          class="rounded-2xl overflow-hidden aspect-video bg-surface-200 mb-6 flex items-center justify-center"
        >
          <Icon name="lucide:image" class="w-20 h-20 text-surface-300" />
        </div>

        <!-- Header -->
        <div class="flex items-start justify-between gap-4 mb-5">
          <div>
            <UiBadge variant="forest" class="mb-3">Luxury Suites</UiBadge>
            <h1 class="font-display text-display-md text-brand-900 font-light">
              {{ room.name }}
            </h1>
          </div>
          <div class="text-right shrink-0">
            <div class="font-display text-3xl font-semibold text-brand-900">
              {{ formatPrice(room.price) }}
            </div>
            <div class="text-xs text-muted font-sans">/night</div>
            <div class="flex items-center gap-1 mt-1 justify-end">
              <Icon
                name="lucide:star"
                class="w-3.5 h-3.5 text-accent fill-accent"
              />
              <span class="text-sm font-sans font-medium">{{
                room.rating.toFixed(1)
              }}</span>
            </div>
          </div>
        </div>

        <!-- Specs -->
        <div
          class="flex items-center gap-6 py-4 border-y border-surface-200 mb-6"
        >
          <span
            class="flex items-center gap-2 text-sm text-brand-800 font-sans"
          >
            <Icon name="lucide:bed-double" class="w-4 h-4 text-forest" />
            {{ room.beds }} Bedroom{{ room.beds > 1 ? "s" : "" }}
          </span>
          <span
            class="flex items-center gap-2 text-sm text-brand-800 font-sans"
          >
            <Icon name="lucide:bath" class="w-4 h-4 text-forest" />
            {{ room.baths }} Bathroom{{ room.baths > 1 ? "s" : "" }}
          </span>
          <span
            class="flex items-center gap-2 text-sm text-brand-800 font-sans"
          >
            <Icon name="lucide:move" class="w-4 h-4 text-forest" />
            {{ room.sqft }} sqft
          </span>
        </div>

        <!-- Description -->
        <p class="text-sm font-sans text-muted leading-relaxed mb-8">
          {{ room.description }}
        </p>

        <!-- Amenities -->
        <h3 class="font-display text-xl text-brand-900 mb-4">Room Amenities</h3>
        <div class="grid grid-cols-2 sm:grid-cols-4 gap-3 mb-8">
          <div
            v-for="a in room.amenities"
            :key="a"
            class="flex items-center gap-2 text-sm font-sans text-brand-800"
          >
            <Icon name="lucide:check" class="w-4 h-4 text-forest shrink-0" />
            {{ a }}
          </div>
        </div>
      </div>

      <!-- Right: booking card -->
      <div class="lg:col-span-1">
        <div class="card p-6 sticky top-24">
          <h3 class="font-display text-xl text-brand-900 mb-1">
            Book This Room
          </h3>
          <p class="text-xs text-muted font-sans mb-5">
            Reserve now, pay on arrival
          </p>

          <div class="space-y-4">
            <UiInput
              v-model="checkIn"
              label="Check In"
              type="date"
              icon="lucide:calendar"
            />
            <UiInput
              v-model="checkOut"
              label="Check Out"
              type="date"
              icon="lucide:calendar"
            />

            <div>
              <label
                class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                >Guests</label
              >
              <select v-model="guests" class="input">
                <option value="1">1 Adult</option>
                <option value="2">2 Adults</option>
                <option value="3">2 Adults, 1 Child</option>
                <option value="4">2 Adults, 2 Children</option>
              </select>
            </div>

            <div class="py-4 border-y border-surface-200 space-y-2">
              <div class="flex justify-between text-sm font-sans">
                <span class="text-muted">Room rate</span>
                <span>{{ formatPrice(room.price) }}/night</span>
              </div>
              <div class="flex justify-between text-sm font-sans">
                <span class="text-muted">Taxes &amp; fees</span>
                <span>{{ formatPrice(Math.round(room.price * 0.15)) }}</span>
              </div>
            </div>
            <div class="flex justify-between font-display text-xl">
              <span>Total</span>
              <span>{{
                formatPrice(room.price + Math.round(room.price * 0.15))
              }}</span>
            </div>

            <UiButton
              @click="book"
              :loading="loading"
              size="md"
              class="w-full justify-center"
            >
              Reserve Now
            </UiButton>
          </div>

          <p class="text-xs text-muted font-sans mt-4 text-center">
            Free cancellation up to 24h before arrival
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
