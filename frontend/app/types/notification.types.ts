export interface Notification {
  id: string;
  user_id: string;
  type: string;
  channel: "in_app" | "email" | "sms" | "push";
  title: string;
  body: string;
  is_read: boolean;
  metadata: Record<string, unknown>;
  created_at: string;
  read_at?: string;
}
