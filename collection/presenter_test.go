package collection

import (
	"fmt"
	"testing"
)

func TestPresenter_AsMap(t *testing.T) {
	pkg := "standard"
	detailName := "Go with Standard"
	detailDescription := "Standard objects in Go, redesigned."
	detail := New().
		Set("name", detailName).
		Set("description", detailDescription)
	version := 1.5
	collection := New().
		Set("pkg", pkg).
		Set("detailCollection", detail).
		Set("version", version).
		Set("isPublic", true)
	object := collection.Present().AsMap()
	detailStrShouldBe := fmt.Sprintf("map[description:%s name:%s]",
		detailDescription, detailName)
	objectStrShouldBe := fmt.Sprintf("map[detailCollection:%v isPublic:%v pkg:%v version:%v]",
		detailStrShouldBe,
		true,
		pkg,
		version)
	objectStr := fmt.Sprint(object)
	if objectStr != objectStrShouldBe {
		t.Error("collection.Present().AsMap() map value does not match Collection input")
		t.Errorf("Expecting %v, got %v", objectStrShouldBe, objectStr)
		return
	}
}
