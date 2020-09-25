// Package modelsold contains the types for schema 'public'.
package model

// Produto represents a row from 'public.produto'.
type Produto struct {
	ID        int        `json:"id"`        // id
	Titulo    string     `json:"titulo"`    // titulo
	Esquema   NullString `json:"esquema"`   // esquema
	Marca     int        `json:"marca"`     // marca
	Categoria int        `json:"categoria"` // categoria
}