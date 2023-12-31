# Simple Go package for fetching string data from common sources

[![GoDoc](https://godoc.org/github.com/rwxrob/get?status.svg)](https://godoc.org/github.com/rwxrob/get)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rwxrob/get)](https://goreportcard.com/report/github.com/rwxrob/get)


Get data from any common place a person might keep it based on the user's preference.

```go
get.String(target string) (string, error)
```

The `target` is in the form of a URL but includes additional schemas to those expected. See Go documentation for more.

## Installation

Most will likely want to cut and paste from this package and "vendor" the code into your own (which is encouraged to keep dependencies down where possible). Please consider mentioning where you got that code or keeping the heading of these files with the original copyright.

Otherwise, you can just import like anything else:

```sh
go get -u github.com/rwxrob/get
```
