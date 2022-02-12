BIN= $(CURDIR)/bin
Q= $(if $(filter 1,$V),,@)
M= $(shell printf "\033[34;1m▶\033[0m")
FUNCTION=get-contact-aws-lambda

all: build

deps: ;(info $(M) installing dependencies...) @
	@go mod tidy

build: ; $(info $(M) building executable...) @
	@env GOOS=linux go build -ldflags="-s -w" -o bin/$(FUNCTION) cmd/main.go

lint:
	@go vet ./...

format:
	@go fmt ./...

clean:
	@rm -rf $(BIN)

zip: build
	@zip -j $(FUNCTION).zip bin/$(FUNCTION)

.PHONY: build
