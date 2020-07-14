// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package object

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
)

// The New() function creates a new Object.
func New() Object {
	return Object{}
}

// The NewFromJsonString() function parses a given JSON string, and returns
// the result as a new Object. It uses "encoding/json" package, so it will
// returns an InvalidUnmarshalError if v is nil or not a pointer.
func NewFromJsonString(v string) (Object, error) {
	var object map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(v))
	if err := decoder.Decode(&object); err != nil {
		return nil, err
	}
	return NewFromMap(object), nil
}

// The NewFromMap() function returns a new Object from a given map.
func NewFromMap(v map[string]interface{}) Object {
	return v
}

// Object defines a Object Type. See "object" package documentation for more
// information.
type Object map[string]interface{}

// The Clear() function removes all elements from the Object.
func (object Object) Clear() Object {
	for key, _ := range object {
		delete(object, key)
	}
	return object
}

// The Delete() function removes the specified element from the Object by key.
func (object Object) Delete(key string) Object {
	if object.Has(key) {
		delete(object, key)
	}
	return object
}

// The ForEach() function executes a provided function once for each Object element.
// Because Object does not remember the insertion order of each elements, for each
// element that is being iterated, the keys are sorted alphabetically.
//
// If the order of iteration does not matter, a common for range loop to the Object
// might be used. Because Object is built on top of Go native map, it will behave the
// same on a for range loop.
func (object Object) ForEach(f ForEachFunc) {
	for _, key := range object.Keys() {
		f(key, object[key])
	}
}

// The Get() function returns a specified element from the Object.
//
// If the element with the given key does not exist, it will returns nil.
func (object Object) Get(key string) interface{} {
	if !object.Has(key) {
		return nil
	}
	return object[key]
}

// The Has() function returns a boolean indicating whether an element with
// the specified key exists or not.
func (object Object) Has(key string) bool {
	_, exists := object[key]
	return exists
}

// The HasAll() function is the same as Has() function, but accepts a slice
// of keys instead of a single key string. If the object has all the element
// for each given key, it will returns true. Otherwise, it will returns
// false.
func (object Object) HasAll(keys ...string) bool {
	for _, k := range keys {
		if !object.Has(k) {
			return false
		}
	}
	return true
}

// The HasSome() function is the same as Has() function, but accepts a slice
// of keys instead of a single key string. If the object has one or more
// element for each given key, it will returns true. Otherwise, it will
// return false.
func (object Object) HasSome(keys ...string) bool {
	for _, k := range keys {
		if object.Has(k) {
			return true
		}
	}
	return false
}

// The Keys() function returns a slice of string that contains the keys
// for each element in the Object. Because an Object does not remember
// the insertion order of an element, the returned slice order will
// be sorted alphabetically.
func (object Object) Keys() []string {
	var keys []string
	for key, _ := range object {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// The Length() function returns the number of elements contained inside the
// Object.
func (object Object) Length() int {
	return len(object)
}

// The Present() function returns an Object Presenter, which capable to
// returns the the collection to other predefined data type.
func (object Object) Present() Presenter {
	return Presenter{ object }
}

// The Reflect() function returns the Object value of the given key as a
// reflect.Value data. If there is no element saved with the given key, it
// will returns reflect.Value of nil.
func (object Object) Reflect(key string) reflect.Value {
	if !object.Has(key) {
		return reflect.ValueOf(nil)
	}
	return reflect.ValueOf(object.Get(key))
}

// The Reflects() function returns the map[string]reflect.Value representation
// of the Object.
func (object Object) Reflects() map[string]reflect.Value {
	reflection := make(map[string]reflect.Value)
	for key, _ := range object {
		reflection[key] = object.Reflect(key)
	}
	return reflection
}

// The Set() function adds or updates an element with a specified key and value
// to the Object. Since the Set() function returns back the same Object, you
// can chain the function call.
//
// If an element with the specified key exists, it will update the value.
// Otherwise, it will add a new element based on the given key and value.
func (object Object) Set(key string, value interface{}) Object {
	object[key] = value
	return object
}

// The String() function returns a string representing the specified Object
// and its elements.
func (object Object) String() string {
	b, _ := json.Marshal(object)
	return string(b)
}

// The Values() function returns a slice of interface{} that contains the
// values for each element in the Object. Because an Object does not
// remember the insertion order of an element, the returned slice order
// will be ordered based on the keys that sorted alphabetically.
func (object Object) Values() []interface{} {
	var keys []string
	for key, _ := range object {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var values []interface{}
	for _, key := range keys {
		values = append(values, object[key])
	}
	return values
}

type ForEachFunc func (key string, value interface{})