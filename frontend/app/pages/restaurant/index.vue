<script setup lang="ts">
import type { Item } from "~/types/item.types";

useHead({ title: "Restaurant & Dining" });

const api = useApi();
const { data: backendItems } = await useAsyncData<Item[]>(
  "restaurant-items",
  () => api.get<Item[]>("/api/v1/items", { status: "active" }).catch(() => []),
);

const foodItems = computed(() =>
  (backendItems.value ?? []).filter(
    (i) =>
      (i.metadata?.type as string | undefined) === "menu" ||
      (i.metadata?.category as string | undefined)?.includes("food"),
  ),
);

const staticMenuItems = [
  {
    id: "1",
    name: "Continental Breakfast",
    desc: "Croissants, seasonal fruit, yoghurt, freshly squeezed orange juice.",
    price: 2800,
    category: "breakfast",
    image: "",
  },
  {
    id: "2",
    name: "Full English Breakfast",
    desc: "Eggs your way, crispy bacon, grilled tomatoes, mushrooms, toast.",
    price: 3500,
    category: "breakfast",
    image: "",
  },
  {
    id: "3",
    name: "Pan-Seared Salmon",
    desc: "Atlantic salmon, lemon butter sauce, seasonal vegetables, wild rice.",
    price: 6800,
    category: "lunch",
    image: "",
  },
  {
    id: "4",
    name: "Grilled Chicken Club",
    desc: "Sourdough, free-range chicken, avocado, smoked bacon, house aioli.",
    price: 4200,
    category: "lunch",
    image: "",
  },
  {
    id: "5",
    name: "Wagyu Tenderloin",
    desc: "200g A5 wagyu, truffle jus, dauphinoise potato, asparagus.",
    price: 18500,
    category: "dinner",
    image: "",
  },
  {
    id: "6",
    name: "Lobster Bisque",
    desc: "Maine lobster, cognac cream, chive oil, artisan bread.",
    price: 7200,
    category: "dinner",
    image: "",
  },
  {
    id: "7",
    name: "Chocolate Fondant",
    desc: "Dark chocolate, warm centre, vanilla bean ice cream, berry coulis.",
    price: 2400,
    category: "desserts",
    image: "",
  },
  {
    id: "8",
    name: "Crème Brûlée",
    desc: "Classic vanilla custard, caramelised sugar crust, fresh raspberries.",
    price: 2200,
    category: "desserts",
    image: "",
  },
];

const staticFoodImages = [
  {
    id: "p1",
    url: "https://images.unsplash.com/photo-1540189549336-e6e99c3679fe?auto=format&fit=crop&w=1400&q=80",
    name: "Fine Dining Table Spread",
  },
  {
    id: "p2",
    url: "https://images.unsplash.com/photo-1606756790138-261d2b21cd75?auto=format&fit=crop&w=1400&q=80",
    name: "Burger & Fries Platter",
  },
  {
    id: "p3",
    url: "https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=1400&q=80",
    name: "Gourmet Plated Dish",
  },
  {
    id: "p4",
    url: "https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=1400&q=80",
    name: "Roasted Meat Dish",
  },
  {
    id: "p5",
    url: "https://images.unsplash.com/photo-1544510808-91bcbee1df55?auto=format&fit=crop&w=1400&q=80",
    name: "Chef Prepared Meal",
  },
  {
    id: "p6",
    url: "https://images.unsplash.com/photo-1550547660-d9450f859349?auto=format&fit=crop&w=1400&q=80",
    name: "Burger & Fries Close-up",
  },
  {
    id: "p7",
    url: "https://images.unsplash.com/photo-1529692236671-f1f6cf9683ba?auto=format&fit=crop&w=1400&q=80",
    name: "Restaurant Plate Dish",
  },
];
const allItems = computed(() => {
  const foods = foodItems.value ?? [];
  const images = staticFoodImages ?? [];
  const menus = staticMenuItems ?? [];

  const fallbackImage = (index: number) =>
    images.length > 0 ? (images[index % images.length]?.url ?? "") : "";

  if (foods.length > 0) {
    return foods.map((i, index) => ({
      id: i.id,
      name: i.name,
      desc: i.description,
      price: i.price,
      category: (i.metadata?.category as string) ?? "dinner",
      image: i.images?.[0] || fallbackImage(index),
    }));
  }

  return menus.map((i, index) => ({
    ...i,
    image: fallbackImage(index),
  }));
});

const menuCategories = [
  { label: "All", id: "all" },
  { label: "Breakfast", id: "breakfast" },
  { label: "Lunch", id: "lunch" },
  { label: "Dinner", id: "dinner" },
  { label: "Desserts", id: "desserts" },
];

const activeCategory = ref("all");
const search = ref("");

const filtered = computed(() => {
  let items =
    activeCategory.value === "all"
      ? allItems.value
      : allItems.value.filter((m) => m.category === activeCategory.value);
  if (search.value) {
    items = items.filter((m) =>
      m.name.toLowerCase().includes(search.value.toLowerCase()),
    );
  }
  return items;
});

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });

const fallback = "/images/fallback.webp";

function onImgError(e: Event) {
  const img = e.target as HTMLImageElement | null;
  if (img) img.src = fallback;
}
</script>

<template>
  <div>
    <!-- Hero -->
    <div
      class="bg-forest py-20 section-px text-center relative overflow-hidden"
    >
      <div
        class="absolute inset-0 opacity-10"
        style="
          background: repeating-linear-gradient(
            45deg,
            white 0,
            white 1px,
            transparent 1px,
            transparent 40px
          );
        "
      />
      <div class="relative z-10">
        <div class="ornament mb-3"><span class="text-accent">✦</span></div>
        <p class="eyebrow text-accent mb-3">Fine Dining</p>
        <h1 class="font-display text-display-lg text-white font-light">
          Restaurant &amp; Dining
        </h1>
        <p
          class="text-white/60 font-sans text-sm mt-4 max-w-md mx-auto leading-relaxed"
        >
          Experience culinary artistry crafted from the finest seasonal
          ingredients in an atmosphere of refined elegance.
        </p>
      </div>
    </div>

    <section class="section-py section-px max-w-7xl mx-auto">
      <!-- Filters + Search -->
      <div
        class="flex flex-col sm:flex-row items-start sm:items-center gap-4 mb-10"
      >
        <div class="flex flex-wrap gap-2">
          <button
            v-for="cat in menuCategories"
            :key="cat.id"
            :class="[
              'px-4 py-1.5 rounded-full text-xs font-sans font-medium uppercase tracking-widest transition-all border',
              activeCategory === cat.id
                ? 'bg-forest text-white border-forest'
                : 'border-surface-200 text-muted hover:border-forest hover:text-forest',
            ]"
            @click="activeCategory = cat.id"
          >
            {{ cat.label }}
          </button>
        </div>
        <div class="sm:ml-auto w-full sm:w-60">
          <SearchInput v-model="search" placeholder="Search menu…" />
        </div>
      </div>

      <!-- Grid -->
      <div
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-5"
      >
        <div
          v-for="item in filtered"
          :key="item.id"
          class="card group cursor-pointer"
        >
          <div
            class="aspect-[4/3] bg-surface-200 relative overflow-hidden rounded-t-2xl"
          >
            <img
              v-if="item.image"
              :src="item.image"
              :alt="item.name"
              class="w-full h-full object-cover"
              loading="lazy"
              @error="onImgError"
            />
            <div
              v-else
              class="absolute inset-0 bg-gradient-to-br from-forest/20 to-brand-900/10 flex items-center justify-center"
            >
              <Icon name="lucide:utensils" class="w-12 h-12 text-surface-300" />
            </div>
          </div>
          <div class="p-4">
            <div class="flex items-start justify-between gap-2 mb-2">
              <h3 class="font-display text-lg text-brand-900 leading-snug">
                {{ item.name }}
              </h3>
              <span
                class="font-display text-base font-semibold text-forest shrink-0"
                >{{ formatPrice(item.price) }}</span
              >
            </div>
            <p class="text-xs text-muted font-sans leading-relaxed">
              {{ item.desc }}
            </p>
            <div class="mt-3">
              <UiBadge variant="muted">{{ item.category }}</UiBadge>
            </div>
          </div>
        </div>

        <div v-if="filtered.length === 0" class="col-span-full">
          <EmptyState
            icon="lucide:utensils"
            title="No dishes found"
            :description="
              search
                ? `No menu items match '${search}'.`
                : 'No items in this category.'
            "
          />
        </div>
      </div>

      <!-- Reservation CTA -->
      <div class="mt-16 bg-forest rounded-2xl p-10 text-center">
        <p class="eyebrow text-accent mb-3">Reserve a Table</p>
        <h2 class="font-display text-display-md text-white font-light mb-4">
          An Unforgettable Dining Experience Awaits
        </h2>
        <p class="text-white/60 font-sans text-sm mb-8 max-w-sm mx-auto">
          Our culinary team is ready to create a bespoke dining experience for
          you and your guests.
        </p>
        <NuxtLink to="/contact">
          <UiButton variant="gold" size="md">Make a Reservation</UiButton>
        </NuxtLink>
      </div>
    </section>
  </div>
</template>
