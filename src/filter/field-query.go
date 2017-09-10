package filter

import "reflect"

type Compare func(reflect.Value) bool

type FieldQuery struct {
	Field   string
	Compare Compare
}
