<script setup lang="ts">
import type { User } from "~/types/user.types";

definePageMeta({
  layout: "admin",
  middleware: ["admin"],
  requiresAuth: true,
});
useHead({ title: "Users" });

const api = useApi();
const search = ref("");

const { data: users, pending: loading } = await useAsyncData<User[]>(
  "admin-users",
  () => api.get<User[]>("/api/v1/users", { limit: "100" }),
);

const filtered = computed(() =>
  (users.value ?? []).filter(
    (u) =>
      !search.value ||
      u.name.toLowerCase().includes(search.value.toLowerCase()) ||
      u.email.toLowerCase().includes(search.value.toLowerCase()),
  ),
);

const formatDate = (d: string) =>
  new Date(d).toLocaleDateString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
  });
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="font-display text-2xl text-brand-900">Users</h2>
        <p class="text-xs text-muted font-sans mt-0.5">
          {{ users?.length ?? 0 }} total users
        </p>
      </div>
    </div>

    <div class="mb-4 max-w-xs">
      <UiInput
        v-model="search"
        placeholder="Search users…"
        icon="lucide:search"
      />
    </div>

    <div v-if="loading" class="space-y-3">
      <UiSkeleton v-for="i in 5" :key="i" class="h-14 rounded-lg" />
    </div>

    <UiTable
      v-else
      :headers="[
        { key: 'name', label: 'Name' },
        { key: 'email', label: 'Email' },
        { key: 'role', label: 'Role' },
        { key: 'status', label: 'Status' },
        { key: 'joined', label: 'Joined' },
      ]"
    >
      <AdminUserRow
        v-for="u in filtered"
        :key="u.id"
        :id="u.id"
        :name="u.name"
        :email="u.email"
        :role="u.role"
        :status="u.is_active ? 'active' : 'inactive'"
        :joined="formatDate(u.created_at)"
        @edit="() => {}"
        @delete="() => {}"
      />
      <tr v-if="!filtered.length">
        <td
          colspan="6"
          class="px-5 py-12 text-center text-muted font-sans text-sm"
        >
          No users found.
        </td>
      </tr>
    </UiTable>
  </div>
</template>
