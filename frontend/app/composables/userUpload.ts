export interface UploadedFile {
  id: string;
  url: string;
  filename: string;
  original_name: string;
  mime_type: string;
  size: number;
}

export const useUpload = () => {
  const uploading = ref(false);
  const error = ref<string | null>(null);

  const upload = async (file: File): Promise<UploadedFile | null> => {
    uploading.value = true;
    error.value = null;

    try {
      const config = useRuntimeConfig();
      const token = import.meta.client
        ? localStorage.getItem("access_token")
        : null;
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
      return body?.data ?? body;
    } catch (e: any) {
      error.value = e.message;
      return null;
    } finally {
      uploading.value = false;
    }
  };

  const deleteUpload = async (id: string) => {
    const api = useApi();
    await api.del(`/api/v1/uploads/${id}`);
  };

  return { uploading, error, upload, deleteUpload };
};
