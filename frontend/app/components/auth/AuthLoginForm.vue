<script setup lang="ts">
const emit = defineEmits<{ success: [] }>();

const email = ref("");
const password = ref("");
const loading = ref(false);
const error = ref("");

const submit = async () => {
  if (!email.value || !password.value) {
    error.value = "Please fill in all fields.";
    return;
  }
  loading.value = true;
  error.value = "";
  // useAuth().login() wired in Phase 2
  await new Promise((r) => setTimeout(r, 800));
  loading.value = false;
  emit("success");
};
</script>

<template>
  <div class="space-y-4">
    <UiAlert v-if="error" type="error" :message="error" />
    <UiInput
      v-model="email"
      label="Email"
      type="email"
      placeholder="you@example.com"
      icon="lucide:mail"
    />
    <UiInput
      v-model="password"
      label="Password"
      type="password"
      placeholder="••••••••"
      icon="lucide:lock"
    />
    <div class="flex justify-end">
      <NuxtLink
        to="/auth/forgot-password"
        class="text-xs text-forest font-sans hover:underline"
        >Forgot password?</NuxtLink
      >
    </div>
    <UiButton
      @click="submit"
      :loading="loading"
      size="md"
      class="w-full justify-center"
    >
      Sign In
    </UiButton>
  </div>
</template>
