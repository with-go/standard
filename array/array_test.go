// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	array := New(1, "a", false)
	if len(array) != 3 {
		t.Error("New(values) failed to instantiate")
		return
	}
}

func TestNewFromSlice(t *testing.T) {
	slice := []int{ 10, 100, 1000 }
	array, err := NewFromSlice(slice)
	if err != nil {
		t.Error(err)
	}
	if len(array) != 3 {
		t.Error("NewFromSlice(v) failed to instantiate")
	}
	if array[0] != slice[0] || array[1] != slice[1] || array[2] != slice[2] {
		t.Error("NewFromSlice(v) does not match with the given slice")
	}
}

func TestArray_Concat(t *testing.T) {
	array1 := New("a", "b", "c")
	array2 := New("d", "e", "f")
	concatArray := New("a", "b", "c", "d", "e", "f")
	resultArray := array1.Concat(array2)
	if len(resultArray) != len(concatArray) {
		t.Error("array.Concat(other) length does not match")
		t.Errorf("Expecting %d, got %d", len(concatArray), len(resultArray))
	}
	for i, v := range resultArray {
		if v != concatArray[i] {
			t.Error("array.Concat(other) values do not match")
			t.Errorf("Expecting %s, got %s", concatArray, resultArray)
			break
		}
	}
}

func TestArray_DeepEqual(t *testing.T) {
	array := New("a", "b", "c")
	other, _ := NewFromSlice([]string{"a", "b", "c"})
	broken := New("d", "e", "f")
	otherEqual := array.DeepEqual(other)
	brokenEqual := array.DeepEqual(broken)
	if !otherEqual {
		t.Error("array.DeepEqual(other) equality check failed")
		t.Errorf("Expecting %v, got %v", true, otherEqual)
		return
	}
	if brokenEqual {
		t.Error("array.DeepEqual(other) equality check failed")
		t.Errorf("Expecting %v, got %v", false, brokenEqual)
		return
	}
}

func TestArray_Equal(t *testing.T) {
	array := New("a", "b", "c")
	other, _ := NewFromSlice([]string{"a", "b", "c"})
	broken := New("d", "e", "f")
	otherEqual := array.Equal(other)
	brokenEqual := array.Equal(broken)
	if !otherEqual {
		t.Error("array.Equal(other) equality check failed")
		t.Errorf("Expecting %v, got %v", true, otherEqual)
		return
	}
	if brokenEqual {
		t.Error("array.Equal(other) equality check failed")
		t.Errorf("Expecting %v, got %v", false, brokenEqual)
		return
	}
}

func TestArray_Filter(t *testing.T) {
	array := New("spray", "limit", "elite", "exuberant", "destruction", "present")
	filteredArray := New("exuberant", "destruction", "present")
	resultArray := array.Filter(func(array Array, index int, value interface{}) bool {
		return len(value.(string)) > 6
	})
	if len(resultArray) != len(filteredArray) {
		t.Error("array.Filter(function) length does not match")
		t.Errorf("Expecting %d, got %d", len(filteredArray), len(resultArray))
	}
	for i, v := range resultArray {
		if v != filteredArray[i] {
			t.Error("array.Filter(function) values do not match")
			t.Errorf("Expecting %s, got %s", filteredArray, resultArray)
			break
		}
	}
}

func TestArray_Find(t *testing.T) {
	shouldBe := 12
	array := New(5, 12, 8, 130, 44)
	found := array.Find(func(array Array, index int, value interface{}) bool {
		return value.(int) > 10
	}).(int)
	if found != shouldBe {
		t.Error("array.Find(function) value does not match")
		t.Errorf("Expecting %d, got %d", shouldBe, found)
	}
}

func TestArray_FindIndex(t *testing.T) {
	shouldBe := 3
	array := New(5, 12, 8, 130, 44)
	foundIndex := array.FindIndex(func(array Array, index int, value interface{}) bool {
		return value.(int) > 15
	})
	if foundIndex != shouldBe {
		t.Error("array.FindIndex(function) index does not match")
		t.Errorf("Expecting %d, got %d", shouldBe, foundIndex)
	}
}

func TestArray_ForEach(t *testing.T) {
	shouldBe := "abc"
	array := New("a", "b", "c")
	str := ""
	array.ForEach(func(array Array, index int, value interface{}) {
		str += fmt.Sprintf("%v", value)
	})
	if str != shouldBe {
		t.Error("array.ForEach(function) iteration does not yield the expected result")
		t.Errorf("Expecting %v, got %v", shouldBe, str)
	}
}

func TestArray_Includes(t *testing.T) {
	array1 := New(1, 2, 3)
	isInclude1 := array1.Includes(2)
	if isInclude1 != true {
		t.Error("array.Includes(value) does not have a value that does exist inside array")
		t.Errorf("Expecting %v, got %v", true, isInclude1)
	}
	array2 := New("cat", "dog", "bat")
	isInclude2 := array2.Includes("at")
	if isInclude2 != false {
		t.Error("array.Includes(value) does have a value that does not exist inside array")
		t.Errorf("Expecting %v, got %v", false, isInclude2)
	}
}

func TestArray_IndexOf(t *testing.T) {
	array := New("ant", "bison", "camel", "duck", "bison")
	indexOfBison := array.IndexOf("bison")
	bisonShouldBe := 1
	if indexOfBison != bisonShouldBe {
		t.Error("array.IndexOf(value) index does not match")
		t.Errorf("Expecting %v, got %v", bisonShouldBe, indexOfBison)
	}
	indexOfGiraffe := array.IndexOf("giraffe")
	giraffeShouldBe := -1
	if indexOfGiraffe != giraffeShouldBe {
		t.Error("array.IndexOf(value) have index of a value that does not exist inside array")
		t.Errorf("Expecting %v, got %v", giraffeShouldBe, indexOfGiraffe)
	}
}

func TestArray_Join(t *testing.T) {
	shouldBe1 := "AirWaterEarthFire"
	shouldBe2 := "Air-Water-Earth-Fire"
	array := New("Air", "Water", "Earth", "Fire")
	joinedStr1 := array.Join("")
	if joinedStr1 != shouldBe1 {
		t.Error("array.Join(separator) does not match expected string")
		t.Errorf("Expecting %v, got %v", shouldBe1, joinedStr1)
	}
	joinedStr2 := array.Join("-")
	if joinedStr2 != shouldBe2 {
		t.Error("array.Join(separator) does not match expected string")
		t.Errorf("Expecting %v, got %v", shouldBe2, joinedStr2)
	}
}

func TestArray_Keys(t *testing.T) {
	keySlice := []int{0, 1, 2}
	array := New("a", "b", "c")
	if !reflect.DeepEqual(array.Keys(), keySlice) {
		t.Error("array.Keys() does not have expected keys")
		t.Errorf("Expecting %v, got %v", keySlice, array.Keys())
	}
}

func TestArray_LastIndexOf(t *testing.T) {
	array := New("ant", "bison", "camel", "duck", "bison")
	lastIndexOfCamel := array.LastIndexOf("camel", -1)
	camelShouldBe := 2
	if lastIndexOfCamel != camelShouldBe {
		t.Error("array.LastIndexOf(value) index does not match")
		t.Errorf("Expecting %v, got %v", camelShouldBe, lastIndexOfCamel)
	}
	lastIndexOfBison := array.LastIndexOf("bison", -1)
	bisonShouldBe := 4
	if lastIndexOfBison != bisonShouldBe {
		t.Error("array.LastIndexOf(value) does not match")
		t.Errorf("Expecting %v, got %v", bisonShouldBe, lastIndexOfBison)
	}
	lastIndexOfGiraffe := array.LastIndexOf("giraffe", -1)
	giraffeShouldBe := -1
	if lastIndexOfGiraffe != giraffeShouldBe {
		t.Error("array.LastIndexOf(value) have index of a value that does not exist inside array")
		t.Errorf("Expecting %v, got %v", giraffeShouldBe, lastIndexOfGiraffe)
	}
}

func TestArray_Length(t *testing.T) {
	array := New("ant", "bison", "camel", "duck", "bison")
	length := array.Length()
	shouldBe := len(array)
	if length != shouldBe {
		t.Error("array.Length() does not match")
		t.Errorf("Expecting %v, got %v", shouldBe, length)
	}
}

func TestArray_Map(t *testing.T) {
	squaredArray := New(2, 8, 18, 32)
	array := New(1, 4, 9, 16)
	resultArray := array.Map(func(array Array, index int, value interface{}) interface{} {
		return value.(int) * 2
	})
	if len(resultArray) != len(squaredArray) {
		t.Error("array.Map(function) length does not match")
		t.Errorf("Expecting %d, got %d", len(squaredArray), len(resultArray))
	}
	for i, v := range resultArray {
		if v != squaredArray[i] {
			t.Error("array.Map(function) values do not match")
			t.Errorf("Expecting %s, got %s", squaredArray, resultArray)
			break
		}
	}
}

func TestArray_Pop(t *testing.T) {
	poppedArray := New("broccoli", "cauliflower", "cabbage", "kale")
	array := New("broccoli", "cauliflower", "cabbage", "kale", "tomato")
	resultArray, lastValue := array.Pop()
	shouldBe := "tomato"
	if lastValue.(string) != shouldBe {
		t.Error("array.Pop() value does not match")
		t.Errorf("Expecting %s, got %s", shouldBe, lastValue.(string))
	}
	if len(resultArray) != len(poppedArray) {
		t.Error("array.Pop() length does not match")
		t.Errorf("Expecting %d, got %d", len(poppedArray), len(resultArray))
	}
	for i, v := range resultArray {
		if v != poppedArray[i] {
			t.Error("array.Pop() values do not match")
			t.Errorf("Expecting %s, got %s", poppedArray, resultArray)
			break
		}
	}
}

func TestArray_Present(t *testing.T) {
	array := New("a", "b", "c")
	presenter := array.Present()
	if len(presenter.array) != len(array) {
		t.Error("array.Present() returns an Array Presenter which have an array, but the array length does not match")
		t.Errorf("Expecting %d, got %d", len(array), len(presenter.array))
	}
	for i, v := range presenter.array {
		if v != array[i] {
			t.Error("array.Present() returns an Array Presenter which have an array, but the array values do not match")
			t.Errorf("Expecting %s, got %s", array, presenter.array)
			break
		}
	}
}

func TestArray_Push(t *testing.T) {
	pushedArray1 := New("pigs", "goats", "sheep", "cows")
	pushedArray2 := New("pigs", "goats", "sheep", "cows", "chickens", "cats", "dogs")
	array := New("pigs", "goats", "sheep")
	resultArray1 := array.Push("cows")
	if len(resultArray1) != len(pushedArray1) {
		t.Error("array.Push() length does not match")
		t.Errorf("Expecting %d, got %d", len(pushedArray1), len(resultArray1))
	}
	for i, v := range resultArray1 {
		if v != pushedArray1[i] {
			t.Error("array.Push() values do not match")
			t.Errorf("Expecting %s, got %s", pushedArray1, resultArray1)
			break
		}
	}
	resultArray2 := resultArray1.Push("chickens", "cats", "dogs")
	if len(resultArray2) != len(pushedArray2) {
		t.Error("array.Push(values) length does not match")
		t.Errorf("Expecting %d, got %d", len(pushedArray2), len(resultArray2))
	}
	for i, v := range resultArray2 {
		if v != pushedArray2[i] {
			t.Error("array.Push(values) values do not match")
			t.Errorf("Expecting %s, got %s", pushedArray2, resultArray2)
			break
		}
	}
}

func TestArray_Reflect(t *testing.T) {
	child := New("Standard", "with", "Go")
	array := New("Hello!", true, 3.14, -10, &child)
	value := array.Reflect(0)
	kind := value.Kind()
	if kind != reflect.String {
		t.Error("array.Reflect(index) reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.String, kind)
		return
	}
	if value.String() != array[0].(string) {
		t.Error("array.Reflect(index) reflect.Value does not match")
		t.Errorf("Expecting %s, got %s", array[0].(string), value.String())
		return
	}
	value = array.Reflect(1)
	kind = value.Kind()
	if kind != reflect.Bool {
		t.Error("array.Reflect(index) reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.Bool, kind)
		return
	}
	if value.Bool() != array[1].(bool) {
		t.Error("array.Reflect(index) reflect.Value does not match")
		t.Errorf("Expecting %v, got %v", array[1].(bool), value.Bool())
		return
	}
	value = array.Reflect(2)
	kind = value.Kind()
	if kind != reflect.Float64 {
		t.Error("array.Reflect(index) reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.Float64, kind)
		return
	}
	if value.Float() != array[2].(float64) {
		t.Error("array.Reflect(index) reflect.Value does not match")
		t.Errorf("Expecting %v, got %v", array[2].(float64), value.Float())
		return
	}
	value = array.Reflect(3)
	kind = value.Kind()
	if kind != reflect.Int {
		t.Error("array.Reflect(index) reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.Int64, kind)
		return
	}
	if value.Int() != int64(array[3].(int)) {
		t.Error("array.Reflect(index) reflect.Value does not match")
		t.Errorf("Expecting %v, got %v", int64(array[3].(int)), value.Int())
		return
	}
	value = array.Reflect(4)
	kind = value.Kind()
	if kind != reflect.Ptr && value.Elem().Kind() != reflect.Slice {
		t.Error("array.Reflect(index) reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.Ptr, kind)
		t.Errorf("Expecting %s, got %s", reflect.Slice, value.Elem().Kind())
		return
	}
	value = value.Elem()
	slice := make([]interface{}, value.Len())
	for index := 0; index < value.Len(); index++ {
		slice[index] = value.Index(index).Interface()
	}
	if !child.DeepEqual(slice) {
		t.Error("array.Reflect(index) reflect.Value does not match")
		t.Errorf("Expecting %v, got %v", child, slice)
		return
	}
}

func TestArray_Reflects(t *testing.T) {
	child := New("Standard", "with", "Go")
	array := New("Hello!", true, 3.14, -10, &child)
	reflections := array.Reflects()
	kind, value := reflections[0].Kind(), reflections[0]
	if kind != reflect.String {
		t.Error("array.Reflects() reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.String, kind)
		return
	}
	if value.String() != array[0].(string) {
		t.Error("array.Reflects() reflect.Value does not match")
		t.Errorf("Expecting %s, got %s", array[0].(string), value.String())
		return
	}
	kind, value = reflections[1].Kind(), reflections[1]
	if kind != reflect.Bool {
		t.Error("array.Reflects() reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.Bool, kind)
		return
	}
	if value.Bool() != array[1].(bool) {
		t.Error("array.Reflects() reflect.Value does not match")
		t.Errorf("Expecting %v, got %v", array[1].(bool), value.Bool())
		return
	}
	kind, value = reflections[2].Kind(), reflections[2]
	if kind != reflect.Float64 {
		t.Error("array.Reflects() reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.Float64, kind)
		return
	}
	if value.Float() != array[2].(float64) {
		t.Error("array.Reflects() reflect.Value does not match")
		t.Errorf("Expecting %v, got %v", array[2].(float64), value.Float())
		return
	}
	kind, value = reflections[3].Kind(), reflections[3]
	if kind != reflect.Int {
		t.Error("array.Reflects() reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.Int64, kind)
		return
	}
	if value.Int() != int64(array[3].(int)) {
		t.Error("array.Reflects() reflect.Value does not match")
		t.Errorf("Expecting %v, got %v", int64(array[3].(int)), value.Int())
		return
	}
	kind, value = reflections[4].Kind(), reflections[4]
	if kind != reflect.Ptr && value.Elem().Kind() != reflect.Slice {
		t.Error("array.Reflects() reflect.Kind does not match")
		t.Errorf("Expecting %s, got %s", reflect.Ptr, kind)
		t.Errorf("Expecting %s, got %s", reflect.Slice, value.Elem().Kind())
		return
	}
	value = value.Elem()
	slice := make([]interface{}, value.Len())
	for index := 0; index < value.Len(); index++ {
		slice[index] = value.Index(index).Interface()
	}
	if !child.DeepEqual(slice) {
		t.Error("array.Reflects() reflect.Value does not match")
		t.Errorf("Expecting %v, got %v", child, slice)
		return
	}
}

func TestArray_Reverse(t *testing.T) {
	reversedArray := New("three", "two", "one")
	array := New("one", "two", "three")
	resultArray := array.Reverse()
	if len(resultArray) != len(reversedArray) {
		t.Error("array.Reverse() length does not match")
		t.Errorf("Expecting %d, got %d", len(reversedArray), len(resultArray))
	}
	for i, v := range resultArray {
		if v != reversedArray[i] {
			t.Error("array.Reverse() values do not match")
			t.Errorf("Expecting %s, got %s", reversedArray, resultArray)
			break
		}
	}
}

func TestArray_Shift(t *testing.T) {
	shouldBe := 1
	shiftedArray := New(2, 3)
	array := New(1, 2, 3)
	resultArray, firstValue := array.Shift()
	if firstValue.(int) != shouldBe {
		t.Error("array.Shift() value does not match")
		t.Errorf("Expecting %d, got %d", shouldBe, firstValue.(int))
	}
	if len(resultArray) != len(shiftedArray) {
		t.Error("array.Shift() length does not match")
		t.Errorf("Expecting %d, got %d", len(shiftedArray), len(resultArray))
	}
	for i, v := range resultArray {
		if v != shiftedArray[i] {
			t.Error("array.Shift() values do not match")
			t.Errorf("Expecting %s, got %s", shiftedArray, resultArray)
			break
		}
	}
}

func TestArray_Sort(t *testing.T) {
	array1 := New(4, 2, 5, 1, 3)
	ascArray1 := New(1, 2, 3, 4, 5)
	descArray1 := New(5, 4, 3, 2, 1)
	array2 := New("d", "c", "a", "e", "f", "b")
	ascArray2 := New("a", "b", "c", "d", "e", "f")
	resultAscArray1 := array1
	resultAscArray1.Sort(func(a interface{}, b interface{}) int {
		return a.(int) - b.(int)
	})
	for i, v := range resultAscArray1 {
		if v != ascArray1[i] {
			t.Error("Sorted array does not match ascending int array")
			t.Errorf("Expecting %s, got %s", ascArray1, resultAscArray1)
			break
		}
	}
	resultDescArray1 := array1
	resultDescArray1.Sort(func(a interface{}, b interface{}) int {
		if  a.(int) > b.(int) {
			return -1
		} else {
			return 1
		}
	})
	for i, v := range resultDescArray1 {
		if v != descArray1[i] {
			t.Error("Sorted array does not match descending int array")
			t.Errorf("Expecting %s, got %s", descArray1, resultDescArray1)
			break
		}
	}
	resultAscArray2 := array2
	resultAscArray2.Sort(func(a interface{}, b interface{}) int {
		if  a.(string) > b.(string) {
			return 1
		} else {
			return -1
		}
	})
	for i, v := range resultAscArray2 {
		if v != ascArray2[i] {
			t.Error("Sorted array does not match ascending string array")
			t.Errorf("Expecting %s, got %s", ascArray2, resultAscArray2)
			break
		}
	}
}

func TestArray_String(t *testing.T) {
	array := New(1, 2, "a", "new array")
	str := fmt.Sprint(array)
	shouldBe := "[1,2,\"a\",\"new array\"]"
	if str != shouldBe {
		t.Error("array.String() does not return the expected string")
		t.Errorf("Expecting %s, got %s", shouldBe, str)
	}
}

func TestArray_Unshift(t *testing.T) {
	unshiftArray := New(4, 5, 1, 2, 3)
	array := New(1, 2, 3)
	resultArray := array.Unshift(4, 5)
	if len(resultArray) != len(unshiftArray) {
		t.Error("array.Unshift() length does not match")
		t.Errorf("Expecting %d, got %d", len(unshiftArray), len(resultArray))
	}
	for i, v := range resultArray {
		if v != unshiftArray[i] {
			t.Error("array.Unshift() values do not match")
			t.Errorf("Expecting %s, got %s", unshiftArray, resultArray)
			break
		}
	}
}

func TestArray_Values(t *testing.T) {
	valueSlice := []interface{}{"a", "b", "c"}
	array := New("a", "b", "c")
	values := array.Values()
	if !reflect.DeepEqual(values, valueSlice) {
		t.Error("array.Values() does not have expected values")
		t.Errorf("Expecting %v, got %v", values, valueSlice)
	}
}