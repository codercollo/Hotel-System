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
  <div class="relative flex items-center w-full">
    <Icon
      name="lucide:search"
      class="absolute left-4 pl-2 text-muted w-4 h-4 pointer-events-none z-10"
    />
    <input
      :value="val"
      :placeholder="placeholder ?? 'Search rooms…'"
      class="input pl-10 pr-28 w-full"
      @input="onInput"
      @keydown.enter="onSearch"
    />
    <button
      @click="onSearch"
      class="absolute right-1.5 btn-primary text-[10px] uppercase font-bold tracking-widest px-4 h-[calc(100%-12px)] flex items-center justify-center"
    >
      SEARCH
    </button>
  </div>
</template>
