default: run

run:
	go run ./cmd/unicode-converter/

build-win:
	GOOS=windows GOARCH=amd64 go build ./cmd/unicode-converter

clean:
	rm -rf unicode-converter.exe

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

tidy:
	go mod tidy