package mongodb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMongoEnvironment_MultipleDatabases(t *testing.T) {
	t.Parallel()

	env := NewMongoEnvironment()

	require.NotPanics(t, func() {
		cleanUp, err := env.Start()
		defer cleanUp()

		require.NotNil(t, cleanUp)
		require.NoError(t, err)

		db1, closer1 := env.NewInstance()
		defer closer1()

		require.NotNil(t, db1)

		db2, closer2 := env.NewInstance()
		defer closer2()

		require.NotNil(t, db2)

		require.NotEqual(t, db1, db2)
	})
}
