GO=go
NPM=npm
ZIP=zip

BUILD_DIR=.build
FUNCS_DIR=api/functions

COVERAGE_FILE=.coverage

FUNC_NAMES=$(patsubst ./$(FUNCS_DIR)/%/main.go, %, $(wildcard ./$(FUNCS_DIR)/*/main.go))

help:
	@echo soon™

build:
	$(NPM) run build
	@for NAME in $(FUNC_NAMES) ; do \
		GOOS=linux GOARCH=amd64 \
		$(GO) build -o "$(BUILD_DIR)/$$NAME" "./$(FUNCS_DIR)/$$NAME" ;\
		$(ZIP) -j "$(BUILD_DIR)/$$NAME.zip" "$(BUILD_DIR)/$$NAME" ;\
	done

test:
	$(NPM) run test
	$(GO) test ./... -cover -race -count=1

coverage:
	$(GO) test ./... -coverprofile=$(COVERAGE_FILE)
	$(GO) tool cover -html=$(COVERAGE_FILE)
