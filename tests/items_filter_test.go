package test

import (
	"testing"
	"filter"
  "queryfunctions"
)

func Test_When_ItemsFilter_Is_Given_An_Item_It_Should_Match_If_Field_Queries_Match(t *testing.T) {
  type Person struct {
    IsReal bool
    Age int
  }
  p := Person{IsReal: true, Age: 22}
  f := filter.NewItemsFilter(&p)
  fieldQueryIsExists := queryfunctions.IsTrue("isReal")
  fieldQueryAge := queryfunctions.IsGreaterThanN("age", 20)
  didMatch := f.IsMatch(&p, []filter.FieldQuery{fieldQueryIsExists, fieldQueryAge})
  if !didMatch {
    t.Fatal("Should have matched the item")
  }
}

func Test_When_ItemsFilter_Is_Given_An_Item_It_Should_Not_Match_If_Any_Field_Queries_Do_Not_Match(t *testing.T) {
  type Person struct {
    IsReal bool
    Age int
  }
  p := Person{IsReal: true, Age: 22}
  f := filter.NewItemsFilter(&p)
  fieldQueryIsExists := queryfunctions.IsTrue("isReal")
  fieldQueryAge := queryfunctions.IsGreaterThanN("age", 30)
  didMatch := f.IsMatch(&p, []filter.FieldQuery{fieldQueryIsExists, fieldQueryAge})
  if didMatch {
    t.Fatal("Should NOT have matched the item")
  }
}
