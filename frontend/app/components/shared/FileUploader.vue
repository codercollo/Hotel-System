<script setup lang="ts">
const props = defineProps<{
  accept?: string;
  maxSizeMb?: number;
  multiple?: boolean;
}>();

const emit = defineEmits<{
  uploaded: [urls: string[]];
  error: [message: string];
}>();

const uploading = ref(false);
const progress = ref(0);
const previews = ref<{ name: string; url: string; id: string }[]>([]);
const dragOver = ref(false);
const inputRef = ref<HTMLInputElement>();

const maxBytes = computed(() => (props.maxSizeMb ?? 10) * 1024 * 1024);
const accept = computed(() => props.accept ?? "image/*");

async function uploadFiles(files: FileList | File[]) {
  const list = Array.from(files);
  if (!list.length) return;

  for (const file of list) {
    if (file.size > maxBytes.value) {
      emit("error", `${file.name} exceeds ${props.maxSizeMb ?? 10}MB limit`);
      return;
    }
  }

  uploading.value = true;
  progress.value = 0;

  const urls: string[] = [];

  try {
    const config = useRuntimeConfig();
    const token = import.meta.client
      ? localStorage.getItem("access_token")
      : null;
    for (const file of list) {
      const form = new FormData();
      form.append("file", file);

      const res = await fetch(`${config.public.apiBase}/api/v1/uploads`, {
        method: "POST",
        headers: token ? { Authorization: `Bearer ${token}` } : {},
        body: form,
      });

      if (!res.ok) {
        const body = await res.json().catch(() => ({}));
        throw new Error(body?.error?.message ?? "Upload failed");
      }

      const body = await res.json();
      const upload = body?.data ?? body;
      previews.value.push({ name: file.name, url: upload.url, id: upload.id });
      urls.push(upload.url);
      progress.value = Math.round((previews.value.length / list.length) * 100);
    }

    emit("uploaded", urls);
  } catch (e: any) {
    emit("error", e.message ?? "Upload failed");
  } finally {
    uploading.value = false;
  }
}

const onFileInput = (e: Event) => {
  const files = (e.target as HTMLInputElement).files;
  if (files) uploadFiles(files);
};

const onDrop = (e: DragEvent) => {
  dragOver.value = false;
  const files = e.dataTransfer?.files;
  if (files) uploadFiles(files);
};

const remove = (id: string) => {
  previews.value = previews.value.filter((p) => p.id !== id);
};
</script>

<template>
  <div class="space-y-3">
    <!-- Drop zone -->
    <div
      :class="[
        'border-2 border-dashed rounded-2xl p-8 text-center transition-all cursor-pointer',
        dragOver
          ? 'border-forest bg-forest/5'
          : 'border-surface-200 hover:border-forest/40 hover:bg-surface-50',
        uploading ? 'pointer-events-none opacity-60' : '',
      ]"
      @click="inputRef?.click()"
      @dragover.prevent="dragOver = true"
      @dragleave="dragOver = false"
      @drop.prevent="onDrop"
    >
      <input
        ref="inputRef"
        type="file"
        class="hidden"
        :accept="accept"
        :multiple="multiple"
        @change="onFileInput"
      />

      <div v-if="!uploading">
        <Icon
          name="lucide:upload-cloud"
          class="w-10 h-10 text-muted mx-auto mb-3"
        />
        <p class="text-sm font-sans text-brand-900 font-medium">
          Drop files here or <span class="text-forest underline">browse</span>
        </p>
        <p class="text-xs text-muted font-sans mt-1">
          {{ accept }} · Max {{ maxSizeMb ?? 10 }}MB
        </p>
      </div>

      <div v-else class="space-y-2">
        <Icon
          name="lucide:loader-2"
          class="w-8 h-8 text-forest mx-auto animate-spin"
        />
        <p class="text-sm font-sans text-muted">Uploading… {{ progress }}%</p>
        <div class="w-full bg-surface-200 rounded-full h-1.5 overflow-hidden">
          <div
            class="bg-forest h-1.5 rounded-full transition-all duration-300"
            :style="{ width: `${progress}%` }"
          />
        </div>
      </div>
    </div>

    <!-- Previews -->
    <div v-if="previews.length" class="grid grid-cols-3 sm:grid-cols-4 gap-2">
      <div
        v-for="p in previews"
        :key="p.id"
        class="relative group rounded-lg overflow-hidden aspect-square bg-surface-200"
      >
        <img :src="p.url" :alt="p.name" class="w-full h-full object-cover" />
        <button
          class="absolute top-1 right-1 w-5 h-5 rounded-full bg-brand-900/70 text-white flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity"
          @click.stop="remove(p.id)"
        >
          <Icon name="lucide:x" class="w-3 h-3" />
        </button>
      </div>
    </div>
  </div>
</template>
