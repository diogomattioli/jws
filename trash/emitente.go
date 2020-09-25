// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
)

// Emitente represents a row from 'public.emitente'.
type Emitente struct {
	ID        int              `json:"id"`        // id
	Nome      string           `json:"nome"`      // nome
	Registro  string           `json:"registro"`  // registro
	Cadastro  string           `json:"cadastro"`  // cadastro
	Im        string           `json:"im"`        // im
	Cnae      string           `json:"cnae"`      // cnae
	Tipo      int16            `json:"tipo"`      // tipo
	Telefone  string           `json:"telefone"`  // telefone
	Descricao model.NullString `json:"descricao"` // descricao
}

func (o *Emitente) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}