<script setup lang="ts">
const props = defineProps<{
  page: number;
  totalPages: number;
}>();
const emit = defineEmits<{ change: [page: number] }>();

const pages = computed(() => {
  const arr: (number | "…")[] = [];
  const { page, totalPages } = props;
  for (let i = 1; i <= totalPages; i++) {
    if (i === 1 || i === totalPages || (i >= page - 1 && i <= page + 1)) {
      arr.push(i);
    } else if (arr[arr.length - 1] !== "…") {
      arr.push("…");
    }
  }
  return arr;
});
</script>

<template>
  <div class="flex items-center gap-1">
    <button
      :disabled="page === 1"
      @click="emit('change', page - 1)"
      class="w-8 h-8 flex items-center justify-center rounded border border-surface-200 disabled:opacity-40 hover:border-forest hover:text-forest transition-colors"
    >
      <Icon name="lucide:chevron-left" class="w-4 h-4" />
    </button>
    <template v-for="p in pages" :key="p">
      <span
        v-if="p === '…'"
        class="w-8 h-8 flex items-center justify-center text-muted text-sm"
        >…</span
      >
      <button
        v-else
        @click="emit('change', p as number)"
        :class="[
          'w-8 h-8 flex items-center justify-center rounded border text-sm font-sans transition-colors',
          p === page
            ? 'bg-forest text-white border-forest'
            : 'border-surface-200 hover:border-forest hover:text-forest',
        ]"
      >
        {{ p }}
      </button>
    </template>
    <button
      :disabled="page === totalPages"
      @click="emit('change', page + 1)"
      class="w-8 h-8 flex items-center justify-center rounded border border-surface-200 disabled:opacity-40 hover:border-forest hover:text-forest transition-colors"
    >
      <Icon name="lucide:chevron-right" class="w-4 h-4" />
    </button>
  </div>
</template>
