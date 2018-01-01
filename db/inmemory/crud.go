package inmemory

import (
	"errors"

	"github.com/turtledev/refract-api/db"
)

var (
	errInvalidID = errors.New("ID must be non-zero")
)

// abstractStore is container type for implementing a CRUD interface.
// please note that since the concrete types are not known by the container, any
// validation needs to be done by the calling code.
type abstractStore struct {
	store  map[uint64]interface{}
	prevID uint64
}

func (r *abstractStore) GetAll() []interface{} {
	var dataset []interface{}
	for _, data := range r.store {
		dataset = append(dataset, data)
	}
	return dataset
}

func (r *abstractStore) Get(id uint64) (interface{}, error) {
	data, ok := r.store[id]
	if !ok {
		return nil, db.ErrNotFound
	}
	return data, nil
}

func (r *abstractStore) Delete(id uint64) error {
	if _, ok := r.store[id]; !ok {
		return db.ErrNotFound
	}
	delete(r.store, id)
	return nil
}

func (r *abstractStore) Create(id uint64, data interface{}) error {
	if r.store == nil {
		r.store = make(map[uint64]interface{})
	}

	if id == 0 {
		return errInvalidID
	} else if _, found := r.store[id]; found {
		return db.ErrAlreadyExists
	}
	r.store[id] = data
	return nil
}

func (r *abstractStore) Update(id uint64, data interface{}) error {
	if id == 0 {
		return errInvalidID
	}
	_, found := r.store[id]
	if !found {
		return db.ErrNotFound
	}
	r.store[id] = data
	return nil
}

func (r *abstractStore) NextID() uint64 {
	for {
		r.prevID++
		if _, found := r.store[r.prevID]; !found {
			return r.prevID
		}
	}
}
