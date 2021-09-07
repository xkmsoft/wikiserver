package main

import (
	"github.com/gorilla/mux"
	"github.com/xkmsoft/wikiserver/server"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/query", server.MakeGzipHandler(server.HandleQuery)).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
