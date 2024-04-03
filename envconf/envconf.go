package envconf

import (
	"fmt"
	"reflect"
	"strconv"

	envs "github.com/caarlos0/env/v7"
)

type parserFunc func(v any, funcMap map[reflect.Type]envs.ParserFunc, opts ...envs.Options) error

// Parse parses the environment variables into the given struct.
//
// It uses the default parsers from caarlos0/env, and adds some custom parsers
// for bool values.
func Parse(v any) error {
	return parse(v, envs.ParseWithFuncs, nil)
}

func parse(v any, parser parserFunc, env map[string]string) error {
	var opts []envs.Options

	if env != nil {
		opts = append(opts, envs.Options{Environment: env})
	}

	if err := parser(v, customParsers(), opts...); err != nil {
		return fmt.Errorf("parse env: %w", err)
	}

	return nil
}

func customParsers() map[reflect.Type]envs.ParserFunc {
	return map[reflect.Type]envs.ParserFunc{
		reflect.TypeOf(true): func(v string) (any, error) {
			return parseBool(v)
		},
	}
}

//nolint:wrapcheck // don't clutter the error
func parseBool(v string) (bool, error) {
	switch v {
	case "on", "yes":
		return true, nil
	case "off", "no":
		return false, nil
	default:
		return strconv.ParseBool(v)
	}
}
