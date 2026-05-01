<script setup lang="ts">
defineProps<{ status: string }>();

const map: Record<
  string,
  { variant: "forest" | "gold" | "muted"; label: string }
> = {
  pending: { variant: "gold", label: "Pending" },
  confirmed: { variant: "forest", label: "Confirmed" },
  processing: { variant: "gold", label: "Processing" },
  completed: { variant: "muted", label: "Completed" },
  cancelled: { variant: "muted", label: "Cancelled" },
};
</script>

<template>
  <UiBadge :variant="map[status]?.variant ?? 'muted'">
    {{ map[status]?.label ?? status }}
  </UiBadge>
</template>
