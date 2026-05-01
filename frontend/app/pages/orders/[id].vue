<script setup lang="ts">
const route = useRoute();
useHead({ title: "Booking Detail" });

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });
const formatDate = (d: string) =>
  new Date(d).toLocaleDateString("en-US", {
    month: "long",
    day: "numeric",
    year: "numeric",
  });

// Placeholder — replaced by useOrders(id) in Phase 4
const order = {
  id: route.params.id as string,
  room: "The Pearl Suite",
  checkIn: "2025-07-15",
  checkOut: "2025-07-18",
  nights: 3,
  guests: 2,
  status: "confirmed",
  subtotal: 135000,
  taxes: 20250,
  total: 155250,
  timeline: [
    { label: "Booking Placed", date: "2025-06-01 09:12", done: true },
    { label: "Payment Confirmed", date: "2025-06-01 09:13", done: true },
    { label: "Booking Confirmed", date: "2025-06-01 09:15", done: true },
    { label: "Check-In", date: "2025-07-15 14:00", done: false },
    { label: "Check-Out", date: "2025-07-18 11:00", done: false },
  ],
};
</script>

<template>
  <div class="max-w-4xl mx-auto section-px py-12">
    <!-- Breadcrumb -->
    <nav class="flex items-center gap-2 text-xs font-sans text-muted mb-8">
      <NuxtLink to="/orders" class="hover:text-forest transition-colors"
        >My Bookings</NuxtLink
      >
      <Icon name="lucide:chevron-right" class="w-3 h-3" />
      <span class="text-brand-900">{{ order.id }}</span>
    </nav>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Left: main details -->
      <div class="lg:col-span-2 space-y-6">
        <!-- Header card -->
        <div class="card p-6">
          <div class="flex items-start justify-between mb-4">
            <div>
              <UiBadge variant="forest" class="mb-3">{{
                order.status
              }}</UiBadge>
              <h1
                class="font-display text-display-sm text-brand-900 font-light"
              >
                {{ order.room }}
              </h1>
              <p class="text-xs font-mono text-muted mt-1">{{ order.id }}</p>
            </div>
            <div class="text-right">
              <div class="font-display text-2xl text-brand-900">
                {{ formatPrice(order.total) }}
              </div>
              <div class="text-xs text-muted font-sans">total charged</div>
            </div>
          </div>

          <div
            class="grid grid-cols-2 sm:grid-cols-4 gap-4 pt-4 border-t border-surface-200"
          >
            <div>
              <p
                class="text-xs font-sans uppercase tracking-widest text-muted mb-1"
              >
                Check In
              </p>
              <p class="text-sm font-sans text-brand-900 font-medium">
                {{ formatDate(order.checkIn) }}
              </p>
            </div>
            <div>
              <p
                class="text-xs font-sans uppercase tracking-widest text-muted mb-1"
              >
                Check Out
              </p>
              <p class="text-sm font-sans text-brand-900 font-medium">
                {{ formatDate(order.checkOut) }}
              </p>
            </div>
            <div>
              <p
                class="text-xs font-sans uppercase tracking-widest text-muted mb-1"
              >
                Nights
              </p>
              <p class="text-sm font-sans text-brand-900 font-medium">
                {{ order.nights }}
              </p>
            </div>
            <div>
              <p
                class="text-xs font-sans uppercase tracking-widest text-muted mb-1"
              >
                Guests
              </p>
              <p class="text-sm font-sans text-brand-900 font-medium">
                {{ order.guests }}
              </p>
            </div>
          </div>
        </div>

        <!-- Timeline -->
        <div class="card p-6">
          <h3 class="font-display text-xl text-brand-900 mb-6">
            Booking Timeline
          </h3>
          <ol class="relative border-l border-surface-200 space-y-6 ml-3">
            <li v-for="(step, i) in order.timeline" :key="i" class="pl-6">
              <div
                :class="[
                  'absolute -left-2 w-4 h-4 rounded-full border-2 border-white',
                  step.done ? 'bg-forest' : 'bg-surface-200',
                ]"
              />
              <p
                class="text-sm font-sans font-medium"
                :class="step.done ? 'text-brand-900' : 'text-muted'"
              >
                {{ step.label }}
              </p>
              <p class="text-xs text-muted font-sans mt-0.5">{{ step.date }}</p>
            </li>
          </ol>
        </div>
      </div>

      <!-- Right: price breakdown & actions -->
      <div class="space-y-4">
        <div class="card p-5">
          <h3 class="font-display text-lg text-brand-900 mb-4">
            Price Breakdown
          </h3>
          <div class="space-y-2.5 text-sm font-sans">
            <div class="flex justify-between">
              <span class="text-muted"
                >Room rate × {{ order.nights }} nights</span
              >
              <span>{{ formatPrice(order.subtotal) }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-muted">Taxes &amp; fees (15%)</span>
              <span>{{ formatPrice(order.taxes) }}</span>
            </div>
            <div
              class="flex justify-between pt-2.5 border-t border-surface-200 font-medium"
            >
              <span>Total</span>
              <span class="font-display text-lg">{{
                formatPrice(order.total)
              }}</span>
            </div>
          </div>
        </div>

        <div class="card p-5 space-y-3">
          <h3 class="font-display text-lg text-brand-900 mb-1">Actions</h3>
          <UiButton variant="outline" size="sm" class="w-full justify-center">
            <Icon name="lucide:download" class="w-3.5 h-3.5" /> Download Receipt
          </UiButton>
          <UiButton
            v-if="order.status === 'confirmed'"
            variant="outline"
            size="sm"
            class="w-full justify-center text-red-600 border-red-200 hover:bg-red-50 hover:border-red-300"
          >
            <Icon name="lucide:x-circle" class="w-3.5 h-3.5" /> Cancel Booking
          </UiButton>
        </div>

        <div class="card p-5">
          <h3 class="font-display text-lg text-brand-900 mb-3">Need Help?</h3>
          <p class="text-xs text-muted font-sans mb-4 leading-relaxed">
            Our concierge team is available 24/7 for any questions about your
            stay.
          </p>
          <a href="tel:+10000000000">
            <UiButton size="sm" class="w-full justify-center">
              <Icon name="lucide:phone" class="w-3.5 h-3.5" /> Call Concierge
            </UiButton>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>
