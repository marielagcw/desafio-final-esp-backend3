package domain

type Turno struct {
	ID           int    `json:"id"`
	Fecha        string `json:"fecha"`
	Hora         string `json:"hora"`
	Descripcion  string `json:"descripcion"`
	OdontologoId int    `json:"odontologo_id"`
	PacienteId   int    `json:"paciente_id"`
}

type RequestTurno struct {
	Fecha        string `json:"fecha"`
	Hora         string `json:"hora"`
	Descripcion  string `json:"descripcion"`
	OdontologoId int    `json:"odontologo_id"`
	PacienteId   int    `json:"paciente_id"`
}
