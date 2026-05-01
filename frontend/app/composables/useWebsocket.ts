// useWebSocket — Phase 9 implements the full real-time layer
export const useWebSocket = () => {
  const config = useRuntimeConfig();
  let ws: WebSocket | null = null;
  const connected = ref(false);
  const handlers = new Map<string, ((payload: any) => void)[]>();

  const connect = (token: string) => {
    if (ws) return;
    ws = new WebSocket(`${config.public.wsBase}/ws?token=${token}`);

    ws.onopen = () => {
      connected.value = true;
    };
    ws.onclose = () => {
      connected.value = false;
      ws = null;
    };
    ws.onerror = () => {
      connected.value = false;
    };

    ws.onmessage = (e) => {
      try {
        const msg = JSON.parse(e.data);
        const fns = handlers.get(msg.type) ?? [];
        fns.forEach((fn) => fn(msg.payload));
      } catch {
        /* ignore malformed */
      }
    };
  };

  const disconnect = () => {
    ws?.close();
    ws = null;
    connected.value = false;
  };

  const on = (type: string, fn: (payload: any) => void) => {
    if (!handlers.has(type)) handlers.set(type, []);
    handlers.get(type)!.push(fn);
    return () => {
      const fns = handlers.get(type) ?? [];
      handlers.set(
        type,
        fns.filter((f) => f !== fn),
      );
    };
  };

  return { connected, connect, disconnect, on };
};
