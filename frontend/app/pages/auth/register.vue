<script setup lang="ts">
definePageMeta({ layout: "auth" });
useHead({ title: "Create Account" });

const { register } = useAuth();

const name = ref("");
const email = ref("");
const password = ref("");
const confirm = ref("");
const loading = ref(false);
const error = ref("");

const onSubmit = async () => {
  error.value = "";

  if (!name.value || !email.value || !password.value || !confirm.value) {
    error.value = "Please fill in all fields.";
    return;
  }

  if (password.value !== confirm.value) {
    error.value = "Passwords do not match.";
    return;
  }

  if (password.value.length < 8) {
    error.value = "Password must be at least 8 characters.";
    return;
  }

  loading.value = true;
  try {
    await register(name.value, email.value, password.value);
    await navigateTo("/");
  } catch (e: unknown) {
    error.value =
      e instanceof Error ? e.message : "Registration failed. Please try again.";
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div>
    <p class="eyebrow mb-3">Get Started</p>
    <h2 class="font-display text-display-md text-brand-900 font-light mb-2">
      Create your account
    </h2>
    <p class="text-sm text-muted font-sans mb-8">
      Already have an account?
      <NuxtLink to="/auth/login" class="text-forest hover:underline"
        >Sign in</NuxtLink
      >
    </p>

    <UiAlert v-if="error" type="error" :message="error" class="mb-5" />

    <div class="space-y-4">
      <UiInput
        v-model="name"
        label="Full Name"
        placeholder="Jane Smith"
        icon="lucide:user"
        :disabled="loading"
      />
      <UiInput
        v-model="email"
        label="Email Address"
        type="email"
        placeholder="you@example.com"
        icon="lucide:mail"
        :disabled="loading"
      />
      <UiInput
        v-model="password"
        label="Password"
        type="password"
        placeholder="At least 8 characters"
        icon="lucide:lock"
        :disabled="loading"
      />
      <UiInput
        v-model="confirm"
        label="Confirm Password"
        type="password"
        placeholder="Repeat password"
        icon="lucide:lock"
        :disabled="loading"
        @keyup.enter="onSubmit"
      />

      <UiButton
        @click="onSubmit"
        :loading="loading"
        size="md"
        class="w-full justify-center mt-2"
      >
        Create Account
      </UiButton>
    </div>

    <p class="text-xs text-muted font-sans mt-6 text-center">
      By creating an account you agree to our
      <a href="#" class="text-forest hover:underline">Terms</a> &amp;
      <a href="#" class="text-forest hover:underline">Privacy Policy</a>
    </p>
  </div>
</template>
