# adfmd

A CLI tool for converting between Atlassian Document Format (ADF) and Markdown.

Built on top of:

- [adf-to-markdown](https://github.com/ajbeck/adf-to-markdown) — ADF JSON to Markdown
- [goldmark-adf](https://github.com/ajbeck/goldmark-adf) — Markdown to ADF JSON

## Installation

Download and install the latest pre-built binary for your platform to `/usr/local/bin`:

### Linux

```bash
curl -Lo adfmd "https://github.com/jhutar/adfmd/releases/latest/download/adfmd-linux-$(uname -m | sed -e 's/x86_64/amd64/' -e 's/aarch64/arm64/')" && chmod +x adfmd && sudo mv adfmd /usr/local/bin/ && adfmd --help
```

### macOS

```bash
curl -Lo adfmd "https://github.com/jhutar/adfmd/releases/latest/download/adfmd-darwin-$(uname -m | sed 's/x86_64/amd64/')" && chmod +x adfmd && sudo mv adfmd /usr/local/bin/ && adfmd --help
```

## Build

```
make build
```

The binary is placed in `bin/adfmd`.

## Usage

Autodetect direction from file extension:

```
adfmd document.adf > document.md
adfmd document.json > document.md
adfmd document.md > document.adf
```

Explicit subcommands with a file or stdin:

```
adfmd to-md document.adf
adfmd to-adf document.md
cat document.adf | adfmd to-md
cat document.md | adfmd to-adf
```

## Releasing

Tags use [Semantic Versioning](https://semver.org/). Pushing a tag triggers a GitHub Action that builds and publishes binaries.

```
git tag v1.0.0
git push origin v1.0.0
```
