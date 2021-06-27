.PHONY: lint
lint:
	golangci-lint  run -c ../.golangci.yml

.PHONY:	gofumpt
gofumpt:
	gofumpt -w .

.PHONY:	test
test:
	go test -v
