<script setup lang="ts">
defineProps<{
  items: {
    label: string;
    icon?: string;
    href?: string;
    danger?: boolean;
    divider?: boolean;
  }[];
  align?: "left" | "right";
}>();

const open = ref(false);
</script>

<template>
  <div class="relative inline-block">
    <div @click="open = !open">
      <slot name="trigger">
        <button class="p-1.5 text-muted hover:text-brand-900 transition-colors">
          <Icon name="lucide:more-vertical" class="w-4 h-4" />
        </button>
      </slot>
    </div>

    <Transition name="dropdown">
      <div
        v-if="open"
        :class="[
          'absolute z-50 mt-1 w-44 bg-white rounded-xl border border-surface-200 shadow-card-hover py-1 overflow-hidden',
          (align ?? 'right') === 'right' ? 'right-0' : 'left-0',
        ]"
      >
        <template v-for="(item, i) in items" :key="i">
          <div v-if="item.divider" class="my-1 border-t border-surface-200" />
          <component
            v-else
            :is="item.href ? 'a' : 'button'"
            :href="item.href"
            :class="[
              'flex items-center gap-2.5 w-full px-3.5 py-2 text-sm font-sans text-left transition-colors',
              item.danger
                ? 'text-red-600 hover:bg-red-50'
                : 'text-brand-800 hover:bg-surface-100',
            ]"
            @click="open = false"
          >
            <Icon
              v-if="item.icon"
              :name="item.icon"
              class="w-3.5 h-3.5 shrink-0"
            />
            {{ item.label }}
          </component>
        </template>
      </div>
    </Transition>

    <div v-if="open" class="fixed inset-0 z-40" @click="open = false" />
  </div>
</template>

<style scoped>
.dropdown-enter-active,
.dropdown-leave-active {
  transition:
    opacity 0.15s ease,
    transform 0.15s ease;
}
.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>
