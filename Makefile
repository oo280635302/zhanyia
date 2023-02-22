export GOOS=linux
export GOARCH=amd64

all:
	make build
	make image
	make compose

build:
	mkdir -p bin
	go build  -o ./bin/demo ./src/main.go

image:
	docker build -t zhanyi:1.0.0 .

compose:
	docker-compose up --remove-orphans --force-recreate -d demo

