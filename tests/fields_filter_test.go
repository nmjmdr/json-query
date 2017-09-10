package test

import (
	"testing"
	"filter"
  "queryfunctions"
	"reflect"
)

func Test_When_FieldsFilter_Is_Given_An_Item_It_Should_Match_If_Field_Queries_Match(t *testing.T) {
  f := filter.NewFieldsFilter()
  type Person struct {
    IsReal bool
    Age int
  }
  p := Person{IsReal: true, Age: 22}
  fieldQueryIsExists := queryfunctions.IsTrue("isReal")
  fieldQueryAge := queryfunctions.IsGreaterThanN("age", 20)
  didMatch := f.IsMatch(&p, []filter.FieldQuery{fieldQueryIsExists, fieldQueryAge}, func(fieldName string) reflect.Value {
		fieldReader := filter.NewFieldReader()
		return fieldReader.Read(&p,fieldName)
  })
  if !didMatch {
    t.Fatal("Should have matched the item")
  }
}

func Test_When_FieldsFilter_Is_Given_An_Item_It_Should_Not_Match_If_Even_One_Field_Querty_Does_Not_Match(t *testing.T) {
  f := filter.NewFieldsFilter()
  type Person struct {
    IsReal bool
    Age int
  }
  p := Person{IsReal: true, Age: 22}
  fieldQueryIsExists := queryfunctions.IsTrue("isReal")
  fieldQueryAge := queryfunctions.IsGreaterThanN("age", 30)
  didMatch := f.IsMatch(&p, []filter.FieldQuery{fieldQueryIsExists, fieldQueryAge}, func(fieldName string) reflect.Value {
		fieldReader := filter.NewFieldReader()
		return fieldReader.Read(&p,fieldName)
  })
  if didMatch {
    t.Fatal("Should NOT have matched the item")
  }
}
