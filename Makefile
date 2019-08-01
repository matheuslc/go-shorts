build:
	docker-compose build

start:
	docker-compose up

test:
	docker-compose run --rm test

test-file:
	docker-compose run --rm test go test $@