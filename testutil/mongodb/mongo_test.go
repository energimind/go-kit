package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func TestMongoEnvironment_URI(t *testing.T) {
	t.Parallel()

	env := NewMongoEnvironment()

	cleanUp, err := env.Start()
	defer cleanUp()

	require.NotNil(t, cleanUp)
	require.NoError(t, err)

	require.NotEmpty(t, env.URI)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.URI))

	require.NotNil(t, client)
	require.NoError(t, err)
}
