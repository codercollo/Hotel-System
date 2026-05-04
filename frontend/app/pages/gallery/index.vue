<script setup lang="ts">
useHead({ title: "Gallery — Luxury Hotel" });

const categories = ["All", "Rooms", "Dining", "Pool", "Spa", "Events"];
const active = ref("All");
const lightbox = ref<{ url: string; label: string; category: string } | null>(
  null,
);

// Fetch real uploads from backend
const api = useApi();
const { data: uploads, pending: loading } = await useAsyncData(
  "gallery-uploads",
  () =>
    api
      .get<
        { id: string; url: string; original_name: string; mime_type: string }[]
      >("/api/v1/uploads")
      .catch(() => []),
);

// Map uploads to gallery items with inferred category from filename
const inferCategory = (name: string): string => {
  const n = name.toLowerCase();
  if (n.includes("room") || n.includes("suite")) return "Rooms";
  if (n.includes("dine") || n.includes("dining") || n.includes("restaurant"))
    return "Dining";
  if (n.includes("pool")) return "Pool";
  if (n.includes("spa")) return "Spa";
  if (n.includes("event") || n.includes("ballroom")) return "Events";
  return "Rooms";
};

// Merge real uploads with static placeholders so gallery is never empty
const staticPlaceholders = [
  {
    id: "p1",
    url: "https://images.unsplash.com/photo-1590490360182-c33d57733427?auto=format&fit=crop&w=1400&q=80",
    original_name: "Minimal Luxury Bedroom",
    category: "Rooms",
    aspect: "landscape",
    isPlaceholder: true,
  },
  {
    id: "p2",
    url: "https://images.unsplash.com/photo-1505693416388-ac5ce068fe85?auto=format&fit=crop&w=900&q=80",
    original_name: "Elegant Suite Bed",
    category: "Rooms",
    aspect: "portrait",
    isPlaceholder: true,
  },
  {
    id: "p3",
    url: "https://images.unsplash.com/photo-1551218808-94e220e084d2?auto=format&fit=crop&w=1400&q=80",
    original_name: "Luxury Bedroom Interior",
    category: "Rooms",
    aspect: "landscape",
    isPlaceholder: true,
  },
  {
    id: "p4",
    url: "https://images.unsplash.com/photo-1520250497591-112f2f40a3f4?auto=format&fit=crop&w=1400&q=80",
    original_name: "Modern Hotel Room",
    category: "Rooms",
    aspect: "landscape",
    isPlaceholder: true,
  },
  {
    id: "p5",
    url: "https://images.unsplash.com/photo-1615874959474-d609969a20ed?auto=format&fit=crop&w=900&q=80",
    original_name: "Cozy Interior Lounge Bed",
    category: "Rooms",
    aspect: "portrait",
    isPlaceholder: true,
  },
  {
    id: "p6",
    url: "https://images.unsplash.com/photo-1505691938895-1758d7feb511?auto=format&fit=crop&w=1400&q=80",
    original_name: "Luxury White Bedding",
    category: "Rooms",
    aspect: "landscape",
    isPlaceholder: true,
  },
];

const galleryItems = computed(() => {
  const real = (uploads.value ?? [])
    .filter((u) => u.mime_type.startsWith("image/"))
    .map((u, i) => ({
      id: u.id,
      url: u.url,
      original_name: u.original_name,
      category: inferCategory(u.original_name),
      aspect: i % 3 === 1 ? "portrait" : "landscape",
      isPlaceholder: false,
    }));

  return real.length > 0 ? real : staticPlaceholders;
});

const filtered = computed(() =>
  active.value === "All"
    ? galleryItems.value
    : galleryItems.value.filter((g) => g.category === active.value),
);
</script>

<template>
  <div>
    <!-- Hero -->
    <div class="bg-brand-900 py-20 section-px text-center">
      <p class="eyebrow text-accent mb-3">Our World</p>
      <h1 class="font-display text-display-lg text-white font-light">
        Gallery
      </h1>
      <p class="text-white/50 font-sans text-sm mt-4 max-w-md mx-auto">
        A visual journey through our spaces, cuisine, and experiences.
      </p>
    </div>

    <section class="section-py section-px max-w-7xl mx-auto">
      <!-- Category filters -->
      <div class="flex flex-wrap justify-center gap-2 mb-10">
        <button
          v-for="cat in categories"
          :key="cat"
          :class="[
            'px-5 py-2 rounded-full text-xs font-sans font-medium uppercase tracking-widest transition-all border',
            active === cat
              ? 'bg-forest text-white border-forest'
              : 'border-surface-200 text-muted hover:border-forest hover:text-forest',
          ]"
          @click="active = cat"
        >
          {{ cat }}
        </button>
      </div>

      <!-- Loading -->
      <div
        v-if="loading"
        class="columns-1 sm:columns-2 lg:columns-3 gap-4 space-y-4"
      >
        <div
          v-for="i in 6"
          :key="i"
          :class="[
            'break-inside-avoid rounded-xl bg-surface-200 animate-pulse',
            i % 3 === 1 ? 'aspect-[3/4]' : 'aspect-[4/3]',
          ]"
        />
      </div>

      <!-- Grid -->
      <div v-else class="columns-1 sm:columns-2 lg:columns-3 gap-4 space-y-4">
        <div
          v-for="item in filtered"
          :key="item.id"
          class="break-inside-avoid rounded-xl overflow-hidden bg-surface-200 relative cursor-pointer group"
          :class="item.aspect === 'portrait' ? 'aspect-[3/4]' : 'aspect-[4/3]'"
          @click="
            lightbox = {
              url: item.url,
              label: item.original_name,
              category: item.category,
            }
          "
        >
          <!-- Real image -->
          <img
            v-if="item.url"
            :src="item.url"
            :alt="item.original_name"
            class="w-full h-full object-cover"
            loading="lazy"
            onerror="this.style.display = 'none'"
          />

          <!-- Placeholder -->
          <div
            v-else
            class="absolute inset-0 bg-gradient-to-br from-forest/30 to-brand-900/20 flex items-center justify-center"
          >
            <Icon name="lucide:image" class="w-12 h-12 text-white/30" />
          </div>

          <!-- Hover overlay -->
          <div
            class="absolute inset-0 bg-brand-900/50 opacity-0 group-hover:opacity-100 transition-all duration-300 flex items-end p-5"
          >
            <div>
              <UiBadge variant="gold" class="mb-2">{{ item.category }}</UiBadge>
              <p class="text-white font-display text-lg">
                {{ item.original_name }}
              </p>
            </div>
          </div>
        </div>
      </div>

      <div
        v-if="!loading && filtered.length === 0"
        class="text-center py-16 text-muted font-sans"
      >
        <Icon name="lucide:images" class="w-10 h-10 mx-auto mb-3 opacity-30" />
        <p>No images in this category yet.</p>
      </div>
    </section>

    <!-- Lightbox -->
    <UiModal
      :open="!!lightbox"
      :title="lightbox?.label ?? ''"
      @close="lightbox = null"
    >
      <div class="rounded-xl overflow-hidden mb-4 bg-surface-200">
        <img
          v-if="lightbox?.url"
          :src="lightbox.url"
          :alt="lightbox.label"
          class="w-full max-h-[60vh] object-contain"
        />
        <div v-else class="aspect-video flex items-center justify-center">
          <Icon name="lucide:image" class="w-16 h-16 text-surface-300" />
        </div>
      </div>
      <p class="text-sm text-muted font-sans">{{ lightbox?.category }}</p>
    </UiModal>
  </div>
</template>
