#!/usr/bin/env bash
set -euo pipefail

# Seed script — runs the Go seed command which inserts default roles,
# permissions, and an admin user into the database.
# Usage (inside backend container): ./scripts/seed.sh

echo "→ Seeding database..."
go run ./cmd/seed/main.go
echo "✓ Seed complete."