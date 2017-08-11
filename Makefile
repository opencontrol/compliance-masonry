########################################################################
# Makefile, ABr
# Project support to build infrastructure

# initial settings
DEBUG ?= 1
VERBOSE ?= 1

# settings for go build
GOBIN ?= "$(GOPATH)/bin"
l_GOPATH="$(shell cd ../../../.. && realpath .):$(CURDIR)"

# boilerplate
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
BIN     = ./bin
GO      = go
GODOC   = godoc
GOFMT   = gofmt
GOLINT  = golint
GLIDE   = glide
TIMEOUT = 15
Q = $(if $(filter 1,$(VERBOSE)),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# debug flag
ifdef DEBUG
DEBUGFLAGS ?= -gcflags="-N -l"
endif

# dependencies
DEPEND=github.com/golang/lint/golint

########################################################################
# standard targets
.PHONY: all build clean rebuild test lint depend env-setup

all: build

build: env-setup
	env GOPATH=$(l_GOPATH) $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		$(DEBUGFLAGS) \
		-o $(BIN)/compliance-masonry \
		./masonry-go.go ./diff.go ./extract.go

clean: env-setup
	@rm -fR $(BIN)

rebuild: clean build

test: env-setup
	@env GOPATH=$(l_GOPATH) $(GO) get -t ./...
	@env GOPATH=$(l_GOPATH) $(GO) test $(shell glide nv)

lint: env-setup
	@if env GOPATH=$(l_GOPATH) $(GOFMT) -l . | grep -v '^vendor/' | grep -e '\.go'; then \
		echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	for pkg in $$(env GOPATH=$(l_GOPATH) $(GO) list ./... |grep -v /vendor/) ; do env GOPATH=$(l_GOPATH) $(GOLINT) $$pkg ; done

depend:
	@go get -v $(DEPEND)


env-setup:
	@mkdir -p "$(BIN)"

