package sqlbuilder

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}, message string) {
	if a != b {
		t.Logf(`Expected "%v" (type %v) - Got "%v" (type %v): %s`, b, reflect.TypeOf(b), a, reflect.TypeOf(a), message)
		t.Fail()
	}
}

func TestAND(t *testing.T) {

	args := []string{`field = 1`, `a.field2 LIKE 'apples'`, `field3 IN(12,23,34,45)`}

	result := AND(args...)

	expected := `field = 1 AND a.field2 LIKE 'apples' AND field3 IN(12,23,34,45)`

	expect(t, result, expected, "should correctly build AND query")

}

func TestEQUAL_STRING(t *testing.T) {
	args := []struct {
		field    string
		expected string
	}{
		{field: `field1`, expected: ` [field1] = ? `},
		{field: `customerID`, expected: ` [customerID] = ? `},
	}

	for _, a := range args {
		result := EQUAL_STRING(a.field)
		expect(t, result, a.expected, "should correctly build EQUAL query")
	}
}

func TestEQUAL_INT(t *testing.T) {
	args := []struct {
		field    string
		expected string
	}{
		{field: `field1`, expected: ` [field1] = ?nnn `},
		{field: `a.customerID`, expected: ` [a.customerID] = ?nnn `},
	}

	for _, a := range args {
		result := EQUAL_INT(a.field)
		expect(t, result, a.expected, "should correctly build EQUAL query")
	}
}

func TestLIKE(t *testing.T) {
	args := []struct {
		field    string
		expected string
	}{
		{field: `field1`, expected: ` [field1] LIKE ? `},
		{field: `a.customerID`, expected: ` [a.customerID] LIKE ? `},
	}

	for _, a := range args {
		result := LIKE(a.field)
		expect(t, result, a.expected, "should correctly build LIKE query")
	}
}

func TestIN_INT(t *testing.T) {
	args := []struct {
		field    string
		args     []int64
		expected string
	}{
		{field: `field1`, args: []int64{1, 2, 3}, expected: ` [field1] IN(1,2,3) `},
	}

	for _, a := range args {
		result := IN_INT(a.field, a.args...)
		expect(t, result, a.expected, "should correctly build LIKE query")
	}
}

func TestCOMPOSE(t *testing.T) {
	args := []struct {
		args     []string
		expected string
	}{
		{args: []string{`a`, `BIG`, `h.ello`}, expected: `aBIGh.ello`},
	}

	for _, a := range args {
		result := COMPOSE(a.args...)
		expect(t, result, a.expected, "should correctly build COMPOSE query")
	}
}

func TestSELECT_FROM(t *testing.T) {
	args := []struct {
		table    string
		columns  []string
		expected string
	}{
		{table: `Customer`, columns: []string{`CustomerID`, `Name`}, expected: `SELECT [CustomerID],[Name] FROM [dbo].[Customer]`},
	}

	for _, a := range args {
		result := SELECT_FROM(a.table, a.columns...)
		expect(t, result, a.expected, "should correctly build LIKE query")
	}
}

func TestORDER_BY(t *testing.T) {
	args := []struct {
		offset, rows int64
		fields       []string
		expected     string
	}{
		{offset: 0, rows: 10, fields: []string{`CustomerID`, `Name`}, expected: ` ORDER BY [CustomerID],[Name] OFFSET 0 ROWS FETCH NEXT 10 ROWS ONLY `},
		{offset: 4000, rows: 0, fields: []string{`SupplierID`}, expected: ` ORDER BY [SupplierID] OFFSET 4000 ROWS `},
	}

	for _, a := range args {
		result := ORDER_BY(a.offset, a.rows, a.fields...)
		expect(t, result, a.expected, "should correctly build ORDER_BY query")
	}
}
