dev:
	scripts/hotreload.sh

run:
	go run main.go

lint:
	golangci-lint run --enable-all
