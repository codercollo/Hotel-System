export interface Payment {
  id: string;
  order_id: string;
  provider: string;
  provider_ref?: string;
  amount: number;
  currency: string;
  status: "pending" | "processing" | "completed" | "failed" | "refunded";
  webhook_payload?: Record<string, unknown>;
  metadata?: Record<string, unknown>;
  created_at: string;
  updated_at: string;
}

export interface InitiateRequest {
  order_id: string;
  provider: string;
  phone?: string;
}

export interface InitiateResponse {
  payment_id: string;
  checkout_url?: string;
  checkout_code?: string;
  status: string;
}
