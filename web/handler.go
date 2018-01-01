package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	refract "github.com/turtledev/refract-api"
	"github.com/turtledev/refract-api/db"
	"github.com/turtledev/refract-api/db/inmemory"
)

var (
	trackRepository db.TrackRepository
	teamRepository  db.TeamRepository
)

var statusMap = errorStatusMap{
	db.ErrNotFound:      http.StatusNotFound,
	db.ErrAlreadyExists: http.StatusBadRequest,
}

func apiHandler() http.Handler {
	// initialise the repository
	trackRepository = new(inmemory.TrackRepository)
	teamRepository = new(inmemory.TeamRepository)

	// can return an error
	teamRepository.Create(&refract.Team{Name: "devs and hackers", Domain: "dev-s"})

	// initialise routes and handlers
	router := mux.NewRouter().PathPrefix("/v0").Subrouter()
	router.Path("/tracks").Methods("GET").HandlerFunc(getAllTracks(trackRepository))
	router.Path("/tracks").Methods("POST").HandlerFunc(createTrack(trackRepository))
	router.Path("/tracks/{id}").Methods("GET").HandlerFunc(getTrackByID(trackRepository))
	router.Path("/tracks/{id}").Methods("DELETE").HandlerFunc(deleteTrackByID(trackRepository))
	router.Path("/tracks/{id}").Methods("POST", "PUT").Handler(updateTrack(trackRepository))

	router.Path("/teams").Methods("GET").HandlerFunc(getAllTeams(teamRepository))

	handler := gorillaHandlers.CORS()(router)
	return handler
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
			writeErrorJSON(w, err, statusMap.StatusForError(err))
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
			writeErrorJSON(w, err, statusMap.StatusForError(err))
			return
		}
		writeJSON(w, JSONResponse{"status": "success"})
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
		writeJSON(w, JSONResponse{"status": "success", "id": strconv.Itoa(int(id))})
	}
}

func updateTrack(repo db.TrackRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			writeErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		track := new(refract.Track)
		if err := json.NewDecoder(r.Body).Decode(track); err != nil {
			writeErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		err = repo.Update(uint64(id), track)
		if err != nil {
			writeErrorJSON(w, err, statusMap.StatusForError(err))
		}
		writeJSON(w, JSONResponse{"status": "success", "id": strconv.Itoa(int(id))})
	}
}

func getAllTeams(repo db.TeamRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, repo.GetAll())
	}
}
