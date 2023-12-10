package utils

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// BuildSearchQuery looks for fields with non-zero value into data and build sql query
func BuildSearchQuery(data any) (query string, values []any, err error) {
	list, err := GetFields(data, "db")
	if err != nil {
		log.Printf("failed to build search query: %v", err)
		return
	}

	queryLocal := make([]string, 0)
	colsMap := list.ToMap()
	for col, d := range colsMap {
		if reflect.ValueOf(d).IsZero() {
			continue
		}

		queryLocal = append(queryLocal, fmt.Sprintf(" %s = ? ", col))
		values = append(values, colsMap[col])
	}

	query = strings.Join(queryLocal, " AND ")
	return
}
