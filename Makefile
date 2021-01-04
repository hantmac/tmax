.PHONY: build
build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o tmax main.go

