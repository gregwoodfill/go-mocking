.PHONY: clean test format lint build initial env

all: clean test format lint build

help:
	@echo "initial - install required tools to system, assumes Mac OS"
	@echo "env - update local environment to latest tool versions"
	@echo "clean, all format, lint, build - all self-explanatory"

clean:
	go clean

test:
	go test ./...

format:
	go fmt ./...

lint:
	go vet ./...

build:
	go build ./...

initial:
	# intial installations required to run
	# assumes mac system with homebrew and python
	@brew install asdf
	@asdf plugin-add golang https://github.com/kennyp/asdf-golang.git
	@pip install pre-commit

.make:
	# directory to store files to track timestamps so that make doesn't run all targets every time
	@mkdir -p .make

install-precommit: .make/install-precommit
.make/install-precommit:
	@pip install pre-commit
	@touch $@

install-precommit-hooks: .make/install-precommit-hooks install-precommit
.make/install-precommit-hooks: .pre-commit-config.yaml
	@pre-commit install
	@touch $@

install-asdf: .make/install-asdf
.make/install-asdf: .tool-versions
	@asdf install
	@touch $@

env: install-precommit-hooks install-asdf
