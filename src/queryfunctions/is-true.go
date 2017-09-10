package queryfunctions

import "filter"
import "reflect"

func IsTrue(fieldName string) filter.FieldQuery {
  f := func(v reflect.Value) bool {
    val := v.Bool()
    return val == true
  }
  return filter.FieldQuery { Field: fieldName, Compare: f }
}
