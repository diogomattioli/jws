// Package modelsold contains the types for schema 'public'.
package model

// Categoria represents a row from 'public.categoria'.
type Categoria struct {
	ID     int    `json:"id"`     // id
	Titulo string `json:"titulo"` // titulo
}
