test:
	go test -cover ./internal/core/repository/postgres/shop/...
rundev:
	go run cmd/main.go -env=development   
runprod:
	sudo docker compose -f docker-compose.prod.yaml --env-file .env.production up --build
