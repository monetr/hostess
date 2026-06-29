.PHONY: build test vet fmt fmt-check lint tidy install clean

# NOTE There is a hostess/ library subdir, so `go build -o hostess .` writes the
# binary INTO that directory instead of producing a ./hostess file. Build into
# bin/ to dodge that collision. The release workflow does not have this problem
# because it builds uniquely named hostess-<tag>-<os>-<arch> artifacts.
build:
	go build ./...
	go build -o bin/hostess .

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

fmt-check:
	@out="$$(gofmt -l .)"; if [ -n "$$out" ]; then echo "gofmt needed:"; echo "$$out"; exit 1; fi

lint: vet fmt-check

tidy:
	go mod tidy

install:
	go build -o bin/hostess .
	sudo mv bin/hostess /usr/local/bin/hostess

clean:
	rm -rf bin dist
	rm -f hostess-v*
