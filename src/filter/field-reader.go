package filter

import "reflect"

type FieldReader interface {
	Read(st interface{}, field string) reflect.Value
}

type fieldReader struct {
}

func NewFieldReader() FieldReader {
	return new(fieldReader)
}

func (f *fieldReader) Read(st interface{}, field string) reflect.Value {
	r := reflect.ValueOf(st)
	v := reflect.Indirect(r).FieldByName(field)
	return v
}
