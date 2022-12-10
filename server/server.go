package server

import "net/http"

func NewServer(port string) *http.Server {
	InitRoutes()
	return &http.Server{
		Addr: port,
	}
}
