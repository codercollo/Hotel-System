type ToastType = "success" | "error" | "warning" | "info";

interface Toast {
  id: string;
  type: ToastType;
  message: string;
}

const toasts = ref<Toast[]>([]);

export const useToast = () => {
  const add = (message: string, type: ToastType = "info", duration = 4000) => {
    const id = Math.random().toString(36).slice(2);
    toasts.value.push({ id, type, message });
    setTimeout(() => remove(id), duration);
  };

  const remove = (id: string) => {
    toasts.value = toasts.value.filter((t) => t.id !== id);
  };

  return {
    toasts: readonly(toasts),
    success: (msg: string) => add(msg, "success"),
    error: (msg: string) => add(msg, "error"),
    warning: (msg: string) => add(msg, "warning"),
    info: (msg: string) => add(msg, "info"),
    remove,
  };
};
