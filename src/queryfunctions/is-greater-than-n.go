package queryfunctions

import "filter"
import "reflect"

func IsGreaterThanN(fieldName string, number int) filter.FieldQuery {
  f := func(v reflect.Value) bool {
    val := int(v.Int())
    return val > number
  }
  return filter.FieldQuery { Field: fieldName, Compare: f }
}
