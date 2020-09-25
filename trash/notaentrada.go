// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
	"time"
)

// Notaentrada represents a row from 'public.notaentrada'.
type Notaentrada struct {
	ID             int               `json:"id"`             // id
	Numero         string            `json:"numero"`         // numero
	Data           time.Time         `json:"data"`           // data
	Arquivo        model.NullString  `json:"arquivo"`        // arquivo
	Descricao      model.NullString  `json:"descricao"`      // descricao
	Transportadora model.NullInt64   `json:"transportadora"` // transportadora
	Frete          model.NullFloat64 `json:"frete"`          // frete
	Remetente      model.NullBool    `json:"remetente"`      // remetente
	Cliente        int               `json:"cliente"`        // cliente
}

func (o *Notaentrada) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}