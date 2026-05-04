<script setup lang="ts">
import type { Item, ItemMetadata } from "~/types/item.types";

const route = useRoute();
const router = useRouter();
const id = route.params.id as string;

const { data: item, pending: loading, error } = await useItemDetail(id);

useHead(
  computed(() => ({
    title: item.value
      ? `${item.value.name} — Luxury Hotel`
      : "Room Detail — Luxury Hotel",
  })),
);

const metadata = computed<ItemMetadata>(() => item.value?.metadata ?? {});
const images = computed<string[]>(() => item.value?.images ?? []);
const badges = computed<string[]>(() => metadata.value.badges ?? []);
const rating = computed<number>(() => metadata.value.rating ?? 0);

const formattedPrice = computed(() => {
  if (!item.value) return "";
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: item.value.currency || "USD",
    minimumFractionDigits: 0,
  }).format(item.value.price / 100);
});

const activeImage = ref(0);
const checkIn = ref("");
const checkOut = ref("");
const guestCount = ref("1");

const cartStore = useCartStore();
const { create } = useOrders();

const booking = ref(false);
const bookingError = ref<string | null>(null);

async function bookNow() {
  if (!item.value || !checkIn.value || !checkOut.value) return;

  const nights = Math.max(
    1,
    Math.ceil(
      (new Date(checkOut.value).getTime() - new Date(checkIn.value).getTime()) /
        86_400_000,
    ),
  );

  booking.value = true;
  bookingError.value = null;

  try {
    const order = await create({
      items: [{ item_id: item.value.id, quantity: nights }],
      notes: "",
      metadata: {
        room_name: item.value.name,
        check_in: checkIn.value,
        check_out: checkOut.value,
        guests: parseInt(guestCount.value),
        nights,
        price_per_night: item.value.price,
      },
    });

    cartStore.add({
      itemId: item.value.id,
      name: item.value.name,
      price: item.value.price / 100,
      quantity: 1,
      checkIn: checkIn.value,
      checkOut: checkOut.value,
      nights,
      image: images.value[0],
    });

    await router.push(`/orders/${order!.id}`);
  } catch (e: any) {
    bookingError.value = e.message ?? "Booking failed. Please try again.";
  } finally {
    booking.value = false;
  }
}
</script>

<template>
  <div>
    <!-- Error -->
    <div
      v-if="error"
      class="section-py section-px max-w-4xl mx-auto text-center text-muted font-sans"
    >
      <Icon
        name="lucide:door-closed"
        class="w-12 h-12 mx-auto mb-4 opacity-40"
      />
      <p class="font-display text-xl text-brand-900 mb-2">Room not found</p>
      <p class="text-sm mb-4">
        {{ (error as any).message ?? "This room could not be loaded." }}
      </p>
      <NuxtLink to="/items">
        <UiButton variant="outline" size="sm">Back to Rooms</UiButton>
      </NuxtLink>
    </div>

    <!-- Loading skeleton -->
    <div
      v-else-if="loading"
      class="section-py section-px max-w-7xl mx-auto animate-pulse"
    >
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-12">
        <div class="aspect-[4/3] bg-surface-200 rounded-2xl" />
        <div class="space-y-4 pt-2">
          <div class="h-4 bg-surface-200 rounded w-1/4" />
          <div class="h-8 bg-surface-200 rounded w-3/4" />
          <div class="h-4 bg-surface-200 rounded w-full" />
          <div class="h-4 bg-surface-200 rounded w-5/6" />
          <div class="h-4 bg-surface-200 rounded w-2/3" />
          <div class="grid grid-cols-4 gap-3 pt-4">
            <div
              v-for="n in 4"
              :key="n"
              class="h-20 bg-surface-200 rounded-xl"
            />
          </div>
          <div class="h-44 bg-surface-200 rounded-2xl mt-6" />
        </div>
      </div>
    </div>

    <!-- Item content -->
    <template v-else-if="item">
      <!-- Breadcrumb -->
      <div class="section-px max-w-7xl mx-auto pt-8 pb-2">
        <nav class="flex items-center gap-2 text-xs font-sans text-muted">
          <NuxtLink to="/" class="hover:text-forest transition-colors"
            >Home</NuxtLink
          >
          <Icon name="lucide:chevron-right" class="w-3 h-3" />
          <NuxtLink to="/items" class="hover:text-forest transition-colors"
            >Rooms</NuxtLink
          >
          <Icon name="lucide:chevron-right" class="w-3 h-3" />
          <span class="text-brand-900">{{ item.name }}</span>
        </nav>
      </div>

      <section class="section-px max-w-7xl mx-auto section-py">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-14">
          <!-- Gallery -->
          <div>
            <div
              class="rounded-2xl overflow-hidden aspect-[4/3] bg-surface-200 mb-3"
            >
              <img
                v-if="images[activeImage]"
                :src="images[activeImage]"
                :alt="item.name"
                class="w-full h-full object-cover"
              />
              <div
                v-else
                class="w-full h-full flex items-center justify-center"
              >
                <Icon name="lucide:image" class="w-16 h-16 text-surface-300" />
              </div>
            </div>
            <div
              v-if="images.length > 1"
              class="flex gap-2 overflow-x-auto pb-1"
            >
              <button
                v-for="(img, i) in images"
                :key="i"
                class="shrink-0 w-20 h-16 rounded-lg overflow-hidden border-2 transition-all"
                :class="
                  i === activeImage
                    ? 'border-accent'
                    : 'border-transparent opacity-60 hover:opacity-100'
                "
                @click="activeImage = i"
              >
                <img
                  :src="img"
                  :alt="`${item.name} ${i + 1}`"
                  class="w-full h-full object-cover"
                />
              </button>
            </div>
          </div>

          <!-- Details -->
          <div>
            <!-- Badges -->
            <div v-if="badges.length" class="flex flex-wrap gap-2 mb-4">
              <span
                v-for="badge in badges"
                :key="badge"
                class="bg-forest/10 text-forest text-[10px] uppercase tracking-widest font-sans px-2.5 py-1 rounded"
              >
                {{ badge }}
              </span>
            </div>

            <h1
              class="font-display text-display-md text-brand-900 font-light mb-3"
            >
              {{ item.name }}
            </h1>

            <!-- Rating -->
            <div v-if="rating" class="flex items-center gap-1.5 mb-5">
              <Icon
                v-for="i in 5"
                :key="i"
                name="lucide:star"
                class="w-4 h-4"
                :class="
                  i <= Math.round(rating)
                    ? 'text-accent fill-accent'
                    : 'text-surface-300'
                "
              />
              <span class="text-sm font-sans text-muted ml-1">
                {{ rating.toFixed(1) }} / 5.0
              </span>
            </div>

            <p class="text-sm text-muted font-sans leading-relaxed mb-8">
              {{ item.description }}
            </p>

            <!-- Room stats -->
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-4 mb-8">
              <div
                v-if="metadata.beds"
                class="text-center p-3 rounded-xl bg-surface-100"
              >
                <Icon
                  name="lucide:bed-double"
                  class="w-5 h-5 text-forest mx-auto mb-1"
                />
                <div class="font-display text-lg text-brand-900">
                  {{ metadata.beds }}
                </div>
                <div
                  class="text-[10px] uppercase tracking-wider text-muted font-sans"
                >
                  Beds
                </div>
              </div>
              <div
                v-if="metadata.baths"
                class="text-center p-3 rounded-xl bg-surface-100"
              >
                <Icon
                  name="lucide:bath"
                  class="w-5 h-5 text-forest mx-auto mb-1"
                />
                <div class="font-display text-lg text-brand-900">
                  {{ metadata.baths }}
                </div>
                <div
                  class="text-[10px] uppercase tracking-wider text-muted font-sans"
                >
                  Baths
                </div>
              </div>
              <div
                v-if="metadata.sqft"
                class="text-center p-3 rounded-xl bg-surface-100"
              >
                <Icon
                  name="lucide:square"
                  class="w-5 h-5 text-forest mx-auto mb-1"
                />
                <div class="font-display text-lg text-brand-900">
                  {{ metadata.sqft }}
                </div>
                <div
                  class="text-[10px] uppercase tracking-wider text-muted font-sans"
                >
                  sqft
                </div>
              </div>
              <div
                v-if="metadata.floor"
                class="text-center p-3 rounded-xl bg-surface-100"
              >
                <Icon
                  name="lucide:building-2"
                  class="w-5 h-5 text-forest mx-auto mb-1"
                />
                <div class="font-display text-lg text-brand-900">
                  {{ metadata.floor }}
                </div>
                <div
                  class="text-[10px] uppercase tracking-wider text-muted font-sans"
                >
                  Floor
                </div>
              </div>
            </div>

            <!-- View -->
            <div
              v-if="metadata.view"
              class="flex items-center gap-2 mb-8 text-sm text-muted font-sans"
            >
              <Icon name="lucide:eye" class="w-4 h-4 text-forest" />
              <span
                >View:
                <strong class="text-brand-900">{{
                  metadata.view
                }}</strong></span
              >
            </div>

            <!-- Booking panel -->
            <div
              class="border border-surface-200 rounded-2xl p-6 bg-surface-50"
            >
              <div class="flex items-baseline gap-1 mb-5">
                <span class="font-display text-3xl text-brand-900">{{
                  formattedPrice
                }}</span>
                <span class="text-sm text-muted font-sans">/ night</span>
              </div>

              <div class="grid grid-cols-2 gap-3 mb-4">
                <div>
                  <label
                    class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                  >
                    Check In
                  </label>
                  <input v-model="checkIn" type="date" class="input" />
                </div>
                <div>
                  <label
                    class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                  >
                    Check Out
                  </label>
                  <input v-model="checkOut" type="date" class="input" />
                </div>
              </div>

              <div class="mb-5">
                <label
                  class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
                >
                  Guests
                </label>
                <select v-model="guestCount" class="input appearance-none">
                  <option value="1">1 Adult</option>
                  <option value="2">2 Adults</option>
                  <option value="3">2 Adults, 1 Child</option>
                  <option value="4">2 Adults, 2 Children</option>
                </select>
              </div>

              <UiAlert
                v-if="bookingError"
                variant="error"
                :message="bookingError"
                class="mb-4 text-xs"
              />

              <UiButton
                class="w-full justify-center"
                :disabled="!checkIn || !checkOut || booking"
                :loading="booking"
                @click="bookNow"
              >
                {{ booking ? "Booking…" : "Reserve Now" }}
              </UiButton>

              <p class="text-[10px] text-muted font-sans text-center mt-3">
                Free cancellation up to 24 hours before check-in
              </p>
            </div>
          </div>
        </div>
      </section>
    </template>
  </div>
</template>
