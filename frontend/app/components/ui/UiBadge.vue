<script setup lang="ts">
defineProps<{
  variant?: "forest" | "gold" | "muted" | "outline";
}>();
</script>

<template>
  <span
    :class="[
      'badge',
      {
        'badge-forest': (variant ?? 'forest') === 'forest',
        'badge-gold': variant === 'gold',
        'badge-muted': variant === 'muted',
        'bg-transparent border border-brand-300 text-brand-700':
          variant === 'outline',
      },
    ]"
  >
    <slot />
  </span>
</template>
