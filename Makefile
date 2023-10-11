default: run

run:
	QT_QPA_PLATFORM=xcb go run ./cmd/unicode-converter-gui/

cli-build:
	go build -o unicode-converter ./cmd/unicode-converter-cli

cli-build-win:
	GOOS=windows go build -o unicode-converter.exe ./cmd/unicode-converter-cli

clean:
	rm -rf unicode-converter.exe

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

tidy:
	go mod tidy

qt-ui:
	goqtuic -ui-file ./internal/desktop/view/unicode_converter.ui -go-ui-dir ./internal/desktop/uigen

gen:
	go generate ./...