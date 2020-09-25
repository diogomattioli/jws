// Package modelsold contains the types for schema 'public'.
package model

// Cliente represents a row from 'public.cliente'.
type Cliente struct {
	ID        int        `json:"id"`        // id
	Nome      string     `json:"nome"`      // nome
	Registro  string     `json:"registro"`  // registro
	Cadastro  string     `json:"cadastro"`  // cadastro
	Tipo      int16      `json:"tipo"`      // tipo
	Descricao NullString `json:"descricao"` // descricao
}

func (o *Cliente) Search() []string {
	return []string{"nome", "registro", "cadastro"}
}
