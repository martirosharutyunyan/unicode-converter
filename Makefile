default: run

run:
	go run ./cmd/unicode-converter-gui/

build-win:
	cd ./cmd/unicode-converter-gui && make build && cd ../../

clean:
	rm -rf unicode-converter.exe

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

tidy:
	go mod tidy