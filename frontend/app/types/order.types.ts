export interface Order {
  id: string;
  user_id: string;
  status: "pending" | "confirmed" | "processing" | "completed" | "cancelled";
  total: number;
  currency: string;
  notes: string;
  metadata: Record<string, unknown>;
  items: OrderItem[];
  created_at: string;
  updated_at: string;
}

export interface OrderItem {
  id: string;
  order_id: string;
  item_id: string;
  name: string;
  price: number;
  quantity: number;
  subtotal: number;
}

export interface CreateOrderRequest {
  items: { item_id: string; quantity: number }[];
  notes?: string;
  metadata?: Record<string, unknown>;
}
