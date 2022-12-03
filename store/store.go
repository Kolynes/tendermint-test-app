package store

import "errors"

type Store map[string]any

func (store *Store) Get(key string) (any, error) {
	if !store.Has(key) {
		return nil, errors.New("not found")
	} else {
		return (*store)[key], nil
	}
}

func (store *Store) Has(key string) bool {
	return (*store)[key] != nil
}

func (store *Store) Put(key string, value any) {
	(*store)[key] = value
}
