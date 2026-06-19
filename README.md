# adfmd

A CLI tool for converting between Atlassian Document Format (ADF) and Markdown.

Built on top of:

- [adf-to-markdown](https://github.com/ajbeck/adf-to-markdown) — ADF JSON to Markdown
- [goldmark-adf](https://github.com/ajbeck/goldmark-adf) — Markdown to ADF JSON

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
