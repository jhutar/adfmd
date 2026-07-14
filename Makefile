.DEFAULT_GOAL := help
.PHONY: help bootstrap check check-all build test vendor tag

# Let Go automatically download the toolchain version required by go.mod.
export GOTOOLCHAIN := auto

# Required by adf-to-markdown and goldmark-adf libraries.
export GOEXPERIMENT := jsonv2

help:
	@echo "Available targets:"
	@echo "  help                 - Show this help message"
	@echo "  bootstrap            - Install all development tools"
	@echo "  check                - Run checks on staged changes"
	@echo "  check-all            - Run checks on all files"
	@echo "  build                - Build the adfmd binary"
	@echo "  test                 - Smoke-test the adfmd binary"
	@echo "  vendor               - Tidy and vendor Go dependencies"
	@echo "  tag TAG=vX.Y.Z       - Create and push a release tag"

bootstrap:
	@echo "==> Installing Python 3.12 (via uv)..."
	uv python install 3.12
	@echo "==> Installing pre-commit..."
	uv tool install pre-commit || uv tool upgrade pre-commit
	@echo "==> Installing golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "==> Installing gosec..."
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	@echo "==> Installing pre-commit hooks..."
	@PATH="$(HOME)/.local/bin:$(PATH)" pre-commit install
	@echo ""
	@echo "==> Bootstrap complete!"
	@echo "    Make sure $(HOME)/.local/bin and $(HOME)/go/bin are on your PATH."

check:
	pre-commit run

check-all:
	pre-commit run --all-files

build:
	go build -o bin/adfmd main.go

test: build
	./bin/adfmd --help > /dev/null
	./bin/adfmd to-md --help > /dev/null
	./bin/adfmd to-adf --help > /dev/null
	./bin/adfmd to-md resources/sample.adf > /dev/null
	./bin/adfmd to-adf resources/sample.md > /dev/null
	cat resources/sample.adf | ./bin/adfmd to-md > /dev/null
	cat resources/sample.md | ./bin/adfmd to-adf > /dev/null
	./bin/adfmd resources/sample.adf > /dev/null
	./bin/adfmd resources/sample.md > /dev/null
	echo '{"type":"doc","version":1,"content":[{"type":"codeBlock","attrs":{"wrap":true},"content":[{"type":"text","text":"hello"}]}]}' | ./bin/adfmd to-md > /dev/null
	echo '{"type":"doc","version":1,"content":[{"type":"taskList","content":[{"type":"taskItem","content":[{"type":"text","text":"item 1"}],"attrs":{"localId":"a","state":"TODO"}},{"type":"taskList","content":[{"type":"taskItem","content":[{"type":"text","text":"nested"}],"attrs":{"localId":"b","state":"TODO"}}],"attrs":{"localId":"c"}}],"attrs":{"localId":"d"}}]}' | ./bin/adfmd to-md > /dev/null
	@echo "All good!"

tag:
ifndef TAG
	@LATEST=$$(git tag --sort=-v:refname | head -1); \
	if [ -z "$$LATEST" ]; then \
		echo "No tags found. Create one with:"; \
		echo "  make tag TAG=v0.1.0"; \
	else \
		echo "Latest tag: $$LATEST"; \
		echo "  make tag TAG=<new_tag>"; \
	fi
else
	git tag $(TAG)
	git push origin $(TAG)
endif

vendor:
	go mod tidy
	go mod vendor
