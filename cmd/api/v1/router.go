package v1

import (
	"encoding/json"
	"fmt"
	"github.com/crashdump/birdview/internal/sync"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// UI
	r.HandleFunc("/", handlerGetRoot).Methods("GET")
	staticFileDirectory := http.Dir(dir + "/ui/")
	staticFileHandler := http.StripPrefix("/ui/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/ui").Handler(staticFileHandler).Methods("GET")

	// API
	r.HandleFunc("/api/v1/sync", handlerPutSync).Methods("PUT")
	r.HandleFunc("/api/v1/resources", handlerGetResources).Methods("GET")
	r.HandleFunc("/api/v1/eol", handlerGetEOLs).Methods("GET")
	r.HandleFunc("/api/v1/events", handlerGetHealthEvents).Methods("GET")
	r.HandleFunc("/api/v1/version", handlerGetVersion).Methods("GET")

	return r
}

func handlerGetRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui/index.html", 302)
}

func handlerGetHealthEvents(w http.ResponseWriter, r *http.Request) {
	s := sync.Init()
	json.NewEncoder(w).Encode(s.GetHealthEvents())
}

func handlerGetResources(w http.ResponseWriter, r *http.Request) {
	s := sync.Init()
	json.NewEncoder(w).Encode(s.GetResources())
}

func handlerGetEOLs(w http.ResponseWriter, r *http.Request) {
	s := sync.Init()
	json.NewEncoder(w).Encode(s.GetServicesEOL())
}

func handlerPutSync(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not implemented yet.")
}

func handlerGetVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "v1.0.0")
}