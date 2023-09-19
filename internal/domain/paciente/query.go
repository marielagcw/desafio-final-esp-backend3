package domain

var (
	QueryInsertPaciente = `INSERT INTO db_desafio_final.paciente (nombre, apellido, dni, domicilio, fecha_alta) VALUES (?,?,?,?,?)`
)
