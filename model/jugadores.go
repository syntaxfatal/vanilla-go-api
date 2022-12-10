package model

type JugadoresStruct struct {
	Id     int
	Nombre string
	Goles  int
}

var Jugadores []*JugadoresStruct = []*JugadoresStruct{}
