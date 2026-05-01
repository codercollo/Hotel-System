<script setup lang="ts">
const route = useRoute();
const mobileOpen = ref(false);

const navLinks = [
  { label: "Home", to: "/" },
  { label: "Rooms & Suites", to: "/items" },
  { label: "Restaurant", to: "/restaurant" },
  { label: "Gallery", to: "/gallery" },
  { label: "About Us", to: "/about" },
  { label: "Contact Us", to: "/contact" },
];

const socialIcons = [
  "lucide:facebook",
  "lucide:twitter",
  "lucide:instagram",
  "lucide:youtube",
];

const isActive = (to: string) =>
  to === "/" ? route.path === "/" : route.path.startsWith(to);

watch(mobileOpen, (val) => {
  if (import.meta.client) {
    document.body.style.overflow = val ? "hidden" : "";
  }
});

onBeforeUnmount(() => {
  if (import.meta.client) {
    document.body.style.overflow = "";
  }
});
</script>

<template>
  <header class="w-full sticky top-0 z-50">
    <!-- Top bar -->
    <div class="bg-forest text-white text-xs font-sans font-light">
      <div
        class="max-w-7xl mx-auto px-4 md:px-8 py-2 flex items-center justify-between gap-4"
      >
        <div class="flex items-center gap-5">
          <a
            href="tel:+10000000000"
            class="flex items-center gap-1.5 text-white/80 hover:text-white transition-colors"
          >
            <Icon name="lucide:phone" class="w-3 h-3" />
            (000) 000-0000
          </a>
          <a
            href="mailto:example@mail.com"
            class="hidden sm:flex items-center gap-1.5 text-white/80 hover:text-white transition-colors"
          >
            <Icon name="lucide:mail" class="w-3 h-3" />
            example@mail.com
          </a>
          <span class="hidden lg:flex items-center gap-1.5 text-white/60">
            <Icon name="lucide:map-pin" class="w-3 h-3" />
            2464 Royal Ln, Mesa, New Jersey 45463
          </span>
        </div>
        <div class="flex items-center gap-3">
          <a
            v-for="icon in socialIcons"
            :key="icon"
            href="#"
            class="text-white/60 hover:text-white transition-colors"
          >
            <Icon :name="icon" class="w-3.5 h-3.5" />
          </a>
        </div>
      </div>
    </div>

    <!-- Main nav -->
    <nav class="bg-white/95 backdrop-blur-md shadow-nav">
      <div
        class="max-w-7xl mx-auto px-4 md:px-8 h-16 flex items-center justify-between"
      >
        <!-- Logo -->
        <NuxtLink to="/" class="flex items-center gap-2.5 shrink-0">
          <div
            class="w-9 h-9 rounded-full bg-forest flex items-center justify-center text-white font-display text-lg font-semibold"
          >
            B
          </div>
          <div class="hidden sm:block">
            <div
              class="font-display text-lg font-semibold text-brand-900 leading-tight"
            >
              Boyelle
            </div>
            <div
              class="text-[9px] font-sans uppercase tracking-[0.2em] text-muted -mt-0.5"
            >
              Hotel & Resort
            </div>
          </div>
        </NuxtLink>

        <!-- Desktop links -->
        <ul class="hidden lg:flex items-center gap-7">
          <li v-for="link in navLinks" :key="link.to">
            <NuxtLink
              :to="link.to"
              :class="['nav-link', { active: isActive(link.to) }]"
            >
              {{ link.label }}
            </NuxtLink>
          </li>
        </ul>

        <!-- CTA + mobile toggle -->
        <div class="flex items-center gap-3">
          <NuxtLink to="/items" class="hidden md:block">
            <UiButton size="sm">Book Now</UiButton>
          </NuxtLink>
          <button
            class="lg:hidden p-2 text-brand-800 hover:text-forest transition-colors"
            @click="mobileOpen = !mobileOpen"
          >
            <Icon
              :name="mobileOpen ? 'lucide:x' : 'lucide:menu'"
              class="w-5 h-5"
            />
          </button>
        </div>
      </div>
    </nav>

    <!-- Mobile menu -->
    <Transition name="slide-down">
      <div
        v-if="mobileOpen"
        class="lg:hidden bg-white border-t border-surface-200 shadow-lg absolute w-full"
      >
        <ul class="px-6 py-4 flex flex-col gap-1">
          <li v-for="link in navLinks" :key="link.to">
            <NuxtLink
              :to="link.to"
              :class="[
                'block py-3 border-b border-surface-200 font-sans text-sm',
                isActive(link.to)
                  ? 'text-forest font-medium'
                  : 'text-brand-800',
              ]"
              @click="mobileOpen = false"
            >
              {{ link.label }}
            </NuxtLink>
          </li>
          <li class="pt-3">
            <NuxtLink to="/items" class="block" @click="mobileOpen = false">
              <UiButton class="w-full justify-center" size="sm"
                >Book Now</UiButton
              >
            </NuxtLink>
          </li>
        </ul>
      </div>
    </Transition>
  </header>
</template>

<style scoped>
.slide-down-enter-active,
.slide-down-leave-active {
  transition:
    transform 0.25s ease,
    opacity 0.25s ease;
}
.slide-down-enter-from,
.slide-down-leave-to {
  transform: translateY(-8px);
  opacity: 0;
}
</style>
