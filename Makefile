SHELL := /bin/ash

build:
	docker build --tag casinowar .

run:
	docker run -d -p 8081:8081 --name casinowar casinowar:latest

test:
	go test ./... -v




