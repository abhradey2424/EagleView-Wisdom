
.PHONY: run test

run:
	@go run cmd/wisdom.go dispense

test:
	@go test ./...