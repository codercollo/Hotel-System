include .env
export
.PHONY: run build test test-integration lint migrate migrate-down tidy clean

MODULE := github.com/codercollo/hotel-system/backend

run:
	go run ./cmd/api

build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/api ./cmd/api
	CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/worker ./cmd/worker

test:
	go test ./... -v -count=1 -race

test-integration:
	go test ./tests/integration/... -v -tags integration

lint:
	golangci-lint run ./...

migrate-up:
	migrate -path database/migrations \
		-database "$(DATABASE_URL)" up

migrate-down:
	migrate -path database/migrations \
		-database "$(DATABASE_URL)" down


tidy:
	go mod tidy

clean:
	rm -rf bin/ tmp/ coverage.out