// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package collection

import (
	"reflect"
)

type Presenter struct {
	collection *Collection
}

// The AsMap() function will parses the Collection and returns a string map of
// interface{} based on the elements inside the Collection. As a map will not
// remember the insertion order (as how a Collection does remember the insertion
// order) the element order of the returned Object will not predictable and may
// not be the same as the Collection.
func (presenter Presenter) AsMap() map[string]interface{} {
	object := make(map[string]interface{})
	for _, pair := range presenter.collection.pairs {
		object[pair.key] = pair.value
		if t := reflect.TypeOf(pair.value); t.Kind() == reflect.Ptr && t.Elem().Name() == "Collection" {
			element := pair.value.(*Collection)
			object[pair.key] = element.Present().AsMap()
		}
	}
	return object
}