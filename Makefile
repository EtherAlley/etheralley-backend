test-common:
	go test -v ./common/...

test-profiles-api:
	go test -v ./profiles-api/...

test-daos-api:
	go test -v ./daos-api/...

build-profiles-api:
	go build -o bin/profiles-api ./profiles-api/

build-daos-api:
	go build -o bin/daos-api ./daos-api/

run-profiles-api:
	ENV=dev go run profiles-api/main.go

run-daos-api:
	ENV=dev go run daos-api/main.go

docker-build-profiles-api:
	docker build -f .docker/DockerfileProfilesApi  -t profiles-api .

docker-build-daos-api:
	docker build -f .docker/DockerfileDaosApi  -t daos-api .

docker-run-profiles-api:
	docker rm -f profiles-api && docker run -d -p 8080:8080 --env-file .env/.env.profiles-api.docker --name profiles-api profiles-api

docker-run-daos-api:
	docker rm -f daos-api && docker run -d -p 8081:8081 --env-file .env/.env.daos-api.docker --name daos-api daos-api

docker-run-mongo:
	docker rm -f mongo && docker run -d -p 27017:27017 -e MONGO_INITDB_ROOT_PASSWORD=secret -e MONGO_INITDB_ROOT_USERNAME=mongoadmin --name mongo mongo

docker-run-redis:
	docker rm -f redis && docker run -d -p 6379:6379 --name redis redis
