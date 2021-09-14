GOCMD       = go
GOBUILD     = $(GOCMD) build
GOVENDOR     = $(GOCMD) mod vendor
GOCLEAN     = $(GOCMD) clean
GOTEST      = $(GOCMD) test
GOGET       = $(GOCMD) get
GOINSTALL   = $(GOCMD) install
GOINSTALL   = $(GOCMD) install
SQLITE3CMD  = sqlite3

BINARY_NAME=bin/

all: test build
db: build_dir rm_db init_db

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	FLIBOOBSTIER_DEBUG=false $(GOTEST) -count=50 ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

install:
	$(GOINSTALL) -v ./...

deps:
	$(GOVENDOR)

rm_db:
	rm -f ./bin/fliboobstier.db

init_db:
	$(SQLITE3CMD) bin/fliboobstier.db < storage/sqlite_schema.sql

build_dir:
	mkdir -p bin