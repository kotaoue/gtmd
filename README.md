# gtmd

[![Go](https://github.com/kotaoue/gtmd/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kotaoue/gtmd/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/kotaoue/gtmd/branch/main/graph/badge.svg)](https://codecov.io/gh/kotaoue/gtmd)
[![Go Report Card](https://goreportcard.com/badge/github.com/kotaoue/gtmd)](https://goreportcard.com/report/github.com/kotaoue/gtmd)
[![License](https://img.shields.io/github/license/kotaoue/gtmd)](https://github.com/kotaoue/gtmd/blob/main/LICENSE)

Get the title tag and make markdown.

## Flags

| Flag | Short | Values | Default | Description |
|------|-------|--------|---------|-------------|
| `--url` | `-u` | any URL | — | Source URL (can also be given as a positional argument) |
| `--format` | `-f` | `link`, `clipboard` | — | Output format. Creates a markdown file when not set. |
| `--source` | `-s` | `bookmeter`, `manual` | — | Source type. Auto-detected from the URL when not set. Use `manual` to bypass auto-detection. |

## Usage

### Basic usage (creates markdown file)

```shell
$ go run . --url https://go.dev/play/

$ cat "Go Playground - The Go Programming Language.md"
# [Go Playground - The Go Programming Language](https://go.dev/play/)
```

### Using short flags

```shell
go run . -u https://go.dev/play/
```

### Using positional arguments

```shell
go run . https://go.dev/play/
```

### Output as markdown link (`--format`)

```shell
# Print to stdout
$ go run . -u https://go.dev/play/ -f link
[Go Playground - The Go Programming Language](https://go.dev/play/)

# Copy to clipboard
$ go run . -u https://go.dev/play/ -f clipboard
Copied to clipboard: [Go Playground - The Go Programming Language](https://go.dev/play/)
```

### Specifying source type (`--source`)

Bookmeter URLs are auto-detected. The source type can also be set explicitly:

```shell
# Auto-detected from URL
go run . -u https://bookmeter.com/books/556977

# Explicit bookmeter source type
go run . -u https://bookmeter.com/books/556977 -s bookmeter

# Bypass auto-detection and create a plain markdown file
go run . -u https://bookmeter.com/books/556977 -s manual
```

## Build

```shell
go build -o gtmd .
```
