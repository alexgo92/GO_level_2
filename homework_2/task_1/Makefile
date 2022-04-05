## сборка приложения под windows 386
.PHONY: build
build:
	GOOS=windows GOARCH=386 go build main.go

## запускает приложение
.PHONY: run
run: 
	go run ./...

## проверка линтерами
.PHONY: lint
lint:
	golangci-lint run ./...