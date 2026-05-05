// app/server/api/demo/rooms-[id].get.ts
import { defineEventHandler, getRouterParam, createError } from "h3";

// Re-use same data — in real app this would be a shared module
const rooms: Record<string, object> = {
  "01JROOM001EXECUTIVESUITE01": {
    id: "01JROOM001EXECUTIVESUITE01",
    name: "Executive Suite",
    description:
      "A sprawling top-floor suite with panoramic city views, a king-size bed, private lounge, and complimentary minibar. Ideal for business travellers and VIP guests.",
    price: 45000,
    currency: "KES",
    status: "active",
    stock: 3,
    images: [
      "https://images.unsplash.com/photo-1631049307264-da0ec9d70304?w=800",
      "https://images.unsplash.com/photo-1578683010236-d716f9a3f461?w=800",
    ],
    metadata: {
      beds: 1,
      baths: 2,
      sqft: 850,
      floor: 18,
      view: "City Panorama",
      rating: 4.9,
      badges: ["VIP", "Sea View", "Free WiFi"],
      amenities: ["King Bed", "Jacuzzi", "Lounge", "Minibar", "Smart TV"],
    },
    created_at: "2025-01-10T08:00:00Z",
    updated_at: "2026-04-01T12:00:00Z",
  },
  "01JROOM002DELUXEDOUBLE0002": {
    id: "01JROOM002DELUXEDOUBLE0002",
    name: "Deluxe Double Room",
    description:
      "Elegantly appointed double room with garden views, premium linens, and a spacious marble bathroom. Perfect for couples.",
    price: 18500,
    currency: "KES",
    status: "active",
    stock: 8,
    images: [
      "https://images.unsplash.com/photo-1618773928121-c32242e63f39?w=800",
    ],
    metadata: {
      beds: 2,
      baths: 1,
      sqft: 420,
      floor: 7,
      view: "Garden",
      rating: 4.7,
      badges: ["Popular", "Garden View"],
      amenities: ["Double Beds", "Rain Shower", "Desk", "Smart TV", "Safe"],
    },
    created_at: "2025-01-10T08:00:00Z",
    updated_at: "2026-04-01T12:00:00Z",
  },
  "01JROOM003STANDARDSINGLE03": {
    id: "01JROOM003STANDARDSINGLE03",
    name: "Standard Single Room",
    description:
      "A cozy, well-appointed single room for the efficient business traveller.",
    price: 9500,
    currency: "KES",
    status: "active",
    stock: 15,
    images: [
      "https://images.unsplash.com/photo-1540518614846-7eded433c457?w=800",
    ],
    metadata: {
      beds: 1,
      baths: 1,
      sqft: 240,
      floor: 3,
      view: "Courtyard",
      rating: 4.5,
      badges: ["Best Value"],
      amenities: ["Single Bed", "Shower", "Work Desk", "Fast WiFi"],
    },
    created_at: "2025-01-10T08:00:00Z",
    updated_at: "2026-04-01T12:00:00Z",
  },
  "01JROOM004PRESIDENTIALSUITE4": {
    id: "01JROOM004PRESIDENTIALSUITE4",
    name: "Presidential Suite",
    description:
      "Two-bedroom suite spanning the east wing — private dining, butler service, ocean-facing terrace.",
    price: 120000,
    currency: "KES",
    status: "active",
    stock: 1,
    images: [
      "https://images.unsplash.com/photo-1582719478250-c89cae4dc85b?w=800",
    ],
    metadata: {
      beds: 2,
      baths: 3,
      sqft: 1800,
      floor: 20,
      view: "Ocean & City",
      rating: 5.0,
      badges: ["Exclusive", "Butler Service", "Ocean View"],
      amenities: [
        "2 King Beds",
        "Private Dining",
        "Butler",
        "Terrace",
        "Jacuzzi",
      ],
    },
    created_at: "2025-01-10T08:00:00Z",
    updated_at: "2026-04-01T12:00:00Z",
  },
  "01JROOM005FAMILYCONNECTED05": {
    id: "01JROOM005FAMILYCONNECTED05",
    name: "Family Connected Room",
    description:
      "Two connecting rooms — master bedroom + twin — with kitchenette and child-safe features.",
    price: 32000,
    currency: "KES",
    status: "active",
    stock: 4,
    images: [
      "https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=800",
    ],
    metadata: {
      beds: 3,
      baths: 2,
      sqft: 680,
      floor: 5,
      view: "Pool",
      rating: 4.8,
      badges: ["Family Friendly", "Pool View", "Kitchenette"],
      amenities: ["King + Twin Beds", "Kitchenette", "Pool View", "Child-Safe"],
    },
    created_at: "2025-01-10T08:00:00Z",
    updated_at: "2026-04-01T12:00:00Z",
  },
};

export default defineEventHandler(async (event) => {
  await new Promise((r) => setTimeout(r, 120));

  const id = getRouterParam(event, "id");
  const room = id ? rooms[id] : null;

  if (!room) {
    throw createError({
      statusCode: 404,
      data: {
        success: false,
        error: { code: "NOT_FOUND", message: "Room not found" },
      },
    });
  }

  return { success: true, data: room };
});
