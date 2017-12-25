package inmemory

import (
	"errors"

	refract "github.com/turtledev/refract-api"
	"github.com/turtledev/refract-api/db"
)

var (
	errInvalidID = errors.New("ID must be non-zero")
)

type TrackRepository struct {
	tracks map[uint64]*refract.Track
	prevID uint64
}

func (t *TrackRepository) GetAll() []*refract.Track {
	tracks := []*refract.Track{}
	for _, track := range t.tracks {
		tracks = append(tracks, track)
	}
	return tracks
}

func (t *TrackRepository) Get(id uint64) (*refract.Track, error) {
	track, ok := t.tracks[id]
	if !ok {
		return nil, db.ErrNotFound
	}
	return track, nil
}

func (t *TrackRepository) Delete(id uint64) error {
	if _, ok := t.tracks[id]; !ok {
		return db.ErrNotFound
	}
	delete(t.tracks, id)
	return nil
}

func (t *TrackRepository) Create(track *refract.Track) (uint64, error) {
	if t.tracks == nil {
		t.tracks = make(map[uint64]*refract.Track)
	}

	if track.ID == 0 {
		for {
			t.prevID++
			if _, found := t.tracks[t.prevID]; !found {
				track.ID = t.prevID
				break
			}
		}
	} else if _, found := t.tracks[track.ID]; found {
		return 0, db.ErrAlreadyExists
	}
	t.tracks[track.ID] = track
	return track.ID, nil
}

func (t *TrackRepository) Update(id uint64, track *refract.Track) error {
	if id == 0 {
		return errInvalidID
	}
	// should this be legal?
	track.ID = id
	_, found := t.tracks[id]
	if !found {
		return db.ErrNotFound
	}
	t.tracks[id] = track
	return nil
}
