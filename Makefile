APP_NAME=bms
CMD_DIR=cmd/bms
BIN_DIR=bin


build:
	go build -o $(BIN_DIR)/$(APP_NAME) ./$(CMD_DIR)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME) ./$(CMD_DIR)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME).exe ./$(CMD_DIR)

run:
	go run ./$(CMD_DIR)

clean:
	rm -rf $(BIN_DIR)/*



.PHONY: test fmt vet

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...


.PHONY: help

help:
	@echo "Available commands:"
	@echo "  make build     - build binary"
	@echo "  make run       - run app"
	@echo "  make test      - run tests"
	@echo "  make fmt       - format code"
	@echo "  make clean     - remove binaries"