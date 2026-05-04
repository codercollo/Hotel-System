<script setup lang="ts">
// v-model maps to the provider string sent to the backend.
// IDs MUST match the backend validator: oneof=mpesa stripe flutterwave paystack
const selected = defineModel<string>({ default: "stripe" });

const methods = [
  { id: "stripe", icon: "lucide:credit-card", label: "Credit / Debit Card" },
  { id: "mpesa", icon: "lucide:smartphone", label: "M-Pesa" },
  { id: "flutterwave", icon: "lucide:wallet", label: "Flutterwave" },
  { id: "paystack", icon: "lucide:banknote", label: "Paystack" },
] as const;

export type PaymentProvider = (typeof methods)[number]["id"];
</script>

<template>
  <div class="space-y-2.5">
    <p class="text-xs uppercase tracking-widest font-sans text-muted mb-3">
      Payment Method
    </p>
    <label
      v-for="m in methods"
      :key="m.id"
      :class="[
        'flex items-center gap-3 p-3.5 rounded-xl border cursor-pointer transition-all',
        selected === m.id
          ? 'border-forest bg-forest/5 shadow-sm'
          : 'border-surface-200 hover:border-forest/40',
      ]"
    >
      <input
        v-model="selected"
        type="radio"
        :value="m.id"
        class="accent-forest"
      />
      <Icon
        :name="m.icon"
        class="w-4 h-4"
        :class="selected === m.id ? 'text-forest' : 'text-muted'"
      />
      <span
        class="text-sm font-sans"
        :class="selected === m.id ? 'text-brand-900 font-medium' : 'text-muted'"
      >
        {{ m.label }}
      </span>
    </label>
  </div>
</template>
