<template>
  <div class="min-h-screen bg-surface-100 flex">
    <!-- Decorative left panel -->
    <div
      class="hidden lg:flex flex-col w-[45%] bg-forest relative overflow-hidden"
    >
      <!-- Background texture -->
      <div
        class="absolute inset-0 opacity-10"
        style="
          background-image: url(&quot;/images/room-placeholder.jpg&quot;);
          background-size: cover;
          background-position: center;
        "
      />
      <div
        class="absolute inset-0 bg-gradient-to-br from-forest-dark/80 to-forest/60"
      />

      <div class="relative z-10 flex flex-col h-full p-14">
        <NuxtLink to="/" class="flex items-center gap-3 mb-auto">
          <div
            class="w-10 h-10 rounded-full bg-white/10 border border-white/20 flex items-center justify-center font-display text-xl font-semibold text-white"
          >
            B
          </div>
          <div>
            <div
              class="font-display text-xl font-semibold text-white leading-tight"
            >
              Boyelle
            </div>
            <div class="text-[9px] uppercase tracking-[0.2em] text-white/50">
              Hotel & Resort
            </div>
          </div>
        </NuxtLink>

        <div class="mb-auto">
          <p class="eyebrow text-accent-light mb-4">Welcome Back</p>
          <h2
            class="font-display text-5xl font-light text-white leading-tight mb-6"
          >
            Where Luxury<br />Meets Excellence
          </h2>
          <p class="text-white/60 text-sm font-sans leading-relaxed max-w-xs">
            Experience unparalleled comfort and timeless elegance. Your perfect
            stay awaits.
          </p>
        </div>

        <!-- Stars -->
        <div class="flex items-center gap-1.5 mb-2">
          <Icon
            v-for="i in 5"
            :key="i"
            name="lucide:star"
            class="w-4 h-4 text-accent fill-accent"
          />
        </div>
        <p class="text-white/50 text-xs font-sans">
          Rated 4.9/5 by over 60,000 guests
        </p>
      </div>
    </div>

    <!-- Right: form -->
    <div class="flex-1 flex items-center justify-center p-8">
      <div class="w-full max-w-md">
        <!-- Mobile logo -->
        <NuxtLink to="/" class="lg:hidden flex items-center gap-2.5 mb-10">
          <div
            class="w-8 h-8 rounded-full bg-forest flex items-center justify-center font-display text-base font-semibold text-white"
          >
            B
          </div>
          <div class="font-display text-lg font-semibold text-brand-900">
            Boyelle Hotel
          </div>
        </NuxtLink>

        <slot />
      </div>
    </div>
  </div>
</template>
