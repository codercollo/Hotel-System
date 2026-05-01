<script setup lang="ts">
definePageMeta({ layout: "admin" });
useHead({ title: "Platform Settings" });

const settings = reactive({
  hotelName: "Boyelle Hotel & Resort",
  hotelEmail: "info@boyelle.com",
  hotelPhone: "(000) 000-0000",
  checkInTime: "14:00",
  checkOutTime: "11:00",
  currency: "USD",
  timezone: "America/New_York",
  maintenanceMode: false,
});

const saved = ref(false);
const loading = ref(false);

const save = async () => {
  loading.value = true;
  await new Promise((r) => setTimeout(r, 700));
  loading.value = false;
  saved.value = true;
  setTimeout(() => (saved.value = false), 3000);
};
</script>

<template>
  <div class="max-w-2xl">
    <div class="mb-6">
      <h2 class="font-display text-2xl text-brand-900">Platform Settings</h2>
      <p class="text-xs text-muted font-sans mt-0.5">
        Configure global hotel settings.
      </p>
    </div>

    <UiAlert
      v-if="saved"
      type="success"
      message="Settings saved successfully."
      class="mb-5"
    />

    <!-- General -->
    <div class="card p-6 mb-5 space-y-4">
      <h3 class="font-display text-lg text-brand-900">General</h3>
      <UiInput
        v-model="settings.hotelName"
        label="Hotel Name"
        icon="lucide:building"
      />
      <UiInput
        v-model="settings.hotelEmail"
        label="Contact Email"
        type="email"
        icon="lucide:mail"
      />
      <UiInput
        v-model="settings.hotelPhone"
        label="Phone Number"
        icon="lucide:phone"
      />
    </div>

    <!-- Booking rules -->
    <div class="card p-6 mb-5 space-y-4">
      <h3 class="font-display text-lg text-brand-900">Booking Rules</h3>
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >Check-In Time</label
          >
          <input v-model="settings.checkInTime" type="time" class="input" />
        </div>
        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >Check-Out Time</label
          >
          <input v-model="settings.checkOutTime" type="time" class="input" />
        </div>
      </div>
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >Currency</label
          >
          <select v-model="settings.currency" class="input">
            <option>USD</option>
            <option>EUR</option>
            <option>GBP</option>
            <option>KES</option>
          </select>
        </div>
        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >Timezone</label
          >
          <select v-model="settings.timezone" class="input">
            <option>America/New_York</option>
            <option>America/Los_Angeles</option>
            <option>Europe/London</option>
            <option>Africa/Nairobi</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Maintenance -->
    <div class="card p-6 mb-6">
      <div class="flex items-center justify-between">
        <div>
          <h3 class="font-display text-lg text-brand-900">Maintenance Mode</h3>
          <p class="text-xs text-muted font-sans mt-0.5">
            When enabled, the public site shows a maintenance page.
          </p>
        </div>
        <label class="relative inline-flex items-center cursor-pointer">
          <input
            v-model="settings.maintenanceMode"
            type="checkbox"
            class="sr-only peer"
          />
          <div
            class="w-10 h-5 bg-surface-200 rounded-full peer peer-checked:bg-forest transition-colors"
          />
          <div
            class="absolute left-0.5 top-0.5 w-4 h-4 bg-white rounded-full shadow transition-transform peer-checked:translate-x-5"
          />
        </label>
      </div>
    </div>

    <div class="flex justify-end">
      <UiButton @click="save" :loading="loading" size="sm"
        >Save Settings</UiButton
      >
    </div>
  </div>
</template>
