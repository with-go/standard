// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package array

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	NonSliceTypeError = errors.New("the given parameter v is a non-slice type, " +
		"parameter v should be a map")
)

// The New() function creates a new Array. If there are values given, it will be
// pushed as the elements the new Array.
func New(values ...interface{}) Array {
	array := Array{}
	array = array.Push(values...)
	return array
}

// The NewFromSlice() function creates a new Array from a native slice. If the given
// parameter v is not a slice, it will returns NonSliceTypeError.
func NewFromSlice(v interface{}) (Array, error) {
	reflection := reflect.ValueOf(v)
	if reflection.Kind() != reflect.Slice {
		return New(), NonSliceTypeError
	}
	newSlice := make([]interface{},reflection.Len())
	for index := 0; index < reflection.Len(); index++ {
		newSlice[index] = reflection.Index(index).Interface()
	}
	return New().Push(newSlice...), nil
}

// Array defines an Array Type. See "array" package documentation for more
// information.
type Array []interface{}

// The Concat() function is used to merge two or more Arrays. This function
// does not change the existing array, but instead returns a new Array.
func (array Array) Concat(others ...Array) Array {
	concatenatedArray := array
	for _, other := range others {
		concatenatedArray = concatenatedArray.Push(other...)
	}
	return concatenatedArray
}

// The DeepEqual() function is the same as Equal() function but, instead of using
// an internal testing mechanism, it would use the DeepEqual() function from the
// "reflect" package.
func (array Array) DeepEqual(other Array) bool {
	return reflect.DeepEqual(array, other)
}

// The Equal() function determines whether the Array has the same elements
// compared to the elements of the other provided Array. Note that the
// element and its order has to be the same for this function to returns true.
func (array Array) Equal(other Array) bool {
	if len(array) != len(other) {
		return false
	}
	for i, v := range array {
		if other[i] != v {
			return false
		}
	}
	return true
}

// The Filter() function creates a new Array with all elements that pass
// the test implemented by the provided function.
func (array Array) Filter(function FilterFunc) Array {
	filteredArray := New()
	for i, v := range array {
		if function(array, i, v) {
			filteredArray = filteredArray.Push(v)
		}
	}
	return filteredArray
}

// The Find() function returns the value of the first element in the provided
// Array that satisfies the provided testing function. Otherwise, it returns
// nil, indicating that no element passed the test.
func (array Array) Find(function FindFunc) interface{} {
	for i, v := range array {
		if function(array, i, v) {
			return v
		}
	}
	return nil
}

// The FindIndex() function returns the index of the first element in the Array
// that satisfies the provided testing function. Otherwise, it returns -1,
// indicating that no element passed the test.
func (array Array) FindIndex(function FindIndexFunc) int {
	for i, v := range array {
		if function(array, i, v) {
			return i
		}
	}
	return -1
}

// The ForEach() function executes a provided function once for each Array element.
func (array Array) ForEach(function ForEachFunc) {
	for i, v := range array {
		function(array, i, v)
	}
}

// The Includes() function determines whether the Array includes a certain value among
// its entries, returning true or false as appropriate.
func (array Array) Includes(value interface{}) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// The IndexOf() function returns the first index at which a given element can be found
// in the Array, or -1 if it is not present.
func (array Array) IndexOf(value interface{}) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// The Join() function creates and returns a new string by concatenating all of the elements
// in the Array, separated by a specified separator string. If the Array has only one item,
// then that item will be returned without using the separator.
func (array Array) Join(separator string) string {
	str := fmt.Sprint()
	for i, v := range array {
		str += fmt.Sprintf("%v", v)
		if i != len(array) - 1 {
			str += fmt.Sprintf("%s", separator)
		}
	}
	return str
}

// The Keys() function returns a new slice of integer that contains the keys for
// each index in the Array.
func (array Array) Keys() []int {
	var s []int
	for i := range array {
		s = append(s, i)
	}
	return s
}


// The LastIndexOf() function returns the last index at which a given element can be found in
// the Array, or -1 if it is not present. The Array is searched backwards, starting at
// fromIndex. If fromIndex is less than 0 or more than the last index of the Array, it
// will automatically be set to the last index of the Array.
func (array Array) LastIndexOf(value interface{}, fromIndex int) int {
	if fromIndex < 0 || fromIndex > len(array) - 1 {
		fromIndex = len(array) - 1
	}
	for i := fromIndex; i >= 0; i-- {
		v := array[i]
		if v == value {
			return i
		}
	}
	return -1
}

// The Length() function returns the number of elements contained inside the Array.
func (array Array) Length() int {
	return len(array)
}

// The Map() function creates a new Array populated with the results of calling
// a provided function on every element in the Array.
func (array Array) Map(function MapFunc) Array {
	newArray := New()
	for i, v := range array {
		newArray = newArray.Push(function(array, i, v))
	}
	return newArray
}

// The Pop() function removes the last element from the Array and returns the new
// Array and that last element. This function does not change the existing array.
func (array Array) Pop() (Array, interface{}) {
	if len(array) == 0 {
		return array, nil
	}
	lastIndex := len(array) - 1
	lastValue := array[lastIndex]
	return array[:lastIndex], lastValue
}

// The Present() function returns an Array Presenter, which capable to
// returns the the Array to other predefined data type.
func (array Array) Present() Presenter {
	return Presenter{ array }
}

// The Push() function adds one or more elements to the end of an array and
// returns the new Array. This function does not change the existing array.
func (array Array) Push(values ...interface{}) Array {
	return append(array, values...)
}

// The Reflect() function returns the array element of the given index as a
// reflect.Value data. If there is no element saved with the given index, it
// will returns reflect.Value of nil.
func (array Array) Reflect(index int) reflect.Value {
	if index < 0 || index >= len(array) {
		return reflect.ValueOf(nil)
	}
	return reflect.ValueOf(array[index])
}

// The Reflects() function returns the slice of reflect.Value
func (array Array) Reflects() []reflect.Value {
	slice := make([]reflect.Value, len(array))
	for i, _ := range array {
		slice[i] = array.Reflect(i)
	}
	return slice
}

// The Reverse() function reverses the Array in place. The first Array element
// becomes the last, and the last Array element becomes the first. This function
// does not change the existing array.
func (array Array) Reverse() Array {
	reversedArray := New()
	for i := len(array) - 1; i >= 0; i-- {
		reversedArray = reversedArray.Push(array[i])
	}
	return reversedArray
}

// The Shift() function removes the first element from the Array and returns
// the new Array and that removed element. This function does not change
// the existing array.
func (array Array) Shift() (Array, interface{}) {
	if len(array) == 0 {
		return array, nil
	}
	firstValue := array[0]
	return array[1:], firstValue
}

// The Sort() function sorts the elements of the Array in place based on a
// provided compare function that compare two elements in array named "a"
// and "b".
//  - If compareFunction(a, b) returns less than 0, a will be sorted
//    to an index lower than b.
//  - If compareFunction(a, b) returns exactly 0, a and b will be
//    leaved unchanged with respect to each other, but sorted with
//    respect to all different elements.
//  - If compareFunction(a, b) returns greater than 0, b will be
//    sorted to an index lower than a.
func (array Array) Sort(function SortFunc) {
	for {
		shiftCount := 0
		for i := 1; i < len(array); i++ {
			a := array[i-1]
			b := array[i]
			sort := function(a, b)
			if sort > 0 {
				array[i-1] = b
				array[i] = a
				shiftCount++
			}
		}
		if shiftCount == 0 {
			break
		}
	}
}
// The String() function returns a string representing the specified Array
// and its elements.
func (array Array) String() string {
	b, _ := json.Marshal(array)
	return string(b)
}

// The Unshift() function adds one or more elements to the beginning
// of the Array and returns the new array. This function does not change
// the existing array.
func (array Array) Unshift(values ...interface{}) Array {
	return append(values, array...)
}

// The Values() function returns a new slice that contains the values for
// each index in the Array.
func (array Array) Values() []interface{} {
	return array
}

type FilterFunc func (array Array, index int, value interface{}) bool
type FindFunc func (array Array, index int, value interface{}) bool
type FindIndexFunc func (array Array, index int, value interface{}) bool
type ForEachFunc func (array Array, index int, value interface{})
type MapFunc func (array Array, index int, value interface{}) interface{}
type SortFunc func (a interface{}, b interface{}) int
