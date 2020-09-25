// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
	"time"
)

// Notavenda represents a row from 'public.notavenda'.
type Notavenda struct {
	ID             int               `json:"id"`             // id
	Numero         string            `json:"numero"`         // numero
	Data           time.Time         `json:"data"`           // data
	Cfop           string            `json:"cfop"`           // cfop
	Arquivo        model.NullString  `json:"arquivo"`        // arquivo
	Descricao      model.NullString  `json:"descricao"`      // descricao
	Transportadora model.NullInt64   `json:"transportadora"` // transportadora
	Frete          model.NullFloat64 `json:"frete"`          // frete
	Remetente      model.NullBool    `json:"remetente"`      // remetente
	Cliente        int               `json:"cliente"`        // cliente
	Clientealt     model.NullInt64   `json:"clientealt"`     // clientealt
	Volume         model.NullString  `json:"volume"`         // volume
	Lacre          model.NullInt64   `json:"lacre"`          // lacre
	Pesoliq        model.NullFloat64 `json:"pesoliq"`        // pesoliq
	Pesobru        model.NullFloat64 `json:"pesobru"`        // pesobru
}

func (o *Notavenda) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}