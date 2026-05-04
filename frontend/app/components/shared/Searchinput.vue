<script setup lang="ts">
import type { SearchResult } from "~/types/search.types";

const props = defineProps<{
  modelValue?: string;
  placeholder?: string;
  loading?: boolean;
  results?: SearchResult[];
  showDropdown?: boolean;
}>();

const emit = defineEmits<{
  "update:modelValue": [v: string];
  select: [r: SearchResult];
  clear: [];
}>();

const router = useRouter();

const handleSelect = (r: SearchResult) => {
  emit("select", r);
  if (r.type === "item") router.push(`/items/${r.id}`);
};

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });
</script>

<template>
  <div class="relative">
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
      <button
        v-if="modelValue"
        class="absolute right-3 top-1/2 -translate-y-1/2 text-muted hover:text-brand-900"
        @click="$emit('clear')"
      >
        <Icon name="lucide:x" class="w-3.5 h-3.5" />
      </button>
    </div>

    <!-- Results dropdown -->
    <div
      v-if="showDropdown && results && results.length > 0"
      class="absolute top-full left-0 right-0 mt-1 bg-white rounded-xl shadow-card-hover border border-surface-200 z-50 overflow-hidden"
    >
      <ul class="divide-y divide-surface-100 max-h-64 overflow-y-auto">
        <li
          v-for="r in results"
          :key="r.id"
          class="flex items-center gap-3 px-4 py-3 hover:bg-surface-50 cursor-pointer transition-colors"
          @click="handleSelect(r)"
        >
          <Icon name="lucide:bed-double" class="w-4 h-4 text-forest shrink-0" />
          <div class="flex-1 min-w-0">
            <p class="text-sm font-sans font-medium text-brand-900 truncate">
              {{ r.name }}
            </p>
            <p class="text-xs text-muted font-sans truncate">
              {{ r.description }}
            </p>
          </div>
          <span
            v-if="r.price"
            class="text-xs font-display text-brand-900 shrink-0"
          >
            {{ formatPrice(r.price) }}
          </span>
        </li>
      </ul>
    </div>
  </div>
</template>
