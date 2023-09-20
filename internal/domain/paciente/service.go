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
	GetAll(ctx context.Context) ([]Paciente, error)
	GetById(ctx context.Context, id int) (Paciente, error)
	GetByDni(ctx context.Context, dni int) (Paciente, error)
	Update(ctx context.Context, id int, requestPaciente RequestPaciente) (Paciente, error)
	Patch(ctx context.Context, id int, requestPaciente RequestPaciente) (Paciente, error)
	Delete(ctx context.Context, id int) error
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

/* --------------------------------- GET ALL -------------------------------- */
func (s *service) GetAll(ctx context.Context) ([]Paciente, error) {
	response, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return []Paciente{}, errors.New("error en el servicio - Método getAll")
	}
	return response, nil
}

/* --------------------------------- GET BY ID ------------------------------- */
func (s *service) GetById(ctx context.Context, id int) (Paciente, error) {
	response, err := s.repository.GetById(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Paciente{}, errors.New("error en el servicio - Método getById")
	}
	return response, nil
}

/* --------------------------------- UPDATE --------------------------------- */
func (s *service) Update(ctx context.Context, id int, requestPaciente RequestPaciente) (Paciente, error) {
	oldPaciente, err := s.repository.GetById(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Paciente{}, errors.New("error en el servicio - Método update")
	}
	newPaciente := requestToPaciente(requestPaciente)
	newPaciente.ID = oldPaciente.ID
	response, err := s.repository.Update(ctx, newPaciente)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Paciente{}, errors.New("error en el servicio - Método update")
	}
	return response, nil
}

/* --------------------------------- PATCH ---------------------------------- */
func (s *service) Patch(ctx context.Context, id int, requestPaciente RequestPaciente) (Paciente, error) {
	oldPaciente, err := s.GetById(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Paciente{}, errors.New("error en el servicio - Método patch")
	}
	newPaciente := requestToPaciente(requestPaciente)
	newPaciente.ID = oldPaciente.ID
	if newPaciente.Nombre == "" {
		newPaciente.Nombre = oldPaciente.Nombre
	}
	if newPaciente.Apellido == "" {
		newPaciente.Apellido = oldPaciente.Apellido
	}
	if newPaciente.Dni == "" {
		newPaciente.Dni = oldPaciente.Dni
	}
	if newPaciente.Domicilio == "" {
		newPaciente.Domicilio = oldPaciente.Domicilio
	}
	if newPaciente.FechaAlta.IsZero() {
		newPaciente.FechaAlta = oldPaciente.FechaAlta
	}
	response, err := s.repository.Update(ctx, newPaciente)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Paciente{}, errors.New("error en el servicio - Método patch")
	}
	return response, nil
}

/* --------------------------------- DELETE --------------------------------- */
func (s *service) Delete(ctx context.Context, id int) error {
	_, err := s.repository.GetById(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return errors.New("error en el servicio - Método delete")
	}
	err = s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return errors.New("error en el servicio - Método delete")
	}
	return nil
}

/* --------------------------------- GET BY ID ------------------------------- */
func (s *service) GetByDni(ctx context.Context, id int) (Paciente, error) {
	response, err := s.repository.GetByDni(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Paciente{}, errors.New("error en el servicio - Método getByDni")
	}
	return response, nil
}

/* --------------------------------- REQUEST TO PACIENTE -------------------------------- */
func requestToPaciente(requestPaciente RequestPaciente) Paciente {
	var paciente Paciente
	paciente.Nombre = requestPaciente.Nombre
	paciente.Apellido = requestPaciente.Apellido
	paciente.Dni = requestPaciente.Dni
	paciente.Domicilio = requestPaciente.Domicilio
	paciente.FechaAlta = requestPaciente.FechaAlta
	return paciente
}
