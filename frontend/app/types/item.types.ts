export interface Item {
  id: string;
  name: string;
  description: string;
  price: number; // smallest currency unit (cents)
  currency: string;
  status: "active" | "inactive" | "archived";
  stock: number;
  images: string[];
  metadata: Record<string, unknown>;
  created_at: string;
  updated_at: string;
}

export interface ItemMetadata {
  beds?: number;
  baths?: number;
  sqft?: number;
  rating?: number;
  badges?: string[];
  floor?: number;
  view?: string;
}

export interface CreateItemRequest {
  name: string;
  description: string;
  price: number;
  currency: string;
  stock: number;
  metadata: ItemMetadata;
}
