// Package modelsold contains the types for schema 'public'.
package trash

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"restProject/model"
	"time"
)

// Notasaida represents a row from 'public.notasaida'.
type Notasaida struct {
	ID             int               `json:"id"`             // id
	Numero         string            `json:"numero"`         // numero
	Emitente       int               `json:"emitente"`       // emitente
	Data           time.Time         `json:"data"`           // data
	Cfop           string            `json:"cfop"`           // cfop
	Arquivo        model.NullString  `json:"arquivo"`        // arquivo
	Descricao      model.NullString  `json:"descricao"`      // descricao
	Transportadora model.NullInt64   `json:"transportadora"` // transportadora
	Frete          model.NullFloat64 `json:"frete"`          // frete
	Sedex          model.NullBool    `json:"sedex"`          // sedex
	Transporte     int16             `json:"transporte"`     // transporte
	Cliente        int               `json:"cliente"`        // cliente
	Clientealt     model.NullInt64   `json:"clientealt"`     // clientealt
	Volume         model.NullString  `json:"volume"`         // volume
	Pesoliq        model.NullFloat64 `json:"pesoliq"`        // pesoliq
	Pesobru        model.NullFloat64 `json:"pesobru"`        // pesobru
}

func (o *Notasaida) Search(db *gorm.DB, values url.Values) *gorm.DB {
	return model.Search(db, values)
}