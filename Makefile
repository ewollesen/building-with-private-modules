.PHONY: build
build: export PRIVATE_MODULE_PATHS=github.com/ewollesen/building-with-private-modules/caution/placeholder
build:
	go generate ./...
	go build -o ./thinger .

.PHONY: both
both: build build-private

.PHONY: generate
generate:
	go generate ./...

.PHONY: build-private
build-private: export GONOPROXY=github.com/ewollesen/optional-private-module
build-private: export GOPRIVATE=github.com/ewollesen/optional-private-module
build-private: export PRIVATE_MODULE_PATH=github.com/ewollesen/optional-private-module/secret
build-private: generate
	go get $(PRIVATE_MODULE_PATH)@latest
	go build -o ./thinger
	go get $(PRIVATE_MODULE_PATH)@none

.PHONY: test
test:
	go test ./...
