package sqlbuilder

import (
	"bytes"
	"fmt"
)

func COMPOSE(args ...string) string {
	var buffer bytes.Buffer

	for _, value := range args {
		buffer.WriteString(value)
	}

	return buffer.String()
}

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
	return fmt.Sprintf(` [%s] = ? `, field)
}

func EQUAL_INT(field string) string {
	return fmt.Sprintf(` [%s] = ?nnn `, field)
}

func LIKE(field string) string {
	return fmt.Sprintf(` [%s] LIKE ? `, field)
}

func IN_INT(field string, args ...int64) string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf(` [%s] IN(`, field))

	length := len(args)
	for i, value := range args {
		buffer.WriteString(fmt.Sprintf(`%d`, value))

		if i < length-1 {
			buffer.WriteString(`,`)
		}
	}

	buffer.WriteString(`) `)

	return buffer.String()

}

func ORDER_BY(offset, rows int64, fields ...string) string {
	query := bytes.NewBufferString(` ORDER BY `)

	length := len(fields)
	for i, value := range fields {
		query.WriteString(`[` + value + `]`)

		if i < length-1 {
			query.WriteString(`,`)
		}
	}

	if rows > 0 {
		query.WriteString(fmt.Sprintf(` OFFSET %d ROWS FETCH NEXT %d ROWS ONLY `, offset, rows))
	} else {
		query.WriteString(fmt.Sprintf(` OFFSET %d ROWS `, offset))
	}

	return query.String()
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
