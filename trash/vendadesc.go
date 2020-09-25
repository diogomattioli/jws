// Package modelsold contains the types for schema 'public'.
package trash

import (
	"restProject/model"
)

// Vendadesc represents a row from 'public.vendadesc'.
type Vendadesc struct {
	ID         int             `json:"id"`         // id
	Titulo     string          `json:"titulo"`     // titulo
	Preco      float32         `json:"preco"`      // preco
	Quantidade int             `json:"quantidade"` // quantidade
	Venda      int             `json:"venda"`      // venda
	Estoque    model.NullInt64 `json:"estoque"`    // estoque
}
