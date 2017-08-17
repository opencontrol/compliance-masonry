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
GODEBUG = dlv
GLIDE   = glide
TIMEOUT = 15
Q = $(if $(filter 1,$(VERBOSE)),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

# debug flag
ifeq ($(DEBUG), 1)
DEBUGFLAGS ?= -gcflags="-N -l"
endif

# dependencies
DEPEND=github.com/golang/lint/golint \
			 github.com/Masterminds/vcs \
			 github.com/blang/semver \
			 github.com/codegangsta/cli \
			 github.com/go-utils/ufs \
			 github.com/tg/gosortmap \
			 github.com/vektra/errors \
			 github.com/laher/goxc \
			 gopkg.in/fatih/set.v0 \
			 gopkg.in/yaml.v2 \
			 vbom.ml/util/sortorder

########################################################################
# standard targets
.PHONY: all build debug clean rebuild test

all: build

build: env-setup
	env GOPATH=$(l_GOPATH) $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		$(DEBUGFLAGS) \
		-o $(BIN)/compliance-masonry \
		./masonry-go.go ./diff.go ./export.go

# example usage:
#   make debug DEBUG_OPTIONS='--debug x -o ~/proj/git/src/lmgitlab.hlsdev.local/demos/ssa-mde-ato/poc/poc-apps/openshift-dev-int/compliance/opencontrols -d /tmp/foo.json -f json -n -k -x FedRAMP-moderate'
debug: build
	@$(GODEBUG) exec $(BIN)/compliance-masonry -- $(DEBUG_OPTIONS)

clean: env-setup
	@rm -fR $(BIN)

rebuild: clean build

test: env-setup
	@env GOPATH=$(l_GOPATH) $(GO) get -t ./...
	env GOPATH=$(l_GOPATH) $(GO) test $(shell glide nv)

########################################################################
# project-specific targets
.PHONY: lint depend env-setup
lint: env-setup
	@if env GOPATH=$(l_GOPATH) $(GOFMT) -l . | grep -v '^vendor/' | grep -e '\.go'; then \
		echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	for pkg in $$(env GOPATH=$(l_GOPATH) $(GO) list ./... |grep -v /vendor/) ; do env GOPATH=$(l_GOPATH) $(GOLINT) $$pkg ; done

depend: env-setup
	@env GOPATH=$(l_GOPATH) $(GO) get -v $(DEPEND)

publish: test
	@if ! hash ./.goxc.local.json ; then echo Missing .goxc.local.json ; false ; else true ; fi
	env GOPATH=$(l_GOPATH) goxc $(PUBLISH_OPTIONS)

env-setup:
	@mkdir -p "$(BIN)"
	@if ! hash $(GLIDE) >/dev/null 2>&1 ; then curl https://glide.sh/get | sh ; fi

########################################################################
# git support
.PHONY: sync commit push
sync:
	@git fetch upstream && git checkout master && git merge upstream/master

commit:
	@git add --all && git commit -a

push:
	@git push origin master

