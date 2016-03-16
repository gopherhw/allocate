/*

The allocate library provides helper functions for allocation/initializing
arbitrary structures.

TODO(cjrd)
Add an allocate.Random() function that assigns random values rather than nil values

@author Colorado Reed (colorado at dotdashpay ... com)

This library was seeded by this discussion in the golang-nuts mailing list:
https://groups.google.com/forum/#!topic/golang-nuts/Wd9jiZswwMU
*/

package allocate

import (
	"fmt"
	"reflect"
)

// Zero allocates an input structure such that all pointer fields
// are fully allocated, i.e. rather than having a nil value,
// the pointer contains a pointer to an initialized value,
// e.g. an *int field will be a pointer to 0 instead of a nil pointer.
//
func Zero(inputIntf interface{}) error {
	indirectVal := reflect.Indirect(reflect.ValueOf(inputIntf))

	if !indirectVal.CanSet() {
		return fmt.Errorf("Input interface is not addressable (can't Set the memory address): %#v",
			inputIntf)
	}
	if indirectVal.Kind() != reflect.Struct {
		return fmt.Errorf("allocate.Zero only works with structs, not type %v",
			indirectVal.Kind())
	}

	// allocate each of the structs fields
	for i := 0; i < indirectVal.NumField(); i++ {
		f := indirectVal.Field(i)
		if f.Kind() == reflect.Ptr && f.IsNil() {
			f.Set(reflect.New(f.Type().Elem()))
		}
		i := reflect.Indirect(f)
		if i.Kind() == reflect.Struct {
			// recursively allocate each of the structs embedded fields
			err := Zero(f.Interface())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
