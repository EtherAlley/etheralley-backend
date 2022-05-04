build:
	go build -o bin/main main.go

test:
	go test -v ./...

run:
	go run main.go

docker-build:
	docker build --tag core-api .

docker-run:
	docker run -d -p 8080:8080 --env-file .env.docker --name core-api core-api

docker-mongo:
	docker run -d -p 27017:27017 -e MONGO_INITDB_ROOT_PASSWORD=secret -e MONGO_INITDB_ROOT_USERNAME=mongoadmin --name mongo mongo

docker-redis:
	docker run -d -p 6379:6379 --name redis redis

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

