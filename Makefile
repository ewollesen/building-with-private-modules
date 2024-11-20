build:
	go build -o ./without-private-code .

both: build build-private

build-private: export GONOPROXY=github.com/ewollesen/optional-private-module
build-private: export GOPRIVATE=github.com/ewollesen/optional-private-module
build-private:
	go build -o ./with-private-code -tags private .
