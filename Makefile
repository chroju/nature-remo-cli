build:
	go build -o $(HOME)/go/bin/remo

GOCMD=go
GOBUILD=gox
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=remo

build:
	$(GOBUILD) -os="linux darwin windows" -arch="386 amd64" -output "bin/remo_{{.OS}}_{{.Arch}}/{{.Dir}}"
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
