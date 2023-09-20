package domain

import (
	"context"
	"errors"
)

/* --------------------------------- ERRORS --------------------------------- */
var (
	// ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("Turno not found")
	ErrStatement = errors.New("error Preparing Statement")
	ErrExec      = errors.New("error Execute Statement")
	ErrLastId    = errors.New("error Getting Last ID")
)

/* ------------------------------- REPOSITORY ------------------------------- */
type Repository interface {
	Create(ctx context.Context, turno Turno) (Turno, error)
	GetAll(ctx context.Context) ([]Turno, error)
	GetAllByPacienteDni(ctx context.Context, pacienteDni int) ([]Turno, error)
	GetById(ctx context.Context, id int) (Turno, error)
	Update(ctx context.Context, turno Turno) (Turno, error)
	Delete(ctx context.Context, id int) error
}
