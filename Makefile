.PHONY: setup setup/*
setup: .env

.env: .env_default
	cp $< $@

.PHONY: up up/detach stop
up: setup
	docker-compose up --build
up/detach: setup
	docker-compose up --build -d
stop:
	docker-compose stop

.PHONY: db/*
db/reset: db/drop db/create db/migrate db/seed

db/drop db/create db/migrate db/seed:
	docker exec -it api-server /bin/sh -c "go run scripts/db/db.go ${@F}"

tidy:
	docker exec -it api-server /bin/sh -c "go mod tidy"

mockgen:
	docker exec -it api-server /bin/sh -c "mockgen -source=./domain/repository/query/${SOURCE} -destination=./infra/reader/mock/${SOURCE}"

.PHONY: lint test
lint:
	docker run -it --rm -v $(PWD)/app:/app -w /app golangci/golangci-lint:v1.52.0 golangci-lint run
lint-fix:
	docker run -it --rm -v $(PWD)/app:/app -w /app golangci/golangci-lint:v1.52.0 golangci-lint run --fix

test-parallel: # make test-parallel OPTION="./package/... -run TestFuncName"
	docker exec -i api-server /bin/sh -c "GO_ENV=test go test -v -tags=parallel ${OPTION}"

test-serial: # make test-parallel OPTION="./package/... -run TestFuncName"
	docker exec -i api-server /bin/sh -c "GO_ENV=test go test -v -tags=serial ${OPTION}"

test-api: # make test-api OPTION="./controllers/... -run TestFuncName"
	docker exec -i api-server /bin/sh -c "GO_ENV=test go test -v -tags=api ${OPTION}"

.PHONY: shell/*
shell/api:
	docker exec -it api-server /bin/bash

shell/redis:
	docker exec -it api-redis /bin/bash -c "redis-cli"

shell/mysql:
	docker exec -it api-db /bin/sh -c "mysql -uroot -proot"

logs/api:
	docker-compose logs -f api-server
