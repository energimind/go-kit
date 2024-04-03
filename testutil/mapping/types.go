package mapping

import (
	"reflect"

	"github.com/stretchr/testify/require"
)

// testingT is an interface wrapper around *testing.T.
type testingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
	Helper()
}

// CheckAllFieldsAreMapped checks if all fields from the 'from' struct are mapped to the 'to' struct.
// It uses reflection to iterate over the fields in the 'from' struct
// and checks if each field is present in the 'to' struct.
func CheckAllFieldsAreMapped(t testingT, from any, to any) {
	t.Helper()

	fromType := reflect.TypeOf(from)
	toType := reflect.TypeOf(to)

	for i := range fromType.NumField() {
		field := fromType.Field(i)
		_, ok := toType.FieldByName(field.Name)

		require.True(t, ok, "field %s is not mapped", field.Name)
	}
}

// CheckAllEnumValuesAreMapped checks if all values from the 'from' enum are mapped to the 'to' enum.
// It converts the enums to slices and then checks if each value in the 'from' slice is present in the 'to' slice.
// It uses the 'mapper' function to convert the 'from' enum value to the 'to' enum value.
func CheckAllEnumValuesAreMapped[T, M comparable](t testingT, from []T, to []M, mapper func(T) M) {
	t.Helper()

	for _, value := range from {
		found := false

		for _, mappedValue := range to {
			if mapper(value) == mappedValue {
				found = true

				break
			}
		}

		require.True(t, found, "enum value %d is not mapped", value)
	}
}
