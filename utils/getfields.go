package utils

import (
	"database/sql/driver"
	"errors"
	"reflect"
)

// ColumnDataPair describes a piece of data that is stored in a database table column.
type ColumnDataPair struct {
	Column string
	Data   interface{}
}

type ListColumnDataPair []ColumnDataPair

func (l ListColumnDataPair) ToMap() map[string]any {
	m := make(map[string]any, 0)
	for idx := range l {
		m[l[idx].Column] = l[idx].Data
	}
	return m
}

// GetFields returns an array of ColumnDataPair.
// It uses the tag struct tab to get the column name.
func GetFields(data interface{}, tag string) (ListColumnDataPair, error) {
	if tag == "" {
		return nil, errors.New("tag should not be empty")
	}
	var row []ColumnDataPair
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		col := field.Tag.Get(tag)
		if col == "" {
			col = field.Name
		}

		val, err := driver.DefaultParameterConverter.ConvertValue(v.Field(i).Interface())
		if err != nil {
			return row, err
		}

		row = append(row, ColumnDataPair{
			Column: col,
			Data:   val,
		})
	}

	return row, nil
}
