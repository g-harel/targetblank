GO=go

COV_FILE=coverage.out

build:
	@echo soon™

test:
	$(GO) test ./... -cover

coverage:
	$(GO) test ./... -coverprofile=$(COV_FILE)
	$(GO) tool cover -html=$(COV_FILE)

clean:
	rm -f $(BIN_FILE) $(COV_FILE) $(ZIP_FILE)