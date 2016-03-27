/*
Simple tests for the allocate library

@author Colorado Reed (colorado at dotdashpay ... com)

This library was seeded by this discussion in the golang-nuts mailing list:
https://groups.google.com/forum/#!topic/golang-nuts/Wd9jiZswwMU
*/

package allocate

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type AllBuiltinTypes struct {
	// pointers
	BoolPtr       *bool
	BytePtr       *byte
	Complex128Ptr *complex128
	Complex64Ptr  *complex64
	ErrorPtr      *error
	Float32Ptr    *float32
	Float64Ptr    *float64
	IntPtr        *int
	Int16Ptr      *int16
	Int32Ptr      *int32
	Int64Ptr      *int64
	Int8Ptr       *int8
	RunePtr       *rune
	StringPtr     *string
	UintPtr       *uint
	Uint16Ptr     *uint16
	Uint32Ptr     *uint32
	Uint64Ptr     *uint64
	Uint8Ptr      *uint8
	UintptrPtr    *uintptr

	// non-pointers
	BoolType       bool
	ByteType       byte
	Complex128Type complex128
	Complex64Type  complex64
	ErrorType      error
	Float32Type    float32
	Float64Type    float64
	IntType        int
	Int16Type      int16
	Int32Type      int32
	Int64Type      int64
	Int8Type       int8
	RuneType       rune
	StringType     string
	UintType       uint
	Uint16Type     uint16
	Uint32Type     uint32
	Uint64Type     uint64
	Uint8Type      uint8
	UintptrType    uintptr
}

// TestZeroWithAllBuiltinTypesStruct uses AllBuiltinTypes to test that all of the pointer struct fields
// in AllBuiltinTypes initialize to the same value as all of the non-pointers
// when the `Zero` function is called.
func TestZeroWithAllBuiltinTypesStruct(t *testing.T) {
	allBuiltins := new(AllBuiltinTypes)
	Zero(allBuiltins)
	checkAllBuiltins(t, allBuiltins)
}

// TestZeroWithEmbeddedBuiltinTypesStruct is the same test as TestZeroWithAllBuiltinTypesStruct
// except the struct is embedded as a pointer field in a wrapper struct
func TestZeroWithEmbeddedBuiltinTypesStruct(t *testing.T) {
	type WrapperStruct struct {
		AllBuiltins *AllBuiltinTypes
	}
	wrapStruct := new(WrapperStruct)
	Zero(wrapStruct)
	checkAllBuiltins(t, wrapStruct.AllBuiltins)
}

// TestZeroWithMapField tests `Zero` with a struct field that is a map
func TestZeroWithMapField(t *testing.T) {
	type MapFieldStruct struct {
		MapField map[string]int
	}
	simpleMapStruct := new(MapFieldStruct)
	Zero(simpleMapStruct)
	// This would panic if executed pre-Zero
	simpleMapStruct.MapField["test"] = 25
}

// TestZeroWithPointerMapField tests `Zero` with a struct field that is a pointer to a map
func TestZeroWithPointerMapField(t *testing.T) {
	type PtrMapFieldStruct struct {
		MapField *map[string]int
	}
	ptrMapStruct := new(PtrMapFieldStruct)
	Zero(ptrMapStruct)
	// This would panic if executed pre-Zero
	(*ptrMapStruct.MapField)["test"] = 25
}

// TestZeroWithPrivateField tests that a private field is not allocated
// using `Zero` but a public field is allocated.
func TestZeroWithPrivateField(t *testing.T) {
	type PrivateFieldStruct struct {
		privField   *int
		PublicField *int
	}
	pfstruct := new(PrivateFieldStruct)

	err := Zero(pfstruct)

	if err != nil {
		t.Errorf("Private field struct produced error: %v", err)
	}

	if pfstruct.privField != nil {
		t.Errorf("Private field is not nil: %v", pfstruct.privField)
	}

	if pfstruct.PublicField == nil || *pfstruct.PublicField != 0 {
		t.Errorf("Public field was not allocated correctly: %v", pfstruct.PublicField)
	}
}

// ExampleZero is a zipmle example to demonstrate the how Zero() performs allocation
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

func checkAllBuiltins(t *testing.T, allBuiltins *AllBuiltinTypes) {
	// for all pointer fields, check that the dereferenced pointer
	// equals the non-pointer fields
	allBuiltinsVal := reflect.Indirect(reflect.ValueOf(allBuiltins))
	for i := 0; i < allBuiltinsVal.NumField(); i++ {
		field := allBuiltinsVal.Field(i)
		if field.Kind() != reflect.Ptr {
			continue
		}
		fieldName := allBuiltinsVal.Type().Field(i).Name
		nonPtrFieldName := strings.TrimSuffix(fieldName, "Ptr") + "Type"
		nonPtrField := allBuiltinsVal.FieldByName(nonPtrFieldName)

		// compare the pointer vs non-pointer init values
		if !reflect.DeepEqual(nonPtrField.Interface(), field.Elem().Interface()) {
			t.Errorf("Builtin pointer to struct field '%s' not initialized to its zero value: %v",
				fieldName, field.Elem())
		}
	}
}
