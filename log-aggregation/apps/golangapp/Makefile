VERSION ?= 0.0.1
PORT ?= 8040

default: run

run:
	GO111MODULE=auto go build
	APP_VERSION=0.0.1 \
	LOAD_BALANCER_URL=http://localhost:80/ \
	PORT=8040 \
		./golangapp