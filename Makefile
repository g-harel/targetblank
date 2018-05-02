GO=go
NPM=npm
ZIP=zip

BUILD_DIR=.build
FUNC_DIR=./functions

COVERAGE_FILE=.coverage

build:
	$(NPM) run build
	@for FILE in $(wildcard $(FUNC_DIR)/*/main.go) ; do \
		NAME=$$(echo $$FILE | grep -oP "(?<=functions\/).+(?=/main.go)") ;\
		GOOS=linux $(GO) build -o "$(BUILD_DIR)/$$NAME" "$$FILE" ;\
		$(ZIP) -j "$(BUILD_DIR)/$$NAME.zip" "$(BUILD_DIR)/$$NAME" ;\
	done

test:
	$(GO) test ./... -cover

coverage:
	$(GO) test ./... -coverprofile=$(COVERAGE_FILE)
	$(GO) tool cover -html=$(COVERAGE_FILE)
