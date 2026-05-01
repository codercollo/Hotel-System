<script setup lang="ts">
useHead({ title: "Rooms & Suites" });
const search = ref("");
const loading = ref(false);

const rooms = [
  {
    id: "r1",
    name: "Deluxe Room",
    price: 25000,
    images: [],
    metadata: {
      beds: 1,
      baths: 1,
      sqft: 300,
      rating: 4.9,
      badges: ["Luxury Room"],
    },
    description: "Elegant room with garden views and modern amenities.",
  },
  {
    id: "r2",
    name: "The Pearl Suite",
    price: 45000,
    images: [],
    metadata: {
      beds: 2,
      baths: 2,
      sqft: 400,
      rating: 5.0,
      badges: ["Luxury Suites"],
    },
    description:
      "Signature suite with panoramic city views and butler service.",
  },
  {
    id: "r3",
    name: "Golden Executive",
    price: 55000,
    images: [],
    metadata: {
      beds: 3,
      baths: 2,
      sqft: 700,
      rating: 4.9,
      badges: ["Premium"],
    },
    description: "Executive suite with private lounge and floor-only access.",
  },
  {
    id: "r4",
    name: "Classic Room",
    price: 18000,
    images: [],
    metadata: {
      beds: 1,
      baths: 1,
      sqft: 240,
      rating: 4.7,
      badges: ["Classic"],
    },
    description: "Comfortable classic room for the business traveller.",
  },
  {
    id: "r5",
    name: "Family Suite",
    price: 38000,
    images: [],
    metadata: { beds: 2, baths: 2, sqft: 550, rating: 4.8, badges: ["Family"] },
    description: "Spacious suite ideal for families with children.",
  },
  {
    id: "r6",
    name: "Honeymoon Suite",
    price: 65000,
    images: [],
    metadata: {
      beds: 1,
      baths: 2,
      sqft: 620,
      rating: 5.0,
      badges: ["Romantic"],
    },
    description: "Secluded romantic suite with private jacuzzi and terrace.",
  },
];

const filtered = computed(() =>
  search.value
    ? rooms.filter((r) =>
        r.name.toLowerCase().includes(search.value.toLowerCase()),
      )
    : rooms,
);
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
      <ItemGrid :items="filtered" :loading="loading" :cols="3" />
      <div
        v-if="!loading && filtered.length === 0"
        class="text-center py-16 text-muted font-sans"
      >
        <Icon
          name="lucide:search-x"
          class="w-10 h-10 mx-auto mb-3 opacity-40"
        />
        <p>No rooms found matching "{{ search }}"</p>
      </div>
    </section>
  </div>
</template>
