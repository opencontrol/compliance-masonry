DEBUG ?= 1
GO ?= go
GODOC ?= godoc
GOFMT ?= gofmt
GOLINT ?= golint
EPOCH_TEST_COMMIT ?= 1cc5a27
PROJECT := github.com/opencontrol/compliance-masonry
PROGNAME ?= masonry
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
PREFIX ?= ${DESTDIR}/usr/local
BINDIR ?= ${PREFIX}/bin
DATAROOTDIR ?= ${PREFIX}/share/masonry

BASHINSTALLDIR=${PREFIX}/share/bash-completion/completions

SELINUXOPT ?= $(shell selinuxenabled 2>/dev/null && echo -Z)

BUILD_DIR ?= ./build
BUILD_INFO := $(shell date --iso-8601=s --utc)
BUILD_PLATFORMS = "windows/amd64" "windows/386" "darwin/amd64" "darwin/386" "linux/386" "linux/amd64" "linux/arm" "linux/arm64"

VERSION := $(shell git for-each-ref --format="%(refname:short)" --sort=-authordate --count=1 refs/tags | cut -d"v" -f2)
SHORTCOMMIT := $(shell git rev-parse --short HEAD)

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

LDFLAGS ?= -ldflags="-s -w -X github.com/opencontrol/compliance-masonry/version.Version=$(VERSION) \
	   -X github.com/opencontrol/compliance-masonry/version.Commit=$(SHORTCOMMIT) \
	   -X github.com/opencontrol/compliance-masonry/version.Date=$(BUILD_INFO)"

build:
	$(GO) build $(DEBUGFLAGS) $(LDFLAGS)\
		-o $(BUILD_DIR)/$(PROGNAME) \
		cmd/masonry/masonry.go
	$(GO) build $(DEBUGFLAGS) $(LDFLAGS)\
                -o $(BUILD_DIR)/compliance-masonry \
                cmd/compliance-masonry/compliance-masonry.go

all: build

platforms:
	@for platform in $(BUILD_PLATFORMS) ; do \
		GOOS="`echo $$platform | cut -d"/" -f1`"; \
		GOARCH="`echo $$platform | cut -d"/" -f2`"; \
		output_name="$(BUILD_DIR)/$$GOOS-$$GOARCH/$(PROGNAME)"; \
		legacy_name="$(BUILD_DIR)/$$GOOS-$$GOARCH/compliance-masonry"; \
		[ $$GOOS = "windows" ] && output_name="$$output_name.exe"; \
		[ $$GOOS = "windows" ] && legacy_name="$$legacy_name.exe"; \
		echo "Building $(PROGNAME) version $(VERSION) for $$GOOS on $$GOARCH"; \
		GOOS=$$GOOS GOARCH=$$GOARCH $(GO) build $(LDFLAGS) -o $$output_name cmd/masonry/masonry.go; \
		GOOS=$$GOOS GOARCH=$$GOARCH $(GO) build $(LDFLAGS) -o $$legacy_name cmd/compliance-masonry/compliance-masonry.go; \
		[ -d $(BUILD_DIR)/$$GOOS-$$GOARCH/ ] && cp {LICENSE.md,README.md} $(BUILD_DIR)/$$GOOS-$$GOARCH/; \
	done

default: help

help:
	@echo "Usage: make <target>"
	@echo
	@echo " * 'install' - Install binaries to system locations"
	@echo " * 'build' - Build compliance-masonry"
	@echo " * 'platforms' - Cross-compile build for multiple environments"
	@echo " * 'archive' - Create tar and zip files for packaging and releases"
	@echo " * 'test' - Execute integration tests"
	@echo " * 'clean' - Clean artifacts"
	@echo " * 'lint' - Execute the source code linter"
	@echo " * 'gofmt' - Verify the source code gofmt"

gofmt:
	find . -name '*.go' ! -path './vendor/*' -exec $(GOFMT) -s -w {} \+
	git diff --exit-code *.go

lint:
	for PKG in $$($(GO) list ./...) ; do $(GOLINT) $$PKG ; done 

clean:
	@rm -fR $(BUILD_DIR)

test:
	$(GO) test -v $(shell $(GO) list ./...)

install: build
	install ${SELINUXOPT} -D -m 755 $(BUILD_DIR)/$(PROGNAME) $(BINDIR)/$(PROGNAME)

uninstall:
	rm -f $(BINDIR)/$(PROGNAME)

archive: platforms
	$(eval PKGS := $(shell find $(BUILD_DIR)/* ! -name archive -type d))
	@mkdir -p $(BUILD_DIR)/archive/source/$(PROGNAME)-$(VERSION)/
	for PKG in $(PKGS); do \
		OS="`echo $$PKG | cut -d"/" -f3 | cut -d"-" -f1`"; \
		[ $$OS = "windows" ] && zip $(BUILD_DIR)/archive/$(PROGNAME)-`basename $$PKG`-$(VERSION).zip -r $$PKG -j ; \
		[ $$OS != "windows" ] && tar cvf $(BUILD_DIR)/archive/$(PROGNAME)-`basename $$PKG`-$(VERSION).tar.gz -C $$PKG . ; \
		true ; \
	done
	$(eval SFILES := $(shell find -maxdepth 1 ! -path '*/\.*' ! -path '\.' ! -name build ! -name appveyor.yml))
	@cp -r $(SFILES) $(BUILD_DIR)/archive/source/$(PROGNAME)-$(VERSION)/
	@cd $(BUILD_DIR)/archive/source/ && tar cvf $(PROGNAME)-$(VERSION).tar.gz $(PROGNAME)-$(VERSION)
	@cd $(BUILD_DIR)/archive/source/ && zip $(PROGNAME)-$(VERSION).zip -r $(PROGNAME)-$(VERSION)
	@echo "Binary archives have been created at $(BUILD_DIR)/archive"
	@echo "Source code has been created at $(BUILD_DIR)/archive/source/"

.PHONY: \
	archive \
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
