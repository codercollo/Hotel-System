<script setup lang="ts">
defineProps<{
  room: string;
  checkIn: string;
  checkOut: string;
  nights: number;
  subtotal: number;
  taxes: number;
  total: number;
}>();

const fmt = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });
const fmtDate = (d: string) =>
  new Date(d).toLocaleDateString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
  });
</script>

<template>
  <div class="card p-5">
    <h3 class="font-display text-xl text-brand-900 mb-4">Booking Summary</h3>

    <div class="flex items-center gap-3 mb-4 pb-4 border-b border-surface-200">
      <div
        class="w-12 h-12 rounded-xl bg-forest/8 flex items-center justify-center"
      >
        <Icon name="lucide:bed-double" class="w-6 h-6 text-forest" />
      </div>
      <div>
        <p class="font-display text-base text-brand-900">{{ room }}</p>
        <p class="text-xs text-muted font-sans">
          {{ fmtDate(checkIn) }} → {{ fmtDate(checkOut) }} ·
          {{ nights }} night{{ nights !== 1 ? "s" : "" }}
        </p>
      </div>
    </div>

    <div class="space-y-2.5 text-sm font-sans">
      <div class="flex justify-between">
        <span class="text-muted">Room rate × {{ nights }} nights</span>
        <span>{{ fmt(subtotal) }}</span>
      </div>
      <div class="flex justify-between">
        <span class="text-muted">Taxes &amp; fees</span>
        <span>{{ fmt(taxes) }}</span>
      </div>
      <div
        class="flex justify-between pt-2.5 border-t border-surface-200 font-medium"
      >
        <span>Total</span>
        <span class="font-display text-lg text-brand-900">{{
          fmt(total)
        }}</span>
      </div>
    </div>
  </div>
</template>
