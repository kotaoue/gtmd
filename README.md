# gtmd
Get the title tag and make markdown.

## Usage
```shell
$ go run . -url=https://go.dev/play/

$ cat "Go Playground - go.dev.md"
# [Go Playground - go.dev](https://go.dev/play/)
```

### for [読書メーター](https://bookmeter.com)
```shell
$ go run main.go -url=https://bookmeter.com/books/556977 -mode=bookmeter
```

### Alternative usage method.
```
$ go run main.go https://go.dev/play/
```
