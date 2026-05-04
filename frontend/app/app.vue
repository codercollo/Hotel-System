<template>
  <ClientOnly>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
    <template #fallback>
      <!-- Blank while store rehydrates — prevents hydration mismatch -->
      <div style="visibility: hidden" />
    </template>
  </ClientOnly>
</template>
