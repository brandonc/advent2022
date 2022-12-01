build:
	go generate internal/commands/init.go
	go build -o bin/advent2022 cmd/main.go

.PHONY: build
