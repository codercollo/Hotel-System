export interface Role {
  id: string;
  name: string;
  description: string;
  permissions: Permission[];
  created_at: string;
}

export interface Permission {
  id: string;
  name: string;
  description: string;
}
