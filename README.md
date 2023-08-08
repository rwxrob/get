# Simple Go package for fetching string data from common sources

Just get a string from any of the common places people keep them.

```go
get.String(source string) (string, error)
get.File(source, path string) (*os.File, error)
get.Config(source, relpath string) (*os.File, error)
get.Cache(source, relpath string) (*os.File, error)
get.Home(source, relpath string) (*os.File, error)
get.Append(source, path string) (string, error)
get.Print(source string) error
get.Log(source string) error
```

The `source` is in the form of a URL but includes additional schemas to those expected. See Go documentation for more.

## Installation

Most will likely want to cut and paste from this package and "vendor" the code into your own (which is encouraged to keep dependencies down where possible). Please consider mentioning where you got that code or keeping the heading of these files with the original copyright.

Otherwise, you can just import like anything else:

```sh
go get -u github.com/rwxrob/get
```
