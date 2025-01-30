COVERAGE_FILE ?= coverage.out

TARGET ?= run # CHANGE THIS TO YOUR BINARY NAME/NAMES

.PHONY: build
build:
	@echo "Выполняется go build для таргета ${TARGET}"
	@mkdir -p .bin
	@go build -o ./bin/${TARGET} ./cmd/${TARGET}

## test: run all tests
.PHONY: test
test:
	@go test -coverpkg='github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/...' --race -count=1 -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'
.PHONY: lint
lint:
	@goimports -w -local github.com/es-debug/backend-academy-2024-go-template .
	@golangci-lint run ./...
