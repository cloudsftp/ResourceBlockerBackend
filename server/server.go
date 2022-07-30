package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/cloudsftp/ResourceBlockerBackend/persist"
	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
}

type resourcesResponse struct {
	ResourceIds []string `json:"resources"`
}

func newResourceResponse(config *Config) *resourcesResponse {
	ids := []string{}
	for id := range config.Resources {
		ids = append(ids, id)
	}
	return &resourcesResponse{ResourceIds: ids}
}

type updateStatusRequest struct {
	Delta int `json:"delta"`
}

var resourceLocks = map[string]*sync.Mutex{}

func StartServer(config *Config) {
	persist.InitializeDatabase()

	for id, res := range config.Resources {
		resourceLocks[id] = &sync.Mutex{}
		persist.InitializeStatusIfNotExists(id, res)
	}

	r := mux.NewRouter()
	r.StrictSlash(true)

	server := &Server{config}

	r.HandleFunc("/", server.homeHandler).Methods("GET")
	r.HandleFunc("/{name}/", server.resourceHandler).Methods("GET", "POST")

	log.Printf("Running server on port %d\n", config.Port)
	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func (server *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	response := newResourceResponse(server.config)
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		internalServerError(w)
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (server *Server) resourceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["name"]
	if !ok {
		internalServerError(w)
		log.Printf("Route resourceHandler wrong configured, vars: %v", vars)
		return
	}

	lock, ok := resourceLocks[id]
	if !ok {
		notFound(w)
		log.Printf("lock for resource %s not found", id)
		return
	}
	lock.Lock()
	defer lock.Unlock()

	status, err := persist.GetStatus(id)
	if err != nil {
		notFound(w)
		log.Printf("resource %s not found", id)
		return
	}

	if r.Method == "POST" {
		var req updateStatusRequest
		json.NewDecoder(r.Body).Decode(&req)
		num := status.Num + req.Delta

		resource, ok := server.config.Resources[id]
		if !ok {
			internalServerError(w)
			log.Printf("resource with id %s not found", id)
			return
		}

		if num < resource.Min ||
			num > resource.Max {

			internalServerError(w)
			log.Printf("num %d out of range for resource with id %s %v", num, id, resource)
			return
		}

		status.Num += req.Delta
		persist.UpdateStatus(id, status)
	}

	jsonBytes, err := json.Marshal(status)
	if err != nil {
		internalServerError(w)
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"error\": \"internal\"}"))
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("{\"error\": \"not found\"}"))
}
