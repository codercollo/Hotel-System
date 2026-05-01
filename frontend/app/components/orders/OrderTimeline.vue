<script setup lang="ts">
defineProps<{
  steps: { label: string; date: string; done: boolean }[];
}>();
</script>

<template>
  <ol class="relative border-l border-surface-200 space-y-6 ml-3">
    <li v-for="(step, i) in steps" :key="i" class="pl-6 relative">
      <div
        :class="[
          'absolute -left-[9px] top-0.5 w-[18px] h-[18px] rounded-full border-2 border-white flex items-center justify-center',
          step.done ? 'bg-forest' : 'bg-surface-200',
        ]"
      >
        <Icon
          v-if="step.done"
          name="lucide:check"
          class="w-2.5 h-2.5 text-white"
        />
      </div>
      <p
        :class="[
          'text-sm font-sans font-medium',
          step.done ? 'text-brand-900' : 'text-muted',
        ]"
      >
        {{ step.label }}
      </p>
      <p class="text-xs text-muted font-sans mt-0.5">{{ step.date }}</p>
    </li>
  </ol>
</template>
