## Allocate
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/cjrd/allocate)](https://goreportcard.com/report/github.com/cjrd/allocate)
[![GoDoc](https://godoc.org/github.com/mkideal/cli?status.svg)](https://godoc.org/github.com/cjrd/allocate)

Allocate provides simple helper functions for allocating go structures so that pointer fields are pointers to zero'd values instead of `nil`.

WARNING: This library is pre-1.0, so I would advise against using this in
production right now (it's great for testing your production code,
though, e.g. allocating structs used in your tests).

See the godoc's for more information: https://godoc.org/github.com/cjrd/allocate
