// Package modelsold contains the types for schema 'public'.
package model

// Impressao represents a row from 'public.impressao'.
type Impressao struct {
	ID       int        `json:"id"`       // id
	Titulo   string     `json:"titulo"`   // titulo
	Logo     NullString `json:"logo"`     // logo
	Entrada  NullString `json:"entrada"`  // entrada
	Venda    NullString `json:"venda"`    // venda
	Endereco NullString `json:"endereco"` // endereco
}
