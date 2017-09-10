package filter

import "reflect"

type ItemsFilter interface {
	IsMatch(items interface{}, fieldQueries []FieldQuery) bool
}

type itemsFilter struct {
	itemType interface{}
}

func NewItemsFilter(itemType interface{}) ItemsFilter {
	i := new(itemsFilter)
	i.itemType = itemType
	return i
}

func (i *itemsFilter) IsMatch(item interface{}, fieldQueries []FieldQuery) bool {
	fieldReader := NewFieldReader()
	fieldsFilter := NewFieldsFilter()
	return fieldsFilter.IsMatch(item, fieldQueries, func(fieldName string) reflect.Value {
		return fieldReader.Read(item, fieldName)
	})
}
