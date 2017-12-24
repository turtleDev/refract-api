package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	refract "github.com/turtledev/refract-api"
	"github.com/turtledev/refract-api/db"
	"github.com/turtledev/refract-api/db/inmemory"
)

var tracks = []*refract.Track{
	{
		Title:    "[DnB] - Priority One & TwoThirds - Hunted (feat. Jonny Rose) [Monstercat Release] - YouTube",
		Duration: 264,
		URL:      "https://www.youtube.com/watch?v=F7Uz_XD3Lhk",
		Type:     refract.TrackTypeYoutube,
	},
	{
		Title:    "TheFatRat - Dancing Naked - YouTube",
		Duration: 278,
		URL:      "https://www.youtube.com/watch?v=oqTHiUgCjDA",
		Type:     refract.TrackTypeYoutube,
	},
	{
		Title:    "Distance Soundtrack - Departure - YouTube",
		Duration: 247,
		URL:      "https://www.youtube.com/watch?v=nMFjeiPY45c",
		Type:     refract.TrackTypeYoutube,
	},
}

var trackRepository db.TrackRepository

func marshalJSON(i interface{}) string {
	raw, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

func apiHandler() http.Handler {
	// initialise the repository
	trackRepository = new(inmemory.TrackRepository)
	for _, track := range tracks {
		trackRepository.Create(track)
	}

	// initialise routes and handlers
	router := mux.NewRouter()
	router.Path("/v0/tracks").Methods("GET").HandlerFunc(GetAllTracksHandler(trackRepository))
	return router
}

func GetAllTracksHandler(repo db.TrackRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, marshalJSON(repo.GetAll()))
	}
}
