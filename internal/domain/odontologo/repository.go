package domain

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

/* --------------------------------- CREATE --------------------------------- */
func (r *repository) Create(ctx context.Context, odontologo Odontologo) (Odontologo, error) {

	statement, err := r.db.Prepare(QueryInsertOdontologo)

	if err != nil {
		return Odontologo{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Nombre,
		odontologo.Apellido,
		odontologo.Matricula,
	)

	if err != nil {
		return Odontologo{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Odontologo{}, err
	}

	odontologo.ID = int(lastId)

	return odontologo, nil
}
