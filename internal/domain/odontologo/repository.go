package domain

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// NewRepository creates a new repository
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

/* --------------------------------- CREATE --------------------------------- */
func (r *repository) Create(ctx context.Context, odontologo Odontologo) (Odontologo, error) {

	statement, err := r.db.Prepare(QueryInsertOdontologo)

	if err != nil {
		return Odontologo{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Nombre,
		odontologo.Apellido,
		odontologo.Matricula,
	)

	if err != nil {
		return Odontologo{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Odontologo{}, ErrLastId
	}

	odontologo.ID = int(lastId)

	return odontologo, nil
}

/* --------------------------------- GET ALL -------------------------------- */
func (r *repository) GetAll(ctx context.Context) ([]Odontologo, error) {
	rows, err := r.db.Query(QueryGetAllOdontologos)
	if err != nil {
		return []Odontologo{}, err
	}

	defer rows.Close()

	var odontologos []Odontologo

	for rows.Next() {
		var odontologo Odontologo
		err := rows.Scan(
			&odontologo.ID,
			&odontologo.Apellido,
			&odontologo.Nombre,
			&odontologo.Matricula,
		)
		if err != nil {
			return []Odontologo{}, err
		}

		odontologos = append(odontologos, odontologo)
	}

	if err := rows.Err(); err != nil {
		return []Odontologo{}, err
	}

	return odontologos, nil
}

/* -------------------------------- GET BY ID ------------------------------- */
func (r *repository) GetById(ctx context.Context, id int) (Odontologo, error) {
	row := r.db.QueryRow(QueryGetByIdOdontologo, id)

	var odontologo Odontologo
	err := row.Scan(
		&odontologo.ID,
		&odontologo.Apellido,
		&odontologo.Nombre,
		&odontologo.Matricula,
	)

	if err != nil {
		return Odontologo{}, err
	}

	return odontologo, nil
}

/* --------------------------------- UPDATE ALL --------------------------------- */
func (r *repository) Update(ctx context.Context, odontologo Odontologo) (Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return Odontologo{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Nombre,
		odontologo.Apellido,
		odontologo.Matricula,
		odontologo.ID,
	)

	if err != nil {
		return Odontologo{}, ErrExec
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Odontologo{}, err
	}

	if rowsAffected < 1 {
		return Odontologo{}, ErrNotFound
	}

	return odontologo, nil
}

/* -------------------------------- UPDATE NAME -------------------------------- */
func (r *repository) UpdateName(ctx context.Context, odontologo Odontologo) (Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateNameOdontologo)
	if err != nil {
		return Odontologo{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Nombre,
		odontologo.ID,
	)

	if err != nil {
		return Odontologo{}, ErrExec
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Odontologo{}, err
	}

	if rowsAffected < 1 {
		return Odontologo{}, ErrNotFound
	}

	return odontologo, nil
}

/* --------------------------------- DELETE --------------------------------- */
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteOdontologo, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}
