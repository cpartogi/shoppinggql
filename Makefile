PROJECT_NAME=checkoutpromo

run:
	go run main.go

generate:
	go run github.com/99designs/gqlgen generate 

compose-up:
	docker-compose up -d --build

compose-down:
	docker-compose down