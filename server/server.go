package server

import (
	"encoding/json"
	"github.com/xkmsoft/wikisearcher/tcpclient"
	"net/http"
)

const (
	Ip = "localhost"
	Port = "3333"
	Network = "tcp"
)

type QueryParams struct {
	Query string `json:"query"`
}

func HandleQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params QueryParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := tcpclient.NewTCPClient(Ip, Port, Network)
	clientResponse, err := client.Query(params.Query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(clientResponse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

