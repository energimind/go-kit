package crud

import (
	"context"
	"strconv"
	"testing"
)

func TestRunTests(t *testing.T) {
	t.Parallel()

	repo := newTestRepo()

	RunTests(t, Setup[testEntity, int]{
		RepoOps: RepoOps[testEntity, int]{
			GetAll: func(ctx context.Context) ([]testEntity, error) {
				return repo.getAll()
			},
			GetByID: func(ctx context.Context, id int) (testEntity, error) {
				return repo.getByID(id)
			},
			Create: func(ctx context.Context, entity testEntity) error {
				return repo.create(entity)
			},
			Update: func(ctx context.Context, entity testEntity) error {
				return repo.update(entity)
			},
			Delete: func(ctx context.Context, id int) error {
				return repo.delete(id)
			},
		},
		EntityOps: EntityOps[testEntity, int]{
			NewEntity: func(key int) testEntity {
				return testEntity{
					id:    key,
					value: "test-" + strconv.Itoa(key),
				}
			},
			ModifyEntity: func(entity testEntity) testEntity {
				entity.value = "test-modified"

				return entity
			},
			UnboundEntity: func() testEntity {
				return testEntity{
					value: "test-unbound",
				}
			},
			ExtractKey: func(entity testEntity) int {
				return entity.id
			},
			MissingKey: func() int {
				return -1
			},
		},
		NotFoundErr: func() any {
			return testError{}
		},
	})
}

type testEntity struct {
	id    int
	value string
}

type testError struct{}

func (t testError) Error() string {
	return "test-error"
}

type testRepo struct {
	entities  map[int]testEntity
	idCounter int
}

func newTestRepo() *testRepo {
	return &testRepo{
		entities: make(map[int]testEntity),
	}
}

func (r *testRepo) getAll() ([]testEntity, error) {
	all := make([]testEntity, 0, len(r.entities))

	for _, entity := range r.entities {
		all = append(all, entity)
	}

	return all, nil
}

func (r *testRepo) getByID(id int) (testEntity, error) {
	entity, ok := r.entities[id]

	if !ok {
		return testEntity{}, testError{}
	}

	return entity, nil
}

func (r *testRepo) create(entity testEntity) error {
	r.idCounter++

	entity.id = r.idCounter
	r.entities[entity.id] = entity

	return nil
}

func (r *testRepo) update(entity testEntity) error {
	if _, ok := r.entities[entity.id]; !ok {
		return testError{}
	}

	r.entities[entity.id] = entity

	return nil
}

func (r *testRepo) delete(id int) error {
	if _, ok := r.entities[id]; !ok {
		return testError{}
	}

	delete(r.entities, id)

	return nil
}
