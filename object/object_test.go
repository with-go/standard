package object

import (
	"fmt"
	"reflect"
	"testing"
)

// string type child-element
var pkg = "standard"
// Object type child-element
var detailName = "Go with Standard"
var detailDescription = "Standard objects in Go, redesigned."
var detailObject Object
var detailStr = fmt.Sprintf("{\"description\":\"%s\",\"name\":\"%s\"}", detailDescription, detailName)
// float64 type child-element
var version = 1.5
// int type child-element
var year = 2020
// the Collection for test
var testObject Object
var testObjectStr = fmt.Sprintf("{\"detail\":%v,\"isPublic\":%v,\"pkg\":\"%s\",\"version\":%v,\"year\":%v}",
	detailStr,
	true,
	pkg,
	version,
	year)

// Reset testObject variable to initial state
func resetTestObject() {
	detailObject = New().
		Set("name", detailName).
		Set("description", detailDescription)
	testObject = New().
		Set("pkg", pkg).
		Set("detail", detailObject).
		Set("version", version).
		Set("year", year).
		// bool type child-element
		Set("isPublic", true)
}

func TestNew(t *testing.T) {
	resetTestObject()
	object := testObject
	if object.Length() != len(object) {
		t.Error("New() Object length does not match")
		t.Errorf("Expecting %d, got %d", len(object), object.Length())
		return
	}
	if fmt.Sprint(object) != testObjectStr {
		t.Error("New() Object value does not match JSON string input")
		t.Errorf("Expecting %s, got %s", testObjectStr, fmt.Sprint(object))
		return
	}
}

func TestNewFromJsonString(t *testing.T) {
	object, err := NewFromJsonString(testObjectStr)
	if err != nil {
		t.Error("NewFromJsonString(v) JSON parsing error")
		t.Errorf("Reason: %s", err.Error())
		return
	}
	if object.Length() != 5 {
		t.Error("NewFromJsonString(v) Object length does not match")
		t.Errorf("Expecting %d, got %d", len(object), object.Length())
		return
	}
	if fmt.Sprint(object) != testObjectStr {
		t.Error("NewFromJsonString(v) Object value does not match JSON string input")
		t.Errorf("Expecting %s, got %s", testObjectStr, fmt.Sprint(object))
		return
	}
}

func TestNewFromMap(t *testing.T) {
	detail := make(map[string]string)
	detail["name"] = detailName
	detail["description"] = detailDescription
	data := make(map[string]interface{})
	data["pkg"] = pkg
	data["detail"] = detail
	data["version"] = version
	data["year"] = year
	data["isPublic"] = true
	object := NewFromMap(data)
	if object.Length() != 5 {
		t.Error("NewFromMap(v) Object length does not match")
		t.Errorf("Expecting %d, got %d", len(object), object.Length())
		return
	}
	if fmt.Sprint(object) != testObjectStr {
		t.Error("NewFromMap(v) Object value does not match map input")
		t.Errorf("Expecting %s, got %s", testObjectStr, fmt.Sprint(object))
		return
	}
}

func TestObject_Clear(t *testing.T) {
	resetTestObject()
	object := testObject
	object.Clear()
	if object.Length() != 0 {
		t.Error("object.Clear() length is not zero")
		t.Errorf("Expecting %v, got %v", 0, object.Length())
		return
	}
	if fmt.Sprint(object) != "{}" {
		t.Error("object.Clear() value does not match")
		t.Errorf("Expecting %v, got %v", "{}", fmt.Sprint(object))
		return
	}
}

func TestObject_Delete(t *testing.T) {
	resetTestObject()
	object := detailObject
	object.Delete("description")
	if object.Length() != 1 {
		t.Error("object.Delete(key) length does not match")
		t.Errorf("Expecting %v, got %v", 1, object.Length())
		return
	}
	objectStr := fmt.Sprintf("{\"name\":\"%s\"}", detailName)
	if fmt.Sprint(object) != objectStr {
		t.Error("object.Delete(key) value does not match")
		t.Errorf("Expecting %v, got %v", objectStr, fmt.Sprint(object))
		return
	}
}

func TestObject_ForEach(t *testing.T) {
	resetTestObject()
	object := testObject
	objectStr := ""
	object.ForEach(func(key string, value interface{}) {
		objectStr = objectStr + fmt.Sprintln(key + ":" + fmt.Sprint(value))
	})
	objectStrShouldBe := fmt.Sprintf("detail:%v\nisPublic:%v\npkg:%v\nversion:%v\nyear:%v\n",
		detailStr,
		true,
		pkg,
		version,
		year)
	if objectStr != objectStrShouldBe {
		t.Error("object.ForEach(function) iteration result does match expectation")
		t.Errorf("Expecting %v, got %v", objectStrShouldBe, objectStr)
		return
	}
}

func TestObject_Get(t *testing.T) {
	resetTestObject()
	object := testObject
	if object.Get("invalid") != nil {
		t.Error("object.Get(key) non-existent key return non-nil value")
		t.Errorf("Expecting %v, got %v", nil, object.Get("invalid"))
		return
	}
	if !reflect.DeepEqual(object.Get("detail"), detailObject) {
		t.Error("object.Get(key) value does not match")
		t.Errorf("Expecting %v, got %v", detailObject, object.Get("detail"))
		return
	}
}

func TestObject_Has(t *testing.T) {
	resetTestObject()
	object := testObject
	if object.Has("invalid") != false {
		t.Error("object.Has(key) returns true on non-existent key")
		t.Errorf("Expecting %v, got %v", false, object.Has("invalid"))
		return
	}
	if object.Has("detail") != true {
		t.Error("object.Has(key) returns false on existent key")
		t.Errorf("Expecting %v, got %v", true, object.Has("detail"))
		return
	}
}

func TestObject_HasAll(t *testing.T) {
	resetTestObject()
	object := testObject
	if object.HasAll("invalid", "detail") != false {
		t.Error("object.HasAll(keys) returns true on all non-existent key")
		t.Errorf("Expecting %v, got %v", false, object.HasAll("invalid", "detail"))
		return
	}
	if object.HasAll("detail", "pkg") != true {
		t.Error("object.HasAll(keys) returns false on all existent key")
		t.Errorf("Expecting %v, got %v", true, object.HasAll("detail", "pkg"))
		return
	}
}

func TestObject_HasSome(t *testing.T) {
	resetTestObject()
	object := testObject
	if object.HasSome("invalid", "not-exist") != false {
		t.Error("object.HasSome(keys) returns true on all non-existent key")
		t.Errorf("Expecting %v, got %v", false, object.HasSome("invalid", "not-exist"))
		return
	}
	if object.HasSome("invalid", "detail") != true {
		t.Error("object.HasSome(keys) returns false on some existent key")
		t.Errorf("Expecting %v, got %v", false, object.HasSome("invalid", "detail"))
		return
	}
	if object.HasSome("detail", "pkg") != true {
		t.Error("object.HasSome(keys) returns false on all existent key")
		t.Errorf("Expecting %v, got %v", true, object.HasSome("detail", "pkg"))
		return
	}
}

func TestObject_Keys(t *testing.T) {
	resetTestObject()
	object := testObject
	keys := []string{"detail", "isPublic", "pkg", "version", "year"}
	if len(object.Keys()) != len(keys) {
		t.Error("object.Keys() length does not match")
		t.Errorf("Expecting %v, got %v", len(keys), len(object.Keys()))
		return
	}
	if !reflect.DeepEqual(object.Keys(), keys) {
		t.Error("object.Keys() slice values do not match")
		t.Errorf("Expecting %v, got %v", keys, object.Keys())
		return
	}
	if fmt.Sprint(object.Keys()) != fmt.Sprint(keys) {
		t.Error("object.Keys() slice String does not match")
		t.Errorf("Expecting %v, got %v", fmt.Sprint(keys), fmt.Sprint(object.Keys()))
		return
	}
}

func TestObject_Length(t *testing.T) {
	resetTestObject()
	object := testObject
	if object.Length() != len(object) {
		t.Error("object.Length() length does not match")
		t.Errorf("Expecting %v, got %v", len(object), object.Length())
		return
	}
}

func TestObject_Present(t *testing.T) {
	resetTestObject()
	object := testObject
	presenter := Presenter{object}
	if !reflect.DeepEqual(object.Present(), presenter) {
		t.Error("object.Present() Presenter does not deeply equal")
		t.Errorf("Expecting %v, got %v", true, false)
		return
	}
}

func TestObject_Reflect(t *testing.T) {
	resetTestObject()
	object := testObject
	if object.Reflect("pkg").Kind() != reflect.String {
		t.Error("object.Reflect(key) Value.Kind() does not deeply equal")
		t.Errorf("Expecting %v, got %v", reflect.String, object.Reflect("pkg").Kind())
		return
	}
}

func TestObject_Reflects(t *testing.T) {
	resetTestObject()
	object := testObject
	if len(object.Reflects()) != object.Length() {
		t.Error("object.Reflects() length does not match")
		t.Errorf("Expecting %v, got %v", object.Length(), len(object.Reflects()))
		return
	}
}

func TestObject_Set(t *testing.T) {
	child := New()
	object := New()
	expecting := ""
	// Test if the key-value pair was added
	object.Set("alpha", "checked")
	expecting = "{\"alpha\":\"checked\"}"
	if fmt.Sprint(object) != expecting {
		t.Error("object.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(object))
		return
	}
	// Test if the element with same key got doubled
	object.Set("alpha", "checked")
	expecting = "{\"alpha\":\"checked\"}"
	if fmt.Sprint(object) != expecting {
		t.Error("object.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(object))
		return
	}
	// Test more add() call to check the insertion order
	object.Set("child", child)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s}",
		"{}")
	if fmt.Sprint(object) != expecting {
		t.Error("object.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(object))
		return
	}
	object.Set("retry", 1)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s,\"retry\":1}",
		"{}")
	if fmt.Sprint(object) != expecting {
		t.Error("object.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(object))
		return
	}
	child.Set("test", "beta")
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s,\"retry\":1}",
		"{\"test\":\"beta\"}")
	if fmt.Sprint(object) != expecting {
		t.Error("object.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(object))
		return
	}
	// Re-add child element to see if it was added in the correct insertion order
	object.Set("child", child)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s,\"retry\":1}",
		"{\"test\":\"beta\"}")
	if fmt.Sprint(object) != expecting {
		t.Error("object.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(object))
		return
	}
}

func TestObject_String(t *testing.T) {
	resetTestObject()
	object := testObject
	objectStr := fmt.Sprint(object)
	if objectStr != testObjectStr {
		t.Error("object.String() does not return the expected string")
		t.Errorf("Expecting %s, got %s", objectStr, testObjectStr)
	}
}

func TestObject_Values(t *testing.T) {
	resetTestObject()
	object := testObject
	values := []interface{}{detailObject, true, pkg, version, year}
	if len(object.Values()) != len(values) {
		t.Error("object.Values() length does not match")
		t.Errorf("Expecting %v, got %v", len(values), len(object.Values()))
		return
	}
	if !reflect.DeepEqual(object.Values(), values) {
		t.Error("object.Values() slice values do not match")
		t.Errorf("Expecting %v, got %v", values, object.Values())
		return
	}
	if fmt.Sprint(object.Values()) != fmt.Sprint(values) {
		t.Error("object.Values() slice String does not match")
		t.Errorf("Expecting %v, got %v", fmt.Sprint(values), fmt.Sprint(object.Values()))
		return
	}
}
