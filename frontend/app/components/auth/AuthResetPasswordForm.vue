<script setup lang="ts">
const emit = defineEmits<{ success: [] }>();

const password = ref("");
const confirm = ref("");
const loading = ref(false);
const error = ref("");

const submit = async () => {
  if (password.value !== confirm.value) {
    error.value = "Passwords do not match.";
    return;
  }
  loading.value = true;
  error.value = "";
  await new Promise((r) => setTimeout(r, 800));
  loading.value = false;
  emit("success");
};
</script>

<template>
  <div class="space-y-4">
    <UiAlert v-if="error" type="error" :message="error" />
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
      @click="submit"
      :loading="loading"
      size="md"
      class="w-full justify-center"
    >
      Set New Password
    </UiButton>
  </div>
</template>
