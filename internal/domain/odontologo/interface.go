package domain

import "context"

/* --------------------------------- ERRORS --------------------------------- */

/* ------------------------------- REPOSITORY ------------------------------- */
type Repository interface {
	Create(ctx context.Context, odontologo Odontologo) (Odontologo, error)
}
