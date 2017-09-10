package filter

import "encoding/json"
import "io"
import "reflect"

type ItemsFilter interface {
	Filter(reader io.Reader, fieldQueries []FieldQuery) error
}

type OnFound func(interface{})

type itemsFilter struct {
	itemType interface{}
	onFound  OnFound
}

func NewItemsFilter(itemType interface{}, onFound OnFound) ItemsFilter {
	i := new(itemsFilter)
	i.itemType = itemType
	i.onFound = onFound
	return i
}

func (i *itemsFilter) Filter(reader io.Reader, fieldQueries []FieldQuery) error {
	decoder := json.NewDecoder(reader)

	_, err := decoder.Token()
	if err != nil {
		return err
	}

	fieldReader := NewFieldReader()
	fieldsFilter := NewFieldsFilter()

	for decoder.More() {
		err = decoder.Decode(i.itemType)
		if err != nil {
			return err
		}
		if fieldsFilter.Filter(i.itemType, fieldQueries, func(fieldName string) reflect.Value {
			return fieldReader.Read(i.itemType, fieldName)
		}) {
			i.onFound(i.itemType)
		}
	}
	_, err = decoder.Token()
	if err != nil {
		return err
	}
	return nil
}
