test-common:
	go test -v ./common/...

test-core:
	go test -v ./core/...

build-core:
	go build -o bin/core ./core/

run-core:
	ENV=dev go run core/main.go

docker-build-core:
	docker build -f .docker/DockerfileCore  -t core-api .

docker-run-core:
	docker rm -f core-api && docker run -d -p 8080:8080 --env-file .env/.env.core.docker --name core-api core-api

docker-run-mongo:
	docker rm -f mongo && docker run -d -p 27017:27017 -e MONGO_INITDB_ROOT_PASSWORD=secret -e MONGO_INITDB_ROOT_USERNAME=mongoadmin --name mongo mongo

docker-run-redis:
	docker rm -f redis && docker run -d -p 6379:6379 --name redis redis
