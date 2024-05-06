package redis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInstance(t *testing.T) {
	inst, closer, err := NewInstance()
	defer closer()

	require.NoError(t, err)
	require.NotNil(t, inst)
	require.NotNil(t, inst.Client)
	require.NotNil(t, inst.Address)
}
