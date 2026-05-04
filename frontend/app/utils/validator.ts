export const isEmail = (v: string) => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(v);

export const isRequired = (v: unknown) =>
  v !== null && v !== undefined && String(v).trim().length > 0;

export const minLength = (min: number) => (v: string) => v.length >= min;

export const maxLength = (max: number) => (v: string) => v.length <= max;

export const isPhone = (v: string) => /^\+?[\d\s\-()]{7,}$/.test(v);

export const isUrl = (v: string) => {
  try {
    new URL(v);
    return true;
  } catch {
    return false;
  }
};

// Compose validators
export const validate = (
  value: unknown,
  rules: ((v: any) => boolean | string)[],
) => {
  for (const rule of rules) {
    const result = rule(value);
    if (result !== true) return result;
  }
  return null;
};
