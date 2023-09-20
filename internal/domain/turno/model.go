package domain

import (
	odontologo "desafio-final/internal/domain/odontologo"
	paciente "desafio-final/internal/domain/paciente"
)

// Turno is a struct that represents how 'Paciente' and 'Odontologo' are schedule
type Turno struct {
	ID          int                   `json:"id"`
	Paciente    paciente.Paciente     `json:"paciente" binding:"required"`
	Odontologo  odontologo.Odontologo `json:"odontologo" binding:"required"`
	Descripcion string                `json:"descripcion"`
	FechaTurno  string                `json:"fecha_turno"`
	HoraTurno   string                `json:"hora_turno"`
}

// RequestTurno represents how request will be handle for 'Turno'
type RequestTurno struct {
	Paciente    int    `json:"paciente" binding:"required"`
	Odontologo  int    `json:"odontologo" binding:"required"`
	Descripcion string `json:"descripcion"`
	FechaTurno  string `json:"fecha_turno"`
	HoraTurno   string `json:"hora_turno"`
}
