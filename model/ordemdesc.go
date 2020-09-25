// Package modelsold contains the types for schema 'public'.
package model

// Ordemdesc represents a row from 'public.ordemdesc'.
type Ordemdesc struct {
	ID         int       `json:"id"`         // id
	Titulo     string    `json:"titulo"`     // titulo
	Preco      float32   `json:"preco"`      // preco
	Quantidade int       `json:"quantidade"` // quantidade
	Ordem      int       `json:"ordem"`      // ordem
	Estoque    NullInt64 `json:"estoque"`    // estoque
}