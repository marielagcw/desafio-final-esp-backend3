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
func (r *repository) Create(ctx context.Context, paciente Paciente) (Paciente, error) {

	statement, err := r.db.Prepare(QueryInsertPaciente)

	if err != nil {
		return Paciente{}, ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		paciente.Nombre,
		paciente.Apellido,
		paciente.Dni,
		paciente.Domicilio,
		paciente.FechaAlta,
	)

	if err != nil {
		return Paciente{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Paciente{}, ErrLastId
	}

	paciente.ID = int(lastId)

	return paciente, nil
}

/* --------------------------------- GET ALL -------------------------------- */
func (r *repository) GetAll(ctx context.Context) ([]Paciente, error) {

	rows, err := r.db.Query(QueryGetAllPaciente)

	if err != nil {
		return []Paciente{}, ErrStatement
	}

	defer rows.Close()

	var pacientes []Paciente

	for rows.Next() {
		var paciente Paciente

		err := rows.Scan(
			&paciente.ID,
			&paciente.Nombre,
			&paciente.Apellido,
			&paciente.Dni,
			&paciente.Domicilio,
			&paciente.FechaAlta,
		)

		if err != nil {
			return []Paciente{}, ErrExec
		}

		pacientes = append(pacientes, paciente)
	}

	if len(pacientes) == 0 {
		return []Paciente{}, ErrEmptyList
	}

	return pacientes, nil
}

/* --------------------------------- GET BY ID ------------------------------- */
func (r *repository) GetById(ctx context.Context, id int) (Paciente, error) {

	row := r.db.QueryRow(QueryGetByIdPaciente, id)

	var paciente Paciente

	err := row.Scan(
		&paciente.ID,
		&paciente.Nombre,
		&paciente.Apellido,
		&paciente.Dni,
		&paciente.Domicilio,
		&paciente.FechaAlta,
	)

	if err != nil {
		return Paciente{}, ErrExec
	}

	return paciente, nil
}

/* --------------------------------- UPDATE --------------------------------- */
func (r *repository) Update(ctx context.Context, paciente Paciente) (Paciente, error) {

	statement, err := r.db.Prepare(QueryUpdatePaciente)

	if err != nil {
		return Paciente{}, ErrStatement
	}

	defer statement.Close()

	_, err = statement.Exec(
		paciente.Nombre,
		paciente.Apellido,
		paciente.Dni,
		paciente.Domicilio,
		paciente.FechaAlta,
		paciente.ID,
	)

	if err != nil {
		return Paciente{}, ErrExec
	}

	return paciente, nil
}

/* --------------------------------- DELETE --------------------------------- */
func (r *repository) Delete(ctx context.Context, id int) error {

	statement, err := r.db.Prepare(QueryDeletePaciente)

	if err != nil {
		return ErrStatement
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		return ErrExec
	}

	return nil
}

/* --------------------------------- GET BY DNI ------------------------------- */
func (r *repository) GetByDni(ctx context.Context, id int) (Paciente, error) {

	row := r.db.QueryRow(QueryGetByDniPaciente, id)

	var paciente Paciente

	err := row.Scan(
		&paciente.ID,
		&paciente.Nombre,
		&paciente.Apellido,
		&paciente.Dni,
		&paciente.Domicilio,
		&paciente.FechaAlta,
	)

	if err != nil {
		return Paciente{}, err
	}

	return paciente, nil
}
