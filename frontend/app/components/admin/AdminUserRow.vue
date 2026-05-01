<script setup lang="ts">
defineProps<{
  id: string;
  name: string;
  email: string;
  role: string;
  status: string;
  joined: string;
}>();
defineEmits<{ edit: [id: string]; delete: [id: string] }>();
</script>

<template>
  <tr
    class="border-b border-surface-200 last:border-0 hover:bg-surface-100 transition-colors"
  >
    <td class="px-5 py-3.5">
      <div class="flex items-center gap-3">
        <div
          class="w-7 h-7 rounded-full bg-forest/10 flex items-center justify-center text-xs font-medium text-forest"
        >
          {{
            name
              .split(" ")
              .map((w: string) => w[0])
              .join("")
          }}
        </div>
        <span class="text-sm font-sans">{{ name }}</span>
      </div>
    </td>
    <td class="px-5 py-3.5 text-sm text-muted">{{ email }}</td>
    <td class="px-5 py-3.5">
      <UiBadge :variant="role === 'admin' ? 'gold' : 'muted'">{{
        role
      }}</UiBadge>
    </td>
    <td class="px-5 py-3.5">
      <span
        :class="[
          'text-xs font-sans font-medium flex items-center gap-1.5',
          status === 'active' ? 'text-green-600' : 'text-muted',
        ]"
      >
        <span
          :class="[
            'w-1.5 h-1.5 rounded-full',
            status === 'active' ? 'bg-green-500' : 'bg-surface-300',
          ]"
        />
        {{ status }}
      </span>
    </td>
    <td class="px-5 py-3.5 text-sm text-muted">{{ joined }}</td>
    <td class="px-5 py-3.5 text-right space-x-3">
      <button
        @click="$emit('edit', id)"
        class="text-muted hover:text-forest transition-colors"
      >
        <Icon name="lucide:pencil" class="w-3.5 h-3.5" />
      </button>
      <button
        @click="$emit('delete', id)"
        class="text-muted hover:text-red-500 transition-colors"
      >
        <Icon name="lucide:trash-2" class="w-3.5 h-3.5" />
      </button>
    </td>
  </tr>
</template>
