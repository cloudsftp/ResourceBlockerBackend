package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
}

func (server *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(server.config)
}

func (server *Server) resourceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	resource := server.config.Resources[name]
	json.NewEncoder(w).Encode(resource)
}

func StartServer(config *Config) {
	r := mux.NewRouter()
	r.StrictSlash(true)

	server := &Server{config}

	r.HandleFunc("/", server.homeHandler).Methods("GET")
	r.HandleFunc("/{name}/", server.resourceHandler).Methods("GET", "POST")

	fmt.Printf("Running server on port %d\n", config.Port)
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
