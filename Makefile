NAME=charlestown
EXE=./dist/charlestown/bin/$(NAME)
VENDOR_MANIFEST=./vendor/modules.txt
SRC=$(shell git ls-files '*.go')
GO=go
RM=rm -rf
LDFLAGS="-s -w"
BUILD_FLAGS=-x -v -ldflags $(LDFLAGS)
CYCLO_FLAGS=-over 15
COVERPROFILE=./dist/tests/cover.out
TEST_FLAGS=-v -count=1 -cover -covermode=set -benchmem -coverprofile=$(COVERPROFILE)

.PHONY: all test clean

all: $(EXE)

clean:
	$(RM) $(basename $(EXE))
	$(RM) ./vendor
	$(RM) $(basename $(COVERPROFILE))

go.sum: go.mod
	$(GO) mod tidy
	$(GO) mod vendor

$(VENDOR_MANIFEST): go.mod go.sum
	$(GO) mod vendor

lint: $(SRC)
	$(GO) get -v -u golang.org/x/lint/golint
	$(GO) get -v -u github.com/fzipp/gocyclo
	$(GO) get -v -u github.com/client9/misspell/cmd/misspell
	misspell -error $(SRC)
	gocyclo $(CYCLO_FLAGS) $(SRC)
	$(GO) list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status
	$(GO) mod tidy

test: $(SRC)
	mkdir -p $(basename $(COVERPROFILE))
	touch $(COVERPROFILE)
	$(GO) vet
	$(GO) test $(TEST_FLAGS) ./...

$(EXE): $(SRC) $(VENDOR_MANIFEST)
	$(GO) build -o $(EXE) $(BUILD_FLAGS)
