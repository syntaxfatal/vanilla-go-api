package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vanilla-go-api/model"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["msg"] = "Hello world!"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error pa. error: %s", err)
	}
	w.Write(jsonResp)
}

func getJugadores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Jugadores)
}

func postJugadores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jugador := &model.JugadoresStruct{}
	var resp map[string]string

	err := json.NewDecoder(r.Body).Decode(jugador)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp["msg"] = "Bad request pa"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error pa. error: %s", err)
		}
		w.Write(jsonResp)
	}

	model.Jugadores = append(model.Jugadores, jugador)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{msg: %s}", "Se agrego el jugador")
}
