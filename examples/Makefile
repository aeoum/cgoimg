BIN=md5hash

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o $(BIN) main.go

.PHONY: clean
clean:
	rm -rf $(BIN)