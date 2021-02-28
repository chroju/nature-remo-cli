build:
	go build -o $(HOME)/go/bin/remo

GOCMD=go
GOBUILD=gox
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=remo

crossbuild:
	$(GOBUILD) -os="linux darwin windows" -arch="386 amd64" -output "bin/remo_{{.OS}}_{{.Arch}}/{{.Dir}}"

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

lint:
	go mod tidy
	gofmt -s -l .
	golint ./...
	go vet ./...

mod:
	go mod download

test: lint
	go test -v ./...

test-coverage: mod
	go test -race -covermode atomic -coverprofile=covprofile ./...
