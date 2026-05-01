<script setup lang="ts">
defineProps<{
  icon?: string;
  title?: string;
  description?: string;
  actionLabel?: string;
  actionTo?: string;
}>();
defineEmits<{ action: [] }>();
</script>

<template>
  <div class="flex flex-col items-center justify-center py-20 text-center px-6">
    <div
      class="w-16 h-16 rounded-full bg-surface-200 flex items-center justify-center mb-5"
    >
      <Icon :name="icon ?? 'lucide:inbox'" class="w-8 h-8 text-muted" />
    </div>
    <h3 class="font-display text-xl text-brand-900 mb-2">
      {{ title ?? "Nothing here yet" }}
    </h3>
    <p class="text-sm text-muted font-sans max-w-xs leading-relaxed mb-6">
      {{ description }}
    </p>
    <NuxtLink v-if="actionTo" :to="actionTo">
      <UiButton size="sm">{{ actionLabel ?? "Get Started" }}</UiButton>
    </NuxtLink>
    <UiButton v-else-if="actionLabel" size="sm" @click="$emit('action')">{{
      actionLabel
    }}</UiButton>
  </div>
</template>
