.PHONY: build
build: export PRIVATE_MODULE_PATHS=github.com/ewollesen/building-with-private-modules/caution/placeholder
build:
	go generate ./...
	go mod tidy
	go build -o ./thinger .

.PHONY: both
both: build build-private

.PHONY: generate
generate:
	go generate ./...

.PHONY: build-private
build-private: export GONOPROXY=github.com/ewollesen/optional-private-module
build-private: export GOPRIVATE=github.com/ewollesen/optional-private-module
build-private: export PRIVATE_MODULE_PATHS=github.com/ewollesen/optional-private-module/secret
build-private: build

.PHONY: test
test:
	go test ./...
