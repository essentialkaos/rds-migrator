################################################################################

# This Makefile generated by GoMakeGen 2.3.0 using next command:
# gomakegen --mod .
#
# More info: https://kaos.sh/gomakegen

################################################################################

export GO111MODULE=on

ifdef VERBOSE ## Print verbose information (Flag)
VERBOSE_FLAG = -v
endif

COMPAT ?= 1.18
MAKEDIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
GITREV ?= $(shell test -s $(MAKEDIR)/.git && git rev-parse --short HEAD)

################################################################################

.DEFAULT_GOAL := help
.PHONY = fmt vet all clean deps update init vendor mod-init mod-update mod-download mod-vendor help

################################################################################

all: rds-migrator ## Build all binaries

rds-migrator:
	go build $(VERBOSE_FLAG) -ldflags="-X main.gitrev=$(GITREV)" rds-migrator.go

install: ## Install all binaries
	cp rds-migrator /usr/bin/rds-migrator

uninstall: ## Uninstall all binaries
	rm -f /usr/bin/rds-migrator

init: mod-init ## Initialize new module

deps: mod-download ## Download dependencies

update: mod-update ## Update dependencies to the latest versions

vendor: mod-vendor ## Make vendored copy of dependencies

mod-init:
ifdef MODULE_PATH ## Module path for initialization (String)
	go mod init $(MODULE_PATH)
else
	go mod init
endif

ifdef COMPAT ## Compatible Go version (String)
	go mod tidy $(VERBOSE_FLAG) -compat=$(COMPAT) -go=$(COMPAT)
else
	go mod tidy $(VERBOSE_FLAG)
endif

mod-update:
ifdef UPDATE_ALL ## Update all dependencies (Flag)
	go get -u $(VERBOSE_FLAG) all
else
	go get -u $(VERBOSE_FLAG) ./...
endif

ifdef COMPAT
	go mod tidy $(VERBOSE_FLAG) -compat=$(COMPAT)
else
	go mod tidy $(VERBOSE_FLAG)
endif

	test -d vendor && rm -rf vendor && go mod vendor $(VERBOSE_FLAG) || :

mod-download:
	go mod download

mod-vendor:
	rm -rf vendor && go mod vendor $(VERBOSE_FLAG)

fmt: ## Format source code with gofmt
	find . -name "*.go" -exec gofmt -s -w {} \;

vet: ## Runs 'go vet' over sources
	go vet -composites=false -printfuncs=LPrintf,TLPrintf,TPrintf,log.Debug,log.Info,log.Warn,log.Error,log.Critical,log.Print ./...

clean: ## Remove generated files
	rm -f rds-migrator

help: ## Show this info
	@echo -e '\n\033[1mTargets:\033[0m\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-14s\033[0m %s\n", $$1, $$2}'
	@echo -e '\n\033[1mVariables:\033[0m\n'
	@grep -E '^ifdef [A-Z_]+ .*?## .*$$' $(abspath $(lastword $(MAKEFILE_LIST))) \
		| sed 's/ifdef //' \
		| awk 'BEGIN {FS = " .*?## "}; {printf "  \033[32m%-14s\033[0m %s\n", $$1, $$2}'
	@echo -e ''
	@echo -e '\033[90mGenerated by GoMakeGen 2.3.0\033[0m\n'

################################################################################
