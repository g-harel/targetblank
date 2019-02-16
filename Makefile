GO=go
NPM=npm

BUILD_DIR=.build
FUNCS_DIR=functions

COVERAGE_FILE=.coverage

FUNC_NAMES=$(patsubst ./$(FUNCS_DIR)/%/main.go, %, $(wildcard ./$(FUNCS_DIR)/*/main.go))

build:
	$(NPM) run build
	@for NAME in $(FUNC_NAMES) ; do \
		GOOS=linux GOARCH=amd64 \
		$(GO) build -o "$(BUILD_DIR)/$$NAME" "./$(FUNCS_DIR)/$$NAME" ;\
	done

cov-report:
	$(GO) test ./... -coverprofile=$(COVERAGE_FILE)
	$(GO) tool cover -html=$(COVERAGE_FILE)
