package crud

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// RepoOps contains the operations that operate on the repository.
type RepoOps[T, K any] struct {
	GetAll  func(ctx context.Context) ([]T, error)
	GetByID func(ctx context.Context, id K) (T, error)
	Create  func(ctx context.Context, model T) error
	Update  func(ctx context.Context, model T) error
	Delete  func(ctx context.Context, id K) error
}

// EntityOps contains the operations that do not operate on the repository, but on the entity itself.
type EntityOps[T, K any] struct {
	NewEntity     func(key int) T // create a new entity with the given key
	ModifyEntity  func(model T) T // modify the given entity
	UnboundEntity func() T        // entity with no key
	ExtractKey    func(T) K       // key extractor
	MissingKey    func() K        // key that does not exist
}

// Setup is the setup for the CRUD tests.
type Setup[T, K any] struct {
	RepoOps     RepoOps[T, K]
	EntityOps   EntityOps[T, K]
	NotFoundErr func() any // return not found error instance
}

// RunTests runs the CRUD tests for the given type.
func RunTests[T, K any](t *testing.T, setup Setup[T, K]) { //nolint:funlen
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	keyCounter := 0

	nextKey := func() int {
		keyCounter++

		return keyCounter
	}

	t.Run("missing-notFound", func(t *testing.T) {
		t.Run("getByID", func(t *testing.T) {
			_, err := setup.RepoOps.GetByID(ctx, setup.EntityOps.MissingKey())

			notFoundErr := setup.NotFoundErr()

			require.ErrorAs(t, err, &notFoundErr)
		})

		t.Run("update", func(t *testing.T) {
			notFoundErr := setup.NotFoundErr()

			require.ErrorAs(t, setup.RepoOps.Update(ctx, setup.EntityOps.UnboundEntity()), &notFoundErr)
		})

		t.Run("delete", func(t *testing.T) {
			notFoundErr := setup.NotFoundErr()

			require.ErrorAs(t, setup.RepoOps.Delete(ctx, setup.EntityOps.MissingKey()), &notFoundErr)
		})
	})

	t.Run("getAll-empty", func(t *testing.T) {
		all1, err := setup.RepoOps.GetAll(ctx)

		require.NoError(t, err)
		require.Empty(t, all1)
	})

	ent := setup.EntityOps.NewEntity(nextKey())
	key := setup.EntityOps.ExtractKey(ent)

	t.Run("create", func(t *testing.T) {
		require.NoError(t, setup.RepoOps.Create(ctx, ent))
	})

	t.Run("getAll-foundOne", func(t *testing.T) {
		all1, err := setup.RepoOps.GetAll(ctx)

		require.NoError(t, err)
		require.Equal(t, []T{ent}, all1)
	})

	t.Run("getByID-found", func(t *testing.T) {
		e2, err := setup.RepoOps.GetByID(ctx, key)

		require.NoError(t, err)
		require.Equal(t, ent, e2)
	})

	entMod := setup.EntityOps.ModifyEntity(ent)

	require.Equal(t, key, setup.EntityOps.ExtractKey(entMod), "key should not change")
	require.NotEqual(t, ent, entMod, "entity should change")

	t.Run("update", func(t *testing.T) {
		require.NoError(t, setup.RepoOps.Update(ctx, entMod))
	})

	t.Run("getByID-updated", func(t *testing.T) {
		fetched, err := setup.RepoOps.GetByID(ctx, key)

		require.NoError(t, err)
		require.Equal(t, entMod, fetched)
	})

	t.Run("delete", func(t *testing.T) {
		require.NoError(t, setup.RepoOps.Delete(ctx, key))
	})

	t.Run("delete-again-notFound", func(t *testing.T) {
		notFoundErr := setup.NotFoundErr()

		require.ErrorAs(t, setup.RepoOps.Delete(ctx, key), &notFoundErr)
	})

	t.Run("getAll-empty", func(t *testing.T) {
		all2, err := setup.RepoOps.GetAll(ctx)

		require.NoError(t, err)
		require.Empty(t, all2)
	})
}
