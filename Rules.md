Yes — and you’re thinking about this the right way.

What you need is a **short, enforceable “architecture prompt”** that every developer (including future you) reads **before writing any backend code** during these phases.

This acts like a **constitution** for the project.

Below is a **drop-in `ARCHITECTURE_RULES.md`** you can place at the root of the backend repo and reference during every phase.

---

# `ARCHITECTURE_RULES.md` — Domain-Agnostic Backend Constitution

## Prime Directive

> This backend is a **platform core**, not a business application.

It must run **unchanged** for:

- hotel system
- e-commerce store
- booking platform
- rentals system
- service marketplace

If code must change for a new domain, the rule has been broken.

---

## 1. Forbidden Words in Core

These words must **never** appear in `internal/`, `pkg/`, `database/`, `events/`:

```
product, room, booking, cart, shop, hotel, customer, reservation
```

Allowed vocabulary:

```
item, order, user, payment, role, notification
```

HTTP layer and frontend may use business words. Core may not.

---

## 2. Generic Domain Model Rule

Core structs must be generic and strongly typed.

```go
type Item struct {
    ID          string
    Name        string
    Description string
    Price       Money
    Metadata    map[string]any
}
```

Business-specific fields go into `Metadata`.

No schema changes for business flavor.

---

## 3. Service Action Rule

Service methods describe **actions**, not stories.

Allowed:

```
CreateItem
PlaceOrder
UpdateOrderStatus
```

Forbidden:

```
BookRoom
AddToCart
ShipProduct
```

---

## 4. HTTP Is the Translator Layer

This is where domain language is allowed.

```
POST /rooms  -> itemService.CreateItem
POST /book   -> orderService.PlaceOrder
```

Handlers translate business → generic.

---

## 5. Modules Must Not Depend on Providers

Modules depend only on interfaces:

```go
type PaymentProvider interface {}
type Storage interface {}
type EventBus interface {}
```

Implementations live outside modules.

---

## 6. Events Are Semantic and Generic

Allowed events:

```
ItemCreated
OrderPlaced
PaymentCompleted
UserRegistered
```

Never business terms.

---

## 7. Database Migration Rule

Migrations must never contain business language.

Forbidden:

```
add_room_number_to_items
```

Use metadata instead.

---

## 8. Tests Must Be Business-Neutral

Tests say:

```
create item
place order
complete payment
```

Not:

```
book room
buy product
```

---

## 9. Frontend Owns Business Vocabulary

Frontend is free to rename:

| Backend | Hotel UI    | Store UI |
| ------- | ----------- | -------- |
| Item    | Room        | Product  |
| Order   | Reservation | Order    |

Backend remains untouched.

---

## 10. Litmus Test Before Commit

Ask:

> “Would this still make sense if this project was a hotel instead of a store?”

If no → refactor.

---

## How This Applies to Your Phases

During your implementation plan:

- Phase 3 (`items`) is **not products or rooms** — it is the canonical pattern
- Phase 4 (`orders`) is **not bookings** — it is orchestration
- Phase 5 (`payments`) is generic money movement
- Every later phase builds on this neutrality

---

## The Outcome

Following this:

- One backend for unlimited projects
- Only frontend, metadata, and seed data change
- You are building a **platform**, not a project backend

---
