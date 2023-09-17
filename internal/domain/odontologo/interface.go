package domain

import (
	"context"
	"errors"
)

/* --------------------------------- ERRORS --------------------------------- */
var (
	ErrEmptyList = errors.New("The list is empty")
	ErrNotFound  = errors.New("Odontologo not found")
	ErrStatement = errors.New("Error Preparing Statement")
	ErrExec      = errors.New("Error Execute Statement")
	ErrLastId    = errors.New("Error Getting Last ID")
)

/* ------------------------------- REPOSITORY ------------------------------- */
type Repository interface {
	Create(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	GetAll(ctx context.Context) ([]Odontologo, error)
	GetById(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, odontologo Odontologo) (Odontologo, error)
}
