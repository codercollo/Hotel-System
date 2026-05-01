<script setup lang="ts">
definePageMeta({ layout: "auth" });
useHead({ title: "Forgot Password" });

const email = ref("");
const sent = ref(false);
const loading = ref(false);

const onSubmit = async () => {
  if (!email.value) return;
  loading.value = true;
  await new Promise((r) => setTimeout(r, 800));
  loading.value = false;
  sent.value = true;
};
</script>

<template>
  <div>
    <NuxtLink
      to="/auth/login"
      class="inline-flex items-center gap-1.5 text-xs font-sans text-muted hover:text-forest transition-colors mb-8"
    >
      <Icon name="lucide:arrow-left" class="w-3.5 h-3.5" /> Back to sign in
    </NuxtLink>

    <template v-if="!sent">
      <p class="eyebrow mb-3">Account Recovery</p>
      <h2 class="font-display text-display-md text-brand-900 font-light mb-2">
        Forgot your password?
      </h2>
      <p class="text-sm text-muted font-sans mb-8">
        Enter your email address and we'll send you a link to reset your
        password.
      </p>

      <div class="space-y-4">
        <UiInput
          v-model="email"
          label="Email Address"
          type="email"
          placeholder="you@example.com"
          icon="lucide:mail"
        />
        <UiButton
          @click="onSubmit"
          :loading="loading"
          size="md"
          class="w-full justify-center"
        >
          Send Reset Link
        </UiButton>
      </div>
    </template>

    <template v-else>
      <div class="text-center py-6">
        <div
          class="w-16 h-16 rounded-full bg-forest/10 flex items-center justify-center mx-auto mb-5"
        >
          <Icon name="lucide:mail-check" class="w-8 h-8 text-forest" />
        </div>
        <h2 class="font-display text-display-sm text-brand-900 font-light mb-3">
          Check your inbox
        </h2>
        <p class="text-sm text-muted font-sans mb-6">
          We've sent a password reset link to
          <strong class="text-brand-900">{{ email }}</strong
          >. The link expires in 30 minutes.
        </p>
        <NuxtLink to="/auth/login">
          <UiButton variant="outline" size="sm">Return to Sign In</UiButton>
        </NuxtLink>
      </div>
    </template>
  </div>
</template>
