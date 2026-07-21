APP_NAME=kakapo

run:
	go run ./cmd/kakapo

dev:
	air

build:
	go build -o bin/$(APP_NAME) ./cmd/kakapo

test:
	go test ./...

test-verbose:
	go test -v ./...

coverage:
	go test -cover ./...

lint:
	golangci-lint run

fmt:
	gofumpt -w .

check:
	make fmt
	make lint
	make test

tidy:
	go mod tidy

docker:
	docker compose up -d

docker-down:
	docker compose down

migrate:
	go run ./cmd/migrate

clean:
	rm -rf ./bin