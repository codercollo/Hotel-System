<script setup lang="ts">
definePageMeta({ layout: "auth" });
useHead({ title: "Reset Password" });

const password = ref("");
const confirm = ref("");
const loading = ref(false);
const error = ref("");
const done = ref(false);

const onSubmit = async () => {
  if (password.value !== confirm.value) {
    error.value = "Passwords do not match.";
    return;
  }
  if (password.value.length < 8) {
    error.value = "Password must be at least 8 characters.";
    return;
  }
  loading.value = true;
  error.value = "";
  await new Promise((r) => setTimeout(r, 800));
  loading.value = false;
  done.value = true;
};
</script>

<template>
  <div>
    <p class="eyebrow mb-3">New Password</p>
    <h2 class="font-display text-display-md text-brand-900 font-light mb-6">
      Reset your password
    </h2>

    <UiAlert v-if="error" type="error" :message="error" class="mb-5" />

    <template v-if="!done">
      <div class="space-y-4">
        <UiInput
          v-model="password"
          label="New Password"
          type="password"
          placeholder="At least 8 characters"
          icon="lucide:lock"
        />
        <UiInput
          v-model="confirm"
          label="Confirm Password"
          type="password"
          placeholder="Repeat new password"
          icon="lucide:lock"
        />
        <UiButton
          @click="onSubmit"
          :loading="loading"
          size="md"
          class="w-full justify-center"
        >
          Set New Password
        </UiButton>
      </div>
    </template>

    <template v-else>
      <UiAlert
        type="success"
        message="Your password has been reset successfully."
        class="mb-6"
      />
      <NuxtLink to="/auth/login">
        <UiButton size="sm" class="w-full justify-center"
          >Sign In with New Password</UiButton
        >
      </NuxtLink>
    </template>
  </div>
</template>
