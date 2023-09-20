package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type repository struct {
	db *sql.DB
}

// NewRepository creates a new repository for 'Turno'
func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

/* Inserts a new record for 'Turno'. Task-> POST: agregar turno.*/
func (r *repository) Create(ctx context.Context, turno Turno) (Turno, error) {

	var err error
	var fechaNula sql.NullString
	var result sql.Result

	statement, err := r.db.Prepare(QueryInsertTurno)

	if err != nil {
		return Turno{}, ErrStatement
	}

	defer statement.Close()

	if turno.FechaTurno != "" {
		result, err = statement.Exec(
			turno.Paciente.ID,
			turno.Odontologo.ID,
			turno.Descripcion,
			turno.FechaTurno,
			turno.HoraTurno,
		)
	} else {
		result, err = statement.Exec(
			turno.Paciente.ID,
			turno.Odontologo.ID,
			turno.Descripcion,
			fechaNula,
			turno.HoraTurno,
		)
	}

	if err != nil {
		return Turno{}, ErrExec
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return Turno{}, ErrLastId
	}

	turno.ID = int(lastId)

	return turno, nil
}

/* Retrieves a collection of records 'Turno' stored in database */
func (r *repository) GetAll(ctx context.Context) ([]Turno, error) {
	var fechaNula sql.NullString
	var horaNula sql.NullString
	var fechaTurno time.Time
	rows, err := r.db.Query(QueryGetAllTurno)

	if err != nil {
		return []Turno{}, err
	}

	defer rows.Close()

	var turnos []Turno

	for rows.Next() {
		var turno Turno

		err := rows.Scan(
			&turno.ID,
			&turno.Paciente.ID,
			&turno.Odontologo.ID,
			&turno.Descripcion,
			&fechaNula,
			&horaNula,
		)

		if err != nil {
			return []Turno{}, err
		}
		if fechaNula.Valid {
			log.Printf("Date value: '%s'", fechaNula.String)
			fechaTurno, err = time.Parse(time.RFC3339, fechaNula.String)
			if err != nil {
				return []Turno{}, errors.New(fmt.Sprintf("Error durante mapeo de fecha en turno: '%s'", err))
			}
			turno.FechaTurno = fechaTurno.Format("2006-01-02")
		} else {
			turno.FechaTurno = ""
		}

		if horaNula.Valid {
			log.Printf("Hour value: '%s'", horaNula.String)
			horaTurno, err := time.Parse("15:04:05", horaNula.String)
			if err != nil {
				return []Turno{}, err
			}
			turno.HoraTurno = horaTurno.Format("15:04:05")
		} else {
			turno.HoraTurno = ""
		}

		turnos = append(turnos, turno)
	}

	return turnos, nil
}

/* Finds a 'Turno' by its ID. Task-> GET: traer turno por ID. */
func (r *repository) GetById(ctx context.Context, id int) (Turno, error) {
	var fechaNula sql.NullString
	var horaNula sql.NullString
	row := r.db.QueryRow(QueryGetByIdTurno, id)

	var turno Turno

	err := row.Scan(
		&turno.ID,
		&turno.Paciente.ID,
		&turno.Odontologo.ID,
		&turno.Descripcion,
		&fechaNula,
		&horaNula,
	)

	if err != nil {
		log.Println(err)
		return Turno{}, err
	}

	if fechaNula.Valid {
		fechaTurno, err := time.Parse(time.RFC3339, fechaNula.String)
		if err != nil {
			return Turno{}, err
		}
		turno.FechaTurno = fechaTurno.Format("2006-01-02")

	} else {
		turno.FechaTurno = ""
	}
	if horaNula.Valid {
		horaTurno, err := time.Parse("15:04:05", horaNula.String)
		if err != nil {
			return Turno{}, err
		}
		turno.HoraTurno = horaTurno.Format("15:04:05")

	} else {
		turno.HoraTurno = ""
	}

	return turno, nil
}

/* Updates a 'Turno' with new values into database. Task-> PUT: actualizar turno. */
func (r *repository) Update(ctx context.Context, turno Turno) (Turno, error) {

	statement, err := r.db.Prepare(QueryUpdateTurno)

	if err != nil {
		return Turno{}, ErrStatement
	}

	defer statement.Close()

	_, err = statement.Exec(
		turno.Paciente.ID,
		turno.Odontologo.ID,
		turno.Descripcion,
		turno.FechaTurno,
		turno.HoraTurno,
		turno.ID,
	)

	if err != nil {
		return Turno{}, ErrExec
	}

	return turno, nil
}

/* Deletes specific record by ID. Task-> DELETE: eliminar turno. */
func (r *repository) Delete(ctx context.Context, id int) error {

	statement, err := r.db.Prepare(QueryDeleteTurno)

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

/* Retrieves a collection of records 'Turno' stored in database filtered by specific 'Paciente' Dni */
func (r *repository) GetAllByPacienteDni(ctx context.Context, pacienteID int) ([]Turno, error) {
	var fechaNula sql.NullString
	var horaNula sql.NullString
	var fechaTurno time.Time

	rows, err := r.db.Query(QueryGetAllTurnoByPaciente, pacienteID)

	if err != nil {
		return []Turno{}, ErrStatement
	}

	defer rows.Close()

	var turnos []Turno

	for rows.Next() {
		var turno Turno

		err := rows.Scan(
			&turno.ID,
			&turno.Paciente.ID,
			&turno.Odontologo.ID,
			&turno.Descripcion,
			&turno.FechaTurno,
			&turno.HoraTurno,
		)

		if err != nil {
			return []Turno{}, ErrExec
		}

		if fechaNula.Valid {
			fechaTurno, err = time.Parse(time.RFC3339, fechaNula.String)
			if err != nil {
				return []Turno{}, err
			}
			turno.FechaTurno = fechaTurno.Format("2006-01-02")
		} else {
			turno.FechaTurno = ""
		}

		if horaNula.Valid {
			log.Printf("Hour value: '%s'", horaNula.String)
			horaTurno, err := time.Parse("15:04:05", horaNula.String)
			if err != nil {
				return []Turno{}, err
			}
			turno.HoraTurno = horaTurno.Format("15:04:05")
		} else {
			turno.HoraTurno = ""
		}

		turnos = append(turnos, turno)
	}

	return turnos, nil
}
