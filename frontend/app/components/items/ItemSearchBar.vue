<script setup lang="ts">
const props = defineProps<{ modelValue?: string; placeholder?: string }>();
const emit = defineEmits<{
  "update:modelValue": [v: string];
  search: [v: string];
}>();
const val = ref(props.modelValue ?? "");

watch(
  () => props.modelValue,
  (v) => {
    val.value = v ?? "";
  },
);

const onInput = (e: Event) => {
  val.value = (e.target as HTMLInputElement).value;
  emit("update:modelValue", val.value);
};
const onSearch = () => emit("search", val.value);
</script>

<template>
  <div class="relative flex items-center">
    <Icon
      name="lucide:search"
      class="absolute left-3.5 text-muted w-4 h-4 pointer-events-none"
    />
    <input
      :value="val"
      :placeholder="placeholder ?? 'Search rooms…'"
      class="input pl-10 pr-24"
      @input="onInput"
      @keydown.enter="onSearch"
    />
    <button
      @click="onSearch"
      class="absolute right-2 btn-primary text-xs px-4 py-1.5"
    >
      Search
    </button>
  </div>
</template>
