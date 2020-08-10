DOCKER_TAG=gowpcli:test

build:
	docker build --tag=${DOCKER_TAG} .

run:
	docker run --rm -it ${DOCKER_TAG}

cli: build run


dc-down:
	docker-compose down -v

dc-up:
	docker-compose up --build

dc-restart: dc-down dc-up
