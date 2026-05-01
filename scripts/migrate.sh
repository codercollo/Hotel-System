#!/usr/bin/env bash
set -euo pipefail

# Migration runner wrapper for golang-migrate.
# Usage (inside backend container):
#   ./scripts/migrate.sh up
#   ./scripts/migrate.sh down
#   ./scripts/migrate.sh status

COMMAND="${1:-up}"
MIGRATIONS_DIR="database/migrations"
DATABASE_URL="${DATABASE_URL:?DATABASE_URL env var is required}"

# Ensure migrate binary is available
if ! command -v migrate &>/dev/null; then
  echo "→ Installing golang-migrate..."
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
fi

case "$COMMAND" in
  up)
    echo "→ Running migrations UP..."
    migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" up
    echo "✓ Migrations applied."
    ;;
  down)
    echo "→ Rolling back last migration..."
    migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" down 1
    echo "✓ Rollback complete."
    ;;
  status)
    echo "→ Migration status:"
    migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" version
    ;;
  drop)
    echo "⚠ Dropping all migrations..."
    migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" drop -f
    echo "✓ All migrations dropped."
    ;;
  *)
    echo "Usage: $0 [up|down|status|drop]"
    exit 1
    ;;
esac