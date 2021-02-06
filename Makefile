.PHONY: build
build:
	go build

.PHONY: install
install:
	go mod download

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -f quickbase-textmate