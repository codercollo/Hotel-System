<script setup lang="ts">
defineProps<{
  items: any[];
  loading?: boolean;
  cols?: 2 | 3 | 4;
}>();
</script>

<template>
  <div
    :class="[
      'grid gap-6',
      {
        'grid-cols-1 sm:grid-cols-2': (cols ?? 3) === 2,
        'grid-cols-1 sm:grid-cols-2 lg:grid-cols-3': (cols ?? 3) === 3,
        'grid-cols-1 sm:grid-cols-2 lg:grid-cols-4': cols === 4,
      },
    ]"
  >
    <template v-if="loading">
      <div v-for="i in 6" :key="i" class="rounded-xl overflow-hidden">
        <UiSkeleton class="h-52 w-full" />
        <div class="p-4 space-y-2">
          <UiSkeleton class="h-5 w-3/4" />
          <UiSkeleton class="h-3 w-full" />
          <UiSkeleton class="h-3 w-1/2" />
        </div>
      </div>
    </template>
    <template v-else>
      <ItemCard
        v-for="item in items"
        :key="item.id"
        v-bind="item"
        :beds="item.metadata?.beds"
        :baths="item.metadata?.baths"
        :sqft="item.metadata?.sqft"
        :rating="item.metadata?.rating"
        :badges="item.metadata?.badges"
      />
    </template>
  </div>
</template>
