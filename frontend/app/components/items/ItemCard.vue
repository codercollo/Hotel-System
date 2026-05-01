<script setup lang="ts">
defineProps<{
  id: string;
  name: string;
  description?: string;
  price: number;
  currency?: string;
  rating?: number;
  images?: string[];
  badges?: string[];
  beds?: number;
  baths?: number;
  sqft?: number;
}>();

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });
</script>

<template>
  <NuxtLink :to="`/items/${id}`" class="card-room block group">
    <!-- Image -->
    <div class="relative overflow-hidden">
      <img
        :src="images?.[0] ?? '/images/room-placeholder.jpg'"
        :alt="name"
        class="card-room-img"
        loading="lazy"
      />
      <!-- Badges overlay -->
      <div class="absolute top-3 left-3 flex gap-1.5">
        <UiBadge
          v-for="b in (badges ?? []).slice(0, 2)"
          :key="b"
          variant="forest"
          >{{ b }}</UiBadge
        >
      </div>
      <!-- Rating pill -->
      <div
        v-if="rating"
        class="absolute top-3 right-3 bg-white/90 backdrop-blur-sm px-2.5 py-1 rounded-full flex items-center gap-1.5"
      >
        <Icon name="lucide:star" class="w-3 h-3 text-accent fill-accent" />
        <span class="text-xs font-sans font-semibold text-brand-900">{{
          rating.toFixed(1)
        }}</span>
      </div>
    </div>

    <!-- Info -->
    <div class="p-4">
      <h3 class="font-display text-lg text-brand-900 leading-snug mb-0.5">
        {{ name }}
      </h3>
      <p
        v-if="description"
        class="text-xs text-muted font-sans line-clamp-2 mb-3"
      >
        {{ description }}
      </p>

      <!-- Specs -->
      <div class="flex items-center gap-3.5 text-xs text-muted font-sans mb-3">
        <span v-if="beds" class="flex items-center gap-1">
          <Icon name="lucide:bed-double" class="w-3.5 h-3.5" /> {{ beds }} Bed
        </span>
        <span v-if="baths" class="flex items-center gap-1">
          <Icon name="lucide:bath" class="w-3.5 h-3.5" /> {{ baths }} Bath
        </span>
        <span v-if="sqft" class="flex items-center gap-1">
          <Icon name="lucide:move" class="w-3.5 h-3.5" /> {{ sqft }} sqft
        </span>
      </div>

      <!-- Price -->
      <div class="flex items-center justify-between">
        <div>
          <span class="font-display text-xl font-semibold text-brand-900">{{
            formatPrice(price)
          }}</span>
          <span class="text-xs text-muted font-sans ml-1">/night</span>
        </div>
        <span
          class="text-forest text-xs font-sans font-medium group-hover:underline underline-offset-2 transition-all"
        >
          View Details →
        </span>
      </div>
    </div>
  </NuxtLink>
</template>
