package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	res := JSONResponse{"error": inferValue(i)}
	fmt.Fprint(w, marshalJSON(res))
}

// errorStatusMap is a mapping between errors returned by repositories
// and their corresponding http statuses
type errorStatusMap map[error]int

func (m errorStatusMap) StatusForError(e error) int {
	status, found := m[e]
	if !found {
		status = http.StatusInternalServerError
	}
	return status
}
