AWS=aws
GO=go
ZIP=zip

BIN_FILE=main.bin
COV_FILE=coverage.out
ZIP_FILE=deployment.zip

lambda:
	GOOS=linux $(GO) build -o $(BIN_FILE)
	$(ZIP) $(ZIP_FILE) $(BIN_FILE)
	$(AWS) lambda create-function \
		--region us-east-1 \
		--function-name helloworld \
		--zip-file fileb://./$(ZIP_FILE) \
		--runtime go1.x \
		--role TODO
		--handler $(BIN_FILE)

test:
	$(GO) test ./... -cover

coverage:
	$(GO) test ./... -coverprofile=$(COV_FILE)
	$(GO) tool cover -html=$(COV_FILE)

clean:
	rm -f $(BIN_FILE) $(COV_FILE) $(ZIP_FILE)