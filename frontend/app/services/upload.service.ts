export const uploadService = {
  upload: (file: File) => {
    const form = new FormData();
    form.append("file", file);
    const config = useRuntimeConfig();
    return fetch(`${config.public.apiBase}/api/v1/uploads`, {
      method: "POST",
      body: form,
    })
      .then((r) => r.json())
      .then((b) => b.data);
  },
  delete: (id: string) => useApi().del(`/api/v1/uploads/${id}`),
};
