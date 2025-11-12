.PHONY: *

run:
	go run ./...

test:
	go test ./...

image:
	podman build . -t smtn:latest
