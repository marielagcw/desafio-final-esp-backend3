package domain

import (
	"context"
)

/* ------------------------------- REPOSITORY ------------------------------- */
type Repository interface {
	Create(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	GetAll(ctx context.Context) ([]Odontologo, error)
	GetById(ctx context.Context, id int) (Odontologo, error)
	Update(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	UpdateName(ctx context.Context, odontologo Odontologo) (Odontologo, error)
	Delete(ctx context.Context, id int) error
}
