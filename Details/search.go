package details

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type response struct {
	title    string    `json:"title"`
	location string    `json:"location"`
	date     time.Time `json:"date"`
	seat     [2][2]int `json:"seat"`
	avail    bool      `json:"avail"`
	details  string    `json:"detail"`
	id       int       `json:"id"`
}

type request struct {
	title string `json:title`
}

var events = []Event{
	{ID: 1, Title: "Spiderman", Location: "PVR Vellore", Date: time.Now(), Avail: true, Details: "Action movie"},
	{ID: 2, Title: "Coldplay Concert", Location: "DLF Arena", Date: time.Now(), Avail: true, Details: "Live concert"},
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	q := strings.ToLower(r.URL.Query().Get("q"))

	var results []Event
	for _, e := range events {
		if q == "" || strings.Contains(strings.ToLower(e.Title), q) {
			results = append(results, e)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		log.Println("encode error:", err)
	}
}

func EventInfo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	for _, e := range events {
		if id == "" {
			continue
		}
		if idStr := strconvItoa(e.ID); idStr == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(e)
			return
		}
	}
	http.Error(w, "event not found", http.StatusNotFound)
}
