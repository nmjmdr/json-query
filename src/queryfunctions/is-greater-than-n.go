package queryfunctions

import "filter"
import "reflect"
import "strings"

func IsGreaterThanN(fieldName string, number int) filter.FieldQuery {
	fieldName = strings.Title(fieldName)
	f := func(v reflect.Value) bool {
		val := int(v.Int())
		return val > number
	}
	return filter.FieldQuery{Field: fieldName, Compare: f}
}
