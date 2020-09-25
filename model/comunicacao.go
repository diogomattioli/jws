// Package modelsold contains the types for schema 'public'.
package model

import (
	"database/sql"
)

// Comunicacao represents a row from 'public.comunicacao'.
type Comunicacao struct {
	ID        int            `json:"id"`        // id
	Valor     string         `json:"valor"`     // valor
	Tipo      int16          `json:"tipo"`      // tipo
	Descricao sql.NullString `json:"descricao"` // descricao
	Cliente   int            `json:"cliente"`   // cliente
}
