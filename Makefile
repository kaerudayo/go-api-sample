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

.PHONY: shell/*
shell/api:
	docker exec -it api-server /bin/bash

shell/redis:
	docker exec -it api-redis /bin/bash -c "redis-cli"

shell/mysql:
	docker exec -it api-db /bin/sh -c "mysql -uroot -proot"

logs/api:
	docker-compose logs -f api-server
