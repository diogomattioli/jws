// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
)

// Transportadoraend represents a row from 'public.transportadoraend'.
type Transportadoraend struct {
	ID             int              `json:"id"`             // id
	Numero         string           `json:"numero"`         // numero
	Complemento    string           `json:"complemento"`    // complemento
	Endereco       string           `json:"endereco"`       // endereco
	Bairro         string           `json:"bairro"`         // bairro
	Cep            string           `json:"cep"`            // cep
	Cidade         string           `json:"cidade"`         // cidade
	Estado         string           `json:"estado"`         // estado
	Descricao      model.NullString `json:"descricao"`      // descricao
	Transportadora int              `json:"transportadora"` // transportadora
}

func (o *Transportadoraend) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}