VERSION ?= 0.1.0
PORT ?= 8040

default: run

run:
	GO111MODULE=auto go build
	APP_VERSION=0.1.0 \
	LOAD_BALANCER_URL=http://localhost:80/ \
	PORT=8040 \
		./golangapp

package:
	rm -f golangapp
	GO111MODULE=auto GOOS=linux GOARCH=arm64 go build
	chmod +x golangapp