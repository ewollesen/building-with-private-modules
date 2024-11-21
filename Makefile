PRIVATE_MODULE_PATH = github.com/ewollesen/building-with-private-modules
PRIVATE_MODULE_PKG_PATH = github.com/ewollesen/building-with-private-modules/placeholder

.PHONY: build
build: generate
	go build -o ./thinger .
	./thinger

.PHONY: both
both: build build-private

.PHONY: generate
generate:
	go generate ./...

.PHONY: build-private
build-private: PRIVATE_MODULE_PATH=github.com/ewollesen/optional-private-module
build-private: export PRIVATE_MODULE_PKG_PATH=github.com/ewollesen/optional-private-module/secret
build-private: export GONOPROXY=$(PRIVATE_MODULE_PATH)
build-private: export GOPRIVATE=$(PRIVATE_MODULE_PATH)
build-private: build
	go get $(PRIVATE_MODULE_PATH)@none

.PHONY: test
test:
	go test ./...
