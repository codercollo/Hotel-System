<script setup lang="ts">
const { toasts, remove } = useToast();

const iconMap: Record<string, string> = {
  success: "lucide:check-circle-2",
  error: "lucide:alert-circle",
  warning: "lucide:alert-triangle",
  info: "lucide:info",
};

const colorMap: Record<string, string> = {
  success: "border-green-200 bg-green-50 text-green-800",
  error: "border-red-200 bg-red-50 text-red-800",
  warning: "border-amber-200 bg-amber-50 text-amber-800",
  info: "border-blue-200 bg-blue-50 text-blue-800",
};
</script>

<template>
  <Teleport to="body">
    <div
      class="fixed bottom-6 right-6 z-[100] flex flex-col gap-2.5 max-w-sm w-full pointer-events-none"
    >
      <TransitionGroup name="toast">
        <div
          v-for="t in toasts"
          :key="t.id"
          :class="[
            'pointer-events-auto flex items-start gap-3 px-4 py-3 rounded-xl border shadow-lg text-sm font-sans',
            colorMap[t.type] ?? colorMap.info,
          ]"
        >
          <Icon
            :name="iconMap[t.type] ?? 'lucide:info'"
            class="w-4 h-4 mt-0.5 shrink-0"
          />
          <span class="flex-1">{{ t.message }}</span>
          <button
            @click="remove(t.id)"
            class="ml-2 opacity-60 hover:opacity-100 transition-opacity"
          >
            <Icon name="lucide:x" class="w-3.5 h-3.5" />
          </button>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<style scoped>
.toast-enter-active {
  transition: all 0.3s cubic-bezier(0.22, 1, 0.36, 1);
}
.toast-leave-active {
  transition: all 0.2s ease;
}
.toast-enter-from {
  opacity: 0;
  transform: translateX(40px);
}
.toast-leave-to {
  opacity: 0;
  transform: translateX(40px);
}
</style>
