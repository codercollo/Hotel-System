// Currency
export const formatCurrency = (
  amount: number,
  currency = "USD",
  locale = "en-US",
): string =>
  (amount / 100).toLocaleString(locale, {
    style: "currency",
    currency,
    maximumFractionDigits: 0,
  });

// Date
export const formatDate = (
  iso: string,
  opts?: Intl.DateTimeFormatOptions,
): string =>
  new Date(iso).toLocaleDateString(
    "en-US",
    opts ?? { year: "numeric", month: "long", day: "numeric" },
  );

// Relative time (basic)
export const timeAgo = (iso: string): string => {
  const diff = Date.now() - new Date(iso).getTime();
  const m = Math.floor(diff / 60_000);
  if (m < 1) return "just now";
  if (m < 60) return `${m}m ago`;
  const h = Math.floor(m / 60);
  if (h < 24) return `${h}h ago`;
  return `${Math.floor(h / 24)}d ago`;
};

// Truncate
export const truncate = (s: string, len: number): string =>
  s.length > len ? `${s.slice(0, len)}…` : s;

// Initials
export const initials = (name: string): string =>
  name
    .split(" ")
    .slice(0, 2)
    .map((w) => w[0])
    .join("")
    .toUpperCase();
