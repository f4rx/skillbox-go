.PHONY: lint
lint:
	golangci-lint  run -c ../.golangci.yml

.PHONY:	gofumpt
gofumpt:
	gofumpt -w .

.PHONY:	test
test:
	go test -v

.PHONY: run1
run1:
	DEBUG=y go run ./main.go file1.txt file2.txt
