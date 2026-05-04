<script setup lang="ts">
import type { AdminStats } from "~/types/admin.types";

useHead({ title: "About Us — Where Luxury Meets Excellence" });

const api = useApi();
// Allow null in the generic type
const { data: stats } = await useAsyncData<AdminStats | null>(
  "about-stats",
  () => api.get<AdminStats>("/api/v1/admin/stats").catch(() => null),
);

const displayStats = computed(() => [
  {
    value: stats.value ? `${stats.value.total_items}+` : "50+",
    label: "Luxury Rooms",
  },
  { value: "60,000+", label: "Happy Guests" },
  { value: "99%", label: "Satisfaction" },
]);

const timeline = [
  {
    year: "2008",
    title: "Our Founding",
    desc: "Jenny Alexander opened our doors with a vision: to create a sanctuary where every guest feels like royalty.",
  },
  {
    year: "2012",
    title: "Spa & Wellness Wing",
    desc: "We expanded with a world-class spa facility, earning our first regional luxury hospitality award.",
  },
  {
    year: "2016",
    title: "The Pearl Suite Launch",
    desc: "Our flagship suite debuted, setting a new standard for luxury accommodation in New Jersey.",
  },
  {
    year: "2020",
    title: "Sustainability Pledge",
    desc: "We committed to carbon-neutral operations and introduced our farm-to-table dining programme.",
  },
  {
    year: "2024",
    title: "60,000 Guests Milestone",
    desc: "We celebrated over sixty thousand happy guests — a testament to our unwavering dedication to excellence.",
  },
];

const team = [
  { name: "Jenny Alexander", role: "Founder & Owner", initials: "JA" },
  { name: "Marcus Reid", role: "Head of Hospitality", initials: "MR" },
  { name: "Sophia Okonkwo", role: "Executive Chef", initials: "SO" },
  { name: "Daniel Voss", role: "Spa Director", initials: "DV" },
];

const values = [
  {
    icon: "lucide:heart",
    label: "Genuine Care",
    desc: "Every interaction is rooted in warmth and personal attention.",
  },
  {
    icon: "lucide:gem",
    label: "Uncompromising Quality",
    desc: "We source only the finest materials, ingredients, and talent.",
  },
  {
    icon: "lucide:leaf",
    label: "Sustainability",
    desc: "Luxury that respects the world we share with future generations.",
  },
  {
    icon: "lucide:shield",
    label: "Integrity",
    desc: "Honest pricing, transparent policies, and promises we keep.",
  },
];
</script>

<template>
  <div>
    <!-- Hero -->
    <section
      class="relative min-h-[55vh] flex items-end pb-16 bg-brand-900 overflow-hidden"
    >
      <div
        class="absolute inset-0 opacity-15"
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
      <div
        class="absolute inset-0"
        style="
          background: linear-gradient(
            135deg,
            rgb(25, 55, 48) 0%,
            rgb(38, 79, 72) 50%,
            rgb(52, 58, 32) 100%
          );
        "
      />
      <div class="relative z-10 max-w-7xl mx-auto section-px w-full pt-32">
        <div class="flex items-center gap-2 mb-4">
          <Icon
            v-for="i in 5"
            :key="i"
            name="lucide:star"
            class="w-3.5 h-3.5 text-accent fill-accent"
          />
        </div>
        <p class="eyebrow text-accent mb-3">Our Story</p>
        <h1
          class="font-display text-display-xl text-white font-light leading-tight max-w-2xl"
        >
          Sixteen Years of Crafting<br />Unforgettable Moments
        </h1>
      </div>
    </section>

    <!-- Intro -->
    <section class="section-py section-px max-w-7xl mx-auto">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-14 items-center">
        <div>
          <p class="eyebrow mb-3">Who We Are</p>
          <h2
            class="font-display text-display-lg text-brand-900 font-light mb-5"
          >
            A Legacy Built on<br />Passion &amp; Precision
          </h2>
          <p class="text-muted font-sans text-sm leading-relaxed mb-4">
            Nestled in the heart of New Jersey, our hotel has stood as a beacon
            of refined hospitality since 2008. What began as a single
            proprietor's dream has grown into a celebrated destination for
            discerning travellers from around the globe.
          </p>
          <p class="text-muted font-sans text-sm leading-relaxed mb-8">
            We believe luxury is not merely about opulent furnishings — it is
            about the feeling of being truly seen, heard, and cared for. From
            the moment you arrive, our dedicated team orchestrates every detail
            so you can simply be present and savour the experience.
          </p>
          <div class="flex flex-wrap gap-8">
            <div v-for="s in displayStats" :key="s.label">
              <div class="stat-number">{{ s.value }}</div>
              <div
                class="text-xs font-sans uppercase tracking-widest text-muted mt-0.5"
              >
                {{ s.label }}
              </div>
            </div>
          </div>
        </div>

        <div class="relative">
          <div
            class="rounded-2xl overflow-hidden aspect-[4/3] bg-surface-200 relative"
          >
            <div
              class="absolute inset-0 bg-gradient-to-br from-forest/20 to-brand-900/10 z-10"
            />

            <NuxtImg
              src="/images/hotel/about-lobby.jpg"
              alt="Luxury hotel lobby with elegant lighting and marble floors"
              format="webp"
              quality="85"
              loading="lazy"
              class="w-full h-full object-cover object-center"
            />
          </div>

          <div
            class="absolute -bottom-5 -left-5 bg-accent text-white rounded-xl px-5 py-4 shadow-lg"
          >
            <div class="font-display text-4xl font-bold leading-none">16</div>
            <div class="text-xs font-sans uppercase tracking-wider mt-0.5">
              Years<br />of Excellence
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Values -->
    <section class="section-py section-px bg-surface-100">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-12">
          <div class="ornament mb-3">✦</div>
          <p class="eyebrow mb-3">What We Stand For</p>
          <h2 class="font-display text-display-lg text-brand-900 font-light">
            Our Core Values
          </h2>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="v in values" :key="v.label" class="card p-6 text-center">
            <div
              class="w-12 h-12 rounded-xl bg-forest/8 flex items-center justify-center mx-auto mb-4"
            >
              <Icon :name="v.icon" class="w-5 h-5 text-forest" />
            </div>
            <h3 class="font-display text-xl text-brand-900 mb-2">
              {{ v.label }}
            </h3>
            <p class="text-xs text-muted font-sans leading-relaxed">
              {{ v.desc }}
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Timeline -->
    <section class="section-py section-px max-w-5xl mx-auto">
      <div class="text-center mb-12">
        <div class="ornament mb-3">✦</div>
        <p class="eyebrow mb-3">Our Journey</p>
        <h2 class="font-display text-display-lg text-brand-900 font-light">
          Milestones That Shaped Us
        </h2>
      </div>

      <div class="relative">
        <div
          class="absolute left-1/2 -translate-x-px top-0 bottom-0 w-px bg-brand-200 hidden lg:block"
        />
        <div class="space-y-10">
          <div
            v-for="(item, i) in timeline"
            :key="item.year"
            class="grid grid-cols-1 lg:grid-cols-2 gap-6 lg:gap-12 items-center"
          >
            <div
              :class="
                i % 2 === 0
                  ? 'lg:text-right lg:pr-10'
                  : 'lg:col-start-2 lg:pl-10'
              "
            >
              <div class="card p-6 inline-block w-full">
                <span class="font-display text-3xl text-accent font-bold">{{
                  item.year
                }}</span>
                <h3 class="font-display text-xl text-brand-900 mt-1 mb-2">
                  {{ item.title }}
                </h3>
                <p class="text-xs text-muted font-sans leading-relaxed">
                  {{ item.desc }}
                </p>
              </div>
            </div>
            <div
              v-if="i % 2 !== 0"
              class="hidden lg:flex lg:col-start-1 lg:row-start-1 items-center justify-end pr-10"
            >
              <div
                class="w-4 h-4 rounded-full bg-accent border-4 border-white shadow"
              />
            </div>
            <div v-else class="hidden lg:flex items-center pl-10">
              <div
                class="w-4 h-4 rounded-full bg-accent border-4 border-white shadow"
              />
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Team -->
    <section class="section-py section-px bg-surface-100">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-12">
          <div class="ornament mb-3">✦</div>
          <p class="eyebrow mb-3">The People Behind the Magic</p>
          <h2 class="font-display text-display-lg text-brand-900 font-light">
            Meet Our Team
          </h2>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div
            v-for="member in team"
            :key="member.name"
            class="card p-6 text-center group"
          >
            <div
              class="w-20 h-20 rounded-full bg-forest flex items-center justify-center mx-auto mb-4 group-hover:bg-accent transition-colors duration-300"
            >
              <span class="font-display text-2xl text-white font-bold">{{
                member.initials
              }}</span>
            </div>
            <h3 class="font-display text-xl text-brand-900 mb-1">
              {{ member.name }}
            </h3>
            <p class="text-xs font-sans uppercase tracking-widest text-muted">
              {{ member.role }}
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="relative py-24 bg-brand-900 overflow-hidden text-center">
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
      <div class="relative z-10 px-6">
        <p class="eyebrow text-accent mb-4">Ready to Experience It?</p>
        <h2
          class="font-display text-display-lg text-white font-light mb-8 max-w-xl mx-auto"
        >
          Your Perfect Stay Awaits
        </h2>
        <NuxtLink to="/items">
          <UiButton size="md" variant="gold">Book Your Room</UiButton>
        </NuxtLink>
      </div>
    </section>
  </div>
</template>
