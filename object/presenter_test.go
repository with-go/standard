package object

import (
	"fmt"
	"testing"
)

type detailStructure struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type dataStructure struct {
	Pkg      	string			`json:"pkg"`
	Detail   	detailStructure	`json:"detail"`
	Version  	float64			`json:"version"`
	Year		int64			`json:"year"`
	IsPublic 	bool			`json:"isPublic"`
}

func TestPresenter_AsStruct(t *testing.T) {
	pkg := "standard"
	detailName := "Go with Standard"
	detailDescription := "Standard objects in Go, redesigned."
	detail := New().
		Set("name", detailName).
		Set("description", detailDescription)
	version := 1.5
	year := 2020
	object := New().
		Set("pkg", pkg).
		Set("detail", detail).
		Set("version", version).
		Set("year", year).
		Set("isPublic", true)
	var data dataStructure
	err := object.Present().AsStruct(&data)
	if err != nil {
		t.Error("object.Present().AsStruct(v) parsing error")
		t.Errorf("Reason: %v", err)
		return
	}
	strShouldBe := "{standard {Go with Standard Standard objects in Go, redesigned.} 1.5 2020 true}"
	if fmt.Sprint(data) != strShouldBe {
		t.Error("object.Present().AsStruct(v) String() value not match")
		t.Errorf("Expecting %s, got %s", strShouldBe, fmt.Sprint(data))
		return
	}
}
