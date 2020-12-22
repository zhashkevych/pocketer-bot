.PHONY:

build-bot:
	go build -o ./.bin/bot cmd/bot/main.go

start-bot: build-bot
	./.bin/bot

build-server:
	go build -o ./.bin/server cmd/server/main.go

start-server: build-server
	./.bin/server