package domain

var (
	QueryInsertTurno            = `INSERT INTO db_desafio_final.turno (fecha, hora, descripcion, odontologo_id, paciente_id) VALUES (?,?,?,?,?)`
	QueryGetAllTurnos           = `SELECT id, fecha, hora, descripcion, odontologo_id, paciente_id FROM db_desafio_final.turno`
	QueryGetByIdTurno           = `SELECT id, fecha, hora, descripcion, odontologo_id, paciente_id FROM db_desafio_final.turno WHERE id = ?`
	QueryUpdateTurno            = `UPDATE db_desafio_final.turno SET fecha = ?, hora = ?, descripcion = ?, odontologo_id = ?, paciente_id = ? WHERE id = ?`
	QueryUpdateDescripcionTurno = `UPDATE db_desafio_final.turno SET descripcion = ? WHERE id = ?`
	QueryDeleteTurno            = `DELETE FROM db_desafio_final.turno WHERE id = ?`
)
