export interface Notification {
  id: string;
  user_id: string;
  type: string;
  channel: string;
  title: string;
  body: string;
  is_read: boolean;
  metadata?: Record<string, unknown>;
  created_at: string;
  read_at?: string;
}
