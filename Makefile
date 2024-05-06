.PHONY: test
test:
	@sh scripts/test-coverage.sh

install:
	GOBIN=$(GOBIN) go install -mod=readonly -v ./cmd/...

mod:
	GOBIN=$(GOBIN) go mod vendor
	GOBIN=$(GOBIN) go mod tidy

run:
	GOBIN=$(GOBIN) go run cmd/server/main.go

.PHONY: generate-mocks
generate-mocks:
	@sh scripts/install-mockery.sh
	@sh scripts/generate-mocks.sh
