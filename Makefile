NAME=charlestown
EXE=./dist/charlestown/bin/$(NAME)
VENDOR_MANIFEST=./vendor/modules.txt
VENDOR_FLAGS=-v
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
	$(RM) $(dir $(EXE))
	$(RM) $(dir $(VENDOR_MANIFEST))
	$(RM) $(dir $(COVERPROFILE))

$(VENDOR_MANIFEST): go.mod go.sum
	echo "[charlestown] building vendor dir"
	$(GO) mod tidy
	$(GO) mod vendor $(VENDOR_FLAGS)

lint: $(SRC)
	echo "[charlestown] linting"
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
	echo "[charlestown] building"
	$(GO) build -o $(EXE) $(BUILD_FLAGS)
