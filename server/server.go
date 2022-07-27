package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func StartServer(config Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello")
	})

	fmt.Printf("Running server on port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
