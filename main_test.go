package main

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
		{field: `field1`, expected: `field1 = ?`},
		{field: `customerID`, expected: `customerID = ?`},
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
		{field: `field1`, expected: `field1 = ?nnn`},
		{field: `a.customerID`, expected: `a.customerID = ?nnn`},
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
		{field: `field1`, expected: `field1 LIKE ?`},
		{field: `a.customerID`, expected: `a.customerID LIKE ?`},
	}

	for _, a := range args {
		result := LIKE(a.field)
		expect(t, result, a.expected, "should correctly build LIKE query")
	}
}
