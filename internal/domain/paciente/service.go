package domain

import (
	"context"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(ctx context.Context, requestPaciente RequestPaciente) (Paciente, error)
}

// NewService creates a new odontologo service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

/* --------------------------------- CREATE --------------------------------- */
func (s *service) Create(ctx context.Context, requestPaciente RequestPaciente) (Paciente, error) {
	odontologo := requestToPaciente(requestPaciente)
	response, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Paciente{}, errors.New("error en el servicio - Método create")
	}
	return response, nil
}

func requestToPaciente(requestPaciente RequestPaciente) Paciente {
	var paciente Paciente
	paciente.Nombre = requestPaciente.Nombre
	paciente.Apellido = requestPaciente.Apellido
	paciente.Dni = requestPaciente.Dni
	paciente.Domicilio = requestPaciente.Domicilio
	paciente.FechaAlta = requestPaciente.FechaAlta
	return paciente
}
