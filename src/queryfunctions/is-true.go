package queryfunctions

import "filter"
import "reflect"
import "strings"

func IsTrue(fieldName string) filter.FieldQuery {
	fieldName = strings.Title(fieldName)
	f := func(v reflect.Value) bool {
		val := v.Bool()
		return val == true
	}
	return filter.FieldQuery{Field: fieldName, Compare: f}
}
