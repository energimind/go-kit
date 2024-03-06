package envconf

import (
	"errors"
	"reflect"
	"testing"

	envs "github.com/caarlos0/env/v7"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()

	val := struct {
		Value string
	}{}

	err := Parse(&val)

	require.NoError(t, err)
}

func Test_parse(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		val := struct {
			Value  string `env:"TEST_KEY"`
			Toggle bool   `env:"TOGGLE"`
			Active bool   `env:"ACTIVE"`
		}{}

		env := map[string]string{
			"TEST_KEY": "test-value-1",
			"TOGGLE":   "on",
			"ACTIVE":   "yes",
		}

		err := parse(&val, envs.ParseWithFuncs, env)

		require.NoError(t, err)
		require.Equal(t, "test-value-1", val.Value)
		require.Equal(t, true, val.Toggle)
		require.Equal(t, true, val.Active)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		val := struct{}{}

		forcedError := errors.New("forced-error")

		err := parse(&val, func(v any, funcMap map[reflect.Type]envs.ParserFunc, opts ...envs.Options) error {
			return forcedError
		}, nil)

		require.ErrorIs(t, err, forcedError)
	})
}

func Test_parseBool(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		value string
		want  bool
	}{
		"true": {
			value: "true",
			want:  true,
		},
		"false": {
			value: "false",
			want:  false,
		},
		"on": {
			value: "on",
			want:  true,
		},
		"off": {
			value: "off",
			want:  false,
		},
		"yes": {
			value: "yes",
			want:  true,
		},
		"no": {
			value: "no",
			want:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseBool(tt.value)

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
