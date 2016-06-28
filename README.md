## Allocate
[![Build Status](https://travis-ci.org/cjrd/allocate.svg?branch=master)](https://travis-ci.org/cjrd/allocate)
[![Coverage Status](https://coveralls.io/repos/github/cjrd/allocate/badge.svg?branch=master)](https://coveralls.io/github/cjrd/allocate?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/cjrd/allocate)](https://goreportcard.com/report/github.com/cjrd/allocate)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/mkideal/cli?status.svg)](https://godoc.org/github.com/cjrd/allocate)

Allocate provides simple helper functions for allocating go structures so that pointer fields are pointers to zero'd values instead of `nil`.

### Example
```go
package main

import (
	"fmt"
	"github.com/cjrd/allocate"
)

func main() {
	type SimplePtrStruct struct {
		PtrField *int
	}

	ptrStruct := SimplePtrStruct{}

	fmt.Printf("pre allocate.Zero: %v\n", ptrStruct.PtrField)
	allocate.Zero(&ptrStruct)
	fmt.Printf("post allocate.Zero: %v\n", *ptrStruct.PtrField)
}

```

```bash
# OUTPUT
pre allocate.Zero: <nil>
post allocate.Zero: 0
```

See the godoc's for more information: https://godoc.org/github.com/cjrd/allocate
