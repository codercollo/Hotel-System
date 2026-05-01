<script setup lang="ts">
defineProps<{
  type?: "success" | "error" | "warning" | "info";
  message?: string;
}>();
</script>

<template>
  <div
    :class="[
      'flex items-start gap-3 px-4 py-3 rounded-lg border text-sm font-sans',
      {
        'bg-green-50 border-green-200 text-green-800':
          (type ?? 'info') === 'success',
        'bg-red-50 border-red-200 text-red-800': type === 'error',
        'bg-amber-50 border-amber-200 text-amber-800': type === 'warning',
        'bg-blue-50 border-blue-200 text-blue-800': type === 'info',
      },
    ]"
  >
    <Icon
      :name="
        {
          success: 'lucide:check-circle',
          error: 'lucide:alert-circle',
          warning: 'lucide:alert-triangle',
          info: 'lucide:info',
        }[type ?? 'info']
      "
      class="w-4 h-4 mt-0.5 shrink-0"
    />
    <span
      ><slot>{{ message }}</slot></span
    >
  </div>
</template>
