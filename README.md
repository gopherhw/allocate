## Allocate

Allocate provides simple helper functions for allocating go structures so that pointer fields are pointers to zero'd values instead of `nil`.

WARNING: This library is pre-1.0, so I would advise against using this in
production right now (it's great for testing your production code,
though, e.g. allocating structs used in your tests).

See the godoc's for more information: https://godoc.org/github.com/cjrd/allocate
