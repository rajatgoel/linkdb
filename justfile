setup:
        brew install go bufbuild/buf/buf golangci-lint sqlc

gen:
        buf generate proto

lint: gen
        buf lint proto
        buf breaking proto --against '.git#branch=main,subdir=proto' || true
        golangci-lint run ./...

test: lint
        go vet ./...
        go test ./...
