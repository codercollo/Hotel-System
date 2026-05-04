import type { Notification } from "~/types/notification.types";

export const useWebSocket = (wsBase: string) => {
  let socket: WebSocket | null = null;
  let reconnectTimer: ReturnType<typeof setTimeout> | null = null;
  let retryCount = 0;
  let currentToken = "";

  const MAX_RETRIES = 10;
  const BASE_DELAY = 1_000;
  const MAX_DELAY = 30_000;

  function backoffDelay(): number {
    const exp = Math.min(BASE_DELAY * 2 ** retryCount, MAX_DELAY);
    return exp + Math.random() * 1_000;
  }

  function dispatch(msg: { type: string; payload: unknown }) {
    switch (msg.type) {
      case "notification": {
        const notifStore = useNotificationStore();
        notifStore.push(msg.payload as Notification);
        break;
      }
      case "order_status": {
        // Broadcast to any listening composable via a custom event
        // so order detail pages can react without polling
        if (import.meta.client) {
          window.dispatchEvent(
            new CustomEvent("ws:order_status", { detail: msg.payload }),
          );
        }
        break;
      }
      case "payment_status": {
        if (import.meta.client) {
          window.dispatchEvent(
            new CustomEvent("ws:payment_status", { detail: msg.payload }),
          );
        }
        break;
      }
      case "pong":
        break;
      default:
        console.debug("[WS] unhandled message type:", msg.type);
    }
  }

  function connect(token: string) {
    currentToken = token;
    if (socket?.readyState === WebSocket.OPEN) return;

    const url = `${wsBase}/ws?token=${encodeURIComponent(token)}`;
    socket = new WebSocket(url);

    let lastFetch = 0;

    socket.onopen = () => {
      console.debug("[WS] connected");
      retryCount = 0;
      if (reconnectTimer) {
        clearTimeout(reconnectTimer);
        reconnectTimer = null;
      }
      // Only fetch notifications if last fetch was >30s ago
      const now = Date.now();
      if (now - lastFetch > 30_000) {
        lastFetch = now;
        useNotifications().fetch();
      }
    };

    socket.onmessage = (event) => {
      try {
        const msg = JSON.parse(event.data as string);
        dispatch(msg);
      } catch {
        console.warn("[WS] non-JSON message", event.data);
      }
    };

    socket.onerror = () => {};

    socket.onclose = (ev) => {
      console.debug(`[WS] closed (code=${ev.code}, retries=${retryCount})`);
      socket = null;
      const isCleanClose = ev.code === 1000 || ev.code === 1001;
      if (isCleanClose || !currentToken) return;
      if (retryCount >= MAX_RETRIES) {
        console.warn(`[WS] giving up after ${MAX_RETRIES} retries`);
        return;
      }
      const delay = backoffDelay();
      retryCount++;
      reconnectTimer = setTimeout(() => connect(currentToken), delay);
    };
  }

  function disconnect() {
    currentToken = "";
    retryCount = 0;
    if (reconnectTimer) {
      clearTimeout(reconnectTimer);
      reconnectTimer = null;
    }
    if (socket) {
      socket.close(1000, "logout");
      socket = null;
    }
  }

  function send(data: unknown) {
    if (socket?.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify(data));
    }
  }

  return { connect, disconnect, send };
};
