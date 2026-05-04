<script setup lang="ts">
definePageMeta({
  layout: "admin",
  middleware: ["admin"],
  requiresAuth: true,
});
useHead({ title: "Roles & Permissions" });

const api = useApi();

const { data: roles, pending: loading } = await useAsyncData(
  "admin-roles",
  () =>
    api.get<
      { id: string; name: string; description: string; permissions: string[] }[]
    >("/api/v1/roles"),
);
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="font-display text-2xl text-brand-900">
          Roles &amp; Permissions
        </h2>
        <p class="text-xs text-muted font-sans mt-0.5">
          {{ roles?.length ?? 0 }} roles defined
        </p>
      </div>
    </div>

    <div v-if="loading" class="grid grid-cols-1 lg:grid-cols-2 gap-5">
      <UiSkeleton v-for="i in 2" :key="i" class="h-40 rounded-2xl" />
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-2 gap-5">
      <div v-for="role in roles ?? []" :key="role.id" class="card p-5">
        <div class="flex items-start justify-between mb-3">
          <div>
            <div class="flex items-center gap-2 mb-1">
              <UiBadge :variant="role.name === 'admin' ? 'gold' : 'muted'">{{
                role.name
              }}</UiBadge>
            </div>
            <p class="text-sm font-sans text-muted">{{ role.description }}</p>
          </div>
        </div>
        <div
          class="flex flex-wrap gap-1.5 mt-4 pt-4 border-t border-surface-200"
        >
          <span
            v-for="p in role.permissions"
            :key="p"
            class="text-[10px] font-mono bg-surface-100 text-brand-700 px-2 py-0.5 rounded border border-surface-200"
            >{{ p }}</span
          >
          <span
            v-if="!role.permissions?.length"
            class="text-xs text-muted font-sans"
            >No permissions defined</span
          >
        </div>
      </div>
    </div>
  </div>
</template>
