<script setup lang="ts">
defineProps<{
  headers: {
    key: string;
    label: string;
    align?: "left" | "right" | "center";
  }[];
  loading?: boolean;
  empty?: string;
  total?: number;
  page?: number;
  pages?: number;
}>();
defineEmits<{ pageChange: [p: number] }>();
</script>

<template>
  <div>
    <UiTable :headers="headers" :loading="loading" :empty="empty">
      <slot />
    </UiTable>
    <div
      v-if="pages && pages > 1"
      class="flex items-center justify-between px-1 pt-4"
    >
      <p class="text-xs text-muted font-sans">{{ total }} total records</p>
      <UiPagination
        :page="page ?? 1"
        :total-pages="pages"
        @change="$emit('pageChange', $event)"
      />
    </div>
  </div>
</template>
