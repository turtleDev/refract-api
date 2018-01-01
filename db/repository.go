package db

import (
	refract "github.com/turtledev/refract-api"
)

type TrackRepository interface {
	Get(uint64) (*refract.Track, error)
	GetAll() []*refract.Track
	Delete(uint64) error
	Create(*refract.Track) (uint64, error)
	Update(uint64, *refract.Track) error
}

type TeamRepository interface {
	GetAll() []*refract.Team
	Create(*refract.Team) (uint64, error)
}
