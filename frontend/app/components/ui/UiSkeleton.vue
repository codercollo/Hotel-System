<script setup lang="ts">
defineProps<{ class?: string; rounded?: boolean }>();
</script>

<template>
  <div
    :class="[
      'animate-pulse bg-surface-200',
      rounded ? 'rounded-full' : 'rounded',
      $props.class,
    ]"
    aria-hidden="true"
  />
</template>
