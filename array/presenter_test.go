// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package array

import (
	"reflect"
	"testing"
)

func TestPresenter_AsFloat64Slice(t *testing.T) {
	floatSlice := []float64{0, 1, 0.55, 3.14}
	array := New(0, 1, 0.55, 3.14)
	values, err := array.Present().AsFloat64Slice()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(values, floatSlice) {
		t.Error("array.Present().AsFloat64Slice() does not have expected values")
		t.Errorf("Expecting %v, got %v", values, floatSlice)
	}
}

func TestPresenter_AsInt64Slice(t *testing.T) {
	intSlice := []int64{-1, 0, 1, 2}
	array := New(-1, 0, 1, 2)
	values, err := array.Present().AsInt64Slice()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(values, intSlice) {
		t.Error("array.Present().AsInt64Slice() does not have expected values")
		t.Errorf("Expecting %v, got %v", values, intSlice)
	}
}

func TestPresenter_AsStringSlice(t *testing.T) {
	stringSlice := []string{"a", "b", "c"}
	array := New("a", "b", "c")
	values, err := array.Present().AsStringSlice()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(values, stringSlice) {
		t.Error("array.Present().AsStringSlice() does not have expected values")
		t.Errorf("Expecting %v, got %v", values, stringSlice)
	}
}

func TestPresenter_AsUint64Slice(t *testing.T) {
	uintSlice := []uint64{0, 1, 2, 4, 5}
	array := New(0, 1, 2, 4, 5)
	values, err := array.Present().AsUint64Slice()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(values, uintSlice) {
		t.Error("array.Present().AsUint64Slice() does not have expected values")
		t.Errorf("Expecting %v, got %v", values, uintSlice)
	}
}