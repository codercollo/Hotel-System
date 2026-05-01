<script setup lang="ts">
defineProps<{
  status: "pending" | "processing" | "completed" | "failed" | "refunded";
  amount: number;
  provider: string;
  ref?: string;
}>();

const fmt = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });

const iconMap = {
  pending: { icon: "lucide:clock", color: "text-amber-500 bg-amber-50" },
  processing: { icon: "lucide:loader-2", color: "text-blue-500 bg-blue-50" },
  completed: {
    icon: "lucide:check-circle",
    color: "text-green-600 bg-green-50",
  },
  failed: { icon: "lucide:x-circle", color: "text-red-500 bg-red-50" },
  refunded: {
    icon: "lucide:refresh-ccw",
    color: "text-purple-500 bg-purple-50",
  },
};
</script>

<template>
  <div class="card p-5">
    <div class="flex items-center gap-3 mb-4">
      <div
        :class="[
          'w-10 h-10 rounded-xl flex items-center justify-center',
          iconMap[status].color,
        ]"
      >
        <Icon
          :name="iconMap[status].icon"
          class="w-5 h-5"
          :class="status === 'processing' ? 'animate-spin' : ''"
        />
      </div>
      <div>
        <p class="text-xs font-sans uppercase tracking-widest text-muted">
          Payment Status
        </p>
        <p class="font-display text-lg text-brand-900 capitalize">
          {{ status }}
        </p>
      </div>
    </div>

    <div class="space-y-2 text-sm font-sans">
      <div class="flex justify-between">
        <span class="text-muted">Amount</span>
        <span class="font-medium">{{ fmt(amount) }}</span>
      </div>
      <div class="flex justify-between">
        <span class="text-muted">Provider</span>
        <span class="capitalize">{{ provider }}</span>
      </div>
      <div v-if="ref" class="flex justify-between">
        <span class="text-muted">Reference</span>
        <span class="font-mono text-xs">{{ ref }}</span>
      </div>
    </div>
  </div>
</template>
