build:
	go build -o ./without-private-code .

both: build build-private

build-private: export GONOPROXY=github.com/ewollesen/optional-private-module
build-private: export GOPRIVATE=github.com/ewollesen/optional-private-module
build-private:
	go build -o ./with-private-code -tags private .

test:
	go test ./...

test-private: export GONOPROXY=github.com/ewollesen/optional-private-module
test-private: export GOPRIVATE=github.com/ewollesen/optional-private-module
test-private:
	go test -tags private ./...

vendor: export GONOPROXY=github.com/ewollesen/optional-private-module
vendor: export GOPRIVATE=github.com/ewollesen/optional-private-module
vendor:
	go mod vendor

test-both: test test-private
