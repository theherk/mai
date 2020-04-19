mai:
	go build

clean:
	rm -f mai
	rm -f cover.out
	find . -type d -name "*mocks" -exec rm -rf {} +
	find . -regex ".*\/\(GPATH\|GRTAGS\|GTAGS\)" -exec rm -rf {} +

cover.out: test

.PHONY: coverage
coverage: cover.out
	go tool cover -func=cover.out

.PHONY: test
test:
	go get ./...
	go get github.com/golang/mock/...
	go generate ./...
	go test --cover -coverprofile=cover.out ./...

.PHONY: image
image:
	docker build -t mai .
