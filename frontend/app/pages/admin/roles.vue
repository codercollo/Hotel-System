<script setup lang="ts">
definePageMeta({ layout: "admin" });
useHead({ title: "Roles & Permissions" });

const roles = [
  {
    id: "r1",
    name: "admin",
    description: "Full platform access.",
    permissions: [
      "items:create",
      "items:read",
      "items:update",
      "items:delete",
      "orders:read",
      "orders:update",
      "users:read",
      "users:update",
      "payments:read",
    ],
    users: 2,
  },
  {
    id: "r2",
    name: "user",
    description: "Standard guest access.",
    permissions: ["items:read", "orders:create", "orders:read"],
    users: 1246,
  },
];
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="font-display text-2xl text-brand-900">
          Roles &amp; Permissions
        </h2>
        <p class="text-xs text-muted font-sans mt-0.5">
          {{ roles.length }} roles defined
        </p>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-5">
      <div v-for="role in roles" :key="role.id" class="card p-5">
        <div class="flex items-start justify-between mb-3">
          <div>
            <div class="flex items-center gap-2 mb-1">
              <UiBadge :variant="role.name === 'admin' ? 'gold' : 'muted'">{{
                role.name
              }}</UiBadge>
              <span class="text-xs text-muted font-sans"
                >{{ role.users }} user{{ role.users !== 1 ? "s" : "" }}</span
              >
            </div>
            <p class="text-sm font-sans text-muted">{{ role.description }}</p>
          </div>
          <button class="text-muted hover:text-forest transition-colors">
            <Icon name="lucide:pencil" class="w-4 h-4" />
          </button>
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
        </div>
      </div>
    </div>
  </div>
</template>
