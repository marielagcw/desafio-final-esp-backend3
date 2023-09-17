package domain

var (
	QueryInsertOdontologo = `INSERT INTO db_desafio_final.odontologo (nombre, apellido, matricula) VALUES (?,?,?)`
	QueryGetAllOdontologos = `SELECT id, apellido, nombre, matricula FROM db_desafio_final.odontologo`
)
