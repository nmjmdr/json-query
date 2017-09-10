package filter

import "encoding/json"
import "io"
import "reflect"

type ItemsFilter interface {
	Filter(reader io.Reader, fieldQueries []FieldQuery) ([]interface{}, error)
}

type itemsFilter struct {
	itemType interface{}
}

func NewItemsFilter(itemType interface{}) ItemsFilter {
	i := new(itemsFilter)
	i.itemType = itemType
	return i
}

func (i *itemsFilter) Filter(reader io.Reader, fieldQueries []FieldQuery) ([]interface{}, error) {
	decoder := json.NewDecoder(reader)

	_, err := decoder.Token()
	if err != nil {
		return nil, err
	}
	items := make([]interface{},0)

	fieldReader := NewFieldReader()
	fieldsFilter := NewFieldsFilter()

	for decoder.More() {
		err = decoder.Decode(i.itemType)
		if err != nil {
			return nil, err
		}
		if(fieldsFilter.Filter(i.itemType, fieldQueries, func(fieldName string) reflect.Value {
			return fieldReader.Read(i.itemType, fieldName)
		})){
			items = append(items, i.itemType)
		}
	}
	_, err = decoder.Token()
	if err != nil {
		return nil, err
	}
	return items, nil
}
