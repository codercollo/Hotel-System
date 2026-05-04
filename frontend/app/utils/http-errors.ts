export class HttpError extends Error {
  constructor(
    public status: number,
    public code: string,
    message: string,
    public detail?: unknown,
  ) {
    super(message);
    this.name = "HttpError";
  }
}

export const isUnauthorized = (e: unknown) =>
  e instanceof HttpError && e.status === 401;
export const isForbidden = (e: unknown) =>
  e instanceof HttpError && e.status === 403;
export const isNotFound = (e: unknown) =>
  e instanceof HttpError && e.status === 404;
export const isValidationError = (e: unknown) =>
  e instanceof HttpError && e.status === 422;

export const extractMessage = (e: unknown): string => {
  if (e instanceof HttpError) return e.message;
  if (e instanceof Error) return e.message;
  return "An unexpected error occurred";
};
