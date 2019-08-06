NAME=charlestown
DIST=./dist
BIN_DIST=$(DIST)/charlestown/bin
EXE=$(BIN_DIST)/$(NAME)
VENDOR_MANIFEST=./vendor/modules.txt
VENDOR_FLAGS=-v
SRC=$(shell git ls-files | grep -e '\.go')
GO=go
RM=rm -rf
LDFLAGS="-s -w -X github.com/dadleyy/charlestown/engine/constants.AppVersion=$(shell ./auto/git-version.sh)"
BUILD_FLAGS=-x -v -ldflags $(LDFLAGS)
CYCLO_FLAGS=-over 25
COVERPROFILE=./dist/tests/cover.out
TEST_FLAGS=-v -count=1 -cover -covermode=set -benchmem -coverprofile=$(COVERPROFILE)

OSX_DIST=$(DIST)/charlestown/osx
OSX_BUNDLE_CONTENTS=$(OSX_DIST)/charlestown.app/Contents
OSX_BUNDLE=$(dir $(OSX_BUNDLE_CONTENTS))
OSX_BUNDLE_SRC=$(wildcard ./auto/osx/*)

.PHONY: all test clean osx

all: $(EXE)

osx: $(OSX_BUNDLE)

files:
	@echo $(SRC)
	@echo $(OSX_BUNDLE_SRC)

clean:
	$(RM) $(dir $(EXE))
	$(RM) $(dir $(VENDOR_MANIFEST))
	$(RM) $(dir $(COVERPROFILE))
	$(RM) $(OSX_DIST)

cleanall:
	$(RM) $(DIST)

$(VENDOR_MANIFEST): go.mod go.sum
	@echo "[charlestown] building vendor dir"
	$(GO) mod tidy
	$(GO) mod vendor $(VENDOR_FLAGS)

lint: $(SRC)
	@echo "[charlestown] getting lint tools"
	$(if $(shell which golint), @echo "  - golint found", $(GO) get golang.org/x/lint/golint)
	$(if $(shell which gocyclo), @echo "  - gocyclo found", $(GO) get github.com/fzipp/gocyclo)
	$(if $(shell which misspell), @echo "  - misspell found", $(GO) get github.com/client9/misspell/cmd/misspell)
	@echo "[charlestown] running misspell"
	misspell -error $(SRC)
	@echo "[charlestown] running gocyclo"
	gocyclo $(CYCLO_FLAGS) $(SRC)
	@echo "[charlestown] running golint"
	$(GO) list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status
	$(GO) mod tidy

test: $(SRC)
	@echo "[charlestown] running tests"
	mkdir -p $(basename $(COVERPROFILE))
	touch $(COVERPROFILE)
	$(GO) vet
	$(GO) test $(TEST_FLAGS) ./...

$(EXE): $(SRC) $(VENDOR_MANIFEST)
	@echo "[charlestown] building"
	$(GO) build -o $(EXE) $(BUILD_FLAGS)

$(OSX_BUNDLE): $(EXE) $(OSX_BUNDLE_SRC)
	@echo "[charlestown] building osx bundle $(OSX_BUNDLE_SRC)"
	mkdir -p $(OSX_BUNDLE_CONTENTS)
	mkdir -p $(OSX_BUNDLE_CONTENTS)/MacOS
	cp $(EXE) $(OSX_BUNDLE_CONTENTS)/MacOS/
	cp -r $(dir $(OSX_BUNDLE_SRC))* "$(OSX_BUNDLE_CONTENTS)/"
