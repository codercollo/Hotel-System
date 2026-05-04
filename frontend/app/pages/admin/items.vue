<script setup lang="ts">
import type { Item } from "~/types/item.types";

definePageMeta({
  layout: "admin",
  middleware: ["admin"],
  requiresAuth: true,
});
useHead({ title: "Room Management" });

const api = useApi();
const {
  data: items,
  pending: loading,
  refresh,
} = await useAsyncData<Item[]>("admin-items", () =>
  api.get<Item[]>("/api/v1/items"),
);

const showForm = ref(false);
const saving = ref(false);
const saveError = ref<string | null>(null);
const editing = ref<Partial<Item> & { metadata_raw?: string }>({});

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });

const openEdit = (item: Item) => {
  editing.value = {
    ...item,
    metadata_raw: JSON.stringify(item.metadata ?? {}, null, 2),
  };
  saveError.value = null;
  showForm.value = true;
};

const openNew = () => {
  editing.value = {
    name: "",
    description: "",
    price: 0,
    currency: "USD",
    stock: 1,
    status: "active",
    metadata_raw: JSON.stringify(
      { beds: 1, baths: 1, sqft: 0, rating: 4.5, badges: [] },
      null,
      2,
    ),
  };
  saveError.value = null;
  showForm.value = true;
};

const save = async () => {
  saving.value = true;
  saveError.value = null;
  try {
    let metadata = {};
    try {
      metadata = JSON.parse(editing.value.metadata_raw ?? "{}");
    } catch {
      metadata = {};
    }

    const payload = {
      name: editing.value.name,
      description: editing.value.description,
      price: editing.value.price,
      currency: editing.value.currency ?? "USD",
      stock: editing.value.stock,
      status: editing.value.status,
      metadata,
    };

    if (editing.value.id) {
      await api.patch(`/api/v1/items/${editing.value.id}`, payload);
    } else {
      await api.post("/api/v1/items", payload);
    }

    await refresh();
    showForm.value = false;
    editing.value = {};
  } catch (e: any) {
    saveError.value = e.message;
  } finally {
    saving.value = false;
  }
};

const deleteItem = async (id: string) => {
  if (!confirm("Delete this room?")) return;
  await api.del(`/api/v1/items/${id}`);
  await refresh();
};

const toggleStatus = async (item: Item) => {
  const newStatus = item.status === "active" ? "inactive" : "active";
  await api.patch(`/api/v1/items/${item.id}`, { status: newStatus });
  await refresh();
};
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="font-display text-2xl text-brand-900">Room Management</h2>
        <p class="text-xs text-muted font-sans mt-0.5">
          {{ items?.length ?? 0 }} rooms listed
        </p>
      </div>
      <UiButton size="sm" @click="openNew">
        <Icon name="lucide:plus" class="w-3.5 h-3.5" /> Add Room
      </UiButton>
    </div>

    <div v-if="loading" class="space-y-3">
      <UiSkeleton v-for="i in 4" :key="i" class="h-14 rounded-lg" />
    </div>

    <UiTable
      v-else
      :headers="[
        { key: 'name', label: 'Room Name' },
        { key: 'price', label: 'Price/Night', align: 'right' },
        { key: 'stock', label: 'Available', align: 'right' },
        { key: 'status', label: 'Status' },
        { key: 'actions', label: '', align: 'right' },
      ]"
    >
      <tr
        v-for="item in items ?? []"
        :key="item.id"
        class="border-b border-surface-200 last:border-0 hover:bg-surface-100 transition-colors"
      >
        <td class="px-5 py-3.5">
          <div class="flex items-center gap-3">
            <div
              class="w-8 h-8 rounded-lg bg-forest/8 flex items-center justify-center"
            >
              <Icon name="lucide:bed-double" class="w-4 h-4 text-forest" />
            </div>
            <div>
              <span class="font-display text-base text-brand-900">{{
                item.name
              }}</span>
              <p class="text-xs text-muted font-mono">{{ item.id }}</p>
            </div>
          </div>
        </td>
        <td class="px-5 py-3.5 text-sm text-right font-sans font-medium">
          {{ formatPrice(item.price) }}
        </td>
        <td class="px-5 py-3.5 text-sm text-right text-muted">
          {{ item.stock }}
        </td>
        <td class="px-5 py-3.5">
          <button @click="toggleStatus(item)">
            <UiBadge :variant="item.status === 'active' ? 'forest' : 'muted'">{{
              item.status
            }}</UiBadge>
          </button>
        </td>
        <td class="px-5 py-3.5 text-right space-x-3">
          <button
            @click="openEdit(item)"
            class="text-muted hover:text-forest transition-colors"
          >
            <Icon name="lucide:pencil" class="w-3.5 h-3.5" />
          </button>
          <button
            @click="deleteItem(item.id)"
            class="text-muted hover:text-red-500 transition-colors"
          >
            <Icon name="lucide:trash-2" class="w-3.5 h-3.5" />
          </button>
          <NuxtLink
            :to="`/items/${item.id}`"
            class="text-muted hover:text-forest transition-colors"
          >
            <Icon name="lucide:external-link" class="w-3.5 h-3.5" />
          </NuxtLink>
        </td>
      </tr>
    </UiTable>

    <UiModal
      :open="showForm"
      :title="editing?.id ? 'Edit Room' : 'Add New Room'"
      @close="showForm = false"
    >
      <div v-if="editing" class="space-y-4">
        <UiInput
          v-model="editing.name"
          label="Room Name"
          placeholder="e.g. Deluxe Suite"
          icon="lucide:bed-double"
        />
        <UiInput
          v-model="editing.description"
          label="Description"
          placeholder="Room description…"
          icon="lucide:file-text"
        />

        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >Price per Night (cents)</label
          >
          <input
            v-model.number="editing.price"
            type="number"
            class="input"
            placeholder="e.g. 25000 = $250"
          />
          <p class="text-xs text-muted font-sans mt-1">
            =
            {{
              ((editing.price ?? 0) / 100).toLocaleString("en-US", {
                style: "currency",
                currency: "USD",
              })
            }}
            per night
          </p>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <div>
            <label
              class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
              >Available Units</label
            >
            <input
              v-model.number="editing.stock"
              type="number"
              min="0"
              class="input"
            />
          </div>
          <div>
            <label
              class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
              >Status</label
            >
            <select v-model="editing.status" class="input">
              <option value="active">Active</option>
              <option value="inactive">Inactive</option>
              <option value="archived">Archived</option>
            </select>
          </div>
        </div>

        <div>
          <label
            class="block text-xs uppercase tracking-widest font-sans text-muted mb-1.5"
            >Metadata (JSON)</label
          >
          <textarea
            v-model="editing.metadata_raw"
            class="input font-mono text-xs"
            rows="5"
          />
          <p class="text-xs text-muted mt-1">
            Keys: beds, baths, sqft, floor, view, rating, badges
          </p>
        </div>

        <UiAlert v-if="saveError" variant="error" :message="saveError" />

        <div class="flex justify-end gap-3 pt-2">
          <UiButton variant="outline" size="sm" @click="showForm = false"
            >Cancel</UiButton
          >
          <UiButton size="sm" :loading="saving" @click="save">
            {{ editing.id ? "Save Changes" : "Create Room" }}
          </UiButton>
        </div>
      </div>
    </UiModal>
  </div>
</template>
