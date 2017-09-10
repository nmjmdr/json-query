package test

import (
	"testing"
	"filter"
)

func Test_When_FieldReader_Is_Given_A_Struct_It_Should_Get_The_Value(t *testing.T) {
  f := filter.NewFieldReader()
  type Person struct {
    Name string
  }
  p := Person{Name: "Hitchens"}
  value := f.Read(&p, "Name")
  if value.String() != "Hitchens" {
    t.Fatal("Should have read the field value")
  }
}
