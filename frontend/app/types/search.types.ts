export interface SearchResult {
  id: string;
  type: string;
  name: string;
  description: string;
  price?: number;
  metadata?: Record<string, unknown>;
  score: number;
  created_at: string;
}
