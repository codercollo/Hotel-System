<script setup lang="ts">
definePageMeta({ layout: "admin" });
useHead({ title: "Room Management" });

const showForm = ref(false);
const editing = ref<any>(null);

const items = ref([
  { id: "r1", name: "Deluxe Room", price: 25000, status: "active", stock: 10 },
  {
    id: "r2",
    name: "The Pearl Suite",
    price: 45000,
    status: "active",
    stock: 5,
  },
  {
    id: "r3",
    name: "Golden Executive",
    price: 55000,
    status: "active",
    stock: 3,
  },
  {
    id: "r4",
    name: "Classic Room",
    price: 18000,
    status: "inactive",
    stock: 8,
  },
  { id: "r5", name: "Family Suite", price: 38000, status: "active", stock: 4 },
  {
    id: "r6",
    name: "Honeymoon Suite",
    price: 65000,
    status: "active",
    stock: 2,
  },
]);

const formatPrice = (p: number) =>
  (p / 100).toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  });

const openEdit = (item: any) => {
  editing.value = { ...item };
  showForm.value = true;
};
const openNew = () => {
  editing.value = { name: "", price: 0, status: "active", stock: 1 };
  showForm.value = true;
};
const save = () => {
  showForm.value = false;
  editing.value = null;
};

const toggleStatus = (item: any) => {
  item.status = item.status === "active" ? "inactive" : "active";
};
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="font-display text-2xl text-brand-900">Room Management</h2>
        <p class="text-xs text-muted font-sans mt-0.5">
          {{ items.length }} rooms listed
        </p>
      </div>
      <UiButton size="sm" @click="openNew">
        <Icon name="lucide:plus" class="w-3.5 h-3.5" /> Add Room
      </UiButton>
    </div>

    <UiTable
      :headers="[
        { key: 'name', label: 'Room Name' },
        { key: 'price', label: 'Price/Night', align: 'right' },
        { key: 'stock', label: 'Available', align: 'right' },
        { key: 'status', label: 'Status' },
        { key: 'actions', label: '', align: 'right' },
      ]"
    >
      <tr
        v-for="item in items"
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
            <span class="font-display text-base text-brand-900">{{
              item.name
            }}</span>
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
          <NuxtLink
            :to="`/items/${item.id}`"
            class="text-muted hover:text-forest transition-colors"
          >
            <Icon name="lucide:external-link" class="w-3.5 h-3.5" />
          </NuxtLink>
        </td>
      </tr>
    </UiTable>

    <!-- Create / Edit modal -->
    <UiModal
      :open="showForm"
      :title="editing?.id ? 'Edit Room' : 'Add New Room'"
      @close="save"
    >
      <div v-if="editing" class="space-y-4">
        <UiInput
          v-model="editing.name"
          label="Room Name"
          placeholder="e.g. Deluxe Suite"
          icon="lucide:bed-double"
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
              (editing.price / 100).toLocaleString("en-US", {
                style: "currency",
                currency: "USD",
              })
            }}
            per night
          </p>
        </div>

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

        <div class="flex justify-end gap-3 pt-2">
          <UiButton variant="outline" size="sm" @click="save">Cancel</UiButton>
          <UiButton size="sm" @click="save">{{
            editing.id ? "Save Changes" : "Create Room"
          }}</UiButton>
        </div>
      </div>
    </UiModal>
  </div>
</template>
