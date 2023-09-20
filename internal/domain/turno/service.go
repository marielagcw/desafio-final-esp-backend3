package domain

import (
	"context"
	"desafio-final/pkg/errores"
	"errors"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	Create(ctx context.Context, requestTurno RequestTurno) (Turno, error)
	GetAll(ctx context.Context) ([]Turno, error)
	GetById(ctx context.Context, id int) (Turno, error)
	Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error)
	UpdateDescripcion(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error)
	Delete(ctx context.Context, id int) error
}

// NewService creates a new turno service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

/* --------------------------------- CREATE --------------------------------- */
func (s *service) Create(ctx context.Context, requestTurno RequestTurno) (Turno, error) {
	turno := requestToTurno(requestTurno)
	response, err := s.repository.Create(ctx, turno)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Turno{}, errors.New("Error en el servicio - Método create")
	}
	return response, nil
}

/* --------------------------------- GET ALL -------------------------------- */
func (s *service) GetAll(ctx context.Context) ([]Turno, error) {
	turnos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("Error en el servicio: ", err.Error())
		return []Turno{}, errores.ErrEmptyList
	}
	return turnos, nil
}

/* -------------------------------- GET BY ID ------------------------------- */
func (s *service) GetById(ctx context.Context, id int) (Turno, error) {
	turno, err := s.repository.GetById(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err.Error())
		return Turno{}, errores.ErrNotFound
	}
	return turno, nil
}

/* ------------------------------- UPDATE ALL ------------------------------- */
func (s *service) Update(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error) {
	turno := requestToTurno(requestTurno)
	turno.ID = id
	response, err := s.repository.Update(ctx, turno)
	if err != nil {
		log.Println("Error en el servicio: ", err.Error())
		return Turno{}, errores.ErrNotFound
	}

	return response, nil
}

/* --------------------------------- UPDATE DESCRIPCION --------------------------------- */
func (s *service) UpdateDescripcion(ctx context.Context, requestTurno RequestTurno, id int) (Turno, error) {
	turno := requestToTurno(requestTurno)
	turno.ID = id
	response, err := s.repository.UpdateDescripcion(ctx, turno)
	if err != nil {
		log.Println("Error en el servicio: ", err.Error())
		return Turno{}, errores.ErrNotFound
	}

	return response, nil
}

/* --------------------------------- DELETE --------------------------------- */
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err.Error())
		return errores.ErrNotFound
	}
	return nil
}

/* --------------------------------- REQUEST -------------------------------- */
func requestToTurno(requestTurno RequestTurno) Turno {
	var turno Turno
	turno.Descripcion = requestTurno.Descripcion
	turno.Fecha = requestTurno.Fecha
	turno.Hora = requestTurno.Hora
	turno.OdontologoId = requestTurno.OdontologoId
	turno.PacienteId = requestTurno.PacienteId

	return turno
}
