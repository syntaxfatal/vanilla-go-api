package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/jugadores", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getJugadores(w, r)
		case http.MethodPost:
			postJugadores(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			JsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error pa. error: %s", err)
			}
			w.Write(JsonResp)
		}
	})
}
