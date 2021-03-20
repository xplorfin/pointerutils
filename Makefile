# variables related to your microservice
# change SERVICE_NAME to whatever you're naming your microservice
SERVICE_NAME=pointer-utils
# name of output binary
BINARY_NAME=main

# latest git commit hash
LATEST_COMMIT_HASH=$(shell git rev-parse HEAD)

# go commands and variables
GO=go
GOB=$(GO) build
GOM=$(GO) mod
GOTEST=$(GO) test
GOGET=$(GO) get -u
GOM_TIDY=$(GOM) tidy
CLEAN_MODS=$(GO) clean -modcache

# git commands
GIT=git

# environment variables related to
# cross-compilation.
GOOS_MACOS=darwin
GOOS_LINUX=linux
GOARCH=amd64

# currently installed/running Go version (full and minor)
GOVERSION=$(shell go version | grep -Eo '[1-2]\.[[:digit:]]{1,3}\.[[:digit:]]{0,3}')
MINORVER=$(shell echo $(GOVERSION) | awk '{ split($$0, array, ".") } {print array[2]}')
GOPATH=$(shell go env GOPATH)
GOPATH_BIN="$(GOPATH)/bin"

# Color code definitions
# Note: everything is bold.
GREEN=\033[1;38;5;70m
BLUE=\033[1;38;5;27m
LIGHT_BLUE=\033[1;38;5;32m
MAGENTA=\033[1;38;5;128m
RESET_COLOR=\033[0m

COLORECHO = $(1)$(2)$(RESET_COLOR)

GOLANGCI_URI="https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh"

default: hook-config help

hook-config:
	git config core.hooksPath .github/hooks --replace-all

setup: hook-config air-install stringer-install ## setup the repository (enables git hooks and installs air/stringer for generation)

bench:  ## Run all benchmarks in the Go application
	@$(GOTEST) -bench=. -benchmem

check-vars:
	@echo $(GOPATH_BIN)

mod-tidy:
	@$(GOM_TIDY)

clean-mods: ## Remove all the Go mod cache
	@$(CLEAN_MODS)

coverage: ## Get the test coverage from go-coverage
	@$(GOTEST) -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

godocs: ## Run a godoc server
	@echo "godoc server running on http://localhost:9000"
	@godoc -http=":9000"

golangci-install:
	@#Travis (has sudo)
	@if [ "$(shell which golangci-lint)" = "" ] && [ $(TRAVIS) ]; then curl -sSfL $(GOLANGCI_URI) | sh -s v1.33.0 && sudo cp ./bin/golangci-lint $(GOPATH_BIN)/; fi;
	@#AWS CodePipeline
	@if [ "$(shell which golangci-lint)" = "" ] && [ "$(CODEBUILD_BUILD_ID)" != "" ]; then curl -sSfL $(GOLANGCI_URI) | sh -s -- -b $(GOPATH_BIN); fi;
	@#Github Actions
	@if [ "$(shell which golangci-lint)" = "" ] && [ "$(GITHUB_WORKFLOW)" != "" ]; then curl -sfL $(GOLANGCI_URI) | sudo sh -s -- -b $(GOPATH_BIN); fi;
	@#Brew - MacOS
	@if [ "$(shell which golangci-lint)" = "" ] && [ "$(shell which brew)" != "" ]; then brew install golangci-lint; fi;

go-acc-install:
	@if [ "$(shell which "go-acc")" = "" ]; then $(GOGET) github.com/ory/go-acc; fi;

air-install:
	@if [ "$(shell which "air")" = "" ]; then $(GOGET) github.com/cosmtrek/air; fi;

stringer-install:
	@if [ "$(shell which "stringer")" = "" ]; then $(GOGET) golang.org/x/tools/cmd/stringer; fi;

ci-lint: golangci-install ## Run the golangci-lint application (install if not found) & fix issues if possible
	@golangci-lint run --fix

# pre-commit hook
pre-commit: lint

test: ## run tests without coverage reporting
	@$(GOTEST) ./...

ci-test: golangci-install go-acc-install # run a test with coverage
	@golangci-lint run
	@go-acc -o profile.cov ./...


help: ## This help dialog.
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
	for help_line in $${help_lines[@]}; do \
		IFS=$$'#' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf "%-30s %s\n" $$help_command $$help_info ; \
	done

