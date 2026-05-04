<script setup lang="ts">
definePageMeta({ layout: "auth" });
useHead({ title: "Sign In" });

const { login } = useAuth();
const route = useRoute();

const email = ref("");
const password = ref("");
const loading = ref(false);
const error = ref("");

const onSubmit = async () => {
  error.value = "";

  if (!email.value || !password.value) {
    error.value = "Please enter your email and password.";
    return;
  }

  loading.value = true;
  try {
    await login(email.value, password.value);

    // Redirect to the originally requested page, or dashboard
    const redirect = route.query.redirect as string | undefined;
    await navigateTo(redirect ?? "/");
  } catch (e: unknown) {
    error.value =
      e instanceof Error
        ? e.message
        : "Invalid email or password. Please try again.";
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div>
    <p class="eyebrow mb-3">Welcome Back</p>
    <h2 class="font-display text-display-md text-brand-900 font-light mb-2">
      Sign in to your account
    </h2>
    <p class="text-sm text-muted font-sans mb-8">
      Don't have an account?
      <NuxtLink to="/auth/register" class="text-forest hover:underline"
        >Create one</NuxtLink
      >
    </p>

    <UiAlert v-if="error" type="error" :message="error" class="mb-5" />

    <div class="space-y-4">
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
        placeholder="••••••••"
        icon="lucide:lock"
        :disabled="loading"
        @keyup.enter="onSubmit"
      />

      <div class="flex items-center justify-between">
        <label
          class="flex items-center gap-2 text-sm font-sans text-brand-800 cursor-pointer"
        >
          <input type="checkbox" class="accent-forest" /> Remember me
        </label>
        <NuxtLink
          to="/auth/forgot-password"
          class="text-sm text-forest hover:underline font-sans"
          >Forgot password?</NuxtLink
        >
      </div>

      <UiButton
        @click="onSubmit"
        :loading="loading"
        size="md"
        class="w-full justify-center mt-2"
      >
        Sign In
      </UiButton>
    </div>

    <div class="mt-6 pt-6 border-t border-surface-200 text-center">
      <p class="text-xs text-muted font-sans">
        By signing in you agree to our
        <a href="#" class="text-forest hover:underline">Terms of Service</a> and
        <a href="#" class="text-forest hover:underline">Privacy Policy</a>
      </p>
    </div>
  </div>
</template>
