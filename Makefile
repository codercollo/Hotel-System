.PHONY: dev down logs ps migrate migrate-down seed generate-keys \
        backend-test frontend-test lint build clean help

# ─── Configuration ───────────────────────────────────────────────────────────
COMPOSE      := docker compose
BACKEND_SVC  := backend
FRONTEND_SVC := frontend
DB_SVC       := postgres

# ─── Primary Dev Commands ────────────────────────────────────────────────────

## Start the full development stack
dev:
	@cp -n .env.example .env 2>/dev/null || true
	$(COMPOSE) up --build

## Start stack in background
dev-detached:
	@cp -n .env.example .env 2>/dev/null || true
	$(COMPOSE) up --build -d

## Stop all services
down:
	$(COMPOSE) down

## Stop and remove volumes (full reset)
reset:
	$(COMPOSE) down -v --remove-orphans

## View logs (optionally: make logs svc=backend)
logs:
	$(COMPOSE) logs -f $(svc)

## Show running containers
ps:
	$(COMPOSE) ps

# ─── Database ────────────────────────────────────────────────────────────────

## Run all pending migrations
migrate:
	$(COMPOSE) exec $(BACKEND_SVC) sh -c "cd /app && ./scripts/migrate.sh up"

## Rollback the last migration
migrate-down:
	$(COMPOSE) exec $(BACKEND_SVC) sh -c "cd /app && ./scripts/migrate.sh down"

## Show migration status
migrate-status:
	$(COMPOSE) exec $(BACKEND_SVC) sh -c "cd /app && ./scripts/migrate.sh status"

## Seed the database with initial data
seed:
	$(COMPOSE) exec $(BACKEND_SVC) sh -c "cd /app && ./scripts/seed.sh"

# ─── Code Generation ─────────────────────────────────────────────────────────

## Generate RSA keys for JWT
generate-keys:
	@bash scripts/generate-keys.sh

# ─── Testing ─────────────────────────────────────────────────────────────────

## Run backend tests
backend-test:
	$(COMPOSE) exec $(BACKEND_SVC) go test ./... -v -count=1

## Run backend tests with coverage
backend-coverage:
	$(COMPOSE) exec $(BACKEND_SVC) go test ./... -coverprofile=coverage.out
	$(COMPOSE) exec $(BACKEND_SVC) go tool cover -html=coverage.out

## Run frontend unit tests
frontend-test:
	$(COMPOSE) exec $(FRONTEND_SVC) pnpm test

# ─── Quality ─────────────────────────────────────────────────────────────────

## Run backend linter
lint-backend:
	$(COMPOSE) exec $(BACKEND_SVC) golangci-lint run ./...

## Run frontend type-check
typecheck-frontend:
	$(COMPOSE) exec $(FRONTEND_SVC) pnpm typecheck

# ─── Utilities ───────────────────────────────────────────────────────────────

## Open a psql shell
psql:
	$(COMPOSE) exec $(DB_SVC) psql -U platform platform_dev

## Open a bash shell in the backend container
backend-shell:
	$(COMPOSE) exec $(BACKEND_SVC) sh

## Open a bash shell in the frontend container
frontend-shell:
	$(COMPOSE) exec $(FRONTEND_SVC) sh

## Show this help
help:
	@echo ""
	@echo "  Platform — available make commands"
	@echo ""
	@grep -E '^## ' Makefile | sed 's/## /    /'
	@echo ""