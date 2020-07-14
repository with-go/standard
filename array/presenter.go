// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package array

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ElementTypeNotFloat64Error = errors.New("array contains element that is not a float64 type")
	ElementTypeNotInt64Error = errors.New("array contains element that is not a int64 type")
	ElementTypeNotUint64Error = errors.New("array contains element that is not a uint64 type")
)

type Presenter struct {
	array Array
}

// The AsFloat64Slice() function returns the Array as a float64 slice.
func (presenter Presenter) AsFloat64Slice() ([]float64, error) {
	slice := make([]float64, len(presenter.array))
	for index, reflection := range presenter.array.Reflects() {
		var value float64
		switch reflection.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = float64(reflection.Uint())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = float64(reflection.Int())
			break
		case reflect.Float32, reflect.Float64:
			value = reflection.Float()
			break
		default:
			return nil, ElementTypeNotFloat64Error
		}
		slice[index] = value
	}
	return slice, nil
}

// The AsInt64Slice() function returns the Array as a int64 slice.
func (presenter Presenter) AsInt64Slice() ([]int64, error) {
	slice := make([]int64, len(presenter.array))
	for index, reflection := range presenter.array.Reflects() {
		var value int64
		switch reflection.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = int64(reflection.Uint())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = reflection.Int()
			break
		default:
			return nil, ElementTypeNotInt64Error
		}
		slice[index] = value
	}
	return slice, nil
}

// The AsStringSlice() function returns the Array as a string slice.
func (presenter Presenter) AsStringSlice() ([]string, error) {
	slice := make([]string, len(presenter.array))
	for index, value := range presenter.array {
		slice[index] = fmt.Sprintf("%v", value)
	}
	return slice, nil
}

// The AsUint64Slice() function returns the Array as a uint64 slice.
func (presenter Presenter) AsUint64Slice() ([]uint64, error) {
	slice := make([]uint64, len(presenter.array))
	for index, reflection := range presenter.array.Reflects() {
		var value uint64 = 0
		switch reflection.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = reflection.Uint()
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if reflection.Int() < 0 {
				return nil, ElementTypeNotUint64Error
			}
			value = uint64(reflection.Int())
			break
		default:
			return nil, ElementTypeNotUint64Error
		}
		slice[index] = value
	}
	return slice, nil
}