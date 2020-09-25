// Package modelsold contains the types for schema 'public'.
package model

// Equipamento represents a row from 'public.equipamento'.
type Equipamento struct {
	ID        int        `json:"id"`        // id
	Titulo    NullString `json:"titulo"`    // titulo
	Codigo    NullString `json:"codigo"`    // codigo
	Marca     NullInt64  `json:"marca"`     // marca
	Categoria NullInt64  `json:"categoria"` // categoria
	Produto   NullInt64  `json:"produto"`   // produto
	Cliente   int        `json:"cliente"`   // cliente
}
