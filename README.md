# gtmd

[![Go](https://github.com/kotaoue/gtmd/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kotaoue/gtmd/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/kotaoue/gtmd/badge.svg?branch=main)](https://coveralls.io/github/kotaoue/gtmd?branch=main)
[![License](https://img.shields.io/github/license/kotaoue/gtmd)](https://github.com/kotaoue/gtmd/blob/main/LICENSE)

Get the title tag and make markdown.

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

### Output markdown link format

```shell
$ go run . -u https://go.dev/play/ -m link
[Go Playground - The Go Programming Language](https://go.dev/play/)
```

### Copy to clipboard

```shell
$ go run . -u https://go.dev/play/ -m clipboard
Copied to clipboard: [Go Playground - The Go Programming Language](https://go.dev/play/)
```

### for [読書メーター](https://bookmeter.com)

```shell
go run . -u https://bookmeter.com/books/556977 -m bookmeter
```

## Build

```shell
go build -o gtmd .
```
