# gtmd
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
$ go run . -u https://go.dev/play/
```

### Using positional arguments
```shell
$ go run . https://go.dev/play/
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
$ go run . --url https://bookmeter.com/books/556977 --mode bookmeter
```
