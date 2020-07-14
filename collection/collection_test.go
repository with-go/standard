package collection

import (
	"fmt"
	"reflect"
	"testing"
)

// string type child-element
var pkg = "standard"
// Collection type child-element
var detailName = "Go with Standard"
var detailDescription = "Standard objects in Go, redesigned."
var detailCollection *Collection
var detailStr = fmt.Sprintf("{\"name\":\"%s\",\"description\":\"%s\"}", detailName, detailDescription)
var detailJsonStr = fmt.Sprintf("{\"description\":\"%s\",\"name\":\"%s\"}", detailDescription, detailName)
// float64 type child-element
var version = 1.5
// int type child-element
var year = 2020
// the Collection for test
var testCollection *Collection
var testCollectionStr = fmt.Sprintf("{\"pkg\":\"%s\",\"detail\":%v,\"version\":%v,\"year\":%v,\"isPublic\":%v}",
	pkg,
	detailStr,
	version,
	year,
	true)
var testCollectionJsonStr = fmt.Sprintf("{\"detail\":%v,\"isPublic\":%v,\"pkg\":\"%s\",\"version\":%v,\"year\":%v}",
	detailJsonStr,
	true,
	pkg,
	version,
	year)

// Reset testCollection variable to initial state
func resetTestCollection() {
	detailCollection = New().
		Set("name", detailName).
		Set("description", detailDescription)
	testCollection = New().
		Set("pkg", pkg).
		Set("detail", detailCollection).
		Set("version", version).
		Set("year", year).
		// bool type child-element
		Set("isPublic", true)
}

func TestNew(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if collection.Length() != 5 {
		t.Error("New() Collection length does not match")
		t.Errorf("Expecting %d, got %d", 5, collection.Length())
		return
	}
	if fmt.Sprint(collection) != testCollectionStr {
		t.Error("New() Collection value does not match JSON string input")
		t.Errorf("Expecting %s, got %s", testCollectionStr, fmt.Sprint(collection))
		return
	}
}

func TestNewFromJsonString(t *testing.T) {
	collection, err := NewFromJsonString(testCollectionStr)
	if err != nil {
		t.Error("NewFromJsonString(v) JSON parsing error")
		t.Errorf("Reason: %s", err.Error())
		return
	}
	if collection.Length() != 5 {
		t.Error("NewFromJsonString(v) Collection length does not match")
		t.Errorf("Expecting %d, got %d", 5, collection.Length())
		return
	}
	if fmt.Sprint(collection) != testCollectionJsonStr {
		t.Error("NewFromJsonString(v) Collection value does not match JSON string input")
		t.Errorf("Expecting %s, got %s", testCollectionJsonStr, fmt.Sprint(collection))
		return
	}
}

func TestNewFromMap(t *testing.T) {
	detail := make(map[string]string)
	detail["name"] = detailName
	detail["description"] = detailDescription
	object := make(map[string]interface{})
	object["pkg"] = pkg
	object["detail"] = detail
	object["version"] = version
	object["year"] = year
	object["isPublic"] = true
	collection, err := NewFromMap(object)
	if err != nil {
		t.Error("NewFromMap(v) failed to create Collection from map")
		t.Errorf("Reason: %s", err.Error())
		return
	}
	if collection.Length() != 5 {
		t.Error("NewFromMap(v) Collection length does not match")
		t.Errorf("Expecting %d, got %d", 5, collection.Length())
		return
	}
	if fmt.Sprint(collection) != testCollectionJsonStr {
		t.Error("NewFromMap(v) Collection value does not match map input")
		t.Errorf("Expecting %s, got %s", testCollectionJsonStr, fmt.Sprint(collection))
		return
	}
}

func TestCollection_Add(t *testing.T) {
	child := New()
	collection := New()
	expecting := ""
	// Test if the key-value pair was added
	collection.Add("alpha", "checked")
	expecting = "{\"alpha\":\"checked\"}"
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Add() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	// Test if the element with same key got doubled
	collection.Add("alpha", "checked")
	expecting = "{\"alpha\":\"checked\"}"
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Add() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	// Test more add() call to check the insertion order
	collection.Add("child", child)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s}",
		"{}")
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Add() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	collection.Add("retry", 1)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s,\"retry\":1}",
		"{}")
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Add() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	child.Add("test", "beta")
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s,\"retry\":1}",
		"{\"test\":\"beta\"}")
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Add() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	// Re-add child element to see if it was added in the correct insertion order
	collection.Add("child", child)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"retry\":1,\"child\":%s}",
		"{\"test\":\"beta\"}")
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Add() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
}

func TestCollection_Clear(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	collection.Clear()
	if collection.Length() != 0 {
		t.Error("collection.Clear() length is not zero")
		t.Errorf("Expecting %v, got %v", 0, collection.Length())
		return
	}
	if fmt.Sprint(collection) != "{}" {
		t.Error("collection.Clear() value does not match")
		t.Errorf("Expecting %v, got %v", "{}", fmt.Sprint(collection))
		return
	}
}

func TestCollection_Delete(t *testing.T) {
	resetTestCollection()
	collection := detailCollection
	collection.Delete("description")
	if collection.Length() != 1 {
		t.Error("collection.Delete(key) length does not match")
		t.Errorf("Expecting %v, got %v", 1, collection.Length())
		return
	}
	collectionStr := fmt.Sprintf("{\"name\":\"%s\"}", detailName)
	if fmt.Sprint(collection) != collectionStr {
		t.Error("collection.Delete(key) value does not match")
		t.Errorf("Expecting %v, got %v", collectionStr, fmt.Sprint(collection))
		return
	}
}

func TestCollection_ForEach(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	collectionStr := ""
	collection.ForEach(func(key string, value interface{}) {
		collectionStr = collectionStr + fmt.Sprintln(key + ":" + fmt.Sprint(value))
	})
	collectionStrShouldBe := fmt.Sprintf("pkg:%v\ndetail:%v\nversion:%v\nyear:%v\nisPublic:%v\n",
		pkg,
		detailStr,
		version,
		year,
		true)
	if collectionStr != collectionStrShouldBe {
		t.Error("collection.ForEach(function) iteration result does match expectation")
		t.Errorf("Expecting %v, got %v", collectionStrShouldBe, collectionStr)
		return
	}
}

func TestCollection_Get(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if collection.Get("invalid") != nil {
		t.Error("collection.Get(key) non-existent key return non-nil value")
		t.Errorf("Expecting %v, got %v", nil, collection.Get("invalid"))
		return
	}
	if collection.Get("detail") != detailCollection {
		t.Error("collection.Get(key) value does not match")
		t.Errorf("Expecting %v, got %v", detailCollection, collection.Get("detail"))
		return
	}
}

func TestCollection_Has(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if collection.Has("invalid") != false {
		t.Error("collection.Has(key) returns true on non-existent key")
		t.Errorf("Expecting %v, got %v", false, collection.Has("invalid"))
		return
	}
	if collection.Has("detail") != true {
		t.Error("collection.Has(key) returns false on existent key")
		t.Errorf("Expecting %v, got %v", true, collection.Has("detail"))
		return
	}
}

func TestCollection_HasAll(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if collection.HasAll("invalid", "detail") != false {
		t.Error("collection.HasAll(keys) returns true on all non-existent key")
		t.Errorf("Expecting %v, got %v", false, collection.HasAll("invalid", "detail"))
		return
	}
	if collection.HasAll("detail", "pkg") != true {
		t.Error("collection.HasAll(keys) returns false on all existent key")
		t.Errorf("Expecting %v, got %v", true, collection.HasAll("detail", "pkg"))
		return
	}
}

func TestCollection_HasSome(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if collection.HasSome("invalid", "not-exist") != false {
		t.Error("collection.HasSome(keys) returns true on all non-existent key")
		t.Errorf("Expecting %v, got %v", false, collection.HasSome("invalid", "not-exist"))
		return
	}
	if collection.HasSome("invalid", "detail") != true {
		t.Error("collection.HasSome(keys) returns false on some existent key")
		t.Errorf("Expecting %v, got %v", false, collection.HasSome("invalid", "detail"))
		return
	}
	if collection.HasSome("detail", "pkg") != true {
		t.Error("collection.HasSome(keys) returns false on all existent key")
		t.Errorf("Expecting %v, got %v", true, collection.HasSome("detail", "pkg"))
		return
	}
}

func TestCollection_IndexOf(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if collection.IndexOf("invalid") != -1 {
		t.Error("collection.IndexOf(key) returns index on a non-existent key")
		t.Errorf("Expecting %v, got %v", -1, collection.IndexOf("invalid"))
		return
	}
	if collection.IndexOf("detail") != 1 {
		t.Error("collection.IndexOf(key) returns invalid index")
		t.Errorf("Expecting %v, got %v", 1, collection.IndexOf("detail"))
		return
	}
}

func TestCollection_Keys(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	keys := []string{"pkg", "detail", "version", "year", "isPublic"}
	if len(collection.Keys()) != len(keys) {
		t.Error("collection.Keys() length does not match")
		t.Errorf("Expecting %v, got %v", len(keys), len(collection.Keys()))
		return
	}
	if !reflect.DeepEqual(collection.Keys(), keys) {
		t.Error("collection.Keys() slice values do not match")
		t.Errorf("Expecting %v, got %v", keys, collection.Keys())
		return
	}
	if fmt.Sprint(collection.Keys()) != fmt.Sprint(keys) {
		t.Error("collection.Keys() slice String does not match")
		t.Errorf("Expecting %v, got %v", fmt.Sprint(keys), fmt.Sprint(collection.Keys()))
		return
	}
}

func TestCollection_Length(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if collection.Length() != 5 {
		t.Error("collection.Length() length does not match")
		t.Errorf("Expecting %v, got %v", 5, collection.Length())
		return
	}
}

func TestCollection_PairOf(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	pkgPair := Pair{"pkg", pkg}
	versionPair := Pair{"version", version}
	if !reflect.DeepEqual(*collection.PairOf("pkg"), pkgPair) {
		t.Error("collection.PairOf(key) Pair does not deeply equal")
		t.Errorf("Expecting %v, got %v", true, false)
		return
	}
	if !reflect.DeepEqual(*collection.PairOf("version"), versionPair) {
		t.Error("collection.PairOf(key) Pair does not deeply equal")
		t.Errorf("Expecting %v, got %v", true, false)
		return
	}
}

func TestCollection_Present(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	presenter := Presenter{collection}
	if !reflect.DeepEqual(collection.Present(), presenter) {
		t.Error("collection.Present() Presenter does not deeply equal")
		t.Errorf("Expecting %v, got %v", true, false)
		return
	}
}

func TestCollection_Reflect(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if collection.Reflect("pkg").Kind() != reflect.String {
		t.Error("collection.Reflect(key) Value.Kind() does not deeply equal")
		t.Errorf("Expecting %v, got %v", reflect.String, collection.Reflect("pkg").Kind())
		return
	}
}

func TestCollection_Reflects(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	if len(collection.Reflects()) != collection.Length() {
		t.Error("collection.Reflects() length does not match")
		t.Errorf("Expecting %v, got %v", collection.Length(), len(collection.Reflects()))
		return
	}
}

func TestCollection_Set(t *testing.T) {
	child := New()
	collection := New()
	expecting := ""
	// Test if the key-value pair was added
	collection.Set("alpha", "checked")
	expecting = "{\"alpha\":\"checked\"}"
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	// Test if the element with same key got doubled
	collection.Set("alpha", "checked")
	expecting = "{\"alpha\":\"checked\"}"
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	// Test more add() call to check the insertion order
	collection.Set("child", child)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s}",
		"{}")
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	collection.Set("retry", 1)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s,\"retry\":1}",
		"{}")
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	child.Set("test", "beta")
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s,\"retry\":1}",
		"{\"test\":\"beta\"}")
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
	// Re-add child element to see if it was added in the correct insertion order
	collection.Set("child", child)
	expecting = fmt.Sprintf("{\"alpha\":\"checked\",\"child\":%s,\"retry\":1}",
		"{\"test\":\"beta\"}")
	if fmt.Sprint(collection) != expecting {
		t.Error("collection.Set() value does not match")
		t.Errorf("Expecting %v, got %v", expecting, fmt.Sprint(collection))
		return
	}
}

func TestCollection_String(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	collectionStr := fmt.Sprint(collection)
	if collectionStr != testCollectionStr {
		t.Error("collection.String() does not return the expected string")
		t.Errorf("Expecting %s, got %s", testCollectionStr, collectionStr)
	}
}

func TestCollection_Values(t *testing.T) {
	resetTestCollection()
	collection := testCollection
	values := []interface{}{pkg, detailCollection, version, year, true}
	if len(collection.Values()) != len(values) {
		t.Error("collection.Values() length does not match")
		t.Errorf("Expecting %v, got %v", len(values), len(collection.Values()))
		return
	}
	if !reflect.DeepEqual(collection.Values(), values) {
		t.Error("collection.Values() slice values do not match")
		t.Errorf("Expecting %v, got %v", values, collection.Values())
		return
	}
	if fmt.Sprint(collection.Values()) != fmt.Sprint(values) {
		t.Error("collection.Values() slice String does not match")
		t.Errorf("Expecting %v, got %v", fmt.Sprint(values), fmt.Sprint(collection.Values()))
		return
	}
}
