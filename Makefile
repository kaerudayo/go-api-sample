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

logs/api:
	docker-compose logs -f api-server
