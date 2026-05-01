<script setup lang="ts">
defineProps<{
  headers: {
    key: string;
    label: string;
    align?: "left" | "right" | "center";
  }[];
  loading?: boolean;
  empty?: string;
}>();
</script>

<template>
  <div class="overflow-x-auto rounded-xl border border-surface-200 bg-white">
    <table class="w-full text-sm font-sans">
      <thead>
        <tr class="border-b border-surface-200">
          <th
            v-for="h in headers"
            :key="h.key"
            :class="[
              'px-5 py-3.5 font-medium text-xs uppercase tracking-widest text-muted whitespace-nowrap',
              h.align === 'right'
                ? 'text-right'
                : h.align === 'center'
                  ? 'text-center'
                  : 'text-left',
            ]"
          >
            {{ h.label }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-if="loading">
          <td
            :colspan="headers.length"
            class="px-5 py-10 text-center text-muted"
          >
            <div class="flex items-center justify-center gap-2">
              <svg class="w-4 h-4 animate-spin" viewBox="0 0 24 24" fill="none">
                <circle
                  class="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  stroke-width="4"
                />
                <path
                  class="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
                />
              </svg>
              Loading…
            </div>
          </td>
        </tr>
        <slot v-else />
        <tr v-if="!loading && !$slots.default">
          <td
            :colspan="headers.length"
            class="px-5 py-10 text-center text-muted"
          >
            {{ empty ?? "No records found." }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
