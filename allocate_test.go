/*
Simple tests for the allocate library

@author Colorado Reed (colorado at dotdashpay ... com)

This library was seeded by this discussion in the golang-nuts mailing list:
https://groups.google.com/forum/#!topic/golang-nuts/Wd9jiZswwMU
*/

package allocate

import (
	"fmt"
)

func ExampleZero() {
	type SimplePtrStruct struct {
		PtrField *int
	}

	ptrStruct := SimplePtrStruct{}

	fmt.Printf("pre allocate.Zero: %v\n", ptrStruct.PtrField)
	Zero(&ptrStruct)
	fmt.Printf("post allocate.Zero: %v\n", *ptrStruct.PtrField)

	// Output:
	// pre allocate.Zero: <nil>
	// post allocate.Zero: 0

}
