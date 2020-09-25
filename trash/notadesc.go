// Package modelsold contains the types for schema 'public'.
package trash

import (
	"restProject/model"
)

// Notadesc represents a row from 'public.notadesc'.
type Notadesc struct {
	ID          int             `json:"id"`          // id
	Titulo      string          `json:"titulo"`      // titulo
	Preco       float32         `json:"preco"`       // preco
	Ncm         string          `json:"ncm"`         // ncm
	Notaentrada int             `json:"notaentrada"` // notaentrada
	Notasaida   model.NullInt64 `json:"notasaida"`   // notasaida
	Ordem       int             `json:"ordem"`       // ordem
}
