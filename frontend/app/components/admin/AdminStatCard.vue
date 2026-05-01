<script setup lang="ts">
defineProps<{
  label: string;
  value: string | number;
  icon: string;
  change?: string;
  up?: boolean;
}>();
</script>

<template>
  <div class="bg-white rounded-xl border border-surface-200 p-5">
    <div class="flex items-start justify-between mb-4">
      <div
        class="w-10 h-10 rounded-lg bg-forest/8 flex items-center justify-center"
      >
        <Icon :name="icon" class="w-5 h-5 text-forest" />
      </div>
      <span
        v-if="change"
        :class="[
          'text-xs font-sans font-medium',
          up ? 'text-green-600' : 'text-red-500',
        ]"
      >
        <Icon
          :name="up ? 'lucide:trending-up' : 'lucide:trending-down'"
          class="w-3 h-3 inline mr-0.5"
        />
        {{ change }}
      </span>
    </div>
    <div class="font-display text-3xl font-semibold text-brand-900 mb-0.5">
      {{ value }}
    </div>
    <div class="text-xs font-sans uppercase tracking-widest text-muted">
      {{ label }}
    </div>
  </div>
</template>
