ENV_FILE=.env
include $(ENV_FILE)
export

run:
	go run ./cmd/auth --config=./config/local.yaml

build:
	go build -o auth-service ./cmd/auth

migrate-up:
	docker run --rm \
		--network tradesim_backend \
		-v "$(PWD)/migrations:/migrations" \
		migrate/migrate \
		-path /migrations \
		-database "$(DATABASE_URL)" up

migrate-down:
	docker run --rm \
		--network tradesim_backend \
		-v "$(PWD)/migrations:/migrations" \
		migrate/migrate \
		-path /migrations \
		-database $(DATABASE_URL) down 1