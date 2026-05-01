<script setup lang="ts">
defineProps<{ modelValue?: string; placeholder?: string; loading?: boolean }>();
defineEmits<{ "update:modelValue": [v: string] }>();
</script>

<template>
  <div class="relative">
    <Icon
      v-if="!loading"
      name="lucide:search"
      class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-muted pointer-events-none"
    />
    <svg
      v-else
      class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-muted animate-spin"
      viewBox="0 0 24 24"
      fill="none"
    >
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
    <input
      :value="modelValue"
      :placeholder="placeholder ?? 'Search…'"
      class="input pl-9"
      @input="
        $emit('update:modelValue', ($event.target as HTMLInputElement).value)
      "
    />
  </div>
</template>
