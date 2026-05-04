<script setup lang="ts">
import type { Item } from "~/types/item.types";

useHead({ title: "Rooms & Suites" });

const search = ref("");

const {
  data: allItems,
  pending: loading,
  error,
  refresh,
} = await useItemList();

// Client-side search filter
const filtered = computed<Item[]>(() => {
  const base = allItems.value ?? [];
  if (!search.value) return base;
  const q = search.value.toLowerCase();
  return base.filter(
    (r) =>
      r.name.toLowerCase().includes(q) ||
      r.description.toLowerCase().includes(q),
  );
});
</script>

<template>
  <div>
    <div class="bg-forest py-16 section-px text-center">
      <p class="eyebrow text-accent mb-3">Rooms & Suites</p>
      <h1 class="font-display text-display-lg text-white font-light">
        Luxury Rooms &amp; Suites
      </h1>
    </div>

    <section class="section-py section-px max-w-7xl mx-auto">
      <div class="mb-8 max-w-md">
        <ItemSearchBar
          v-model="search"
          placeholder="Search rooms & suites…"
          @search="search = $event"
        />
      </div>

      <!-- Error state -->
      <div v-if="error" class="text-center py-12 text-muted font-sans text-sm">
        <Icon
          name="lucide:wifi-off"
          class="w-10 h-10 mx-auto mb-3 opacity-40"
        />
        <p class="mb-3">Could not load rooms. Please try again.</p>
        <!-- Fix: wrap refresh in arrow function so PointerEvent is not passed as opts -->
        <button
          class="text-forest underline underline-offset-2 text-sm"
          @click="() => refresh()"
        >
          Retry
        </button>
      </div>

      <ItemGrid v-else :items="filtered" :loading="loading" :cols="3" />

      <div
        v-if="!loading && !error && filtered.length === 0"
        class="text-center py-16 text-muted font-sans"
      >
        <Icon
          name="lucide:search-x"
          class="w-10 h-10 mx-auto mb-3 opacity-40"
        />
        <p>
          {{
            search
              ? `No rooms found matching "${search}"`
              : "No rooms available right now."
          }}
        </p>
      </div>
    </section>
  </div>
</template>
