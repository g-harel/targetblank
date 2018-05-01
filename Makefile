AWS=aws
GO=go
ZIP=zip

BIN_FILE=main.bin
COV_FILE=coverage.out
ZIP_FILE=deployment.zip

build:
	@echo soon™

test:
	$(GO) test ./... -cover

coverage:
	$(GO) test ./... -coverprofile=$(COV_FILE)
	$(GO) tool cover -html=$(COV_FILE)

clean:
	rm -f $(BIN_FILE) $(COV_FILE) $(ZIP_FILE)