<script setup lang="ts">
definePageMeta({ requiresAuth: true });
useHead({ title: "Account Settings" });

const notifications = reactive({
  bookingConfirmations: true,
  paymentReceipts: true,
  promotions: false,
  newsletter: false,
});

const saved = ref(false);
const loading = ref(false);

const save = async () => {
  loading.value = true;
  await new Promise((r) => setTimeout(r, 600));
  loading.value = false;
  saved.value = true;
  setTimeout(() => (saved.value = false), 3000);
};
</script>

<template>
  <div class="max-w-2xl mx-auto section-px py-12">
    <div class="mb-8">
      <p class="eyebrow mb-2">Account</p>
      <h1 class="font-display text-display-md text-brand-900 font-light">
        Settings
      </h1>
    </div>

    <UiAlert
      v-if="saved"
      type="success"
      message="Settings saved."
      class="mb-6"
    />

    <!-- Notification preferences -->
    <div class="card p-6 mb-6">
      <h3 class="font-display text-xl text-brand-900 mb-5">
        Notification Preferences
      </h3>
      <div class="space-y-4">
        <label
          v-for="(val, key) in notifications"
          :key="key"
          class="flex items-center justify-between py-3 border-b border-surface-200 last:border-0 cursor-pointer"
        >
          <div>
            <p class="text-sm font-sans font-medium text-brand-900 capitalize">
              {{ key.replace(/([A-Z])/g, " $1") }}
            </p>
            <p class="text-xs text-muted font-sans mt-0.5">
              {{
                {
                  bookingConfirmations:
                    "Receive confirmation emails for every booking.",
                  paymentReceipts: "Get receipts when a payment is processed.",
                  promotions: "Hear about exclusive deals and seasonal offers.",
                  newsletter: "Monthly newsletter with hotel updates.",
                }[key]
              }}
            </p>
          </div>
          <input
            type="checkbox"
            :checked="(notifications as any)[key]"
            @change="
              (notifications as any)[key] = (
                $event.target as HTMLInputElement
              ).checked
            "
            class="w-4 h-4 accent-forest cursor-pointer"
          />
        </label>
      </div>
    </div>

    <!-- Danger zone -->
    <div class="card p-6 border-red-100">
      <h3 class="font-display text-xl text-brand-900 mb-3">Danger Zone</h3>
      <p class="text-sm text-muted font-sans mb-4">
        Deleting your account is permanent and cannot be undone.
      </p>
      <UiButton
        class="border-red-200 text-red-600 hover:bg-red-50 bg-transparent"
        variant="outline"
        size="sm"
      >
        Delete Account
      </UiButton>
    </div>

    <div class="flex justify-end mt-6">
      <UiButton @click="save" :loading="loading" size="sm"
        >Save Settings</UiButton
      >
    </div>
  </div>
</template>
