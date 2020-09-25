// Package modelsold contains the types for schema 'public'.
package model

// Endereco represents a row from 'public.endereco'.
type Endereco struct {
	ID          int        `json:"id"`          // id
	Numero      string     `json:"numero"`      // numero
	Complemento string     `json:"complemento"` // complemento
	Endereco    string     `json:"endereco"`    // endereco
	Bairro      string     `json:"bairro"`      // bairro
	Cep         string     `json:"cep"`         // cep
	Cidade      string     `json:"cidade"`      // cidade
	Estado      string     `json:"estado"`      // estado
	Descricao   NullString `json:"descricao"`   // descricao
	Cliente     int        `json:"cliente"`     // cliente
}
