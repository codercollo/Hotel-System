export interface Payment {
  id: string;
  order_id: string;
  provider: "mpesa" | "stripe" | "flutterwave" | "paystack";
  provider_ref?: string;
  amount: number;
  currency: string;
  status: "pending" | "processing" | "completed" | "failed" | "refunded";
  metadata: Record<string, unknown>;
  created_at: string;
  updated_at: string;
}
