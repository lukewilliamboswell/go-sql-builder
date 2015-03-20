package sqlbuilder

import (
	"bytes"
	"fmt"
)

func AND(args ...string) string {
	query := bytes.NewBufferString("")

	for i, value := range args {
		if i > 0 {
			query.WriteString(` AND ` + value)
		} else {
			query.WriteString(value)
		}
	}

	return query.String()
}

func EQUAL_STRING(field string) string {
	return fmt.Sprintf(`%s = ?`, field)
}

func EQUAL_INT(field string) string {
	return fmt.Sprintf(`%s = ?nnn`, field)
}

func LIKE(field string) string {
	return fmt.Sprintf(`%s LIKE ?`, field)
}

func SELECT_FROM(table string, columns ...string) string {
	query := bytes.NewBufferString(`SELECT `)

	length := len(columns)
	for i, value := range columns {
		query.WriteString(`[` + value + `]`)

		if i < length-1 {
			query.WriteString(`,`)
		}
	}

	query.WriteString(` FROM [dbo].[` + table + `]`)

	return query.String()
}
