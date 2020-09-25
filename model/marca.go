// Package modelsold contains the types for schema 'public'.
package model

// Marca represents a row from 'public.marca'.
type Marca struct {
	ID     int    `json:"id"`     // id
	Titulo string `json:"titulo"` // titulo
}
