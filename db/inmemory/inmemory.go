package inmemory

import (
	refract "github.com/turtledev/refract-api"
)

type TrackRepository struct {
	store abstractStore
}

func (t *TrackRepository) GetAll() []*refract.Track {
	dataset := t.store.GetAll()
	tracks := make([]*refract.Track, 0)
	for _, data := range dataset {
		// can panic
		track := data.(*refract.Track)
		tracks = append(tracks, track)
	}
	return tracks
}

func (t *TrackRepository) Get(id uint64) (*refract.Track, error) {
	data, err := t.store.Get(id)
	if err != nil {
		return nil, err
	}
	track := data.(*refract.Track)
	return track, err
}

func (t *TrackRepository) Delete(id uint64) error {
	return t.store.Delete(id)
}

func (t *TrackRepository) Create(track *refract.Track) (uint64, error) {
	if track.ID == 0 {
		track.ID = t.store.NextID()
	}
	return track.ID, t.store.Create(track.ID, track)
}

func (t *TrackRepository) Update(id uint64, track *refract.Track) error {
	track.ID = id
	return t.store.Update(id, track)
}

type TeamRepository struct {
}
