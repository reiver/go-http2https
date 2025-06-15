# go-http2https

Package **http2https** provides tools to redirect from HTTP to HTTPS, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-http2https

[![GoDoc](https://godoc.org/github.com/reiver/go-http2https?status.svg)](https://godoc.org/github.com/reiver/go-http2https)

## Examples

Here is a simple example:

```golang
import "github.com/reiver/go-http2https"

// ...

err := http.ListenAndServe(tcpaddr, http2https.Handler)
```

Here is an example that provide the `http2https` http.Hanlder with a function to call to log errors:

```golang
import "github.com/reiver/go-http2https"

// ...

func logError(a ...any) {
	fmt.Println(a...)
}

// ...

var handler http.Handler = http2https.HTTPHandler{
	LogError: logError,
}

err := http.ListenAndServe(tcpaddr, handler)
```

## Import

To import package **http2https** use `import` code like the following:
```
import "github.com/reiver/go-http2https"
```

## Installation

To install package **http2https** do the following:
```
GOPROXY=direct go get github.com/reiver/go-http2https
```

## Author

Package **http2https** was written by [Charles Iliya Krempeaux](http://reiver.link)
