package archives

import (
	"database/sql"
	"fmt"
	"reflect"
)


func getTable[T any](rows *sql.Rows) (out []T) {
	var table []T
	for rows.Next() {
		var data T
		s := reflect.ValueOf(&data).Elem()
		numCols := s.NumField()
		columns := make([]any, numCols)

		for i := range numCols {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		if err := rows.Scan(columns...); err != nil {
			fmt.Println("Case Read Error ", err)
		}

		table = append(table, data)
	}
	return table
}
