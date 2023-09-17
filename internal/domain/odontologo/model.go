package domain

// Odontologo is a struct that represents a Odontologo
type Odontologo struct {
	ID        int    `json:"id"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
	Matricula string `json:"matricula"`
}

// RequestOdontologo is a struct that represents a Request
type RequestOdontologo struct {
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
	Matricula string `json:"matricula"`
}
