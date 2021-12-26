build:
	go build -o bin/main main.go

run:
	go run main.go

docker-build:
	docker build --tag go-example-app .

docker-run:
	docker run -d -p 8080:8080 go-example-app

docker-pg:
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres

docker-redis:
	docker run -d -p 6379:6379 redis

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

