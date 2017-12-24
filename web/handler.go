package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	refract "github.com/turtledev/refract-api"
	"github.com/turtledev/refract-api/db"
	"github.com/turtledev/refract-api/db/inmemory"
)

var trackRepository db.TrackRepository

func marshalJSON(i interface{}) string {
	raw, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

func inferValue(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case error:
		return v.Error()
	default:
		return fmt.Sprint(v)
	}
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, marshalJSON(i))
}

func writeErrorJSON(w http.ResponseWriter, i interface{}, status int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	res := map[string]string{"error": inferValue(i)}
	fmt.Fprint(w, marshalJSON(res))
}

func apiHandler() http.Handler {
	// initialise the repository
	trackRepository = new(inmemory.TrackRepository)

	// initialise routes and handlers
	router := mux.NewRouter().PathPrefix("/v0").Subrouter()
	router.Path("/tracks").Methods("GET").HandlerFunc(getAllTracks(trackRepository))
	router.Path("/tracks").Methods("POST").HandlerFunc(createTrack(trackRepository))
	router.Path("/tracks/{id}").Methods("GET").HandlerFunc(getTrackByID(trackRepository))
	router.Path("/tracks/{id}").Methods("DELETE").HandlerFunc(deleteTrackByID(trackRepository))
	return router
}

func getAllTracks(repo db.TrackRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, repo.GetAll())
	}
}

func getTrackByID(repo db.TrackRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		track, err := repo.Get(uint64(id))
		if err != nil {
			var status int
			if err == db.ErrNotFound {
				status = http.StatusNotFound
			} else {
				status = http.StatusBadRequest
			}
			writeErrorJSON(w, err, status)
			return
		}
		writeJSON(w, track)
	}
}

func deleteTrackByID(repo db.TrackRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		err = repo.Delete(uint64(id))
		if err != nil {
			var status int
			if err == db.ErrNotFound {
				status = http.StatusNotFound
			} else {
				status = http.StatusBadRequest
			}
			writeErrorJSON(w, err, status)
			return
		}
		writeJSON(w, map[string]string{"status": "success"})
	}
}

func createTrack(repo db.TrackRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		track := new(refract.Track)
		if err := json.NewDecoder(r.Body).Decode(track); err != nil {
			writeErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		id, err := repo.Create(track)
		if err != nil {
			writeErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		writeJSON(w, map[string]string{"status": "success", "id": strconv.Itoa(int(id))})
	}
}
