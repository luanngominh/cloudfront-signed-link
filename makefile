.PHONY: build dev

build:
	go build -o kick-start main.go

dev:
	go run main.go
