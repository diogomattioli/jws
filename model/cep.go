// Package modelsold contains the types for schema 'public'.
package model

// Cep represents a row from 'public.cep'.
type Cep struct {
	Endereco string `json:"endereco"` // endereco
	Bairro   string `json:"bairro"`   // bairro
	Cep      string `json:"cep"`      // cep
	Cidade   string `json:"cidade"`   // cidade
	Estado   string `json:"estado"`   // estado
}

func (o *Cep) Search() []string {
	return []string{"cep"}
}
