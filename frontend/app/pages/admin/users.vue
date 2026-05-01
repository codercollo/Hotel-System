<script setup lang="ts">
definePageMeta({ layout: "admin" });
useHead({ title: "Users" });

const search = ref("");
const showAdd = ref(false);

const users = [
  {
    id: "u1",
    name: "Alice Johnson",
    email: "alice@example.com",
    role: "user",
    status: "active",
    joined: "2025-01-12",
  },
  {
    id: "u2",
    name: "Bob Martinez",
    email: "bob@example.com",
    role: "user",
    status: "active",
    joined: "2025-02-20",
  },
  {
    id: "u3",
    name: "Carol White",
    email: "carol@example.com",
    role: "admin",
    status: "active",
    joined: "2024-11-05",
  },
  {
    id: "u4",
    name: "David Lee",
    email: "david@example.com",
    role: "user",
    status: "inactive",
    joined: "2025-03-01",
  },
];

const filtered = computed(() =>
  users.filter(
    (u) =>
      !search.value ||
      u.name.toLowerCase().includes(search.value.toLowerCase()) ||
      u.email.toLowerCase().includes(search.value.toLowerCase()),
  ),
);
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="font-display text-2xl text-brand-900">Users</h2>
        <p class="text-xs text-muted font-sans mt-0.5">
          {{ users.length }} total users
        </p>
      </div>
      <UiButton size="sm" @click="showAdd = true">
        <Icon name="lucide:user-plus" class="w-3.5 h-3.5" /> Add User
      </UiButton>
    </div>

    <!-- Search -->
    <div class="mb-4 max-w-xs">
      <UiInput
        v-model="search"
        placeholder="Search users…"
        icon="lucide:search"
      />
    </div>

    <UiTable
      :headers="[
        { key: 'name', label: 'Name' },
        { key: 'email', label: 'Email' },
        { key: 'role', label: 'Role' },
        { key: 'status', label: 'Status' },
        { key: 'joined', label: 'Joined' },
        { key: 'actions', label: '', align: 'right' },
      ]"
    >
      <tr
        v-for="u in filtered"
        :key="u.id"
        class="border-b border-surface-200 last:border-0 hover:bg-surface-100 transition-colors"
      >
        <td class="px-5 py-3.5">
          <div class="flex items-center gap-3">
            <div
              class="w-7 h-7 rounded-full bg-forest/10 flex items-center justify-center text-xs font-medium text-forest"
            >
              {{
                u.name
                  .split(" ")
                  .map((w) => w[0])
                  .join("")
              }}
            </div>
            <span class="text-sm font-sans">{{ u.name }}</span>
          </div>
        </td>
        <td class="px-5 py-3.5 text-sm text-muted">{{ u.email }}</td>
        <td class="px-5 py-3.5">
          <UiBadge :variant="u.role === 'admin' ? 'gold' : 'muted'">{{
            u.role
          }}</UiBadge>
        </td>
        <td class="px-5 py-3.5">
          <span
            :class="[
              'text-xs font-sans font-medium',
              u.status === 'active' ? 'text-green-600' : 'text-muted',
            ]"
          >
            <span
              :class="[
                'inline-block w-1.5 h-1.5 rounded-full mr-1.5',
                u.status === 'active' ? 'bg-green-500' : 'bg-surface-300',
              ]"
            />
            {{ u.status }}
          </span>
        </td>
        <td class="px-5 py-3.5 text-sm text-muted">{{ u.joined }}</td>
        <td class="px-5 py-3.5 text-right">
          <button class="text-muted hover:text-forest transition-colors mr-3">
            <Icon name="lucide:pencil" class="w-3.5 h-3.5" />
          </button>
          <button class="text-muted hover:text-red-500 transition-colors">
            <Icon name="lucide:trash-2" class="w-3.5 h-3.5" />
          </button>
        </td>
      </tr>
    </UiTable>

    <!-- Add user modal -->
    <UiModal :open="showAdd" title="Add New User" @close="showAdd = false">
      <div class="space-y-4">
        <UiInput label="Full Name" placeholder="Jane Doe" icon="lucide:user" />
        <UiInput
          label="Email Address"
          type="email"
          placeholder="jane@example.com"
          icon="lucide:mail"
        />
        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >Role</label
          >
          <select class="input">
            <option value="user">User</option>
            <option value="admin">Admin</option>
          </select>
        </div>
        <div class="flex justify-end gap-3 pt-2">
          <UiButton variant="outline" size="sm" @click="showAdd = false"
            >Cancel</UiButton
          >
          <UiButton size="sm" @click="showAdd = false">Create User</UiButton>
        </div>
      </div>
    </UiModal>
  </div>
</template>
