import { defineStore } from "pinia";

export interface CartItem {
  itemId: string;
  name: string;
  price: number;
  quantity: number;
  checkIn: string;
  checkOut: string;
  nights: number;
  image?: string;
}

interface CartState {
  items: CartItem[];
}

export const useCartStore = defineStore("cart", {
  state: (): CartState => ({ items: [] }),

  getters: {
    count: (s): number => s.items.length,

    subtotal: (s): number =>
      s.items.reduce((acc, i) => acc + i.price * i.nights * i.quantity, 0),

    taxes(): number {
      return Math.round(this.subtotal * 0.15);
    },

    total(): number {
      return this.subtotal + this.taxes;
    },
  },

  actions: {
    add(item: CartItem) {
      const existing = this.items.find((i) => i.itemId === item.itemId);
      if (existing) {
        existing.quantity++;
      } else {
        this.items.push(item);
      }
    },
    remove(itemId: string) {
      this.items = this.items.filter((i) => i.itemId !== itemId);
    },
    clear() {
      this.items = [];
    },
  },

  persist: true,
});
