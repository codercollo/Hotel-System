// useUpload — Phase 7 wires this to the uploads API
export const useUpload = () => {
  const api = useApi();
  const loading = ref(false);
  const progress = ref(0);
  const error = ref<string | null>(null);

  const upload = async (file: File): Promise<{ url: string; id: string }> => {
    loading.value = true;
    progress.value = 0;
    error.value = null;

    const form = new FormData();
    form.append("file", file);

    try {
      // Phase 7: replace with presigned URL flow
      const config = useRuntimeConfig();
      const res = await fetch(`${config.public.apiBase}/api/v1/uploads`, {
        method: "POST",
        body: form,
      });

      if (!res.ok) throw new Error("Upload failed");
      const body = await res.json();
      progress.value = 100;
      return body.data;
    } catch (e: any) {
      error.value = e.message;
      throw e;
    } finally {
      loading.value = false;
    }
  };

  return { loading, progress, error, upload };
};
