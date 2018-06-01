DEBUG ?= 1
GO ?= go
GODOC ?= godoc
GOFMT ?= gofmt
GOLINT ?= golint
EPOCH_TEST_COMMIT ?= 1cc5a27
PROJECT := github.com/opencontrol/compliance-masonry
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
PREFIX ?= ${DESTDIR}/usr/local
BINDIR ?= ${PREFIX}/bin
DATAROOTDIR ?= ${PREFIX}/share/compliance-masonry

BASHINSTALLDIR=${PREFIX}/share/bash-completion/completions

SELINUXOPT ?= $(shell selinuxenabled 2>/dev/null && echo -Z)

BUILD_INFO := $(shell date +%s)

# If GOPATH not specified, use one in the local directory
ifeq ($(GOPATH),)
export GOPATH := $(CURDIR)
unexport GOBIN
endif
GOPKGDIR := $(GOPATH)/src/$(PROJECT)
GOPKGBASEDIR := $(shell dirname "$(GOPKGDIR)")

# debug flag
ifeq ($(DEBUG), 1)
DEBUGFLAGS ?= -gcflags="-N -l"
endif

build:
	$(GO) build \
		$(DEBUGFLAGS) \
		-o ./build/compliance-masonry \
		./masonry-go.go ./diff.go

all: build 

platforms:
	./build-all.sh $(GOPATH) $(GO) $(PROJECT)

default: help

help:
	@echo "Usage: make <target>"
	@echo
	@echo " * 'install' - Install binaries to system locations"
	@echo " * 'build' - Build compliance-masonry"
	@echo " * 'platforms' - Cross-compile build for multiple environments"
	@echo " * 'test' - Execute integration tests"
	@echo " * 'clean' - Clean artifacts"
	@echo " * 'lint' - Execute the source code linter"
	@echo " * 'gofmt' - Verify the source code gofmt"
	@echo " * 'uninstall' - Uninstall binaries to system locations"

gofmt:
	find . -name '*.go' ! -path './vendor/*' -exec $(GOFMT) -s -w {} \+
	git diff --exit-code

lint:
	for PKG in $$($(GO) list ./...) ; do $(GOLINT) $$PKG ; done 

clean:
	@rm -fR ./build

test:
	$(GO) test -v $(shell $(GO) list ./...)

install: build
	install ${SELINUXOPT} -D -m 755 ./build/compliance-masonry $(BINDIR)/compliance-masonry

uninstall:
	rm -f $(BINDIR)/compliance-masonry

.PHONY: \
	build \
	clean \
	default \
	gofmt \
	help \
	install \
	lint \
	platforms \
	test \
	test-suite \
	uninstall \
