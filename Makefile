GO=go
COV_FILE=coverage.out

test:
	$(GO) test ./... -cover

coverage:
	$(GO) test ./... -coverprofile=$(COV_FILE)
	$(GO) tool cover -html=$(COV_FILE)
