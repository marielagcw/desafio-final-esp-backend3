package domain

var (
	QueryInsertTurno           = `INSERT INTO db_desafio_final.turno (paciente_id, odontologo_id, descripcion, fecha, hora) VALUES (?,?,?,?,?)`
	QueryGetAllTurno           = `SELECT id, paciente_id, odontologo_id, descripcion, fecha, hora FROM db_desafio_final.turno`
	QueryGetAllTurnoByPaciente = `SELECT id, paciente_id, odontologo_id, descripcion, fecha, hora FROM db_desafio_final.turno WHERE paciente_id = ?`
	QueryGetByIdTurno          = `SELECT id, paciente_id, odontologo_id, descripcion, fecha, hora FROM db_desafio_final.turno WHERE id = ?`
	QueryUpdateTurno           = `UPDATE db_desafio_final.turno SET paciente_id = ?, odontologo_id = ?, descripcion = ?, fecha = ?, hora = ? WHERE id = ?`
	QueryDeleteTurno           = `DELETE FROM db_desafio_final.turno WHERE id = ?`
)
