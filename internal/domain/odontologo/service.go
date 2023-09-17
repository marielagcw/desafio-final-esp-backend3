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
	Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error)
	GetAll(ctx context.Context) ([]Odontologo, error)
	GetById(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error)
}

// NewService creates a new odontologo service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

/* --------------------------------- CREATE --------------------------------- */
func (s *service) Create(ctx context.Context, requestOdontologo RequestOdontologo) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	response, err := s.repository.Create(ctx, odontologo)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Odontologo{}, errors.New("Error en el servicio - Método create")
	}
	return response, nil
}

func requestToOdontologo(requestOdontologo RequestOdontologo) Odontologo {
	var odontologo Odontologo
	odontologo.Nombre = requestOdontologo.Nombre
	odontologo.Apellido = requestOdontologo.Apellido
	odontologo.Matricula = requestOdontologo.Matricula
	return odontologo
}

/* --------------------------------- GET ALL -------------------------------- */
func (s *service) GetAll(ctx context.Context) ([]Odontologo, error) {
	odontologos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("Error en el servicio: ", err.Error())
		return []Odontologo{}, ErrEmptyList
	}
	return odontologos, nil
}

/* -------------------------------- GET BY ID ------------------------------- */
func (s *service) GetById(ctx context.Context, id int) (Odontologo, error) {
	odontologo, err := s.repository.GetById(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err.Error())
		return Odontologo{}, ErrNotFound
	}
	return odontologo, nil
}

/* ------------------------------- UPDATE ALL ------------------------------- */
func (s *service) Update(ctx context.Context, requestOdontologo RequestOdontologo, id int) (Odontologo, error) {
	odontologo := requestToOdontologo(requestOdontologo)
	odontologo.ID = id
	response, err := s.repository.Update(ctx, odontologo)
	if err != nil {
		log.Println("Error en el servicio: ", err.Error())
		return Odontologo{}, ErrNotFound
	}

	return response, nil
}
