// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
)

// Transportadora represents a row from 'public.transportadora'.
type Transportadora struct {
	ID        int              `json:"id"`        // id
	Nome      string           `json:"nome"`      // nome
	Registro  string           `json:"registro"`  // registro
	Cadastro  string           `json:"cadastro"`  // cadastro
	Tipo      int16            `json:"tipo"`      // tipo
	Descricao model.NullString `json:"descricao"` // descricao
}

func (o *Transportadora) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}