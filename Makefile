.PHONY: build
build_m1:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=arm64 go build -ldflags '-extldflags "-static"' -o tmax main.go
build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o tmax main.go
.PHONY: lint
lint:
	golangci-lint run
