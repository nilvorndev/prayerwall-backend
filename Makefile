.PHONY: fmt test vet build run compose-up compose-down ci

fmt:
	gofmt -w $$(find cmd internal -name '*.go' -type f)

test:
	go test ./...

vet:
	go vet ./...

build:
	go build -trimpath -o bin/prayerwall-api ./cmd/api

run:
	go run ./cmd/api

compose-up:
	docker compose -f deploy/compose.yaml up --build

compose-down:
	docker compose -f deploy/compose.yaml down

ci: fmt vet test build
