.PHONY: build run up all kill

build:
	docker build -t tick -f Dockerfile-tick .
	docker build -t prometheus -f Dockerfile-prometheus .
	docker build -t grafana -f Dockerfile-grafana .

up:
	docker-compose up

all:
	$(MAKE) build
	$(MAKE) up

kill:
	docker-compose down
	