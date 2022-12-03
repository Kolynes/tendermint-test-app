package store

type Batch struct {
	*Store
	parent *Store
}

func NewBatch(store *Store) *Batch {
	batch := &Batch{
		parent: store,
		Store:  &Store{},
	}
	batch.parent = store
	for key, value := range *store {
		(*batch.Store)[key] = value
	}
	return batch
}

func (batch *Batch) Commit() {
	for key, value := range *batch.Store {
		(*batch.parent)[key] = value
	}
}
