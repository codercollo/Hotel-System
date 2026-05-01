export const ROLES = {
  ADMIN: "admin",
  USER: "user",
} as const;

export const ORDER_STATUSES = [
  "pending",
  "confirmed",
  "processing",
  "completed",
  "cancelled",
] as const;
export const PAYMENT_STATUSES = [
  "pending",
  "processing",
  "completed",
  "failed",
  "refunded",
] as const;

export const DEFAULT_PAGINATION_LIMIT = 20;
export const MAX_UPLOAD_SIZE_MB = 10;
