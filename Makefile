build: generate
	go build -o bin/advent2022 cmd/main.go

generate:
	go generate internal/commands/init.go

.PHONY: build

