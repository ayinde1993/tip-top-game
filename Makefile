build:
	@go build -o bin/the-tip-top-game cmd/main.go

test:
	@go test -v ./...


run: build
	@./bin/the-tip-top-game
# migration:
# 	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# migrate-up:
# 	@go run cmd/migrate/main.go up

# migrate-down:
# 	@go run cmd/migrate/main.go down
