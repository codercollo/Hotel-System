<script setup lang="ts">
useHead({ title: "My Profile" });

const name = ref("Jane Doe");
const email = ref("jane@example.com");
const phone = ref("+1 555 000 0000");
const loading = ref(false);
const saved = ref(false);

const save = async () => {
  loading.value = true;
  await new Promise((r) => setTimeout(r, 800));
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
        My Profile
      </h1>
    </div>

    <!-- Avatar -->
    <div class="card p-6 mb-6 flex items-center gap-5">
      <div
        class="w-16 h-16 rounded-full bg-forest flex items-center justify-center text-white font-display text-2xl font-semibold shrink-0"
      >
        {{
          name
            .split(" ")
            .map((w) => w[0])
            .join("")
        }}
      </div>
      <div>
        <p class="font-display text-xl text-brand-900">{{ name }}</p>
        <p class="text-sm text-muted font-sans">{{ email }}</p>
        <p class="text-xs text-accent font-sans mt-0.5">
          Loyal Guest · 2,400 Points
        </p>
      </div>
    </div>

    <!-- Form -->
    <div class="card p-6 space-y-5">
      <h3 class="font-display text-xl text-brand-900">Personal Information</h3>

      <UiAlert
        v-if="saved"
        type="success"
        message="Your profile has been updated."
      />

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <UiInput
          v-model="name"
          label="Full Name"
          placeholder="Jane Doe"
          icon="lucide:user"
        />
        <UiInput
          v-model="phone"
          label="Phone Number"
          placeholder="+1 555 000 0000"
          icon="lucide:phone"
        />
      </div>
      <UiInput
        v-model="email"
        label="Email Address"
        type="email"
        placeholder="jane@example.com"
        icon="lucide:mail"
      />

      <div class="pt-2 flex justify-end">
        <UiButton @click="save" :loading="loading" size="sm"
          >Save Changes</UiButton
        >
      </div>
    </div>

    <!-- Change password -->
    <div class="card p-6 mt-6 space-y-4">
      <h3 class="font-display text-xl text-brand-900">Change Password</h3>
      <UiInput
        label="Current Password"
        type="password"
        placeholder="••••••••"
        icon="lucide:lock"
      />
      <UiInput
        label="New Password"
        type="password"
        placeholder="At least 8 characters"
        icon="lucide:lock"
      />
      <UiInput
        label="Confirm Password"
        type="password"
        placeholder="Repeat new password"
        icon="lucide:lock"
      />
      <div class="flex justify-end">
        <UiButton variant="outline" size="sm">Update Password</UiButton>
      </div>
    </div>
  </div>
</template>
