.PHONY: build-main build-live run-main run-live watch-main watch-live install-tools

build-main:
	go build -o bin/main-server ./cmd/main-server/

build-live:
	go build -o bin/live-server ./cmd/live-server/

run-main: build-main
	./bin/main-server

run-live: build-live
	./bin/live-server

install-tools:
	go install github.com/cespare/reflex@latest

watch-main:
	reflex -r '\.go$$' -s -- sh -c 'make build-main && ./bin/main-server'

watch-live:
	reflex -r '\.go$$' -s -- sh -c './bin/live-server'