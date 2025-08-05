test:
	go test -cover ./internal/core/repository/postgres/shop/...
rundev:
	go run cmd/main.go -env=development
runprod:
	go run cmd/main.go -env=production
