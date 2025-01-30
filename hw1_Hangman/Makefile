COVERAGE_FILE ?= coverage.out
FILTERED_COVERAGE_FILE ?= coverage_filtered.out
TARGET ?= run

.PHONY: build
build:
	@echo "Выполняется go build для таргета ${TARGET}"
	@mkdir -p bin
	@go build -o ./bin/${TARGET} ./cmd/${TARGET}

.PHONY: test
test:
	@rm -f $(COVERAGE_FILE)_*.out $(FILTERED_COVERAGE_FILE)
	@echo "mode: atomic" > $(FILTERED_COVERAGE_FILE)
	@for pkg in $(shell go list ./... | grep -Ev '/mocks/' | grep -v '/cmd/run'); do \
		go test -coverpkg=./... -coverprofile=$(COVERAGE_FILE)_$$(echo $$pkg | tr "/" "_").out $$pkg || exit 1; \
		grep -h -v "^mode:" $(COVERAGE_FILE)_$$(echo $$pkg | tr "/" "_").out >> $(FILTERED_COVERAGE_FILE); \
	done
	@grep -Ev 'github.com/es-debug/backend-academy-2024-go-template/internal/application/mocks/|github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/io/mocks/|github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/visualization/mocks/' $(FILTERED_COVERAGE_FILE) > $(COVERAGE_FILE)
	@rm $(COVERAGE_FILE)_*.out
	@go tool cover -func=$(COVERAGE_FILE) | grep ^total | tr -s '\t' ' '
	@rm $(COVERAGE_FILE)

.PHONY: lint
lint:
	@goimports -w -local github.com/es-debug/backend-academy-2024-go-template .
	@golangci-lint run ./...

.PHONY: mocks
mocks:
	@mockery --name=WordRepository --dir=./internal/application --output=./internal/application/mocks --outpkg=mocks --with-expecter
	@mockery --name=InputReader --dir=./internal/infrastructure/io --output=./internal/infrastructure/io/mocks --outpkg=mocks --with-expecter
	@mockery --name=OutputWriter --dir=./internal/infrastructure/io --output=./internal/infrastructure/io/mocks --outpkg=mocks --with-expecter
	@mockery --name=HangmanVisualizer --dir=./internal/infrastructure/visualization --output=./internal/infrastructure/visualization/mocks --outpkg=mocks --with-expecter
	@mockery --config .mockery.yaml --name=WordServiceInterface --dir=./internal/application --output=./internal/application/mocks --outpkg=mocks --with-expecter

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: all
all: build test lint mocks
