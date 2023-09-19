package domain

import (
	"context"
	"database/sql"

	"health-center/pkg/errores"
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
func (r *repository) Create(ctx context.Context, turno Turno) (Turno, error) {

	statement, err := r.db.Prepare(QueryInsertTurno)

	if err != nil {
		return Turno{}, errores.ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.Fecha,
		turno.Hora,
		turno.Descripcion,
		turno.OdontologoId,
		turno.PacienteId,
	)

	if err != nil {
		return Turno{}, errores.ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Turno{}, errores.ErrLastId
	}

	turno.ID = int(lastId)

	return turno, nil
}

/* --------------------------------- GET ALL -------------------------------- */
func (r *repository) GetAll(ctx context.Context) ([]Turno, error) {
	rows, err := r.db.Query(QueryGetAllTurnos)
	if err != nil {
		return []Turno{}, err
	}

	defer rows.Close()

	var turnos []Turno

	for rows.Next() {
		var turno Turno
		err := rows.Scan(
			&turno.ID,
			&turno.Fecha,
			&turno.Hora,
			&turno.Descripcion,
			&turno.OdontologoId,
			&turno.PacienteId,
		)
		if err != nil {
			return []Turno{}, err
		}

		turnos = append(turnos, turno)
	}

	if err := rows.Err(); err != nil {
		return []Turno{}, err
	}

	return turnos, nil
}

/* -------------------------------- GET BY ID ------------------------------- */
func (r *repository) GetById(ctx context.Context, id int) (Turno, error) {
	row := r.db.QueryRow(QueryGetByIdTurno, id)

	var turno Turno
	err := row.Scan(
		&turno.ID,
		&turno.Fecha,
		&turno.Hora,
		&turno.Descripcion,
		&turno.OdontologoId,
		&turno.PacienteId,
	)

	if err != nil {
		return Turno{}, err
	}

	return turno, nil
}

/* --------------------------------- UPDATE ALL --------------------------------- */
func (r *repository) Update(ctx context.Context, turno Turno) (Turno, error) {
	statement, err := r.db.Prepare(QueryUpdateTurno)
	if err != nil {
		return Turno{}, errores.ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.Fecha,
		turno.Hora,
		turno.Descripcion,
		turno.OdontologoId,
		turno.PacienteId,
	)

	if err != nil {
		return Turno{}, errores.ErrExec
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Turno{}, err
	}

	if rowsAffected < 1 {
		return Turno{}, errores.ErrNotFound
	}

	return turno, nil
}

/* -------------------------------- UPDATE DESCRIPCION -------------------------------- */
func (r *repository) UpdateDescripcion(ctx context.Context, turno Turno) (Turno, error) {
	statement, err := r.db.Prepare(QueryUpdateDescripcionTurno)
	if err != nil {
		return Turno{}, errores.ErrStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.Descripcion,
		turno.ID,
	)

	if err != nil {
		return Turno{}, errores.ErrExec
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Turno{}, err
	}

	if rowsAffected < 1 {
		return Turno{}, errores.ErrNotFound
	}

	return turno, nil
}

/* --------------------------------- DELETE --------------------------------- */
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteTurno, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errores.ErrNotFound
	}

	return nil
}
