package domain

import "time"

// Paciente is a struct that represents a Paciente
type Paciente struct {
	ID        int       `json:"id"`
	Apellido  string    `json:"apellido"`
	Nombre    string    `json:"nombre"`
	Dni       string    `json:"dni"`
	Domicilio string    `json:"domicilio"`
	FechaAlta time.Time `json:"fecha_alta"`
}

// RequestPaciente is a struct that represents a Request
type RequestPaciente struct {
	Apellido  string    `json:"apellido"`
	Nombre    string    `json:"nombre"`
	Dni       string    `json:"dni"`
	Domicilio string    `json:"domicilio"`
	FechaAlta time.Time `json:"fecha_alta"`
}
