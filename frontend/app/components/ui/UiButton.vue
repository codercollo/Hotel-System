<script setup lang="ts">
defineProps<{
  variant?: "primary" | "outline" | "ghost" | "gold";
  size?: "sm" | "md" | "lg";
  loading?: boolean;
  disabled?: boolean;
  type?: "button" | "submit" | "reset";
}>();
</script>

<template>
  <button
    :type="type ?? 'button'"
    :disabled="disabled || loading"
    :class="[
      'inline-flex items-center justify-center gap-2 font-sans font-medium uppercase tracking-widest transition-all duration-200 rounded-md border disabled:opacity-50 disabled:cursor-not-allowed',
      {
        'text-xs px-6 py-2.5': size === 'sm' || !size,
        'text-sm px-8 py-3.5': size === 'md',
        'text-sm px-10 py-4': size === 'lg',
      },
      {
        'bg-forest text-white border-transparent hover:bg-forest-dark hover:-translate-y-px':
          (variant ?? 'primary') === 'primary',
        'bg-transparent text-brand-900 border-brand-900 hover:bg-brand-900 hover:text-white':
          variant === 'outline',
        'bg-transparent text-white border-transparent underline underline-offset-4 decoration-white/40 hover:decoration-white px-0':
          variant === 'ghost',
        'bg-gradient-to-br from-accent to-accent-light text-white border-transparent hover:opacity-90 hover:-translate-y-px shadow-sm':
          variant === 'gold',
      },
    ]"
  >
    <svg
      v-if="loading"
      class="w-4 h-4 animate-spin"
      viewBox="0 0 24 24"
      fill="none"
    >
      <circle
        class="opacity-25"
        cx="12"
        cy="12"
        r="10"
        stroke="currentColor"
        stroke-width="4"
      />
      <path
        class="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
      />
    </svg>
    <slot />
  </button>
</template>
