// Package modelsold contains the types for schema 'public'.
package model

import (
	"github.com/jinzhu/gorm"
	"net/url"
	"time"
)

// Ordem represents a row from 'public.ordem'.
type Ordem struct {
	ID           int        `json:"id"`            // id
	Cliente      int        `json:"cliente"`       // cliente
	Equipamento  int        `json:"equipamento"`   // equipamento
	Usuario      int        `json:"usuario"`       // usuario
	Garantia     int16      `json:"garantia"`      // garantia
	Arquivo      NullString `json:"arquivo"`       // arquivo
	NotaEntrada  NullString `json:"nota_entrada"`  // nota_entrada
	NotaSaida    NullString `json:"nota_saida"`    // nota_saida
	NotaGarantia NullString `json:"nota_garantia"` // nota_garantia
	Reclamado    NullString `json:"reclamado"`     // reclamado
	Defeito      NullString `json:"defeito"`       // defeito
	Contato      NullString `json:"contato"`       // contato
	Entrada      time.Time  `json:"entrada"`       // entrada
	Conserto     NullTime   `json:"conserto"`      // conserto
	Cobranca     NullTime   `json:"cobranca"`      // cobranca
	Saida        NullTime   `json:"saida"`         // saida
	Observacao   NullString `json:"observacao"`    // observacao
	Descricao    NullString `json:"descricao"`     // descricao
	Pagamento    NullInt64  `json:"pagamento"`     // pagamento
	Status       int16      `json:"status"`        // status
	Situacao     int16      `json:"situacao"`      // situacao
	Aviso        bool       `json:"aviso"`         // aviso
	Lt1          NullTime   `json:"lt1"`           // lt1
	Lt2          NullTime   `json:"lt2"`           // lt2
	Lt3          NullTime   `json:"lt3"`           // lt3
	Lt4          NullTime   `json:"lt4"`           // lt4
	Lt5          NullTime   `json:"lt5"`           // lt5
	Lt6          NullTime   `json:"lt6"`           // lt6
	Lt7          NullTime   `json:"lt7"`           // lt7
}

func (o Ordem) View(db *gorm.DB, obj interface{}, values url.Values) *gorm.DB {
	return db
}