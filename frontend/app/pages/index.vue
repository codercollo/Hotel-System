<script setup lang="ts">
useHead({ title: "Home — Where Luxury Meets Excellence" });

interface Facility {
  icon: string;
  label: string;
  desc: string;
}

const facilities: Facility[] = [
  {
    icon: "lucide:coffee",
    label: "Breakfast Included",
    desc: "Start your day with our premium buffet breakfast, included with every stay.",
  },
  {
    icon: "lucide:waves",
    label: "Swimming Pool",
    desc: "Dive into our temperature-controlled outdoor pool with stunning views.",
  },
  {
    icon: "lucide:wifi",
    label: "High Speed Wifi",
    desc: "Stay connected with complimentary high-speed internet throughout.",
  },
  {
    icon: "lucide:sparkles",
    label: "Spa & Wellness",
    desc: "Rejuvenate your body and mind at our world-class spa and wellness centre.",
  },
  {
    icon: "lucide:car",
    label: "Pick Up & Drop",
    desc: "Complimentary airport transfers in our luxury fleet vehicles.",
  },
  {
    icon: "lucide:dumbbell",
    label: "Fitness Hub",
    desc: "State-of-the-art gym open 24/7 for our guests.",
  },
];

const featuredFacility = computed<Facility>(() => facilities[0]!);

const tickerItems = [
  "Breakfast Included",
  "Swimming Pool",
  "High Speed Wifi",
  "Spa & Wellness",
  "Airport Transfer",
  "Fitness Hub",
];

const stats = [
  { value: "50+", label: "Luxury Rooms" },
  { value: "60,000+", label: "Happy Guests" },
  { value: "99%", label: "Guest Satisfaction" },
];

// Simulated rooms — replaced by real API in Phase 3
const featuredRooms = [
  {
    id: "room-1",
    name: "Deluxe Room",
    description: "Spacious room with garden views and premium amenities.",
    price: 25000,
    images: [],
    rating: 4.9,
    badges: ["Luxury Room"],
    metadata: {
      beds: 1,
      baths: 1,
      sqft: 300,
      rating: 4.9,
      badges: ["Luxury Room"],
    },
  },
  {
    id: "room-2",
    name: "The Pearl Suite",
    description:
      "Our signature suite with panoramic city views and butler service.",
    price: 45000,
    images: [],
    rating: 5.0,
    badges: ["Luxury Suites"],
    metadata: {
      beds: 2,
      baths: 2,
      sqft: 400,
      rating: 5.0,
      badges: ["Luxury Suites"],
    },
  },
  {
    id: "room-3",
    name: "Golden Executive",
    description:
      "Executive suite with workspace, lounge and exclusive floor access.",
    price: 55000,
    images: [],
    rating: 4.9,
    badges: ["Premium"],
    metadata: {
      beds: 3,
      baths: 2,
      sqft: 700,
      rating: 4.9,
      badges: ["Premium"],
    },
  },
];

// Booking widget state
const checkIn = ref("");
const checkOut = ref("");
const rooms = ref("1");
const guests = ref("1");
</script>

<template>
  <div>
    <!-- ══ HERO ══════════════════════════════════════════════════════════ -->
    <section class="relative min-h-[90vh] flex items-center overflow-hidden">
      <!-- Background image -->
      <div class="absolute inset-0 bg-brand-900">
        <div
          class="absolute inset-0 opacity-40"
          style="
            background: linear-gradient(
              135deg,
              rgb(25, 55, 48) 0%,
              rgb(38, 79, 72) 50%,
              rgb(52, 58, 32) 100%
            );
          "
        />
        <!-- Pattern overlay -->
        <div
          class="absolute inset-0 opacity-5"
          style="
            background-image: repeating-linear-gradient(
              45deg,
              white 0px,
              white 1px,
              transparent 1px,
              transparent 50%
            );
            background-size: 30px 30px;
          "
        />
      </div>

      <div
        class="relative z-10 max-w-7xl mx-auto section-px w-full pt-20 pb-32"
      >
        <div class="max-w-2xl">
          <div class="flex items-center gap-2 mb-6 animate-fade-up">
            <Icon
              v-for="i in 5"
              :key="i"
              name="lucide:star"
              class="w-4 h-4 text-accent fill-accent"
            />
          </div>

          <p class="eyebrow text-accent mb-4 animate-fade-up delay-100">
            Where Luxury Meets Excellence
          </p>

          <h1
            class="font-display text-display-xl text-white font-light leading-tight mb-6 animate-fade-up delay-200"
          >
            The Ultimate Luxury Hotel Experience in New Jersey
          </h1>

          <p
            class="text-white/60 text-sm font-sans leading-relaxed mb-10 max-w-lg animate-fade-up delay-300"
          >
            Indulge in unrivalled comfort, world-class dining, and breathtaking
            surroundings. Every detail crafted for your perfect stay.
          </p>

          <div
            class="flex flex-wrap items-center gap-4 animate-fade-up delay-400"
          >
            <NuxtLink to="/items">
              <UiButton size="md" variant="gold">Discover More</UiButton>
            </NuxtLink>
            <button class="btn-ghost flex items-center gap-2">
              <div
                class="w-10 h-10 rounded-full border-2 border-white/40 flex items-center justify-center"
              >
                <Icon name="lucide:play" class="w-4 h-4 ml-0.5" />
              </div>
              Watch Video
            </button>
          </div>
        </div>
      </div>

      <!-- ── Booking widget ───────────────────────────────────────────── -->
      <div class="absolute bottom-0 left-0 right-0 translate-y-1/2 z-20 px-4">
        <div class="max-w-5xl mx-auto bg-white rounded-xl shadow-2xl px-6 py-5">
          <div
            class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-5 gap-4 items-end"
          >
            <!-- Check-in -->
            <div class="lg:col-span-1">
              <label
                class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5 flex items-center gap-1.5"
              >
                <Icon name="lucide:calendar" class="w-3.5 h-3.5" /> Check In
              </label>
              <input v-model="checkIn" type="date" class="input" />
            </div>
            <!-- Check-out -->
            <div class="lg:col-span-1">
              <label
                class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5 flex items-center gap-1.5"
              >
                <Icon name="lucide:calendar" class="w-3.5 h-3.5" /> Check Out
              </label>
              <input v-model="checkOut" type="date" class="input" />
            </div>
            <!-- Rooms -->
            <div>
              <label
                class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5 flex items-center gap-1.5"
              >
                <Icon name="lucide:bed-double" class="w-3.5 h-3.5" /> Rooms
              </label>
              <select v-model="rooms" class="input appearance-none">
                <option v-for="n in 10" :key="n" :value="n">
                  {{ n < 2 ? `0${n} Room` : `0${n} Rooms` }}
                </option>
              </select>
            </div>
            <!-- Guests -->
            <div>
              <label
                class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5 flex items-center gap-1.5"
              >
                <Icon name="lucide:users" class="w-3.5 h-3.5" /> Guests
              </label>
              <select v-model="guests" class="input appearance-none">
                <option value="1">1 Adult, 0 Child</option>
                <option value="2">2 Adults, 0 Child</option>
                <option value="3">2 Adults, 1 Child</option>
                <option value="4">2 Adults, 2 Children</option>
              </select>
            </div>
            <!-- CTA -->
            <NuxtLink to="/items" class="lg:col-span-1">
              <UiButton size="sm" class="w-full justify-center"
                >Check Availability</UiButton
              >
            </NuxtLink>
          </div>
        </div>
      </div>
    </section>

    <!-- Spacer for booking widget overflow -->
    <div class="h-24" />

    <!-- ══ TICKER STRIP ════════════════════════════════════════════════ -->
    <div class="ticker-strip py-3.5">
      <div class="ticker-track">
        <template
          v-for="(item, i) in [...tickerItems, ...tickerItems]"
          :key="`${item}-${i}`"
        >
          <span
            class="flex items-center gap-2 text-xs font-sans uppercase tracking-widest text-white/80"
          >
            <Icon name="lucide:star" class="w-3 h-3 text-accent fill-accent" />
            {{ item }}
          </span>
        </template>
      </div>
    </div>

    <!-- ══ ABOUT ══════════════════════════════════════════════════════ -->
    <section class="section-py section-px max-w-7xl mx-auto">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-14 items-center">
        <!-- Image block -->
        <div class="relative">
          <div
            class="rounded-2xl overflow-hidden aspect-[4/3] bg-surface-200 relative"
          >
            <div
              class="absolute inset-0 bg-gradient-to-br from-forest/20 to-brand-900/10"
            />
            <div class="absolute inset-0 flex items-center justify-center">
              <Icon name="lucide:image" class="w-16 h-16 text-surface-300" />
            </div>
          </div>
          <!-- Years badge -->
          <div
            class="absolute bottom-6 left-6 bg-accent text-white rounded-xl px-5 py-4 shadow-lg"
          >
            <div class="font-display text-4xl font-bold leading-none">16</div>
            <div class="text-xs font-sans uppercase tracking-wider mt-0.5">
              Years<br />of Services
            </div>
          </div>
          <!-- Floating stars -->
          <div class="absolute -top-4 -right-4 text-accent opacity-40">
            <Icon name="lucide:star" class="w-12 h-12" />
          </div>
        </div>

        <!-- Text -->
        <div>
          <p class="eyebrow mb-3">About Us</p>
          <h2
            class="font-display text-display-lg text-brand-900 font-light mb-5"
          >
            Luxurious Comfort,<br />Timeless Elegance Awaits
          </h2>
          <p class="text-muted font-sans text-sm leading-relaxed mb-8">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore. Our commitment to
            excellence spans over a decade and a half of curating unforgettable
            experiences.
          </p>

          <!-- Stats -->
          <div class="flex flex-wrap gap-8 mb-8">
            <div v-for="s in stats" :key="s.label">
              <div class="stat-number">{{ s.value }}</div>
              <div
                class="text-xs font-sans uppercase tracking-widest text-muted mt-0.5"
              >
                {{ s.label }}
              </div>
            </div>
          </div>

          <div class="mb-6">
            <div class="font-display text-2xl text-brand-800 italic">
              Jenny Alexander
            </div>
            <div class="text-xs font-sans text-muted mt-1">
              Jenny Alexander • Owner
            </div>
          </div>

          <NuxtLink to="/about">
            <UiButton size="sm">Learn More</UiButton>
          </NuxtLink>
        </div>
      </div>
    </section>

    <!-- ══ TICKER STRIP (bottom) ══════════════════════════════════════ -->
    <div class="ticker-strip py-3.5">
      <div class="ticker-track">
        <template
          v-for="(item, i) in [...tickerItems, ...tickerItems]"
          :key="`b-${item}-${i}`"
        >
          <span
            class="flex items-center gap-2 text-xs font-sans uppercase tracking-widest text-white/80"
          >
            <Icon name="lucide:star" class="w-3 h-3 text-accent fill-accent" />
            {{ item }}
          </span>
        </template>
      </div>
    </div>

    <!-- ══ ROOMS ══════════════════════════════════════════════════════ -->
    <section class="section-py section-px max-w-7xl mx-auto">
      <div class="text-center mb-12">
        <div class="ornament mb-3">✦</div>
        <p class="eyebrow mb-3">Rooms & Suites</p>
        <h2 class="font-display text-display-lg text-brand-900 font-light">
          Luxury Rooms &amp; Suites
        </h2>
      </div>

      <ItemGrid :items="featuredRooms" :cols="3" />

      <div class="text-center mt-10">
        <NuxtLink to="/items">
          <UiButton variant="outline" size="md">View All Rooms</UiButton>
        </NuxtLink>
      </div>
    </section>

    <!-- ══ FACILITIES ═════════════════════════════════════════════════ -->
    <section class="section-py section-px bg-surface-100">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-12">
          <div class="ornament mb-3">✦</div>
          <p class="eyebrow mb-3">Facilities</p>
          <h2 class="font-display text-display-lg text-brand-900 font-light">
            Experience with<br />Premium Facilities
          </h2>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <!-- Featured facility (large) -->
          <div
            class="lg:row-span-2 bg-forest rounded-2xl p-8 text-white flex flex-col justify-between"
          >
            <div>
              <div
                class="w-12 h-12 rounded-xl bg-white/10 flex items-center justify-center mb-6"
              >
                <Icon
                  :name="featuredFacility.icon"
                  class="w-6 h-6 text-accent"
                />
              </div>
              <h3 class="font-display text-2xl font-light mb-3">
                {{ featuredFacility.label }}
              </h3>
              <p class="text-white/60 text-sm font-sans leading-relaxed">
                {{ featuredFacility.desc }}
              </p>
            </div>
            <NuxtLink to="/items" class="mt-8">
              <UiButton variant="ghost">Learn More</UiButton>
            </NuxtLink>
          </div>

          <!-- Other facilities -->
          <div
            v-for="facility in facilities.slice(1)"
            :key="facility.label"
            class="card p-6 flex items-start gap-4"
          >
            <div
              class="w-11 h-11 rounded-xl bg-forest/8 shrink-0 flex items-center justify-center"
            >
              <Icon :name="facility.icon" class="w-5 h-5 text-forest" />
            </div>
            <div>
              <h4 class="font-display text-lg text-brand-900 mb-1">
                {{ facility.label }}
              </h4>
              <p class="text-xs text-muted font-sans leading-relaxed">
                {{ facility.desc }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ══ VIDEO BANNER ═══════════════════════════════════════════════ -->
    <section class="relative py-24 bg-brand-900 overflow-hidden">
      <div
        class="absolute inset-0 opacity-10"
        style="
          background:
            repeating-linear-gradient(
              0deg,
              white 0,
              white 1px,
              transparent 1px,
              transparent 60px
            ),
            repeating-linear-gradient(
              90deg,
              white 0,
              white 1px,
              transparent 1px,
              transparent 60px
            );
        "
      />

      <div class="relative z-10 text-center px-6">
        <p class="eyebrow text-accent mb-4">Watch the Video</p>
        <h2 class="font-display text-display-lg text-white font-light mb-8">
          Take a Virtual Tour of Our Hotel
        </h2>
        <button
          class="w-16 h-16 rounded-full border-2 border-white/40 flex items-center justify-center mx-auto hover:border-accent hover:bg-accent/10 transition-all"
        >
          <Icon name="lucide:play" class="w-6 h-6 text-white ml-1" />
        </button>
      </div>
    </section>
  </div>
</template>
