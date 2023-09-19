package errores

import (
	"errors"
)

/* --------------------------------- ERRORS --------------------------------- */
var (
	ErrEmptyList = errors.New("The list is empty")
	ErrNotFound  = errors.New("Odontologo not found")
	ErrStatement = errors.New("Error Preparing Statement")
	ErrExec      = errors.New("Error Execute Statement")
	ErrLastId    = errors.New("Error Getting Last ID")
)
