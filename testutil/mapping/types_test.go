package mapping

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckAllFieldsAreMapped(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		s1 := struct {
			A int
			B string
		}{}

		s2 := struct {
			A int
			B string
		}{}

		CheckAllFieldsAreMapped(t, s1, s2)
	})

	t.Run("fail", func(t *testing.T) {
		s1 := struct {
			A int
			B string
		}{}

		s2 := struct {
			A int
		}{}

		require.PanicsWithValue(t, "field B is not mapped", func() {
			CheckAllFieldsAreMapped(&testingTMock{}, s1, s2)
		})
	})
}

func TestCheckAllEnumValuesAreMapped(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		from := []int{1, 2, 3}
		to := []int{11, 12, 13}

		mapper := func(i int) int {
			return i + 10
		}

		CheckAllEnumValuesAreMapped(t, from, to, mapper)
	})

	t.Run("fail", func(t *testing.T) {
		from := []int{1, 2, 3}
		to := []int{1, 2}

		mapper := func(i int) int {
			return i
		}

		require.PanicsWithValue(t, "enum value 3 is not mapped", func() {
			CheckAllEnumValuesAreMapped(&testingTMock{}, from, to, mapper)
		})
	})
}
