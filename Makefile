default: run

run:
	go run ./cmd/unicode-converter/

build-win:
	GOOS=windows GOARCH=amd64 go build ./cmd/unicode-converter

clean:
	rm -rf unicode-converter.exe