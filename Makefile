start-app:
	reflex -s -r '\.go$$' -- godotenv -f .env go run oxmud.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -r '\.qtpl$$' -- qtc -dir=view

oxmud:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o oxmud oxmud.go

db-migrate:
	migrate -path ./migrations -database "postgres://localhost:5432/oxmud?sslmode=disable" up

db-schema-dump:
	pg_dump --schema-only -O oxmud > database/schema.sql

sqlc-gen:
	sqlc --experimental generate

.PHONY: start-app start-view db-migrate db-schema-dump sqlc-gen
