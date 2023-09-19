package domain

import (
	"context"
)

/* ------------------------------- REPOSITORY ------------------------------- */
type Repository interface {
	Create(ctx context.Context, turno Turno) (Turno, error)
	GetAll(ctx context.Context) ([]Turno, error)
	GetById(ctx context.Context, id int) (Turno, error)
	Update(ctx context.Context, turno Turno) (Turno, error)
	UpdateDescripcion(ctx context.Context, turno Turno) (Turno, error)
	Delete(ctx context.Context, id int) error
}
