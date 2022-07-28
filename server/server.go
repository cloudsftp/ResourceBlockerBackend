package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cloudsftp/ResourceBlockerBackend/resource"
	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
}

type AddRequest struct {
	Add int `json:"add"`
}

var status = map[string]*resource.ResourceStatus{}

func StartServer(config *Config) {
	r := mux.NewRouter()
	r.StrictSlash(true)

	server := &Server{config}

	for id, res := range config.Resources {
		status[id] = resource.NewStatus(res)
	}

	r.HandleFunc("/", server.homeHandler).Methods("GET")
	r.HandleFunc("/{name}/", server.resourceHandler).Methods("GET", "POST")

	log.Printf("Running server on port %d\n", config.Port)
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("127.0.0.1:%d", config.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux))
}

func (server *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(server.config)
	if err != nil {
		internalServerError(w)
		log.Print(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (server *Server) resourceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		internalServerError(w)
		log.Printf("Route resourceHandler wrong configured, vars: %v", vars)
		return
	}

	resStatus, ok := status[name]
	if !ok {
		notFound(w)
		log.Printf("resource %s not found", name)
		return
	}

	if r.Method == "POST" {
		var addRequest AddRequest
		json.NewDecoder(r.Body).Decode(&addRequest)
		resStatus.Num += addRequest.Add
	}

	jsonBytes, err := json.Marshal(resStatus)
	if err != nil {
		internalServerError(w)
		log.Print(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}