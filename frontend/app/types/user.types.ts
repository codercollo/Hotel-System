export interface User {
  id: string;
  email: string;
  name: string;
  role: string;
  is_active: boolean;
  metadata: Record<string, unknown>;
  created_at: string;
  updated_at: string;
}
