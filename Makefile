include .env

up:
	docker-compose up -d mysql
	sleep 8
	docker-compose up -d migrate

up-db:
	docker-compose up -d mysql

up-migrate:
	docker-compose up -d migrate

rm-db:
	sudo rm -rf .var