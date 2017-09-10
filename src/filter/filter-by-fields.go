package filter

import "reflect"

type GetFieldValue func(fieldName string) reflect.Value

type FieldsFilter interface {
	Filter(value interface{}, fieldQueries []FieldQuery, getFieldValue GetFieldValue) bool
}

type fieldsFilter struct {
}

func NewFieldsFilter() FieldsFilter {
	return new(fieldsFilter)
}

func (f *fieldsFilter) Filter(value interface{}, fieldQueries []FieldQuery, getFieldValue GetFieldValue) bool {
	for _, query := range fieldQueries {
		value := getFieldValue(query.Field)
		if !query.Compare(value) {
			return false
		}
	}
	return true
}
