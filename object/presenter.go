// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package object

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

type Presenter struct {
	object Object
}

var (
	NonPtrTypeError = errors.New("the given parameter v is a non-pointer type, " +
		"parameter v should be pointer to a struct")
	PtrToNonStructTypeError = errors.New("the given parameter v is a pointer type but not pointed to a struct, " +
		"parameter v should be pointer to a struct")
)

// The AsStruct() function will parses the Object and stores the result in
// the value pointed to by v. Because it using "encoding/json" package, it
// returns an InvalidUnmarshalError if v is nil or not a pointer.
func (presenter Presenter) AsStruct(v interface{}) error {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Ptr {
		return NonPtrTypeError
	}
	if t.Elem().Kind() != reflect.Struct {
		return PtrToNonStructTypeError
	}
	decoder := json.NewDecoder(strings.NewReader(presenter.object.String()))
	return decoder.Decode(v)
}