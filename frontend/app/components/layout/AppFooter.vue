<script setup lang="ts">
const socialIcons = [
  "lucide:facebook",
  "lucide:twitter",
  "lucide:instagram",
  "lucide:youtube",
];

const quickLinks = [
  "Home",
  "Rooms & Suites",
  "Restaurant",
  "Gallery",
  "About Us",
  "Contact Us",
];

const services = [
  "Luxury Rooms",
  "Fine Dining",
  "Spa & Wellness",
  "Swimming Pool",
  "Airport Transfer",
  "Event Hosting",
];
</script>

<template>
  <footer style="background-color: #0f1f1a" class="text-white">
    <div
      class="max-w-7xl mx-auto px-6 md:px-10 pt-16 pb-10 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-10"
    >
      <!-- Brand -->
      <div>
        <div class="flex items-center gap-2.5 mb-5">
          <div
            class="w-9 h-9 rounded-full bg-amber-500 flex items-center justify-center text-white font-semibold text-lg"
          >
            B
          </div>
          <div>
            <div class="text-lg font-semibold text-white leading-tight">
              Boyelle
            </div>
            <div
              class="text-[9px] uppercase tracking-[0.2em] text-white/50 -mt-0.5"
            >
              Hotel & Resort
            </div>
          </div>
        </div>
        <p class="text-sm leading-relaxed text-white/60 mb-5">
          Where luxury meets excellence. Timeless elegance and comfort await you
          in New Jersey.
        </p>
        <div class="flex items-center gap-3">
          <a
            v-for="icon in socialIcons"
            :key="icon"
            href="#"
            class="w-8 h-8 rounded-full border border-white/20 flex items-center justify-center hover:border-amber-400 hover:text-amber-400 transition-colors"
          >
            <Icon :name="icon" class="w-3.5 h-3.5" />
          </a>
        </div>
      </div>

      <!-- Quick links -->
      <div>
        <h5
          class="text-[10px] uppercase tracking-[0.2em] font-sans text-amber-400 mb-5"
        >
          Quick Links
        </h5>
        <ul class="space-y-2.5">
          <li v-for="link in quickLinks" :key="link">
            <a
              href="#"
              class="text-sm text-white/60 hover:text-amber-400 transition-colors flex items-center gap-2"
            >
              <span class="text-amber-400/60">›</span> {{ link }}
            </a>
          </li>
        </ul>
      </div>

      <!-- Services -->
      <div>
        <h5
          class="text-[10px] uppercase tracking-[0.2em] font-sans text-amber-400 mb-5"
        >
          Our Services
        </h5>
        <ul class="space-y-2.5">
          <li v-for="service in services" :key="service">
            <a
              href="#"
              class="text-sm text-white/60 hover:text-amber-400 transition-colors flex items-center gap-2"
            >
              <span class="text-amber-400/60">›</span> {{ service }}
            </a>
          </li>
        </ul>
      </div>

      <!-- Contact -->
      <div>
        <h5
          class="text-[10px] uppercase tracking-[0.2em] font-sans text-amber-400 mb-5"
        >
          Contact Us
        </h5>
        <ul class="space-y-3.5">
          <li class="flex items-start gap-3 text-sm text-white/60">
            <Icon
              name="lucide:map-pin"
              class="w-4 h-4 mt-0.5 text-amber-400 shrink-0"
            />
            2464 Royal Ln, Mesa, New Jersey 45463
          </li>
          <li class="flex items-center gap-3 text-sm text-white/60">
            <Icon name="lucide:phone" class="w-4 h-4 text-amber-400 shrink-0" />
            (000) 000-0000
          </li>
          <li class="flex items-center gap-3 text-sm text-white/60">
            <Icon name="lucide:mail" class="w-4 h-4 text-amber-400 shrink-0" />
            example@mail.com
          </li>
          <li class="flex items-start gap-3 text-sm text-white/60">
            <Icon
              name="lucide:clock"
              class="w-4 h-4 mt-0.5 text-amber-400 shrink-0"
            />
            <span>Check-in: 2:00 PM<br />Check-out: 11:00 AM</span>
          </li>
        </ul>
      </div>
    </div>

    <div class="border-t border-white/10">
      <div
        class="max-w-7xl mx-auto px-6 md:px-10 py-5 flex flex-col sm:flex-row items-center justify-between gap-3"
      >
        <p class="text-xs text-white/40 font-sans">
          © {{ new Date().getFullYear() }} Boyelle Hotel & Resort. All rights
          reserved.
        </p>
        <div class="flex items-center gap-5 text-xs text-white/40">
          <a href="#" class="hover:text-white/70 transition-colors"
            >Privacy Policy</a
          >
          <a href="#" class="hover:text-white/70 transition-colors"
            >Terms of Service</a
          >
        </div>
      </div>
    </div>
  </footer>
</template>
