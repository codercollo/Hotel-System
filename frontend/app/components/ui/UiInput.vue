<script setup lang="ts">
defineProps<{
  modelValue?: string | number;
  placeholder?: string;
  type?: string;
  label?: string;
  error?: string;
  disabled?: boolean;
  icon?: string;
}>();

defineEmits<{ "update:modelValue": [value: string] }>();
</script>

<template>
  <div class="flex flex-col gap-1.5">
    <label
      v-if="label"
      class="font-sans text-xs font-500 uppercase tracking-widest text-brand-700"
    >
      {{ label }}
    </label>
    <div class="relative">
      <Icon
        v-if="icon"
        :name="icon"
        class="absolute left-3 top-1/2 -translate-y-1/2 text-muted w-4 h-4 pointer-events-none"
      />
      <input
        :type="type ?? 'text'"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :class="[
          'input',
          { 'pl-9': icon },
          { 'border-red-400 focus:border-red-400': error },
        ]"
        @input="
          $emit('update:modelValue', ($event.target as HTMLInputElement).value)
        "
      />
    </div>
    <p v-if="error" class="text-xs text-red-500 font-sans">{{ error }}</p>
  </div>
</template>
