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
GLIDE   = glide
TIMEOUT = 15
Q = $(if $(filter 1,$(VERBOSE)),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# debug flag
ifdef DEBUG
DEBUGFLAGS ?= -gcflags="-N -l"
endif

########################################################################
# standard targets
.PHONY: all build clean rebuild test env-setup

all: build

build: env-setup
	env GOPATH=$(l_GOPATH) $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		$(DEBUGFLAGS) \
		-o $(BIN)/compliance-masonry \
		./masonry-go.go ./diff.go ./extract.go

clean:
	@rm -fR $(BIN)

rebuild: clean build

test:
	@env GOPATH=$(l_GOPATH) $(GO) get -t ./...
	@env GOPATH=$(l_GOPATH) $(GO) test $(shell glide nv)

env-setup:
	@mkdir -p "$(BIN)"

