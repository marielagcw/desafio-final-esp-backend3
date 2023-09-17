package domain

import (
	"context"
	"errors"
)

/* --------------------------------- ERRORS --------------------------------- */
var (
	ErrEmptyList = errors.New("the list is empty")
	ErrNotFound  = errors.New("product not found")
	ErrStatement = errors.New("Error Preparing Statement")
	ErrExec      = errors.New("Error Execute Statement")
	ErrLastId    = errors.New("Error getting last ID")
)

/* ------------------------------- REPOSITORY ------------------------------- */
type Repository interface {
	Create(ctx context.Context, odontologo Odontologo) (Odontologo, error)
}
