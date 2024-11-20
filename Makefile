export GONOPROXY=github.com/ewollesen/optional-private-module
export GOPRIVATE=github.com/ewollesen/optional-private-module

both: build build-private

build:
	go build -o ./without-private-code .

build-private:
	go build -o ./with-private-code -tags private .
