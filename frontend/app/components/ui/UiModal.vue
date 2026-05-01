<script setup lang="ts">
defineProps<{ open: boolean; title?: string }>();
defineEmits<{ close: [] }>();
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="open"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
      >
        <div
          class="absolute inset-0 bg-brand-900/50 backdrop-blur-sm"
          @click="$emit('close')"
        />
        <div
          class="relative bg-white rounded-xl shadow-2xl w-full max-w-lg animate-fade-up"
        >
          <div
            class="flex items-center justify-between px-6 py-4 border-b border-surface-200"
          >
            <h3 v-if="title" class="font-display text-xl text-brand-900">
              {{ title }}
            </h3>
            <button
              @click="$emit('close')"
              class="text-muted hover:text-brand-900 transition-colors"
            >
              <Icon name="lucide:x" class="w-5 h-5" />
            </button>
          </div>
          <div class="px-6 py-5">
            <slot />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.2s ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
