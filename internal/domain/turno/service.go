package domain

import (
	"context"
	odontologo "desafio-final/internal/domain/odontologo"
	paciente "desafio-final/internal/domain/paciente"
	"errors"
	"fmt"
	"log"
	"time"
)

type service struct {
	repository        Repository
	pacienteService   paciente.Service
	odontologoService odontologo.Service
}

type Service interface {
	Create(ctx context.Context, requestTurno RequestTurno) (Turno, error)
	GetAll(ctx context.Context) ([]Turno, error)
	GetAllByPacienteDni(ctx context.Context, pacienteID int) ([]Turno, error)
	GetById(ctx context.Context, id int) (Turno, error)
	Update(ctx context.Context, id int, requestTurno RequestTurno) (Turno, error)
	Delete(ctx context.Context, id int) error
}

// NewService creates a new 'Turno' service
func NewService(repository Repository, pacienteService paciente.Service, odontologoService odontologo.Service) Service {
	return &service{
		repository:        repository,
		pacienteService:   pacienteService,
		odontologoService: odontologoService,
	}
}

/* --------------------------------- CREATE --------------------------------- */
func (s *service) Create(ctx context.Context, requestTurno RequestTurno) (Turno, error) {
	turno, err := requestToTurno(s, ctx, requestTurno)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Turno{}, errors.New(fmt.Sprintf("error en el servicio - Método create: %s", err))
	}

	response, err := s.repository.Create(ctx, turno)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Turno{}, errors.New(fmt.Sprintf("error en el servicio - Método create: %s", err))
	}
	return response, nil
}

/* --------------------------------- GET ALL -------------------------------- */
func (s *service) GetAll(ctx context.Context) ([]Turno, error) {
	response, err := s.repository.GetAll(ctx)
	if err != nil {
		return []Turno{}, errors.New(fmt.Sprintf("error en el servicio - Método getAll. %s", err))
	}

	var turnos []Turno
	for _, t := range response {
		appendTurno, err := requestToTurno(s, ctx, turnoToRequest(t))
		if err != nil {
			return nil, err
		}
		appendTurno.ID = t.ID
		turnos = append(turnos, appendTurno)
	}

	return turnos, nil
}

/* --------------------------------- GET BY ID ------------------------------- */
func (s *service) GetById(ctx context.Context, id int) (Turno, error) {
	response, err := s.repository.GetById(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Turno{}, errors.New("error en el servicio - Método getById")
	}

	turno, err := requestToTurno(s, ctx, turnoToRequest(response))
	if err != nil {
		return Turno{}, err
	}
	turno.ID = id
	return turno, nil
}

/* --------------------------------- UPDATE --------------------------------- */
func (s *service) Update(ctx context.Context, id int, requestTurno RequestTurno) (Turno, error) {
	turno, err := s.repository.GetById(ctx, id)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Turno{}, errors.New("error en el servicio - Método update")
	}
	newTurno, err := requestToTurno(s, ctx, requestTurno)
	if err != nil {
		return Turno{}, err
	}

	newTurno.ID = turno.ID
	response, err := s.repository.Update(ctx, newTurno)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return Turno{}, errors.New("error en el servicio - Método update")
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

/* --------------------------------- GET ALL BY 'Paciente' Dni -------------------------------- */
func (s *service) GetAllByPacienteDni(ctx context.Context, pacienteDni int) ([]Turno, error) {
	paciente, err := s.pacienteService.GetByDni(ctx, pacienteDni)
	if err != nil {
		return nil, err
	}

	response, err := s.repository.GetAllByPacienteDni(ctx, paciente.ID)
	if err != nil {
		log.Println("Error en el servicio: ", err)
		return []Turno{}, errors.New("error en el servicio - Método getAll")
	}

	var turnos []Turno
	for _, t := range response {
		appendTurno, err := requestToTurno(s, ctx, turnoToRequest(t))
		if err != nil {
			return nil, err
		}
		appendTurno.ID = t.ID
		turnos = append(turnos, appendTurno)
	}

	return turnos, nil
}

/* ---------------------------------MAPPING REQUEST TO TURNO -------------------------------- */
func requestToTurno(s *service, ctx context.Context, requestTurno RequestTurno) (Turno, error) {
	var err error
	var turno Turno

	mapperPaciente, errPaciente := s.pacienteService.GetById(ctx, requestTurno.Paciente)
	if errPaciente != nil {
		log.Println("Error mapeando turno: ", errPaciente)
		return turno, errors.New("error durante el mapeo del Paciente correspondiente al Turno")
	}

	mapperOdontologo, errOdontologo := s.odontologoService.GetById(ctx, requestTurno.Odontologo)
	if errOdontologo != nil {
		log.Println("Error mapeando turno: ", errOdontologo)
		return turno, errors.New("error durante el mapeo del Odontologo correspondiente al Turno")
	}
	if turno.FechaTurno != "" {
		_, err = time.Parse("2006-01-02", requestTurno.FechaTurno)
		if err != nil {
			log.Println("Error mapeando turno: ", err)
			return turno, errors.New("error durante el mapeo de la fecha(yyyy-MM-dd) correspondiente al Turno")
		}

	}

	if turno.HoraTurno != "" {
		_, err = time.Parse("15:04:05", requestTurno.HoraTurno)
		if err != nil {
			log.Println("Error mapeando turno: ", err)
			return turno, errors.New("error durante el mapeo de la hora(hh:mm:ss) correspondiente al Turno")
		}

	}

	turno.Paciente = mapperPaciente
	turno.Odontologo = mapperOdontologo
	turno.Descripcion = requestTurno.Descripcion
	turno.FechaTurno = requestTurno.FechaTurno
	turno.HoraTurno = requestTurno.HoraTurno
	return turno, nil
}

/* ---------------------------------MAPPING TURNO TO REQUEST -------------------------------- */
func turnoToRequest(turno Turno) RequestTurno {
	var requestTurno RequestTurno

	requestTurno.Paciente = turno.Paciente.ID
	requestTurno.Odontologo = turno.Odontologo.ID
	requestTurno.Descripcion = turno.Descripcion
	requestTurno.FechaTurno = turno.FechaTurno
	requestTurno.HoraTurno = turno.HoraTurno
	return requestTurno

}
