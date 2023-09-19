package domain

import (
	"context"
	"errors"
)

/* --------------------------------- ERRORS --------------------------------- */
var (
	ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("paciente not found")
	ErrStatement = errors.New("error Preparing Statement")
	ErrExec      = errors.New("error Execute Statement")
	ErrLastId    = errors.New("error Getting Last ID")
)

/* ------------------------------- REPOSITORY ------------------------------- */
type Repository interface {
	Create(ctx context.Context, paciente Paciente) (Paciente, error)
	GetAll(ctx context.Context) ([]Paciente, error)
}
