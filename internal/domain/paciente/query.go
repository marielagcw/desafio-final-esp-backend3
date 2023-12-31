package domain

var (
	QueryInsertPaciente   = `INSERT INTO db_desafio_final.paciente (nombre, apellido, dni, domicilio, fecha_alta) VALUES (?,?,?,?,?)`
	QueryGetAllPaciente   = `SELECT id, nombre, apellido, dni, domicilio, fecha_alta FROM db_desafio_final.paciente`
	QueryGetByIdPaciente  = `SELECT id, nombre, apellido, dni, domicilio, fecha_alta FROM db_desafio_final.paciente WHERE id = ?`
	QueryGetByDniPaciente = `SELECT id, nombre, apellido, dni, domicilio, fecha_alta FROM db_desafio_final.paciente WHERE dni = ?`
	QueryUpdatePaciente   = `UPDATE db_desafio_final.paciente SET nombre = ?, apellido = ?, dni = ?, domicilio = ?, fecha_alta = ? WHERE id = ?`
	QueryDeletePaciente   = `DELETE FROM db_desafio_final.paciente WHERE id = ?`
)
